package main

import (
	"fmt"
	m "todolist/models"    // alias m
	t "todolist/models/control/task"  // alias t
)

func main()  {
	fmt.Println(m.Name)
	fmt.Println(t.Task)
	fmt.Println(t.Version)
}