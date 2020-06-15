package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV *ov.OVClient
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

	timeConfigs, err := ovc.GetTimeConfigs()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get Time Config list----#")
		for i := range timeConfigs.Members {
			fmt.Println(timeConfigs.Members[i].DisplayName)
		}
	}

}
