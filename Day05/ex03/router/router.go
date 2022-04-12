package router

import (
	"encoding/json"
	"ex03/dict"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	port = "6505"
	ip   = "10.12.7.14"
)

var logger = logrus.New()

func init_logger() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only logger the warning severity or above.
	logger.SetLevel(logrus.InfoLevel)
}

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

func GetIp(req *http.Request) string {

	ip := req.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return string(netIP)
	}

	ips := req.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip
		}
	}

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return ""
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip
	}

	return ("No valid ip found")
}

var GetRequest = http.HandlerFunc(
	func(writer http.ResponseWriter, req *http.Request) {
		logger.Info(GetIp(req) + " trying to access of a define to a word.")
		key := req.FormValue("WORD")
		val, err := g_dict.Read(key)

		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			logger.Info(GetIp(req) + " not find a definition for the word  " + key)
		} else {
			fmt.Fprint(writer, val)
			logger.Info(GetIp(req) + " get access to the definition of " + key)

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

		logger.Warn(GetIp(req) + " try to add word to the dictionnary")

		err := g_dict.Create(data.Key, data.Value)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			logger.Warn(GetIp(req) + " try with no success to upload a definition ")
			success = Success{false}
		} else {
			writer.WriteHeader(http.StatusCreated)
			success = Success{true}
			logger.Warn(GetIp(req) + " add with success " + data.Key + ":" + data.Value)
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

		logger.Info(GetIp(req) + " try to modify a word in the dictionnary")

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			success = Success{false}
			logger.Info(GetIp(req) + " try with no success to modify a definition ")

		} else {
			success = Success{true}
			logger.Warn(GetIp(req) + " modify with success the define " + data.Key + "for " + data.Value)
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

		logger.Info(GetIp(req) + " trying to delete an entry")

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			success = Success{false}
			logger.Info(GetIp(req) + " failed to delete an entry ")

		} else {
			logger.Warn(GetIp(req) + " delete with success the following entry : " + key)
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
	init_logger()

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
