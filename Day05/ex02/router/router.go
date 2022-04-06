package router

import (
	"encoding/json"
	"ex02/dict"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	port = "6505"
	ip   = "localhost"
)

type Success struct {
	Success bool `json:"Success"`
}

type KeyValue struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

var (
	g_dict = dict.NewDict()
)

// func GetMiddleware(next http.Handler) http.Handler {
// 	return
// 	http.Error(w)
// }

var GetRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {
		key := req.FormValue("WORD")
		val, err := g_dict.Read(key)

		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			fmt.Fprint(writer, val)
		}
	},
)

var PostRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {

		var data KeyValue
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&data)

		encoder := json.NewEncoder(writer)
		var success Success

		err := g_dict.Create(data.Key, data.Value)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			success = Success{false}
		} else {
			writer.WriteHeader(http.StatusCreated)
			success = Success{true}
		}
		encoder.Encode(&success)
	},
)

var PutRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {
		var data KeyValue
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&data)
		err := g_dict.Update(data.Key, data.Value)

		encoder := json.NewEncoder(writer)
		var success Success

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			success = Success{false}
		} else {
			success = Success{true}
		}
		encoder.Encode(success)
	},
)

var DeleteRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {
		key := req.FormValue("WORD")
		err := g_dict.Delete(key)

		encoder := json.NewEncoder(writer)
		var success Success

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			success = Success{false}
		} else {
			success = Success{true}
		}
		encoder.Encode(success)
	},
)

func Handler() http.Handler {
	router := mux.NewRouter()
	// router.Methods("GET")
	router.Handle("/dict", GetRequest).Methods("GET")
	router.Handle("/dict", PostRequest).Methods("POST")
	router.Handle("/dict", PutRequest).Methods("PUT")
	router.Handle("/dict", DeleteRequest).Methods("DELETE")

	return router
}

func StartServ() {

	srv := &http.Server{
		Addr:         ip + ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      Handler(),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
