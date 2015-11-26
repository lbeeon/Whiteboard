package router

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/gorilla/securecookie"
	"github.com/hunterpraska/Whiteboard/auth"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Password  []byte
	UserClass string
}

// Secure Cookie variables
var hashKey = securecookie.GenerateRandomKey(32)
var blockKey = securecookie.GenerateRandomKey(32)
var s = securecookie.New(hashKey, blockKey)

// Database connection
var db *bolt.DB

func OpenDB() error {
	var err error
	db, err = bolt.Open("whiteboard.db", 0600, nil)
	return err
}

func CloseDB() {
	db.Close()
}

// Handles home page requests
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/home.html")
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}

// Handles user login. If user is logged in, redirects to '/'.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if auth.LoggedIn(w, r, s) {
		http.Redirect(w, r, "/auth-check", 302)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles("views/login.html")
		if err != nil {
			log.Println(err)
			return
		}
		t.Execute(w, nil)
	} else {
		// Get values from html form
		user := r.FormValue("user")
		pass := r.FormValue("password")

		// Attempt to validate user, if incorrect info, send user back to login page
		if auth.ValidateLogin(user, pass, db) {
			cookie, err := auth.CreateCookie(s)
			if err != nil {
				log.Println(err)
				http.Redirect(w, r, "/login", 302)
				return
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/auth-check", 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}

// Handles user logout and redirects to '/login'
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.LoggedIn(w, r, s) {
		http.Redirect(w, r, "/login", 302)
		return
	}

	cookie := auth.DeleteCookie()
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", 302)
	return
}

// Allow users to register
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if auth.LoggedIn(w, r, s) {
		http.Redirect(w, r, "/auth-check", 302)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles("views/registration.html")
		if err != nil {
			log.Println(err)
			return
		}
		t.Execute(w, nil)
	} else {
		// Get values from html form
		user := r.FormValue("user")
		password := r.FormValue("password")
		userClass := r.FormValue("user-class")

		err := db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("users"))
			if err != nil {
				return err
			}

			// Encrypt password with bcrypt
			passwordCrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			// Create user
			u := User{passwordCrypt, userClass}

			// Marshal `u` to json object
			uByte, err := json.Marshal(u)
			if err != nil {
				return err
			}

			err = bucket.Put([]byte(user), uByte)
			if err != nil {
				return err
			}

			cookie, err := auth.CreateCookie(s)
			if err != nil {
				return err
			}
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "auth-check", 302)
			return nil
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
	http.Redirect(w, r, "/register", 302)
}

// Test of user authentication. Redirects user to login page if not logged in.
func AuthCheck(w http.ResponseWriter, r *http.Request) {
	if !auth.LoggedIn(w, r, s) {
		http.Redirect(w, r, "/login", 302)
		return
	}

	t, err := template.ParseFiles("views/authCheck.html")
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}
