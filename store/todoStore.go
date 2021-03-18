package store

import (
	"database/sql"
	"github.com/keremavci/todo-api/model"
	"github.com/pkg/errors"
)

type TodoStore struct {
	sqlStore *SqlStore
}

func newTodoStore(sqlStore *SqlStore) *TodoStore {

	ts := &TodoStore{sqlStore}
	dbMap := sqlStore.GetConnection()
	table := dbMap.AddTableWithName(model.Todo{}, "Todo").SetKeys(false, "Id")
	table.ColMap("Id").SetMaxSize(36)
	table.ColMap("CreateAt").SetMaxSize(64)
	table.ColMap("UpdateAt").SetMaxSize(64)
	table.ColMap("Text")
	table.ColMap("Active")

	return ts

}


func (ts *TodoStore) Save(todo *model.Todo)(*model.Todo, error){

	todo.Id = model.NewId()
	todo.CreateAt = model.GetMillis()
	todo.UpdateAt = todo.CreateAt


	tx, err := ts.sqlStore.GetConnection().Begin()
	if err != nil{
		return nil, errors.Wrap(err, "Todo Save Begin Transaction")
	}
	err = tx.Insert(todo)
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "Unable to save todo")
	} else if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "Unable to commit todo transction")
	}

	return todo, nil

}

func (ts *TodoStore) Update(todo *model.Todo)(*model.Todo, error){

	oldResult, err := ts.sqlStore.GetConnection().Get(model.Todo{}, todo.Id)
	if err != nil || oldResult == nil {
		return nil, errors.Wrapf(err, "Failed to get Todo with id=%s", todo.Id)
	}

	oldOrder := oldResult.(*model.Todo)
	todo.CreateAt = oldOrder.CreateAt
	todo.UpdateAt = model.GetMillis()

	tx, err := ts.sqlStore.GetConnection().Begin()
	if err != nil{
		return nil, errors.Wrap(err, "Todo Update Begin Transaction")

	}
	count, err := tx.Update(todo)
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrapf(err, "Unable to update todo with id=%s", todo.Id)
	} else if count > 1 {
		_ = tx.Rollback()
		return nil, errors.Wrapf(err, "Failed trying to update multiple todo updated with id=%s", todo.Id)
	} else if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "Unable to commit todo transction")
	}

	return todo, nil
}

func (ts *TodoStore) Delete(todoId string)(error){

	todo, err := ts.GetById(todoId);
	if err != nil || todo == nil {
		return errors.Wrapf(err, "Failed to get Todo with id=%s", todo.Id)
	}

	tx, err := ts.sqlStore.GetConnection().Begin()
	if err != nil{
		return errors.Wrap(err, "Todo Update Begin Transaction")

	}
	count, err := tx.Delete(todo)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrapf(err, "Unable to delete todo with id=%s", todo.Id)
	} else if count > 1 {
		_ = tx.Rollback()
		return errors.Wrapf(err, "Failed trying to delete multiple todo updated with id=%s", todo.Id)
	} else if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "Unable to commit todo transction")
	}

	return nil
}

func (ts *TodoStore) GetById(todoId string)(*model.Todo, error){
	var dbTodo *model.Todo

	err := ts.sqlStore.GetConnection().SelectOne(&dbTodo, "Select * from Todo Where Id = :Id",
		map[string]interface{}{"Id": todoId})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrapf(err,"Failed to get Todo. TodoId: %s", todoId)

	}
	return dbTodo, nil
}

func (ts *TodoStore) GetAll()([]*model.Todo, error){
	var dbTodoList []*model.Todo

	if _, err := ts.sqlStore.GetConnection().Select(&dbTodoList, "SELECT * FROM Todo ORDER BY CreateAt DESC", map[string]interface{}{}); err != nil {
		return nil, errors.Wrap(err, "failed to find Todos")
	}

 	return dbTodoList,nil

}