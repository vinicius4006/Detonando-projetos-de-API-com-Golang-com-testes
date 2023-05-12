package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

/*
cert:

	openssl genrsa -out ./key 4096
	openssl rsa -in ./key -pubout -out ./key.pub
*/
var (
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
)

func WriteResponse(status int, body interface{}, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(body)
	w.Write(payload)
}

type User struct {
	Username string
	Password string
}

func CreateToken(username string) string {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	claims["authorized"] = true
	claims["user"] = username
	tokenString, err := token.SignedString(PrivateKey)
	if err != nil {
		fmt.Println(err)
		panic(err.Error)
	}
	return tokenString
}

func Login(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var user User
	if err := decoder.Decode(&user); err != nil {
		WriteResponse(http.StatusBadRequest, map[string]string{"error": err.Error()}, w)
		return
	}
	// You should check username and password against database or other resource:
	tokenString := ""
	if user.Username == "user1" && user.Password == "password1" {
		tokenString = CreateToken(user.Username)
		cookie := http.Cookie{
			Name:     "Token",
			Value:    tokenString,
			MaxAge:   1800,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		WriteResponse(http.StatusOK, map[string]string{"AccessToken": tokenString}, w)
	} else {
		WriteResponse(http.StatusUnauthorized, map[string]string{"Error": "Invalid credentials"}, w)
	}

}

func ValidateToken(r *http.Request) bool {

	tokenCookie, err := r.Cookie("Token")

	if err != nil {
		fmt.Println("Error ocurred while reading cookie")
		return false
	}

	tkn, err1 := jwt.Parse(tokenCookie.Value, func(jwtToken *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})

	fmt.Println(tkn.Claims)

	if err1 != nil {
		fmt.Println("Error: ", err1)
		return false
	}

	if tkn == nil {
		fmt.Println("Invalid token - Not parsed")
		return false
	}

	if !tkn.Valid {
		fmt.Println("Invalid token - Marked as invalid")
		return false
	}

	return true
}

func UnprotectedResource(w http.ResponseWriter, r *http.Request) {
	WriteResponse(http.StatusOK, map[string]string{"status": "ok - unprotected resource"}, w)
}

func ProtectedResource(w http.ResponseWriter, r *http.Request) {
	if !(ValidateToken(r)) {
		WriteResponse(http.StatusForbidden, map[string]string{"status": "Not authorized"}, w)
		return
	}
	fmt.Println(ValidateToken(r))
	WriteResponse(http.StatusOK, map[string]string{"status": "ok - PROTECTED resource"}, w)
}

func GetKeys() {
	publicKeyPath := "../jwtkey.pub"
	os.Setenv("API_PUBLIC_KEY", publicKeyPath)
	if pubKey, hasValue := os.LookupEnv("API_PUBLIC_KEY"); hasValue {
		publicKeyPath = pubKey
	} else {
		panic(("You should create env vars API_PUBLIC_KEY and API_PRIVATE_KEY"))
	}

	privateKeyPath := "../jwtkey"
	os.Setenv("API_PRIVATE_KEY", privateKeyPath)
	if privKey, hasValue := os.LookupEnv("API_PRIVATE_KEY"); hasValue {
		privateKeyPath = privKey
	} else {
		panic(("You should create env vars API_PRIVATE_KEY and API_PUBLIC_KEY"))
	}

	prvKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic("Error reading private key file")
	}
	block, _ := pem.Decode(prvKey)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	PrivateKey = key

	pubKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		panic("Error reading public key file")
	}
	block2, _ := pem.Decode(pubKey)
	pkey, err := x509.ParsePKIXPublicKey(block2.Bytes)
	if err != nil {
		panic(err)
	}

	rsaKey, ok := pkey.(*rsa.PublicKey)
	if !ok {
		panic("got unexpected public key type")
	}
	PublicKey = rsaKey
}

func main() {

	GetKeys()

	router := mux.NewRouter()
	router.HandleFunc("/api/unprotected", UnprotectedResource).Methods("GET")
	router.HandleFunc("/api/protected", ProtectedResource).Methods("GET")
	router.HandleFunc("/api/login", Login).Methods("POST")
	err := http.ListenAndServe(fmt.Sprintf(":8888"), router)
	fmt.Printf("Error server to listen: \n%v", err)
}
