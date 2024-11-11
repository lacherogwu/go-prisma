package main

import (
	"database/sql"
	"fmt"
	"go-prisma/.gen/go_test/public/model"
	. "go-prisma/.gen/go_test/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_test sslmode=disable")
	if err != nil {
		fmt.Println("Here")
		panic(err)
	}

	defer db.Close()
	stmt := SELECT(
		Post.ID,
		Post.Title,
		// Author.Name,
	).FROM(Post)

	var posts []model.Post

	// Author struct {
	// 	model.Author
	// }
	err = stmt.Query(db, &posts)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range posts {
		fmt.Printf("Title: %s, ID: %d\n", p.Title, p.ID)
	}

}
