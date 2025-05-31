package main

import "fmt"

type TodoList interface {
	addItem() int
	printItem()
}
type Todo struct {
	items []string
}

func (t *Todo) addItem(value string) {
	t.items = append(t.items, value)

}
func (t *Todo) printItem() {
	for index, value := range t.items {
		fmt.Printf("%d. %s\n", index+1, value)
	}
}
func newTodo() *Todo {
	return &Todo{
		items: []string{},
	}
}
