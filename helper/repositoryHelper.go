package helper

import "github.com/nazzarr03/To-Do-Rest-Api/config"

func GetMaxListID() uint {
	maxListID := uint(0)
	for _, item := range config.ToDoLists {
		if item.ListID > maxListID {
			maxListID = item.ListID
		}
	}
	return maxListID + 1
}
