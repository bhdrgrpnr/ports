package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"saasteamtest/backend/models"
	"testing"

	"github.com/go-chi/chi"
)

// shortcut for creating all the JSON objects we will send and receive.
type MSI map[string]interface{}

func TestCreatePort(t *testing.T) {

	port := models.Port{Name: "somename", City: "istanbul", Province: "turkey", Country: "turkey", Coordinates: []float64{10.1815316, 36.8064948},
		TimeZone: "asia/istabul", Unlocks: []string{"TNTUN"}}

	portCreateRequest := models.PortRequest{Index: "ABCDEF", Port: port}

	body, err := json.Marshal(portCreateRequest)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/ports", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("PUT", "/ports", BaseHandler(CreatePort(portService)))
	r.ServeHTTP(rec, req)

	expectedResult, err := json.Marshal(port)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(rec.Code, 201) {
		t.Errorf("unexpected response code")
	}
	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestUpdatePortFail(t *testing.T) {

	port := models.Port{Name: "somename", City: "istanbul", Province: "turkey", Country: "turkey", Coordinates: []float64{10.1815316, 36.8064948},
		TimeZone: "asia/istabul", Unlocks: []string{"TNTUN"}}

	portUpdateRequest := models.PortRequest{Index: "ABCDEF", Port: port} // nonexistent index

	body, err := json.Marshal(portUpdateRequest)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/ports", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("PUT", "/ports", BaseHandler(UpdatePort(portService)))
	r.ServeHTTP(rec, req)

	_, err = json.Marshal(port)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(rec.Code, 400) {
		t.Errorf("unexpected response code")
	}
}

func TestUpdatePortSuccess(t *testing.T) {

	port := models.Port{Name: "somename", City: "istanbul", Province: "turkey", Country: "turkey", Coordinates: []float64{10.1815316, 36.8064948},
		TimeZone: "asia/istabul", Unlocks: []string{"TNTUN"}}

	portUpdateRequest := models.PortRequest{Index: "AEAJM", Port: port} // existent index

	body, err := json.Marshal(portUpdateRequest)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/ports", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("PUT", "/ports", BaseHandler(UpdatePort(portService)))
	r.ServeHTTP(rec, req)

	expectedResult, err := json.Marshal(port)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(rec.Code, 201) {
		t.Errorf("unexpected response code")
	}
	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}
