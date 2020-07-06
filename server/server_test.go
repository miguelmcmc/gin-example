package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

var address = "localhost"
var port = "8080"

func initServer() {
	HTTPServer(address, port).Init()
}
func Test_httpServer_Init(t *testing.T) {
	go initServer()
}

func Test_Root(t *testing.T) {
	server := httpServer{"", port, nil}
	server.setUp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	server.router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}

func Test_GetRandom(t *testing.T) {
	server := httpServer{"", port, nil}
	server.setUp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/random", nil)
	server.router.ServeHTTP(w, req)

	if 200 != w.Code {
		t.Error("Code Incorrect", w.Body.String())
	}
	fmt.Println("Response (", w.Code, "): ", w.Body.String())

}
func Test_PostRandomLenght(t *testing.T) {

	server := httpServer{"", port, nil}
	server.setUp()

	//Client part
	type bodyS struct {
		Lenght int
	}

	body, err := json.Marshal(bodyS{150})
	if err != nil {
		t.Error("Error creating body", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test/random", bytes.NewBuffer(body))
	server.router.ServeHTTP(w, req)

	// assert.Equal(t, 200, w.Code)
	if 200 != w.Code {
		t.Error("Code Incorrect", w.Body.String())
	}
	fmt.Println("Response (", w.Code, "): ", w.Body.String())
}
