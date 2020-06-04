package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"login/models"

	"github.com/gin-gonic/gin"
)

// Register will jobs
func Register(c *gin.Context) {

	body := c.Request.Body
	bytes, _ := ioutil.ReadAll(body)
	var user models.User
	err := json.Unmarshal(bytes, &user)
	if err != nil {
		log.Println("error parsing page " + err.Error())
	}

	log.Println("user : ", user)
	// userID, err := models.CreateUser(job)
	userID := 1

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "userID": userID})
	}
}

// LogIN will jobs
func LogIN(c *gin.Context) {

	body := c.Request.Body
	bytes, _ := ioutil.ReadAll(body)
	var user models.User
	err := json.Unmarshal(bytes, &user)
	if err != nil {
		log.Println("error parsing page " + err.Error())
	}

	log.Println("user : ", user)
	// userID, err := models.GetJob(job)
	userID := 1

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "userID": userID})
	}
}
