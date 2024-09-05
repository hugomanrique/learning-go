package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Name        string
	Description string
	Done        bool
}

type ListTask struct {
	Tasks []Task
}

func (l *ListTask) addTask(t Task) {
	l.Tasks = append(l.Tasks, t)
}

func (l *ListTask) taskDone(idx int) {
	l.Tasks[idx].Done = true
}

func (l *ListTask) editTask(idx int, t Task) {
	l.Tasks[idx] = t
}

func (l *ListTask) removeTask(idx int) {
	l.Tasks = append(l.Tasks[:idx], l.Tasks[idx+1:]...)
}

func (l *ListTask) printTasks() {
	for _, task := range l.Tasks {
		fmt.Println(task.Name, task.Description, task.Done)
	}
}

func main() {
	list := ListTask{}
	reader := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Println("1. Add Task")
		fmt.Println("2. Edit Task")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Print Tasks")
		fmt.Println("5. Status Task")
		fmt.Println("6. Exit")
		fmt.Scan(&option)
		switch option {
			case 1:
				var task Task
				fmt.Print("Name: ")
				task.Name, _ = reader.ReadString('\n')
				fmt.Print("Description: ")
				task.Description, _ = reader.ReadString('\n')
				list.addTask(task)
			case 2:
				var idx int
				fmt.Println("Index:")
				fmt.Scan(&idx)
				var task Task
				fmt.Print("Name:")
				task.Name,_ = reader.ReadString('\n')
				fmt.Print("Description:")
				task.Description,_ = reader.ReadString('\n')
				list.editTask(idx, task)
			case 3:
				var idx int
				fmt.Println("Index:")
				fmt.Scanln(&idx)
				list.removeTask(idx)
			case 4:
				list.printTasks()
			case 5:
				var idx int
				list.printTasks()
				fmt.Println("Index:")
				fmt.Scanln(&idx)
				list.removeTask(idx)
				return
			case 6:
				return
		}
		list.printTasks()
	}
}