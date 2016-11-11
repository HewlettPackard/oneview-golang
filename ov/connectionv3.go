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

// Package ov -
package ov


import (
	//"github.com/HewlettPackard/oneview-golang/utils"
	"fmt"
)

func (c *OVClient) ManageI3SConnections(connections []Connection) ([]Connection, error) {

	/*
	deployNet, err := c.GetEthernetConnectionByName("deploy.net")
	if err != nil || deployent.URI.IsNil(){
		return connections, fmt.Errorf("Could not find deployment ethernet network name: deploy.net")
	}*/

	availablePortIds := []string{"Mezz 3:1-a", "Mezz 3:1-b", "Mezz 3:1-c", "Mezz 3:1-d",
								  "Mezz 3:2-a", "Mezz 3:2-b", "Mezz 3:2-c", "Mezz 3:2-d"}
	deployConnections := make([]Connection, 2)
	for i := 0; i < len(connections); i++{
		if(connections[i].Name == "Deployment Network A") {
			deployConnections[0] = connections[i]
		} else if (connections[i].Name == "Deployment Network B") {
			deployConnections[1] = connections[i]
		}
		for j := 0; j < len(availablePortIds); j++ {
			if connections[i].PortID == availablePortIds[j] {
				availablePortIds = append(availablePortIds[:i], availablePortIds[i+1:]...)
			}
		}
	}
	return connections, fmt.Errorf("%+v", availablePortIds)
	/*
	if deployConnections[0] = nil {
		boot1 := BootOption{
			Priority: "Primary",
		}
		connection1 := Connection {
			ID: 1,
			Name: "Deployment Network A",
			FunctionType: "Ethernet",
			RequestedMbps: "2500",
			NetworkURI: deployNet.URI,
			Boot: boot1,
			PortID: "Mezz 3:1-a",
		}
	}*/

	/*
	if len(connections) < 4 {
		boot1 := BootOption{
			Priority: "Primary",
		}
		connection1 := Connection {
			ID: 1,
			Name: "Deployment Network A",
			FunctionType: "Ethernet",
			RequestedMbps: "2500",
			NetworkURI: utils.NewNstring("/rest/ethernet-networks/8d4a7e7f-2c9d-4d79-be86-defd0dd5e8cf"),
			Boot: boot1,
			PortID: "Mezz 3:1-a",
		}

		boot2 := BootOption{
			Priority: "Secondary",
		}
		connection2 := Connection {
			ID: 2,
			Name: "Deployment Network B",
			FunctionType: "Ethernet",
			RequestedMbps: "2500",
			NetworkURI: utils.NewNstring("/rest/ethernet-networks/8d4a7e7f-2c9d-4d79-be86-defd0dd5e8cf"),
			Boot: boot2,
			PortID: "Mezz 3:2-a",
		}
		connections = append(connections, connection1)
		connections = append(connections, connection2)
		connections[0].PortID = "Mezz 3:1-b"
		connections[1].PortID = "Mezz 3:2-b"
	} else {
		connections[0].Boot.Priority = "Primary"
		connections[1].Boot.Priority = "Secondary"
	} */

	return connections, nil

}