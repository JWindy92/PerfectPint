package main

import (
	"log"
	"net/http"
	"time"

	"github.com/JWindy92/PerfectPint/jwt-poc/pkg/auth"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

// type User struct {
// 	gorm.Model
// 	Name     string `json:"name"`
// 	Email    string `gorm:"unique" json:"email"`
// 	Password string `json:"password"`
// 	Role     string `json:"role"`
// }

// type Authentication struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type Token struct {
// 	Role        string `json:"role"`
// 	Email       string `json:"email"`
// 	TokenString string `json:"token"`
// }

var router *mux.Router

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRoute() {
	router.HandleFunc("/signup", auth.SignUp).Methods("POST")
	router.HandleFunc("/signin", auth.SignIn).Methods("POST")
	router.HandleFunc("/admin", auth.IsAuthorized(AdminIndex)).Methods("GET")
	router.HandleFunc("/user", auth.IsAuthorized(UserIndex)).Methods("GET")
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}

func main() {
	CreateRouter()
	InitializeRoute()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
}
