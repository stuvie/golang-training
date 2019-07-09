package main

import (
	// "io"
	// "os"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type Setting struct {
	Code  string	`json:"code"`
	// Value string	`json:"value"`
}

type Response []struct {
	Data Setting
}

type Response2 []Setting


func main() {
	// resp, err := http.Get("http://reddit.com/r/golang.json")
	resp, err := http.Get("http://shah-staging.fywss.com/catalog/api/v1/settings")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	// _, err = io.Copy(os.Stdout, resp.Body)
	r := new(Response2)
	err = json.NewDecoder(resp.Body).Decode(r)
	// err = json.Unmarshal(resp.Body, r)
	if err != nil {
		log.Fatal(err)
	}
  	fmt.Println(*r)
	// for _, child := range r { fmt.Println(child.Code) }
}
