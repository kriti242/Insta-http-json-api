package tests

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/kriti242/Insta-http-json-api/controllers"
	"gopkg.in/mgo.v2"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.POST("/users", uc.CreateUsers)
	return router
}

func TestCreateUsers(t *testing.T) {
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	//response recorder to record the response
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	//checking the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("Unexpected body returned: got %v want %v",
			rr.Body.String(), expected)
	}
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return session
}
