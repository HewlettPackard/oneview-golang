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
	"github.com/HewlettPackard/oneview-golang/utils"
)

func (c *OVClient) ManageI3SConnections(connections []Connection) []Connection {

	if len(connections) < 4 {
		boot1 := BootOption{
			Priority: "Primary",
		}
		connection1 := Connection{
			ID:            1,
			Name:          "Deployment Network A",
			FunctionType:  "Ethernet",
			RequestedMbps: "2500",
			NetworkURI:    utils.NewNstring("/rest/ethernet-networks/8d4a7e7f-2c9d-4d79-be86-defd0dd5e8cf"),
			Boot:          boot1,
			PortID:        "Mezz 3:1-a",
		}

		boot2 := BootOption{
			Priority: "Secondary",
		}
		connection2 := Connection{
			ID:            2,
			Name:          "Deployment Network B",
			FunctionType:  "Ethernet",
			RequestedMbps: "2500",
			NetworkURI:    utils.NewNstring("/rest/ethernet-networks/8d4a7e7f-2c9d-4d79-be86-defd0dd5e8cf"),
			Boot:          boot2,
			PortID:        "Mezz 3:2-a",
		}
		connections = append(connections, connection1)
		connections = append(connections, connection2)
		connections[0].PortID = "Mezz 3:1-b"
		connections[1].PortID = "Mezz 3:2-b"
	} else {
		connections[0].Boot.Priority = "Primary"
		connections[1].Boot.Priority = "Secondary"
	}

	return connections

}
