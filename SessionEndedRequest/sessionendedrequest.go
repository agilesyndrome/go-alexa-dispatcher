package goodbye

import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-staticintent/staticintent"
)

func Handler(request alexa.Request) (alexa.Response, error) {
  retValue := staticintent.Simple(request, "goodbye")
  return retValue, nil

}
