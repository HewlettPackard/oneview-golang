package main

import (
	"fmt"
	"strings"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {

	var (
		ClientOV *ov.OVClient
	)
	// Use configuratin file to set the ip and  credentails
	config, config_err := ov.LoadConfigFile("config.json")
	if config_err != nil {
		fmt.Println(config_err)
	}
	ovc := ClientOV.NewOVClient(
		config.OVCred.UserName,
		config.OVCred.Password,
		config.OVCred.Domain,
		config.OVCred.Endpoint,
		config.OVCred.SslVerify,
		config.OVCred.ApiVersion,
		config.OVCred.IfMatch)
	// Get all tasks present
	sort := "name:desc"
	count := "5"
	fmt.Println("\nGetting all tasks present: \n")
	task_list, err := ovc.GetTasks("", sort, count, "", "", "")
	if err != nil {
		fmt.Println("Error getting the tasks ", err)
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

	fmt.Println("\nGet a tree of tasks")
	task_list_tree, err := ovc.GetTasks("taskState='Completed'", "", "10", "tree", "", "")
	if err != nil {
		fmt.Println("Error getting the task details with tree view ", err)
	}
	for i := 0; i < len(task_list_tree.Members); i++ {
		fmt.Println(task_list_tree.Members[i].Name, task_list_tree.Members[i].URI)
	}

	fmt.Println("\nGet a aggregate tree\n")
	task_list_atree, err := ovc.GetTasks("taskState='Completed'", "", "", "aggregatedTree", "2", "2")
	if err != nil {
		fmt.Println("Error getting the task with status Completed and aggregatedTree", err)
	}
	// Printing first 10 tasks
	for i := 0; i < 10; i++ {
		fmt.Println(task_list_atree.Members[i].Name, task_list_atree.Members[i].URI)
	}

	fmt.Println("\nGet a flat tree")
	task_list_flat, err := ovc.GetTasks("taskState='Completed'", "", "1", "flat-tree", "", "")
	if err != nil {
		fmt.Println("Error getting the task details with status Completed", err)
	}

	for i := 0; i < len(task_list_flat.Members); i++ {
		fmt.Println(task_list_flat.Members[i].Name, task_list_flat.Members[i].URI)
	}

	fmt.Println("\nPerform Patch operation")
	task_list_patch, err := ovc.GetTasks("taskState='Running'", "", "", "", "", "")

	for i := 0; i < len(task_list_patch.Members); i++ {
		if task_list_patch.Members[i].IsCancellable {
			task_uri := task_list_patch.Members[0].URI
			err = ovc.PatchTask(task_uri.String())
			if err != nil {
				fmt.Println("Error updating the task details ", err)
			}
			fmt.Println(task.URI, task.TaskState)

		}
	}

}
