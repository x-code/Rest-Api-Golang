package main

//developed by R.Firandika | x-c0de

import (
	"github.com/gin-gonic/gin"
	"./controllers"
	"gopkg.in/mgo.v2"
)

func main() {

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Get a user resource
	router := gin.Default()
	router.Static("/assets", "./assets")
    router.GET("/users", uc.UsersList)
    router.GET("/users/:id", uc.GetUser)
    router.DELETE("/users/:id", uc.RemoveUser)
    router.POST("/users", uc.CreateUser)
    router.PUT("/users/:id", uc.UpdateUser)
    router.PUT("/users/:id/uploadImage", uc.UploadImage)
    
    router.Run(":8000")
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
