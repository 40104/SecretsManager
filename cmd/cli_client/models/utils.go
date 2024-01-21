package models

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"log"
	"os"
	
	
)
func (env *Env) Connect(request_url string, json_string string) []byte{
	json_bytes := bytes.NewBuffer([]byte(json_string))
	bearer := fmt.Sprintf("Bearer %s", os.Getenv("JWT_KEY"))

	client := &http.Client{}

	req, err := http.NewRequest("POST", request_url, json_bytes)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		log.Fatal()
	}
	if res.StatusCode == 401{
		fmt.Printf("Unauthorized requset. \n")
		log.Fatal()
	}
	if res.StatusCode == 403{
		fmt.Printf("Forbidden access. \n")
		log.Fatal()
	}
	
	ResBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		log.Fatal()
	}
	
	return ResBody
}	