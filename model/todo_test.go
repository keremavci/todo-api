package model

import "testing"

func TestTodoJson(t *testing.T){
	todo := &Todo{Id: NewId(),CreateAt: GetMillis(),UpdateAt: GetMillis(),Text: "Test Content",Active: true}
	todo2 := &Todo{Id: NewId(),CreateAt: GetMillis(),UpdateAt: GetMillis(),Text: "Test Content 2",Active: false}

	var todos []*Todo;
	todos=append(todos,todo)
	todos=append(todos,todo2)

	if ! todo.IsValid()  {
		t.Fatalf("Didn't give valid todo ")
	}

	todoStr := todo.ToJson()

	if len(todoStr) == 0 {
		t.Fatalf("Didn't give valid todo ")
	}

	if len(TodoListToJson(todos)) == 0 {
		t.Log(TodoListToJson(todos))
		t.Fatalf("Didn't give valid todo json")
	}

}
