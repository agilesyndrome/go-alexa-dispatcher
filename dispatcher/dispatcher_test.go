package dispatcher

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/arienmalec/alexa-go"
	"github.com/agilesyndrome/go-alexa-i18n/alexai18n"
//	"github.com/davecgh/go-spew/spew"
)

//HELPER FUNCTIONS
func checkResponse(testFile string) (alexa.Response) {
   test_data, json_err := ioutil.ReadFile(testFile)
   check(json_err)
   testRequest := new(alexa.Request)
   err := json.Unmarshal(test_data, &testRequest)
   response, err := Dispatch(*testRequest )
   check(err)

   return response

}

func assertSimpleResponse(testFile, expectedTitleText, expectedResponseText string, t *testing.T) {
   response := checkResponse(testFile)
   expectedResponse := alexa.NewSimpleResponse(expectedTitleText, expectedResponseText)
   assert.IsType(t, alexa.Response{}, response)
   assert.Equal(t, expectedResponse, response)
   assert.Equal(t, true, response.Body.ShouldEndSession)
}

func assertSimpleInternationalResponse(testFile, expectedTitle, expectedText string, t *testing.T) {
    reqEnUs := alexai18n.CultureRequest("en-US")
    var expectedTitleText = alexai18n.WorldString(reqEnUs, expectedTitle)
    var expectedResponseText = alexai18n.WorldString( reqEnUs, expectedText)
    assertSimpleResponse(testFile, expectedTitleText, expectedResponseText, t)
}


func TestLaunchRequest(t *testing.T) {
    assertSimpleResponse("../tests/LaunchRequest.json", "My Carpenter", "World", t)

}

func BasicTests(t *test.T) {
  assert.Equal(t, true, welcome.Test())
}

func TestSessionEndedRequest(t *testing.T) {
  assertSimpleResponse("../tests/SessionEndedRequest.json", "Things For Life", "World", t)
}

