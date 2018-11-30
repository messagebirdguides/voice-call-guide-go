package main

import (
	"fmt"
	"log"

	"github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/voice"
)

func main() {
	//client := messagebird.New(os.Getenv("MESSAGEBIRD_API_KEY"))
	client := messagebird.New("")

	source, dest := "", ""
	callflow := voice.CallFlow{
		Title: "Test flow",
		Steps: []voice.CallFlowStep{
			&voice.CallFlowSayStep{
				Voice:    "female",
				Payload:  "Hey you, a little bird told me you wanted a call!",
				Language: "en-GB",
			},
		},
	}

	call, err := voice.InitiateCall(
		client,
		source,
		dest,
		callflow,
		nil,
	)
	if err != nil {
		switch errResp := err.(type) {
		case messagebird.ErrorResponse:
			for _, mbError := range errResp.Errors {
				fmt.Printf("Error: %#v\n", mbError)
			}
		}

		return
	}
	// We're logging call for development purposes. You can safely discard this variable in production.
	log.Println(call)
}
