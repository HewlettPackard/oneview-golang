package icsp

/*
(c) Copyright [2015] Hewlett Packard Enterprise Development LP

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

// NetConfigInterface - part of NetCustomization type , describes interface configuration
type NetConfigInterface struct {
	Name           string   `json:"name,omitempty"`           // optional name of the nic for the interface, also known as slot or interface name
	WINSServers    []string `json:"winsServers,omitempty"`    // optional [8.8.8.8, 8.8.4.4] optional list of wins servers to configure
	DNSServers     []string `json:"dnsServers,omitempty"`     // optional [8.8.8.8, 8.8.4.4] optional list of dns servers to configure
	DNSSearch      []string `json:"dnsSearch,omitempty"`      // optional [corp.net, my.corp.net] A list of period (.) separated character strings separated by spaces or commas.
	MACAddr        string   `json:"macAddress,omitempty"`     // mac address to configure for the interface
	Enabled        bool     `json:"enabled,omitempty"`        // boolean flag to enable network configuration or leave as is.
	DHCPv4         bool     `json:"dhcpv4,omitempty"`         // boolean flag when set to true no other options are needed, will use dhcp to configure
	IPv6Autoconfig bool     `json:"ipv6Autoconfig,omitempty"` // boolean to automatically configure ipv6
	IPv4Gateway    string   `json:"ipv4gateway,omitempty"`    // when dhcpv4 is false , assume ipv4 config, specify the gateway
	StaticNetworks []string `json:"staticNetworks,omitempty"` // static ips to assign for the server, ie; 172.0.0.2/255.255.255.0
}

// NetConfig - create network customization objects that can be serialized and deserialized
// for configuration to run when executing to be saved on the server as hpsa_netconfig:
// Proliant SW - Post Install Network Personalization build plans
type NetConfig struct {
	Hostname      string               `json:"hostname,omitempty"`  // host1, optional hostname option
	Workgroup     string               `json:"workgroup,omitempty"` // ams, optional WINS workgroup for windows only
	Domain        string               `json:"domain,omitempty"`    // corp.net, optional This is a period (.) separated string and is usually the right side part of a fully qualified host name.
	WINSList      utils.Nstring        // comma seperated list of wins servers
	DNSNameList   utils.Nstring        // comma seperated list of dns servers
	DNSSearchList utils.Nstring        // comma seperated list of dns search servers
	Interfaces    []NetConfigInterface `json:"interfaces,omitempty"` // list of network interfaces to customize
}

// SplitSep - split seperator
const SplitSep = ","

// NewNetConfig - create a new netconfig object without interfaces
func NewNetConfig(hostname utils.Nstring, workgroup utils.Nstring, domain utils.Nstring,
	winslist utils.Nstring, dnsnamelist utils.Nstring, dnssearchlist utils.Nstring) NetConfig {
	var netconfig NetConfig
	netconfig = NetConfig{
		WINSList:      winslist,
		DNSNameList:   dnsnamelist,
		DNSSearchList: dnssearchlist,
	}
	if !hostname.IsNil() {
		netconfig.Hostname = hostname.String()
	}
	if !workgroup.IsNil() {
		netconfig.Workgroup = workgroup.String()
	}
	if !domain.IsNil() {
		netconfig.Domain = domain.String()
	}
	return netconfig
}

// NewNetConfigInterface - creates an interface object for NetConfig
func (n NetConfig) NewNetConfigInterface(
	enable bool,
	macaddr string,
	isdhcp bool,
	isipv6 bool,
	ipv4gateway utils.Nstring, // ipv4 gateway, required when isdhcp is false
	staticnets utils.Nstring, // comma seperated list of ip's, required when isdhcp is false
	name utils.Nstring, // optional name
	wins utils.Nstring, // comma seperated list of wins servers
	dnsservers utils.Nstring, // comma seperated list of dns servers
	dnssearch utils.Nstring) NetConfigInterface { // comma seperated list of dns search

	var inetconfig NetConfigInterface

	if macaddr == "" {
		log.Fatal("Network configuration (NetConfigInterface) requires a MAC Address to create a new interface object.")
	}
	inetconfig = NetConfigInterface{
		Enabled:        enable,
		MACAddr:        macaddr,
		DHCPv4:         isdhcp,
		IPv6Autoconfig: isipv6,
	}
	if !isdhcp {
		if ipv4gateway.IsNil() {
			log.Fatal("Static ipv4 configuration requires a gateway configured (IPv4Gateway)")
		}
		inetconfig.IPv4Gateway = ipv4gateway.String()
		if staticnets.IsNil() {
			log.Fatal("Static ipv4 configuration requires static network list")
		}
		inetconfig.StaticNetworks = strings.Split(staticnets.String(), SplitSep)
	}
	if !name.IsNil() {
		inetconfig.Name = name.String()
	}
	if !wins.IsNil() {
		inetconfig.WINSServers = strings.Split(wins.String(), SplitSep)
	}
	if !dnsservers.IsNil() {
		inetconfig.DNSServers = strings.Split(dnsservers.String(), SplitSep)
	}
	if !dnssearch.IsNil() {
		inetconfig.DNSSearch = strings.Split(dnssearch.String(), SplitSep)
	}
	return inetconfig
}

// AddAllDHCP - make all the netconfig interfaces setup for dhcp
func (n NetConfig) AddAllDHCP(interfaces []Interface, isipv6 bool) {
	var emptystring utils.Nstring
	var emptyinterfaces []NetConfigInterface
	emptystring.Nil()
	n.Interfaces = emptyinterfaces
	for _, iface := range interfaces {
		n.Interfaces = append(n.Interfaces, n.NewNetConfigInterface(true, iface.MACAddr, true, isipv6,
			emptystring, emptystring, utils.Nstring(iface.Slot),
			n.WINSList, n.DNSNameList, n.DNSSearchList))
	}
}

// SetStaticInterface - converts an interface from NetConfig.Interface from dhcp to static
func (n NetConfig) SetStaticInterface(iface Interface, ipv4gateway utils.Nstring, staticiplist utils.Nstring, isipv6 bool) {
	var inet NetConfigInterface
	var bupdated bool
	bupdated = false
	inet = n.NewNetConfigInterface(true, iface.MACAddr, false, isipv6,
		ipv4gateway, staticiplist, utils.Nstring(iface.Slot),
		n.WINSList, n.DNSNameList, n.DNSSearchList)
	// update the existing interfaces in the list
	for i, neti := range n.Interfaces {
		if neti.MACAddr == iface.MACAddr {
			n.Interfaces[i] = inet
			bupdated = true
		}
	}
	// append to the interfaces if we didn't update it
	if !bupdated {
		n.Interfaces = append(n.Interfaces, inet)
	}
}

// toJSON - convert object to JSON string
func (n NetConfig) toJSON() (string, error) {
	data, err := json.Marshal(n)
	return fmt.Sprintf("%s", bytes.NewBuffer(data)), err
}

// Save - save the netconfig to hpsa_netconfig
func (n NetConfig) Save(s Server) {
	data, err := n.toJSON()
	if err != nil {
		log.Fatalf("Unable to save hpsa_netconfig for server, %s", err)
	}
	s.SetCustomAttribute("hpsa_netconfig", "server", data)
}
