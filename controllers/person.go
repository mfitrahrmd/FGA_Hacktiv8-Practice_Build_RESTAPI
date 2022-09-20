package controllers

import (
	"FGA_Hacktiv8-Practice_Build_RESTAPI/structs"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func (idb InDB) CreatePerson(ctx *gin.Context) {
	var person structs.Person

	err := ctx.BindJSON(&person)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = idb.DB.Create(&person).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"result": err.Error(),
			})
	} else {
		ctx.JSON(http.StatusCreated,
			gin.H{
				"result": person,
			})
	}
}

func (idb InDB) GetPerson(ctx *gin.Context) {
	id := ctx.Param("id")

	var person structs.Person

	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"result": err.Error(),
				"count":  0,
			})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": person,
			"count":  1,
		})
	}

}

func (idb InDB) UpdatePerson(ctx *gin.Context) {
	id := ctx.Query("id")

	var person structs.Person
	var updatedPerson structs.Person

	err := idb.DB.First(&person, id).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": "data not found",
		})
	}

	ctx.BindJSON(&updatedPerson)

	fmt.Println(updatedPerson)

	err = idb.DB.Model(&person).Updates(updatedPerson).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": "update failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "succesfully update data",
		})
	}
}

func (idb InDB) DeletePerson(ctx *gin.Context) {
	id := ctx.Param("id")

	var person structs.Person

	err := idb.DB.First(&person, id).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": "data not found",
		})
	}

	err = idb.DB.Delete(&person).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": "updated failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "successfully delete data",
		})
	}
}
