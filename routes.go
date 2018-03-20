package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Service string
		Version string
		Commit  string
	}
	response := &Response{
		Service: Service,
		Version: Version,
		Commit:  Commit,
	}
	j, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(j))
	fmt.Printf("%+v\n", config)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://" + config.ImaginHost + ":" +
		strconv.Itoa(config.ImaginPort) + "/health")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		fmt.Fprintf(w, "Error!")
	} else {
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "not ok")
			log.Fatal(err)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "ok")
		}
	}
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profile := vars["profile"]
	source := vars["source"]

	if profile == "" || source == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	imageReader, err := ImaginaryOps(profile, source)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to process image %s", r.URL)
		fmt.Fprint(w, "%v", err)
		log.Fatal(err)
	} else {
		defer imageReader.Body.Close()
		_, err = io.Copy(w, imageReader.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", err)
			log.Fatal(err)
		}
	}
}
