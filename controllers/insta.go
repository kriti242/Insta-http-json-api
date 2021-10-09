package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kriti242/insta-http-json-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//Create an User
func (uc UserController) CreateUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User

	//decodes the json values entered
	json.NewDecoder(r.Body).Decode(&user)

	//id creation for the user
	user.ID = bson.NewObjectId()

	uc.session.DB("http-json-api").C("users").Insert(user)

	user_json, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("User Account Created")
	fmt.Fprintf(w, "%s\n", user_json)

}

//Get a user using id
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//parameter being passed
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	user := models.User{}

	if err := uc.session.DB("http-json-api").C("users").FindId(oid).One(&user); err != nil {
		w.WriteHeader(404)
		fmt.Println("User not found")
		return
	}

	user_json, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("User : ")
	fmt.Fprintf(w, "%s\n", user_json)
}

//Create a Post
func (uc UserController) CreatePosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	post := models.Post{}

	//decodes the json values entered
	json.NewDecoder(r.Body).Decode(&post)

	//id creation for the user
	post.ID = bson.NewObjectId()

	uc.session.DB("http-json-api").C("posts").Insert(post)

	post_json, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Post Created !")
	fmt.Fprintf(w, "%s\n", post_json)
}

//Get a post using id
func (uc UserController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	var post models.Post
	if err := uc.session.DB("http-json-api").C("posts").FindId(oid).One(&post); err != nil {
		w.WriteHeader(404)
		fmt.Println("Post not found")
		return
	}

	post_json, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("Post : ")
	fmt.Fprintf(w, "%s\n", post_json)
}

//List all posts of a user
func (uc UserController) GetUserPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("user_id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	//declare the instances
	var user models.User
	var post models.Post

	//check for error
	if err := uc.session.DB("http-json-api").C("users").FindId(oid).One(&user); err != nil {
		w.WriteHeader(404)
		fmt.Println("User not found")
		return
	}
	if err := uc.session.DB("http-json-api").C("posts").FindId(oid).One(&post); err != nil {
		w.WriteHeader(404)
		return
	}

	user_json, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	post_json, err1 := json.Marshal(post)
	if err1 != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("User Posts:")
	fmt.Fprintf(w, "%s\n", user_json)
	fmt.Fprintf(w, "%s\n", post_json)

}
