package models
// import packages
import (
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
	"encoding/base64"
)

// Encoding
func (m *DBModel) Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
// Decoding
func (m *DBModel) Decode(s string) string{
   data, err := base64.StdEncoding.DecodeString(s)
   if err != nil {
    log.Fatal(err) //check error
   }
   return string(data)
}
// Encryption
func (m *DBModel) Encrypt(plaintext string) string {
	// AES cipher
    aes, err := aes.NewCipher([]byte(m.Key))
    if err != nil {
        log.Fatal(err) //check error
    }
	// GCM mode
    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        log.Fatal(err) //check error
    }
	// Randomizer
    nonce := make([]byte, gcm.NonceSize())
    _, err = rand.Read(nonce)
    if err != nil {
        log.Fatal(err) //check error
    }
	// Encryption
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

    return m.Encode(ciphertext)
}
// Decryption
func (m *DBModel) Decrypt(ciphertext string) string {
	// AES cipher
    aes, err := aes.NewCipher([]byte(m.Key))
    if err != nil {
        log.Fatal(err) //check error
    }
	// GCM mode
    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        log.Fatal(err) //check error
    }
	// Decoder
	ciphertext = m.Decode(ciphertext)
    nonceSize := gcm.NonceSize()
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	// Decryption
    plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
    if err != nil {
        log.Fatal(err) //check error
    }

    return string(plaintext)
}