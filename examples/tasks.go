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
		2200,
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
		fmt.Println(task_list.Members[i].Name, task_list.Members[i].URI)
	}

	id := "cc0d18ca-ff2d-4da4-8b16-ee0ad864fb2c"

	fmt.Println("Getting Details of the requested Task")
	task, err := ovc.GetTasksById("", "", "", "", id)
	if err != nil {
		fmt.Println("Error getting the task details ", err)
	}
	fmt.Println(task)

}
