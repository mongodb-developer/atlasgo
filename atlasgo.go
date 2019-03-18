package main

import (
	"encoding/json"
	"fmt"
	dac "github.com/xinsnake/go-http-digest-auth-client"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const baseURL = "https://cloud.mongodb.com/api/atlas/v1.0/"

func main() {
	username, usernameexists := os.LookupEnv("ATLAS_USER")
	userkey, userkeyexists := os.LookupEnv("ATLAS_USER_KEY")

	if !usernameexists || !userkeyexists {
		log.Fatalln("Env vars missing")
	}

	t := dac.NewTransport(username, userkey)
	req, err := http.NewRequest("GET", baseURL, nil)

	resp, err := t.RoundTrip(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatalln(readErr)
	}

	var response interface{}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	s, _ := json.MarshalIndent(response, "", " ")

	fmt.Println(string(s))

}
