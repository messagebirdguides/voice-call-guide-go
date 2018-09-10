# Make your first Voice Call with MessageBird
### ⏱ 10 min build time

It's time to make your first voice call using the [MessageBird Voice
API](https://developers.messagebird.com/docs/voice-calling)! Before we
get started, have you set up your Go development environment and
project directory with the MessageBird SDK? 

- **No** - make sure you [read this MessageBird Developer Guide](https://developers.messagebird.com/guides/setup-local-dev-environment) before getting started. 
- **Yes!** - Great! Now you can make your first API request and make a voice call with MessageBird using Go.

## Getting started

First, let's create a new Go project directory. You can do this anywhere, but best practice says that you should make this folder in your `$GOPATH/src/github.com/<your-username>` folder (e.g. your project folder can be `$GOPATH/src/github.com/birddev/send-sms`).

If you're not sure where your `$GOPATH` is located, running `go env GOPATH` will display it in your terminal.

In your project root, create a `main.go` file and open it in your text editor of choice and stub out our program by writing the following code:

````go
package main

import (
    "log"

    "github.com/messagebird/go-rest-api"
)

func main(){
    client := messagebird.New("<your-api-key>")
    log.Println("Hello!")
}
````

Here, we've imported the MessageBird SDK for Go and created a `client` object 
that we'll use to interact with the MessageBird REST API. To create our `client`
object, we make a `messagebird.New()` call which expects a single argument — your [API key](https://dashboard.messagebird.com/en/developers/access). Note that you can create either a **test** or **live** API key:

- **test** API keys simulate responses from the MessageBird server, allowing you to test your code flow and error handling, and before sending real messages. (You still must have an internet connection to use MessageBird REST API features.)
- **live** API keys allow you to send actual messages to your recipients. We recommend that you do not publish this key anywhere.

In order to start using the SDK, replace `<your-api-key>` with your API
key. 

**Pro-tip:** Here, we're hardcoding your API key in your program to keep the guides straightforward. But for production applications, we recommended storing
the key in a configuration file or environment variable instead and pass
this variable with the key to the require function. You'll see this in
practice later in our MessageBird Developer Guides for building common
use cases.

## Making a voice call

The SDK has a `voice` library that allows you to initiate a call using MessageBird's REST API. To make a call:

1. Add the `voice` package to your import statement in `main.go`:

    ````go
    import (
        "log"

        "github.com/messagebird/go-rest-api"
        "github.com/messagebird/go-rest-api/voice"
    )
    ````

2. We'll stub out and log the output of our `voice.InitiateCall()` function call. Modify the `main()` block in `main.go` to look like this:

    ````go
    func main(){
        client := messagebird.New("<your-api-key>")

        call, err := voice.InitiateCall(
            client,
            source,
            destination,
            callflow,
            nil,
        )
        if err != nil {
            log.Println(err)
        }
        // We're logging call for development purposes. You can safely discard this variable in production.
        log.Println(call)
    }
    ````

Here, we're calling `voice.InitiateCall()` with these parameters:

- `client`: This is the client object that we've created earlier.
- `source`: The "originator" which is displayed as the caller's ID on the receiving end.
- `destination`: The "destination" of the call, which is a phone number written in international format including its country code. For example, "+319876543210".
- `callflow`: A call flow object that describes to the MessageBird servers the events that should occur during a call.
- `nil`: Optional attributes that you can specify for this phone call. Here, we're setting it as `nil` because we have no additional attributes to add. You can find more information about these optional attributes in the [Voice Calling API Reference](https://developers.messagebird.com/docs/voice-calling#calls).

Now that we've got `voice.InitiateCall()` stubbed out, we can define these parameters. Add these lines of code just on top of the `voice.InitiateCall()` block:

````go
// We've already defined client

source, destination := "+319876543211", "+319876543210"

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
````

We've defined our `source` and `destination` using example phone numbers. For your application to work, you should replace these numbers with working phone numbers.

After that, we're writing our call flow. The call flow
specifies one or more steps that will execute once the receiver picks up
the call. They can be different actions, including playing audio
files, reading a block of text with a synthesized voice (TTS), or
forwarding the call to another phone.

Inside our `voice.CallFlow{}` struct, we have the following keys:

- `Title`: Each call flow must have a title, and should describe what the call flow does. The title itself is not executed as part of the call flow.
The flow needs an attribute called `title` that describes the flow. It
can help you understand the flow but does not affect its execution.
- `Steps`: The rest of the call flow is an array of "steps", or `voice.CallFlowStep` structs. The MessageBird SDK for Go provides a different struct type for each action type that a call flow step makes. Here, we're adding only one call flow step to the array: a `voice.CallFlowSayStep` struct to tell MessageBird to make a call that speaks a given line of text.

In our `voice.CallFlowSayStep` struct, we need to set three parameters:

- `Voice`: The voice to use when making the call. This can be set to either "male" or "female".
- `Payload`: This is the text will be spoken during the call.
- `Language`: The language of the text to be spoken, written as "en_gb" or similar.

These three parameters are required when setting a `voice.CallFlowSayStep` call flow step.

To learn more about calls, call flows and steps, especially the types of
actions that are available and which options are required or optionally
available for them, read [the respective section in the Voice Calling
API reference documentation](https://developers.messagebird.com/docs/voice-calling).

## Finishing your program

Once you've done all that, you have a fully functioning Go program that makes a call to your set destination phone number when run.

To test your application, run in the terminal:

````go
go run main.go
````

If everything worked, you should see the API response logged in the terminal, signalling that you've successfully made a call.

If you used a live API key and had verified your number or added credits to your account,
your phone will ring shortly and deliver the message when you pick up the phone. 
Congratulations, you just initiated your first voice call with MessageBird!

If you see an error from the script, try to read and understand the
message and fix the problem.

Next steps
----------

Let's head over to the next **MessageBird Developer Guide** learn how
to [set up and handle Incoming Voice
Calls]()<https://developers.messagebird.com/guides/handle-incoming>).

Want to start building your solution but not quite sure how to get
started? Please feel free to let us know at support@messagebird.com,
we'd love to help!
