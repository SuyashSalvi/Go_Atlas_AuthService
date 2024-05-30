package handlers

import (
	"encoding/json"
	"net/http"
	"authservice2/db"
	"authservice2/models"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Password       string `json:"password" bson:"password"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	SessionExpires string `json:"sessionExpires" bson:"sessionExpires"`
	SessionToken   string `json:"sessionToken" bson:"sessionToken"`
	IsAdmin        bool   `json:"isAdmin" bson:"isAdmin"`
}

type Claims struct {
	Name string `json:"username"`
	jwt.StandardClaims
}

func Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling signup request...")

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request payload:", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Println("Error generating hashed password:", err)
		return
	}

	collection := db.Client.Database("authdb").Collection("users")
	_, err = collection.InsertOne(r.Context(), bson.M{"name": creds.Name, "password": string(hashedPassword),"email":creds.Email})
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Println("Error inserting user into database:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("Signup request handled successfully.")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	collection := db.Client.Database("authdb").Collection("users")
	var user models.User
	err = collection.FindOne(r.Context(), bson.M{"name": creds.Name}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Name: creds.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func AccessToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No token", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error reading cookie", http.StatusBadRequest)
		return
	}

	tokenStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid || err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
