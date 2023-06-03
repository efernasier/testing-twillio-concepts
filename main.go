package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	// Find your Account SID (TWILIO_ACCOUNT_SID) and Auth Token (TWILIO_AUTH_TOKEN) at twilio.com/console
	// and set the environment variables. See http://twil.io/secure

	twilioPhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	destinationPhoneNumber := os.Getenv("DESTINATION_PHONE_NUMBER")

	sendSMS(twilioPhoneNumber, destinationPhoneNumber)
	sendMMS(twilioPhoneNumber, destinationPhoneNumber)
	call(twilioPhoneNumber, destinationPhoneNumber)

}

// Consider that is not possible send messages to unverified numbers.
// Trials accounts only can send SMS to verified accounts.

/*
	Status: 400 - ApiError 21608: The number  is unverified.
	Trial accounts cannot send messages to unverified numbers;
	verify  at twilio.com/user/account/phone-numbers/verified,
	or purchase a Twilio number to send messages to unverified numbers.
	(null) More info: https://www.twilio.com/docs/errors/21608
*/
func sendSMS(twilioPhoneNumber, destinationPhoneNumber string) {

	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetFrom(twilioPhoneNumber)
	params.SetBody("Ahoy! world")
	params.SetTo(destinationPhoneNumber)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}

func sendMMS(twilioPhoneNumber, destinationPhoneNumber string) {
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}

	params.SetFrom(twilioPhoneNumber)
	params.SetBody("Ahoy! world!")
	params.SetMediaUrl([]string{"https://demo.twilio.com/owl.png"})
	params.SetTo(destinationPhoneNumber)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}

// Example has been taken from:
// https://www.twilio.com/docs/voice/api/call-resource
func call(twilioPhoneNumber, destinationPhoneNumber string) {
	client := twilio.NewRestClient()

	params := &api.CreateCallParams{}
	params.SetFrom(twilioPhoneNumber)
	params.SetTo(destinationPhoneNumber)
	params.SetTwiml("<Response><Say>Ahoy there!</Say></Response>")

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}
