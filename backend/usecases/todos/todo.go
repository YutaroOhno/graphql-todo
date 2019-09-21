package todos

import (
	"fmt"
	"graphql-todo/backend/entities"
	"graphql-todo/backend/infrastructure/db"
	"log"
)

type TodoUsecase struct {
	DB *db.DB
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

func (usecase *TodoUsecase) CreateTodo(input *entities.NewTodo) (*entities.Todo, error) {
	var todo entities.Todo
	var err error
	rows, err := usecase.DB.SqlxDB.DB.Query("insert into todos(id, title, body) values($1, $2, $3)", 3, input.Title, input.Body)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&todo)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(todo)
	return &todo, nil
}
