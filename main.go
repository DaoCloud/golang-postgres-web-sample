package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	host := os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	port := os.Getenv("POSTGRES_PORT_5432_TCP_PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	dbname := os.Getenv("INSTANCE_NAME")

	connection_info := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	fmt.Println(connection_info)
	db, err := sql.Open("postgres", connection_info)
	if err != nil {
		fmt.Println("Open Error")
		return
	}

	_, _ = db.Exec("CREATE TABLE Persons(Name varchar(255))")

	_, _ = db.Exec(fmt.Sprintf("INSERT INTO Persons VALUES (%s)", "DaoCloud"))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		rows, _ := db.Query(`SELECT * FROM Persons`)
		c.JSON(200, rows)
	})

	r.Run(":8080")
}
