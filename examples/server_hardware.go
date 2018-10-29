package main 
import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"time"
//	"github.com/HewlettPackard/oneview-golang/rest"
)

func main() {
	var ( 	
		ClientOV  *ov.OVClient
	)
	//var serverName string 
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		300)

	ovVer, _ :=ovc.GetAPIVersion()
	fmt.Println(ovVer)
	//print("Get list of all server hardware resources")
	println("Get server hardware list by name")
	time.Sleep(2)
	//serverName,err := ovc.GetServerHardwareByName("eco-dl360-3-ilo")
	serverName,_ := ovc.GetServerHardwareByName("eco-dl360-2-ilo")

	time.Sleep(2)
	fmt.Println("******************")
	fmt.Println(serverName.Model)
	fmt.Println(serverName.URI)
	fmt.Println(serverName.UUID)
	fmt.Println("******************")

	fmt.Println("Get server hardware list by url")
	ServerId,_ := ovc.GetServerHardware("/rest/server-hardware/32353537-3935-5355-4536-303733503533")
	fmt.Println(ServerId.URI)
	time.Sleep(2)

	fmt.Println("Get server-hardware list statistics specifying parameters")
	filters := []string{"name matches 'eco-dl360-3-ilo'"}
	ServerList, _:= ovc.GetServerHardwareList(filters,"name:ascending")
	fmt.Println("Total server list :",ServerList.Total)

	print("Get server-hardware whoes profiles are unassigned ")
	serverHarware,_ := ovc.GetAvailableHardware("/rest/server-hardware-types/9F529AA5-2021-4A10-93ED-DDC5BD80E949","/rest/enclosure-groups/293e8efe-c6b1-4783-bf88-2d35a8e49071")
	fmt.Println(serverHarware.Type)


}
	
