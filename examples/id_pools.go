package main

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV     *ov.OVClient
		poolTypeVmac = "vmac"
		poolTypeVsn = "vsn"
		poolTypevWWN = "vwwn"
		poolTypeIPV4 = "ipv4"
		enable       = true
	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")

	// Gets Pool Type
	idPool, err := ovc.GetPoolType(poolTypeVmac)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get ID Pool Type----------------#")
		jsonResponse, _ := json.MarshalIndent(idPool, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	// Updates Pool Type
	enable_pool := ov.IdPool{
		Type:      "Pool",
		Enabled:   &enable,
		RangeUris: idPool.RangeUris,
	}
	res, err := ovc.UpdatePoolType(enable_pool, poolTypeVmac)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------ID Pool Updated----------------#")
		jsonResponse, _ := json.MarshalIndent(res, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	// Allocates Ids
	allocate := ov.UpdateAllocatorList{
		Count: 5,
	}
	response, err := ovc.Allocator(allocate, poolTypeVsn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------ID Allocated----------------#")
                jsonResponse, _ := json.MarshalIndent(response, "", "  ")
                fmt.Print(string(jsonResponse), "\n\n")
	}

	// Checks the range availability in the ID pool.
	responseAvail, err := ovc.GetRangeAvailibility(poolTypeIPV4,[]string{"10.10.10.8"})
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println("#-------------Check Range Availibility----------------#")
                jsonResponse, _ := json.MarshalIndent(responseAvail, "", "  ")
                fmt.Print(string(jsonResponse), "\n\n")
        }

	// Collects Ids
	CollectId := ov.UpdateCollectorList{
		IdList:	[]utils.Nstring{"10.1.20.16"},
	}
	responseCol, err := ovc.Collector(CollectId, poolTypeIPV4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------ID Collected----------------#")
                jsonResponse, _ := json.MarshalIndent(responseCol, "", "  ")
                fmt.Print(string(jsonResponse), "\n\n")
	}

	responseIds, err := ovc.Generate(poolTypeVmac)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println("#-------------Fragments----------------#")
                jsonResponse, _ := json.MarshalIndent(responseIds, "", "  ")
                fmt.Print(string(jsonResponse), "\n\n")
        }

	// Validates a set of IDs to reserve in the pool.
	responseValidateIds, err := ovc.GetValidateIds(poolTypeIPV4,[]string{"10.10.10.8"})
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println("#-------------Validate IDs----------------#")
                jsonResponse, _ := json.MarshalIndent(responseValidateIds, "", "  ")
                fmt.Print(string(jsonResponse), "\n\n")
        }

	// Validate a set of user specified IDs to reserve in the pool whose poolType is vWWN 
	data := ov.UpdateAllocatorList{
		IdList: []utils.Nstring{"10:00:2c:6c:28:80:00:00", "10:00:2c:6c:28:80:00:01"},
	}
	responseData, err := ovc.UpdateValidateIds(data, poolTypevWWN)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Update Validated IDs----------------#")
                jsonResponse, _ := json.MarshalIndent(responseData, "", "  ")
                fmt.Print(string(jsonResponse), "\n\n")
	}
}
