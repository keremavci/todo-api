package model

import (
	"encoding/json"
	"io"
	"strings"
)

type Todo struct {
	Id	string `json:"id"`
	CreateAt 	int64 `json:"create_at,omitempty"`
	UpdateAt           int64     `json:"update_at,omitempty"`
	Text	string `json:"text"`
	Active	bool `json:"active"`
}


func (todo *Todo) IsValid() bool {

	if len(strings.TrimSpace(todo.Text)) == 0{
		return false
	}
	return true
}

func (todo *Todo) ToJson() string {
	t,_ := json.Marshal(todo)
	return string(t);
}


func TodoFromJson(data io.Reader) *Todo{
	var todo *Todo
	json.NewDecoder(data).Decode(&todo)
	return todo
}

func TodoListToJson(todos []*Todo) string {
	t,_ := json.Marshal(todos)
	return string(t);
}

