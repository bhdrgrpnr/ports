package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"saasteamtest/saasbackend/data"
	"saasteamtest/saasbackend/domain"
	"saasteamtest/saasbackend/models"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	testHelper  TestHelper
	portService domain.PortServiceInterface
)

func TestMain(m *testing.M) {
	setUp()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func setUp() {
	testHelper = TestHelper{}
	db := data.DB{PortTable: map[string]models.Port{}}
	portHandler := data.NewPortHandler(db)
	portService = domain.NewPortService(portHandler)
}

type TestHelper struct{}

// JSONBytesEqual compares the JSON in two byte slices.
func (s *TestHelper) JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, fmt.Errorf("unmarshal: %w", err)
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, fmt.Errorf("unmarshal: %w", err)
	}

	return reflect.DeepEqual(j2, j), nil
}

func (s *TestHelper) MapJSONBodyIsEqualString(t *testing.T, responsString string, myPB string) {
	testBytes := []byte(myPB)
	responseBytes := []byte(responsString)
	eq, err := s.JSONBytesEqual(responseBytes, testBytes)
	if err != nil {
		t.Errorf("error JSONBytesEqual: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			t.Errorf("syntax error at byte offset %d", e.Offset)
		}
		t.Errorf("JSONBytesEqual responseBytes: %q", responseBytes)
		t.Error(err)
	}
	if !eq {
		t.Errorf("\n...JSONBytesEqual expected = %v\n...obtained = %v", string(testBytes), string(responseBytes))
	}

	testMap := map[string]interface{}{}
	err = json.Unmarshal(testBytes, &testMap)
	if err != nil {
		t.Error(err)
	}
	responseMap := map[string]interface{}{}
	err = json.Unmarshal(responseBytes, &responseMap)
	if err != nil {
		t.Error(err)
	}
	// compare the json ojbects without respect to order a second time with a different function
	if !cmp.Equal(testMap, responseMap) {
		diffstring := cmp.Diff(testMap, responseMap)
		t.Errorf("\n...cmp Diff string  = %v", diffstring)
		t.Errorf("\n...cmp Equal false expected = %v\n...obtained = %v", testMap, responseMap)
	}
}
