// Notification handlers
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	palantir "github.com/welcome-to-shire/palantir-go"
)

func createNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subject := vars["subject"]

	log.Printf("creating new notification for %s", subject)

	payload, err := parseJsonBody(r)
	if err != nil {
		log.Print(err)
		errorResponse(w, 400, "Bad Payload")
		return
	}

	title, found := payload["title"]
	if !found {
		errorResponse(w, 400, "title is required")
		return
	}
	content, found := payload["content"]
	if !found {
		errorResponse(w, 400, "content is required")
		return
	}

	storage := GetStorageInstance()
	ticket, err := storage.Create(subject, title.(string), content.(string))
	if err != nil {
		log.Print(err)
		errorResponse(w, 500, "%q", err)
		return
	}

	w.WriteHeader(201)
	w.Write(palantir.Ticket{ticket}.MustMarshal())
}

func getNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	log.Printf("sending notification from %s", vars["subject"])

	storage := GetStorageInstance()
	message, err := storage.GetOne(vars["subject"])
	if err != nil {
		log.Print(err)
		errorResponse(w, 404, "404 Not Found")
		return
	}

	rv, err := json.Marshal(message)
	if err != nil {
		log.Print(err)
		errorResponse(w, 500, "%q", err)
		return
	}

	w.WriteHeader(200)
	w.Write(rv)
}

func errorResponse(w http.ResponseWriter, status int, reason_fmt string, a ...interface{}) {
	w.WriteHeader(status)

	rv := palantir.Error{fmt.Sprintf(reason_fmt, a...)}.MustMarshal()
	w.Write(rv)
}

func parseJsonBody(r *http.Request) (map[string]interface{}, error) {
	var reader io.Reader = r.Body
	rawBody, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}
	if err = json.Unmarshal(rawBody, &body); err != nil {
		return nil, err
	}

	return body, err
}
