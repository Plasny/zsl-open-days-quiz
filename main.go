package main

import (
	// "context"
	// "database/sql"
	// "log"
	"zsl-demo/controller"
	"zsl-demo/model"
)

func main() {
	model.Init()
	controller.Init()

	// q := model.New(model.DB)
	// id, err := q.AddUserReturnId(context.Background(), sql.NullString{Valid: false})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(id)
	//
	// q.AddTime(context.Background(), model.AddTimeParams{Time: 10, UserID: id})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// q.AddTime(context.Background(), model.AddTimeParams{Time: 20, UserID: id})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// id, err = q.AddUserReturnId(context.Background(), sql.NullString{Valid: true, String: "test"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// q.AddTime(context.Background(), model.AddTimeParams{Time: 20, UserID: id})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// id, err = q.AddUserReturnId(context.Background(), sql.NullString{Valid: true, String: "abc"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// q.AddTime(context.Background(), model.AddTimeParams{Time: 15, UserID: id})
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
