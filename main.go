package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kriti242/Insta-http-json-api/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("Instagram API")
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.POST("/users", uc.CreateUsers)
	router.GET("/users/:id", uc.GetUser)
	router.POST("/posts", uc.CreatePosts)
	router.GET("/posts/:id", uc.GetPost)
	router.GET("/posts/users/id", uc.GetUserPosts)
	http.ListenAndServe("localhost:9090", router)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return session
}
