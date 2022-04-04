package router

import (
	"ex02/dict"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	port = "4242"
	ip   = "localhost"
)

type KeyValue struct {
	Key   string
	Value string
}

var (
	g_dict = dict.NewDict()
)

var GetRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {
		key := req.FormValue("WORD")
		g_dict.Read(key)
		fmt.Fprint(writer, key)
	},
)

var PostRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {
		// var data KeyValue
		// // fmt.Fprintln(writer, req.Body)
		// decoder := json.NewDecoder(req.Body)
		// decoder.Decode(&data)
		fmt.Fprintln(writer, req.FormValue("d"))
	},
)

func Handler() http.Handler {
	router := mux.NewRouter()
	// router.Methods("GET")
	router.Handle("/dict", GetRequest).Methods("GET")
	router.Handle("/dict", PostRequest).Methods("POST")

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
