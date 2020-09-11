package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {

	var (
		ClientOV *ov.OVClient
	)

	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		2000,
		"*")
	// Get all tasks present
	sort := "name:desc"
	count := "5"
	fmt.Println("\nGetting all tasks present: \n")
	task_list, err := ovc.GetTasks("", sort, count, "")
	if err != nil {
		fmt.Println("Error getting the storage volumes ", err)
	}
	for i := 0; i < len(task_list.Members); i++ {
		fmt.Println(task_list.Members[i].Name)
	}

	id := "a3af4d5e-7114-4e7a-a6c4-f97b707ec87c"

	fmt.Println("Getting Details of the requested Task")
	task, err := ovc.GetTasksById("", "", "", "", id)
	if err != nil {
		fmt.Println("Error getting the task details ", err)
	}
	fmt.Println(task)
}
