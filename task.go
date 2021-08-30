package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

type taskList struct {
	tasks []*task
}

func (t *taskList) addList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) deleteList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *task) markComplate() {
	t.completed = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *taskList) printTask() {
	for _, task := range t.tasks {
		fmt.Println("Name", task.name)
		fmt.Println("Description", task.description)
	}
}

func (t *taskList) printTaskComplate() {
	for _, task := range t.tasks {
		if task.completed {
			fmt.Println("Name", task.name)
			fmt.Println("Description", task.description)
		}
	}
}

func main() {
	t1 := &task{
		name:        "complete my course of go",
		description: "Complete my course of go in this week",
	}
	t2 := &task{
		name:        "complete my course of React",
		description: "Complete my course of React in this week",
	}

	t3 := &task{
		name:        "complete my course of NodeJs",
		description: "Complete my course of NodeJs in this week",
	}

	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	list.addList(t3)
	list.printTask()
	list.tasks[0].markComplate()
	fmt.Println("Task completed !!!")
	list.printTaskComplate()
}
