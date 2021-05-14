package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HewlettPackard/oneview-golang/ov"
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
	localelist, err := ovc.GetLocales()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Got the supported appliance locales.----#")
		for i := range localelist.Members {
			fmt.Println("Member " + strconv.Itoa(i+1))
			fmt.Print("DisplayName : " + localelist.Members[i].DisplayName + "\n")
			fmt.Print("Locale : " + localelist.Members[i].Locale + "\n")
			fmt.Print("LocaleName : " + localelist.Members[i].LocaleName + "\n")
		}
	}
}
