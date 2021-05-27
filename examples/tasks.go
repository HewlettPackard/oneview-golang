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

	fmt.Println("Get a tree of tasks")
	task_list_tree, err := ovc.GetTasks("taskState='Completed'", "", "10", "tree", "", "")
	if err != nil {
		fmt.Println("Error getting the task details with tree view ", err)
	}

	fmt.Println(task_list_tree.Members)

	fmt.Println("Get a aggregate tree")
	task_list_atree, err := ovc.GetTasks("taskState='Completed'", "", "", "aggregatedTree", "2", "2")
	if err != nil {
		fmt.Println("Error getting the task with status Completed and aggregatedTree", err)
	}

	fmt.Println(task_list_atree.Members)

	fmt.Println("Get a flat tree")
	task_list_flat, err := ovc.GetTasks("taskState='Completed'", "", "1", "flat-tree", "", "")
	if err != nil {
		fmt.Println("Error getting the task details with status Completed", err)
	}

	fmt.Println(task_list_flat.Members)

	fmt.Println("Perform Patch operation")
	task_list_patch, err := ovc.GetTasks("taskState='Running'", "", "", "", "", "")

	for i := 0; i < len(task_list_patch.Members); i++ {
		if task_list_patch.Members[i].IsCancellable {
			task_uri := task_list_patch.Members[0].URI
			task, err = ovc.PatchTask(task_uri.String())
			if err != nil {
				fmt.Println("Error updating the task details ", err)
			}
			fmt.Println(task)

		}
	}

}
