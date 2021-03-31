package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPerson(content *gin.Context) {
	content.JSON(http.StatusAccepted, gin.H{
		"id": "test",
	})
}

// variable一定要大寫, 後面要描述
type Person struct {
	Name string 	`json:"name"`
	Id string 		`json:"id"`
	Content string 	`json:"content"`
}
func AddPerson(content *gin.Context)  {
	var person Person
	if err:=content.ShouldBindJSON(&person); err == nil{
		content.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content.JSON(http.StatusAccepted, gin.H{
		"id": person.Id,
		"name": person.Name,
		"content": person.Content,
	})
}