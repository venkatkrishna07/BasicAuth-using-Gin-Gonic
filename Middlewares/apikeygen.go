package Middlewares

import (
	"BasicAuth/Models"

	"math/rand"
	"net/http"
	"strings"
	"time"

	"crypto/hmac"
	"crypto/sha256"

	"encoding/base64"

	"github.com/gin-gonic/gin"
)

//Generate API key and secret
func Gen(c *gin.Context) {
	email := c.Param("email")
	key := "api_" + GenerateApiKeys(email)
	sec := GenerateSecret(key, email)
	result := map[string]string{"key": key, "secret": sec, "Note": "API secret is not stored on the server.If API secret is forgotten then the API key must be regenatred "}
	c.JSON(http.StatusOK, result)
}

func GenerateRune(word string) []rune {
	res := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" + word +
		"0123456789")

	return res
}

//Generate API key
func GenerateApiKeys(email string) string {
	rand.Seed(time.Now().UnixNano())
	chars := GenerateRune(email)
	length := 16
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

//Generate Secret
func GenerateSecret(key string, email string) string {
	secret := "ServerSecret"
	rand.Seed(time.Now().UnixNano())
	chars := GenerateRune(key)
	length := 16
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	ComputeHmac256(email, key, str, secret)

	return str
}

//Compute the HMAC of secret and store email, api key and HMAC value in db
func ComputeHmac256(email string, api_key string, message string, secret string) {
	var UserEmail Models.Email
	key_sec := []byte(secret)
	h := hmac.New(sha256.New, key_sec)
	h.Write([]byte(message))
	hash_value := base64.StdEncoding.EncodeToString(h.Sum(nil))
	result1, err1 := Models.DB.Query("Select email from api_keys where email = ? ", email)
	if err1 != nil {
		panic(err1)
	}

	for result1.Next() {
		err := result1.Scan(&UserEmail.Email)
		if err != nil {
			panic(err)
		}

	}
	//Check if the API keys for the given email is present on DB
	if UserEmail.Email == email {
		//If present replace the old key and hashvalue with new key and hash value
		_, err := Models.DB.Query("UPDATE api_keys SET api_key = ?, hash_value = ? WHERE email = ?", api_key, hash_value, email)
		if err != nil {
			panic(err)
		}

	} else {
		//Else generate new API keys
		_, err := Models.DB.Query("INSERT into api_keys (email,api_key,hash_value) values (?,?,?)", email, api_key, hash_value)
		if err != nil {
			panic(err)
		}

	}
}
