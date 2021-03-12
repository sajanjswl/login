package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	// get configuration
	address := flag.String("server", "http://localhost:8085", "HTTP gateway url, e.g. http://localhost:8085")
	flag.Parse()

	// t := time.Now().In(time.UTC)
	// pfx := t.Format(time.RFC3339Nano)

	var body string

	// Call Create
	// resp, err := http.Post(*address+"/v1/register", "application/json", strings.NewReader(fmt.Sprintf(`
	// 	{
	// 		"apiVersion":"v1",
	// 		"user": {
	// 			"emailID":"sjnjaiswl@rest",
	// 			"password":"rest@pass",
	// 			"firstName":"rest",
	// 			"lastName":"something rest",
	// 			"mobileNumber":"rest8989797"
	// 		}
	// 	}
	// `)))
	// if err != nil {
	// 	log.Fatalf("failed to call Register method: %v", err)
	// }
	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read Create response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("Create response: Code=%d, Body=%s\n\n", resp, body)

	resp, err := http.Post(*address+"/v1/login", "application/json", strings.NewReader(fmt.Sprintf(`
		{
			"apiVersion":"v1",
			"emailID": "sjnjaiswal@gmail.com",
			"password":"password"
		}
	`)))
	if err != nil {
		log.Fatalf("failed to call Register method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Login response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Loging response: Code=%d, Body=%s\n\n", resp, body)

	// parse ID of created ToDo
	// var created struct {
	// 	API string `json:"api"`
	// 	ID  string `json:"id"`
	// }
	// err = json.Unmarshal(bodyBytes, &created)
	// if err != nil {
	// 	log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
	// 	fmt.Println("error:", err)
	// }

	// Call Read
	// resp, err = http.Get(fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID))
	// if err != nil {
	// 	log.Fatalf("failed to call Read method: %v", err)
	// }
	// bodyBytes, err = ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read Read response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("Read response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// Call Update
	// req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID),
	// 	strings.NewReader(fmt.Sprintf(`
	// 	{
	// 		"api":"v1",
	// 		"toDo": {
	// 			"title":"title (%s) + updated",
	// 			"description":"description (%s) + updated",
	// 			"reminder":"%s"
	// 		}
	// 	}
	// `, pfx, pfx, pfx)))
	// req.Header.Set("Content-Type", "application/json")
	// resp, err = http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatalf("failed to call Update method: %v", err)
	// }
	// bodyBytes, err = ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read Update response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("Update response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// // Call ReadAll
	// resp, err = http.Get(*address + "/v1/todo/all")
	// if err != nil {
	// 	log.Fatalf("failed to call ReadAll method: %v", err)
	// }
	// bodyBytes, err = ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read ReadAll response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("ReadAll response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// Call Delete
	// req, err = http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID), nil)
	// resp, err = http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatalf("failed to call Delete method: %v", err)
	// }
	// bodyBytes, err = ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read Delete response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("Delete response: Code=%d, Body=%s\n\n", resp.StatusCode, body)
}
