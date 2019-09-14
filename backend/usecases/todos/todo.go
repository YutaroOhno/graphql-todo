package todos

import (
	"todo-graphql-api/entities"
	"todo-graphql-api/infrastructure/db"
	"fmt"
)

type TodoUsecase struct {
	DB             *db.DB
}

func NewTodoUsecase(db *db.DB) *TodoUsecase {
	return &TodoUsecase{
		DB: db,
	}
}

func (usecase *TodoUsecase) GetTodos() ([]*entities.Todo, error) {
	// todos := []entities.Todo{}
	var todos []*entities.Todo

	// rows, err :=  usecase.DB.SqlxDB.Queryx("SELECT id, title, body FROM todos")
    // if err != nil {
	// 	panic(err)
	// }
	
    // var todo entities.Todo
    // for rows.Next() {

    //     //rows.Scanの代わりにrows.StructScanを使う
    //     err := rows.StructScan(&todo)
    //     if err != nil {
	// 		panic(err)
    //     }
    //     todos = append(todos, &todo)
    // }


	err := usecase.DB.SqlxDB.Select(&todos, "select id, title, body from todos")
	if err != nil {
		panic(err)
	}

	fmt.Println(todos)
	return todos, nil
}

