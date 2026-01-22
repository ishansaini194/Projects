package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ishansaini194/Projects/mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// Create base Mongo session
	session := getSession()
	defer session.Close()

	// Router
	r := httprouter.New()

	// Controller
	uc := controllers.NewUserController(session)

	// Routes
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	log.Println("Server running on http://localhost:9000")

	// Start server
	log.Fatal(http.ListenAndServe(":9000", r))
}

func getSession() *mgo.Session {
	s, err := mgo.DialWithTimeout("mongodb://localhost:27017", 10*time.Second)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	// Recommended session settings
	s.SetMode(mgo.Monotonic, true)

	return s
}
