package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dyxj/loan-plan/api"
	"github.com/dyxj/loan-plan/loan"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGenerateLoanPlan(t *testing.T) {
	rt := http.NewServeMux()
	api.InitAPIRoutes(rt)

	srv := httptest.NewServer(rt)
	defer srv.Close()

	data := map[string]interface{}{
		"loanAmount":  "5000",
		"nominalRate": "5.0",
		"duration":    24,
		"startDate":   "2018-01-01T00:00:01Z",
	}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(data)
	if err != nil {
		t.Fatalf("could not encode data: %v", err)
	}

	res, err := http.Post(fmt.Sprintf("%s/generate-plan", srv.URL), "application/json", b)
	if err != nil {
		t.Fatalf("could not send Post Request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	d := json.NewDecoder(res.Body)
	var loanResp []loan.RepayMonthDollars
	err = d.Decode(&loanResp)
	if err != nil {
		t.Fatalf("could not decode res body: %v", err)
	}

	var ExpData []loan.RepayMonthDollars
	err = json.Unmarshal([]byte(expDataStr), &ExpData)
	if err != nil {
		t.Fatalf("could not unmarshal expected results: %v", err)
	}

	dataCheck := reflect.DeepEqual(loanResp, ExpData)
	if !dataCheck {
		t.Fatalf("response does not match expected results")
	}
}

func TestGenerateLoanPlanMethodNotAllowed(t *testing.T) {
	rt := http.NewServeMux()
	api.InitAPIRoutes(rt)

	srv := httptest.NewServer(rt)
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/generate-plan", srv.URL))
	if err != nil {
		t.Fatalf("could not send Get Request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status Method Not Allowed; got %v", res.Status)
	}
}
