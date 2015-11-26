package auth

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	Password  []byte
	UserClass string
}

func ValidateLogin(user, password string, db *bolt.DB) bool {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("Bucket pastes not found!")
		}

		var uByte []byte
		if uByte = bucket.Get([]byte(user)); uByte == nil {
			return fmt.Errorf("User not found!")
		}

		var u User
		err := json.Unmarshal(uByte, &u)
		if err != nil {
			return err
		}

		err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		return err
	})
	if err != nil {
		return false
	}
	return true
}

func LoggedIn(w http.ResponseWriter, r *http.Request, s *securecookie.SecureCookie) bool {
	if cookie, err := r.Cookie("whiteboard"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("whiteboard", cookie.Value, &value); err == nil {
			return true
		}
		return false
	}
	return false
}

func CreateCookie(s *securecookie.SecureCookie) (*http.Cookie, error) {
	var err error

	// Create secure cookie with login info
	value := map[string]string{
		"authenticated": "true",
	}
	if encoded, err := s.Encode("whiteboard", value); err == nil {
		cookie := &http.Cookie{
			Name:  "whiteboard",
			Value: encoded,
			Path:  "/",
		}
		cookie.MaxAge = 10000
		return cookie, err
	}

	return nil, err
}

func DeleteCookie() *http.Cookie {
	// Create cookie with info removed
	cookie := &http.Cookie{
		Name: "whiteboard",
	}
	cookie.MaxAge = -1
	return cookie
}
