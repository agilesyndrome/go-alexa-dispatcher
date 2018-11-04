package stop

import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-staticintent/staticintent"
)

func Test() (bool) {
  return true
}

func Handler(request alexa.Request) (alexa.Response, error) {
  retValue, err := staticintent.Simple(request, "stop")
  return retValue, err

}
