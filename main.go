package main

import (
	"FGA_Hacktiv8-Practice_Build_RESTAPI/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	c := controllers.NewPostgresDB()

	r.GET("/person/:id", c.GetPerson)

	r.POST("/person", c.CreatePerson)

	r.PUT("/person", c.UpdatePerson)

	r.DELETE("/person/:id", c.DeletePerson)

	err := r.Run(":80")
	if err != nil {
		panic(err.Error())
	}
}
