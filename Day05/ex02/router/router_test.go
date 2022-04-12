package router

//curl -X GET "localhost:6505/dict?WORD=my_login" -v
//curl -X POST localhost:6505/dict -H 'Content-Type: application/json' -d '{"Key":"my_login","Value":"my_password"}'
//curl -X GET "localhost:6505/dict?WORD=my_login" -v
//curl -X PUT localhost:6505/dict -H 'Content-Type: application/json' -d '{"Key":"my_login","Value":"DEUXOC"}'
//curl -X GET "localhost:6505/dict?WORD=my_login" -v
//
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	// "io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func CreateRequest(method string, url string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {

	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func CreatePostPutRequest(method string, req_url string, key string, value string) *http.Response {
	client := &http.Client{}

	values := KeyValue{key, value}

	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest(method, req_url, bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {

	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func TestRouterDict(T *testing.T) {
	assert := assert.New(T)

	wg.Add(1)
	go func() { StartServ() }()

	fmt.Println("launch")
	res := CreateRequest("GET", "http://localhost:6505/dict?WORD=Cedrat")

	assert.Equal("404 Not Found", res.Status, "Word wasnt implemented in the dict")

	res = CreatePostPutRequest("POST", "http://localhost:6505/dict", "define", "define meaning defining")

	assert.Equal("201 Created", res.Status, "Post a new Key result in a created answer")
	res.Body.Close()

	res = CreateRequest("GET", "http://localhost:6505/dict?WORD=define")

	assert.Equal("200 OK", res.Status, "define should be created")

	b, _ := ioutil.ReadAll(res.Body)
	assert.Equal("The definition of the word define is : define meaning defining", string(b), "definition of define must be set")
	res.Body.Close()

	res = CreatePostPutRequest("PUT", "http://localhost:6505/dict", "define", "define meaning define")
	assert.Equal("200 OK", res.Status, "Put a current key update this")
	res.Body.Close()

	res = CreateRequest("GET", "http://localhost:6505/dict?WORD=define")

	assert.Equal("200 OK", res.Status, "define should be created")

	b, _ = ioutil.ReadAll(res.Body)
	assert.Equal("The definition of the word define is : define meaning define", string(b), "definition of define must be set")
	res.Body.Close()

	res = CreatePostPutRequest("PUT", "http://localhost:6505/dict", "defineeee", "define meaning define")
	assert.Equal("400 Bad Request", res.Status, "Put a key who no exist will fail")
	res.Body.Close()
}
