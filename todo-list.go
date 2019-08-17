package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

type DeleteTodo struct {
	ID int `json: "id"`
}

type TodoItem struct {
	Name        string `json: "name"`
	Description string `json: "description"`
}

type TransformedTodo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoListItem struct {
	id          int
	title       string
	description string
}

func init() {
	db, err := sql.Open("mysql", "root:password@/notes?charset=utf8")
	checkErr(err)

	database = db

	fmt.Println(database)
}

func getAllItem(c *gin.Context) {

	rows, err := database.Query("SELECT id, title, description FROM todo_list")
	checkErr(err)
	defer rows.Close()

	var finalList []TransformedTodo

	for rows.Next() {
		item := TodoListItem{}
		rows.Scan(&item.id, &item.title, &item.description)
		transformedTodo := TransformedTodo{
			ID:          item.id,
			Title:       item.title,
			Description: item.description}

		finalList = append(finalList, transformedTodo)
	}

	c.JSON(200, gin.H{
		"data":         finalList,
		"message":      "Success",
		"responseCode": 200,
	})

}

func addItem(c *gin.Context) {

	var item TodoItem
	c.BindJSON(&item)

	fmt.Println(item)

	statement, error := database.Prepare("INSERT INTO todo_list(title, description) VALUES (?, ?);")
	checkErr(error)

	result, error := statement.Exec(item.Name, item.Description)
	checkErr(error)

	lastInsertedId, error := result.LastInsertId()
	checkErr(error)

	c.JSON(200, gin.H{
		"insertedId":   lastInsertedId,
		"message":      "AddItem",
		"responseCode": 200,
	})
}

func deleteItem(c *gin.Context) {

	var item DeleteTodo
	c.BindJSON(&item)

	fmt.Println(item)

	statement, error := database.Prepare("DELETE FROM todo_list WHERE id=?;")
	checkErr(error)

	result, error := statement.Exec(item.ID)
	checkErr(error)

	RowsAffected, error := result.RowsAffected()
	checkErr(error)

	c.JSON(200, gin.H{
		"rowsAffected": RowsAffected,
		"message":      "DeleteItem",
		"responseCode": 200,
	})
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/getAllItems", getAllItem)

	r.POST("/addItem", addItem)

	r.POST("/deleteItem", deleteItem)

	r.Run()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
