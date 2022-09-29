package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"gitlab.com/Yinebeb-01/ethiopiandateconverter/ethioGrego"

	"github.com/cucumber/godog"
)

var host = "http://localhost:8080"

var res *http.Response

type response struct {
	Msg string `json:"response"`
}
type notFound struct {
	Msg string
}

var expectedEthio, expectedGrego string

// sample unit-test
func TestEthiopianDate(t *testing.T) {
	ethiopianDate := "2015-01-18 00:00:00 +0000 UTC"
	gregorianDate := "2022-09-28 00:00:00 +0000 UTC"

	time, err := ethioGrego.To_ethiopian(2022, 9, 28)
	if err == nil {
		expectedEthio = time.String()
	}
	time, err = ethioGrego.To_gregorian(2015, 1, 18)
	if err == nil {
		expectedGrego = time.String()
	}
	if !reflect.DeepEqual(ethiopianDate, expectedEthio) && !reflect.DeepEqual(gregorianDate, expectedGrego) {
		t.Errorf("Date conversion not work got %v want %v", expectedEthio, ethiopianDate)
	}
}

//feature file steps definition- BDD

// sending a request
func aRequestIsSentToTheEndpoint(method, endpoint string) error {
	var reader = strings.NewReader("")
	var request, err = http.NewRequest(method, host+endpoint, reader)
	if err != nil {
		return fmt.Errorf("could not create request %s", err.Error())
	}

	res, err = http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("could not send request %s", err.Error())
	}
	return nil
}

func theHTTPresponseCodeShouldBe(expectedCode int) error {
	if expectedCode != res.StatusCode {
		return fmt.Errorf("status code not as expected! Expected '%d', got '%d'", expectedCode, res.StatusCode)
	}
	return nil
}

func theResponseContentShouldBe(expectedContent string) error {
	body, _ := ioutil.ReadAll(res.Body)
	resp := response{}
	json.Unmarshal(body, &resp)
	if resp.Msg == "" {

	}

	if expectedContent != resp.Msg && expectedContent != string(body) {
		return fmt.Errorf("response content not as expected! Expected '%s', got '%s'", expectedContent, string(body))
	} else {
		return nil
	}
}

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^a "([^"]*)" request is sent to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)
	s.Step(`^the HTTP-response code should be "(\d+)"$`, theHTTPresponseCodeShouldBe)
	s.Step(`^the response content should be "([^"]*)"$`, theResponseContentShouldBe)
}
