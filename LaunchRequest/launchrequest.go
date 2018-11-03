package dispatcher

import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-staticintent/staticintent"
)

func Welcome(request alexa.request) (alexa.response) {
  return staticintent.Simple(request, "welcome")
}
