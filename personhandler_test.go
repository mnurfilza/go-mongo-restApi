package weinvite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dummyData = []*Person{
	&Person{Nama: "faisal", Alamat: "Kayuringin"},
	&Person{Nama: "akbar", Alamat: "Kayuringin"},
	&Person{Nama: "addien", Alamat: "Kayuringin"},
}

func TestHandlerPost(t *testing.T) {
	router := setupRouter()
	for _, item := range dummyData {
		body, err := json.MarshalIndent(item, "", " ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/person", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	}

}

func TestHandlerGet(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/person/filza", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)
	// assert.Equal(t, 200, w.Code)

}
