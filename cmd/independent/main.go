package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"bytes"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"github.com/dchest/uniuri"
	"github.com/urfave/cli/v2"
	"github.com/joho/godotenv"
	"strconv"
)

// Encryption key
var secretKey string

// Secret structure
type Secret struct {
	Name string
	Username  string
	Secret string
	URL string
}
// Encoding
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
// Decoding
func Decode(s string) string{
   data, err := base64.StdEncoding.DecodeString(s)
   if err != nil {
		panic(err)
   }
   return string(data)
}
// Encryption
func Encrypt(plaintext string) string {
	// AES cipher
    aes, err := aes.NewCipher([]byte(secretKey))
    if err != nil {
        panic(err)
    }
	// GCM mode
    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        panic(err)
    }
	// Randomizer
    nonce := make([]byte, gcm.NonceSize())
    _, err = rand.Read(nonce)
    if err != nil {
        panic(err)
    }
	// Encryption
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

    return Encode(ciphertext)
}
// Decryption
func Decrypt(ciphertext string) string {
	// AES cipher
    aes, err := aes.NewCipher([]byte(secretKey))
    if err != nil {
        panic(err)
    }
	// GCM mode
    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        panic(err)
    }
	// Decoder
	ciphertext = Decode(ciphertext)
    nonceSize := gcm.NonceSize()
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	// Decryption
    plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
    if err != nil {
        panic(err)
    }

    return string(plaintext)
}
// Open Data File
func OpenFile() *os.File {
	// Check file
	if _, err := os.Stat("SecretsManager.json"); os.IsNotExist(err) {
		log.Fatal("tasks file does not exist")
		return nil
	}
	// Open file
	file, err := os.OpenFile("SecretsManager.json", os.O_APPEND|os.O_CREATE|os.O_RDWR,0666)
	if err != nil {
		panic(err)
	}

	return file
}
// Add secret function
func AddSecret(secret Secret){
	// To json formatter
	json, err := json.Marshal(secret)
	if err != nil {
		panic(err)
	}
	// Write to file
	json = append(json, "\n"...)
	file:= OpenFile()
	if _, err := file.Write(json); err != nil {
		panic(err)
	}
	// Close file
	defer file.Close()
}
// Read secret function
func ReadSecret(name string) {
	// Open file
	file := OpenFile()
	defer file.Close()
	// Scan file
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		// To json formatter
		j := scanner.Text()
		data := Secret{}
		err := json.Unmarshal([]byte(j), &data)
		if err != nil {
			panic(err)
		}
		// Print data with given name
		if Decrypt(data.Name) == name {
			fmt.Printf("[%d] %s\n", i, fmt.Sprintf("Name: %s Username: %s Password: %s URL: %s ", 
				Decrypt(data.Name),Decrypt(data.Username), Decrypt(data.Secret), Decrypt(data.URL)))
			i++
		}
	}
}
// Read all secrets with encrypted keys function
func ListSecrets() {
	// Open file
	file := OpenFile()
	defer file.Close()
	// Scan file
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		// To json formatter
		j := scanner.Text()
		data := Secret{}
		err := json.Unmarshal([]byte(j), &data)
		if err != nil {
			panic(err)
		}
		// Print data with encrypted secret
		fmt.Printf("[%d] %s\n", i, fmt.Sprintf("Name: %s Username: %s Password: %s URL: %s ", 
			Decrypt(data.Name),Decrypt(data.Username), data.Secret, Decrypt(data.URL)))
		i++
	}
}
// Remove secret function
func RemoveSecret(name string) {
	// Open file
	file := OpenFile()
	// Scan file
	scanner := bufio.NewScanner(file)
	var bs []byte
    buf := bytes.NewBuffer(bs)

	for scanner.Scan() {
		// To json formatter
		j := scanner.Text()
		data := Secret{}
		err := json.Unmarshal([]byte(j), &data)
		if err != nil {
			panic(err)
		}
		// Safe data without deleted secret
		if Decrypt(data.Name) != name {
			buf.WriteString(j + "\n")
		}
	}
	// Write to file new data
	file.Truncate(0)
    file.Seek(0, 0)
    buf.WriteTo(file)
	// Close file
	defer file.Close()
}
// Import env function
func envVariable(key string) string{
    if err := godotenv.Load("app.env"); err != nil {
    	log.Fatal(err)
    }
	return os.Getenv(key)
}
// Main function
func main() {
	// Global env
	secretKey = envVariable("SECRET_KEY")
	var name string
	var username string
	var secret string
	var url string
	length, err := strconv.Atoi(envVariable("BASE_LENGTH"))
	if err != nil {
        panic(err)
    }
	// CLI parser
	app := &cli.App{
		Commands: []*cli.Command{
			{
				// Add new secret 
				Name:    "add",
				Usage:   "add a new secret",
				Aliases: []string{"a"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Aliases: []string{"n"},
						Usage:       "Set name for the secret",
						Destination: &name,
					},
					&cli.StringFlag{
						Name:        "username",
						Aliases: []string{"u"},
						Usage:       "Set username for the secret",
						Destination: &username,
					},
					&cli.StringFlag{
						Name:        "secret",
						Aliases: []string{"s"},
						Usage:       "Set secret for the secret",
						Destination: &secret,
					},
					&cli.StringFlag{
						Name:        "link",
						Aliases: []string{"l"},
						Usage:       "Set url for the secret",
						Destination: &url,
					},
				},
				Action: func(c *cli.Context) error {
					new_secret:= Secret{Name:Encrypt(name), Username:Encrypt(username),
						Secret:Encrypt(secret), URL:Encrypt(url)}
							
					AddSecret(new_secret)
					return nil
				},		
			},
			{
				// Generate new secret
				Name:    "generate",
				Usage:   "generate a new secret",
				Aliases: []string{"g"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Aliases: []string{"n"},
						Usage:       "Set name for the secret",
						Destination: &name,
					},
					&cli.StringFlag{
						Name:        "username",
						Aliases: []string{"u"},
						Usage:       "Set username for the secret",
						Destination: &username,
					},
					&cli.StringFlag{
						Name:        "link",
						Aliases: []string{"l"},
						Usage:       "Set url for the secret",
						Destination: &url,
					},
					&cli.IntFlag{
						Name:        "length",
						Usage:       "Set length for the secret",
						Value:        length,
						Destination: &length,
					},
					
				},
				Action: func(c *cli.Context) error {
					new_secret:= Secret{Name:Encrypt(name), Username:Encrypt(username), 
						Secret:Encrypt(uniuri.NewLen(length)), URL:Encrypt(url)}	
					AddSecret(new_secret)
					return nil
				},	
			},
			{
				// Read a secret
				Name:  "read",
				Usage: "read an existing secret",
				Aliases: []string{"r"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Aliases: []string{"n"},
						Usage:       "Set name for the secret",
						Destination: &name,
					},
				},
				Action: func(c *cli.Context) error {
					ReadSecret(name)
					return nil
				},
			},
			{
				// Remove a secret
				Name:  "delete",
				Usage: "delete an existing secret",
				Aliases: []string{"d"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Aliases: []string{"n"},
						Usage:       "Set name for the secret",
						Destination: &name,
					},
				},
				Action: func(c *cli.Context) error {
					RemoveSecret(name)
					return nil
				},
			},
			{
				// Print list of secrets
				Name:  "list",
				Usage: "list of the existing secrets",
				Aliases: []string{"l"},
				Action: func(c *cli.Context) error {
					ListSecrets()
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

