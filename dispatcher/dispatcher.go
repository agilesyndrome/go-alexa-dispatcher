package dispatcher

import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-dispatcher/LaunchRequest"
)

var (
  RequestMap map[string]interface{}
)

func Dispatch(request alexa.Request) (alexa.Response) {
  switch request.Body.Type {
    case "LaunchRequest":
      response = staticintent.Simple(request, "welcome"
    case "IntentRequest":
      response = staticintent
  } 
}

func init() {
  RequestMap['LaunchRequest'] = LaunchRequest.Welcome
}
