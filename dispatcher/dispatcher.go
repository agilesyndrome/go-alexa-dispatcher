package dispatcher

import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-dispatcher/LaunchRequest"
  "github.com/agilesyndrome/go-alexa-dispatcher/SessionEndedRequest"
)

type AlexaHandler func (alexa.Request) (alexa.Response, error)

var (
  RequestMap map[string] AlexaHandler = map[string] AlexaHandler {}
  IntentMap map[string] AlexaHandler = map[string] AlexaHandler {}
)

func check(e error) {
  if ( e != nil ) {
    panic(e)
  }
}

func Dispatch(request alexa.Request) (alexa.Response, error) {
  vf:= RequestMap[request.Body.Type]
  response, err := vf(request)
  check(err)
  return response, nil
}

func init() {
  //Four main types of Requests: Launch, CanFulfill, SessionEnded, and Intent
  RequestMap["LaunchRequest"] = welcome.Handler
  //RequestMap["CanFulfillIntentRequest"] = dispatcher.NotImplemented
  RequestMap["SessionEndedRequest"] = goodbye.Handler
  //RequestMap["IntentRequest"] = dispatcher.Intent

  welcome.Test()
  //goodbye.Test()

}
