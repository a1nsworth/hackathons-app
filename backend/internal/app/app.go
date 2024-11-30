package app

import (
	_ "context"
	_ "database/sql"
	_ "fmt"
	_ "log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	// dbConn, err := sql.Open("sqlite3", "main.sqlite")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer dbConn.Close()
	//
	// queries := postgresql.New(dbConn)
	//
	// user, err := queries.GetUserById(context.Background(), 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("User: %+v\n", user)
	r := gin.Default()
	r.GET(
		"/ping", func(c *gin.Context) {
			c.JSON(
				http.StatusOK, gin.H{
					"message": "pong",
				},
			)
		},
	)
	err := r.Run(":4242")
	if err != nil {
		panic(err)
	}
}
