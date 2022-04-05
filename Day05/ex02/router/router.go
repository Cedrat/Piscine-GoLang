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
	Success bool
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
		// fmt.Fprint(writer, val)
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
		fmt.Fprintln(writer, data)
		g_dict.Create(data.Key, data.Value)
	},
)

var PutRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {
		var data KeyValue
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&data)
		fmt.Fprintln(writer, data)
		g_dict.Update(data.Key, data.Value)
	},
)

var DeleteRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {
		key := req.FormValue("WORD")
		err := g_dict.Delete(key)
		if err != nil {
			// fmt.Fprint(writer, err.Error())
			writer.WriteHeader(http.StatusNotFound)

		}
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
