package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Main is the function implementing the action
func Main(params map[string]interface{}) map[string]interface{} {

	action := params["action"].(string)
	twilioNumber := params["twilioNumber"].(string)
	recipientNumber := params["recipientNumber"].(string)

	// only invoke Twilio message service if the GitHub PR action = assigned
	if action == "assigned" {

		fmt.Println("pull request assigned")

		// set account info
		accountSid := params["accountSid"].(string)
		authToken := params["authToken"].(string)
		urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

		// text message being sent to recipient
		textMsg := "New pull request assignee"

		// package the data values
		msgData := url.Values{}
		msgData.Set("To", recipientNumber)
		msgData.Set("From", twilioNumber)
		msgData.Set("Body", textMsg)
		msgDataReader := *strings.NewReader(msgData.Encode())

		msg := request(authToken, accountSid, urlStr, msgDataReader)

		return msg
	}

	fmt.Println("Pull request action = ", action)
	msg := make(map[string]interface{})
	msg["action"] = action

	// return the output JSON
	return msg
}

func request(authToken, accountSid, urlStr string, msgDataReader strings.Reader) map[string]interface{} {
	// create HTTP client, req & set req headers
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// HTTP POST request and return message SID to the console
	resp, _ := client.Do(req)
	var msg = make(map[string]interface{})
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
			msg["status"] = "sent"
		}
	} else {
		fmt.Println(resp.Status)
		msg["status"] = "not sent"
	}
	return msg
}
