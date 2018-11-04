package dispatcher

import (
  "errors"
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-dispatcher/LaunchRequest"
  "github.com/agilesyndrome/go-alexa-dispatcher/SessionEndedRequest"
  "github.com/agilesyndrome/go-alexa-dispatcher/ErrorResponse"
  "github.com/agilesyndrome/go-alexa-staticintent/staticintent"
)

type AlexaHandler func (alexa.Request) (alexa.Response, error)

var (
  RequestMap map[string] AlexaHandler = map[string] AlexaHandler {}
  IntentMap map[string] AlexaHandler = map[string] AlexaHandler {}
)


func Dispatch(request alexa.Request) (alexa.Response, error) {
  vf := RequestMap[request.Body.Type]

  /*Amazon could potentially add new things here and break this on a whim
    Hopefully this response will keep things sluding along, but you should
    make sure something like IOPipes or other logging tool records this error
  */

  var response alexa.Response
  var err error

  if (vf == nil) {
    vf = kaput.Handler
    response,err = vf(request)

    //We don't need any errors from kaput, we'll make our own at this point
    //If kaput is broken, we're all fucked anyway
    err = errors.New("I don't know how to handle a " + request.Body.Type + "request type")
  } else {
    response, err = vf(request)
  }
  
  return response, err
}

func DispatchIntent(request alexa.Request) (alexa.Response, error) {
  var response alexa.Response

  vf := IntentMap[request.Body.Intent.Name]

  //You didn't register an intent, so we'll use the static intent to just read text
  if ( vf == nil ) {
   vf = staticintent.Handler
  }

  response, err := vf(request)
  return response, err

}

func Validate() {
  /* Nice try compiler, but yes, I actually do use the welcome module I import...
     you just don't notice it because it's called virtually from the RequestMap section.
     But since you're being so persnickety, I'll just run a stupid .Test() function on each
     import and return true, just to make you happy. M'kay pumpkin?
  */
  welcome.Test()
  goodbye.Test()
}

func init() {
  //Four main types of Requests: Launch, CanFulfill, SessionEnded, and Intent
  RequestMap["LaunchRequest"] = welcome.Handler
  //RequestMap["CanFulfillIntentRequest"] = dispatcher.NotImplemented
  RequestMap["SessionEndedRequest"] = goodbye.Handler
  RequestMap["IntentRequest"] = DispatchIntent
  Validate()


}
