package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
	"strings"
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
	// Get all tasks present
	sort := "name:desc"
	count := "5"
	fmt.Println("\nGetting all tasks present: \n")
	task_list, err := ovc.GetTasks("", sort, count, "", "", "")
	if err != nil {
		fmt.Println("Error getting the storage volumes ", err)
	}
	for i := 0; i < len(task_list.Members); i++ {
		fmt.Println(task_list.Members[i].Name, task_list.Members[i].URI)
	}

	id := strings.Replace(string(task_list.Members[0].URI), "/rest/tasks/", "", 1)

	fmt.Println("Getting Details of the requested Task")
	task, err := ovc.GetTasksById("", "", "", "", id)
	if err != nil {
		fmt.Println("Error getting the task details ", err)
	}
	fmt.Println(task)

	fmt.Println("Get a tree of tasks")
	filter := []string{"taskState='Completed'"}
	task_list_tree, err := ovc.GetTasks(filter, "", 10, "tree", "", "")
	fmt.Println(task_list_tree.Members)

	fmt.Println("Get a aggregate tree")
	task_list_atree, err := ovc.GetTasks(filter, "", "", "aggregatedTree", 2, 2)
	fmt.Println(task_list_atree.Members)

	fmt.Println("Get a flat tree")
	filter = []string{"status=Warning OR status=OK"}
	task_list_flat, err := ovc.GetTasks(filter, "", 1, "flat-tree")
	fmt.Println(task_list_flat.Members)

	fmt.Println("Perform Patch operation")
	filter = []string{"taskState='Running'", "isCancellable='true'"}
	task_list_patch, err := ovc.GetTasks(filter, "", "", "", "", "")
	task_uri := tasks_list_patch.Members[0].URI
    	task, err = tasks.patch(task_uri)
	if err != nil {
		fmt.Println("Error updating the task details ", err)
	}
	fmt.Println(task)

}
