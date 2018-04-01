package entities

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
	_ "github.com/lib/pq"
	"crypto/aes"
	"io"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Database instance
var db *sql.DB
var key []byte

// Initialize
func init() {
	file, err := ioutil.ReadFile("./db.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	type dbCredentials struct {
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
		Port     string `json:"port"`
		Key      string `json:"key"`
	}
	var credentials dbCredentials
	json.Unmarshal(file, &credentials)

	connection := "user=" + credentials.User +
		" password='" + credentials.Password +
		"' dbname= '" + credentials.DBName +
		"' port= '" + credentials.Port +
		"' sslmode=disable"

	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	key = []byte(credentials.Key)
	return
}

func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func Render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}