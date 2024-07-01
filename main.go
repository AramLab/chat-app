package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Working directory", wd)
	db, err := sql.Open("sqlite", wd+"/database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer func(d *sql.DB) {
		closeErr := d.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(db)

	// Чтение и выполнение содержимого schema.sql
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	if err != nil {
		log.Fatal(err)
	}

	r.POST("/users", func(c *gin.Context) { createUser(c, db) })
	r.POST("/channels", func(c *gin.Context) { createChannel(c, db) })
	r.POST("/messages", func(c *gin.Context) { createMessage(c, db) })

	r.GET("/channels", func(c *gin.Context) { listChannels(c, db) })
	r.GET("/messages", func(c *gin.Context) { listMessages(c, db) })

	r.POST("/login", func(c *gin.Context) { login(c, db) })

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
