package api

import (
	"bytes"
	"github.com/keremavci/todo-api/model"
	"net/http"
	"testing"
)



func TestAddGetTodo(t *testing.T) {

	todo := model.Todo{Text:"Test ToDo",Active: true}
	server:=SetupTest()


	resp, err := http.Post(server.URL + "/api/todo","application/json", bytes.NewBuffer([]byte(todo.ToJson())))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Cannot receive %d response. Response %d",http.StatusCreated, resp.StatusCode)
	}

	//respBody,_:=ioutil.ReadAll(resp.Body)
	dbTodo:=model.TodoFromJson((resp.Body))
	dbTodo.Active=false

	reqPut, err := http.NewRequest("PUT", server.URL+"/api/todo", bytes.NewBuffer([]byte(dbTodo.ToJson())))
	if err != nil {
		t.Fatal(err)
	}

	putResp, err := server.Client().Do(reqPut)
	if err != nil {
		t.Fatal(err)
	}

	if putResp.StatusCode != http.StatusOK {
		t.Fatalf("Cannot receive %d response. Response %d",http.StatusOK, putResp.StatusCode)
	}

	putTodo:=model.TodoFromJson((putResp.Body))
	if putTodo.Id != dbTodo.Id {
		t.Fatalf("Cannot receive expected response. Response.")
	}


	reqDelete, err := http.NewRequest("DELETE", server.URL+"/api/todo/" + dbTodo.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	deleteResp, err := server.Client().Do(reqDelete)
	if err != nil {
		t.Fatal(err)
	}

	if deleteResp.StatusCode != http.StatusOK {
		t.Fatalf("Cannot receive %d response. Response %d",http.StatusOK, deleteResp.StatusCode)
	}

}

func TestGetAllTodo(t *testing.T) {

	server:=SetupTest()
	resp, err := http.Get(server.URL + "/api/todo")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non 200 response %d", resp.StatusCode)
	}
}
