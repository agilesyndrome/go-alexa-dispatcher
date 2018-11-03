package dispatcher

import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-staticintent/staticintent"
)

func Welcome(request alexa.Request) (alexa.Response) {
  return staticintent.Simple(request, "welcome")
}
