# Create an IBM Cloud Function

## Step 1 - Sign Up or Login to IBM Cloud

[Login or Sing Up]() to IBM Cloud

## Step 2 - Create a Cloud Functions Action

Once logged in, in the top search bar enter `cloud functions` and select the option with the `f` symbol.

![ibm cloud functions search](../workshop-assets/ibm-cloud/functions-search.png "Functions Search")

This will take you to the Functions dashboard. Select `Actions` from the side bar.

![functions dashboard](../workshop-assets/ibm-cloud/functions-dashboard.png "Functions Dashboard")

Here you will have 5 single entities to choose from. Select `Action`.

![create action dashboard](../workshop-assets/ibm-cloud/create-dashboard.png "Create Action Dashboard")

Give the `Action` a name and select the Runtime. There is no need to change the package so this can be left as `(Default Package)`. For the purpose of this workshop I have chosen `Go`. Feel free to mix it up and choose a language you are most familiar with. The principals of this workshop are very much the same across the board.

![create action](../workshop-assets/ibm-cloud/create-action.png "Create Action")

## Step 3 - Set up the Action

As it stands this Action is not public and this prevents Webhooks and other public HTTP actions from interacting with it. To change this you need to make it a `Web Action`. Select `Endpoints` from the side bar.

![select endpoints](../workshop-assets/ibm-cloud/boilerplate-code.png "Select Endpoints")

Select the checkbox `Enable as Web Action` and click `Save`. You will notice the `Web Action` icon change and you will be able to see the pub HTTP URL.

![enable web action](../workshop-assets/ibm-cloud/enable-web-action.png "Enable Web Action")

Now the `Action` is public and we can hit the endpoint from external sources, we need to set up some parameters for the code to use. This is essentially an environment variable for the function.

For this, you will need to add the following:

`recipientNumber` = The number you wish to send a text message too

`authToken` = Your Twilio Account Auth Token (Found on your dashboard)

`accountSid` = Your Twilio Account SID (Found on your dashboard)

`twilioNumber` = Your number associated with your Twilio Account

![action parameters](../workshop-assets/ibm-cloud/action-params.png "Action Parameters")

## Step 4 - Create the function code

Read the code below line by line to understand what is happening. If you are not using `Go`, then do not panic as this can easily be translated and it's fairly simple to understand. 

```go
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
```

## **Have you spotted anything odd in the code?**

`action := params["action"].(string)`

We haven't specified an `action` parameter for our function. The reason we haven't is because this will be within the payload of the `POST` request that hits this endpoint. 

Let me explain.

When the GitHub Webhook sends a `POST` request to this endpoint, it will carry with it a payload of information. This payload will tell us much more information about the changes being made on GitHub. 

Visit [GitHub Webhook events](https://docs.github.com/en/developers/webhooks-and-events/webhook-events-and-payloads#webhook-payload-object-common-properties) for more information about each event payload.

For this workshop, I am using the [pull_request](https://docs.github.com/en/developers/webhooks-and-events/webhook-events-and-payloads#pull_request) event. You can see here, the `key` is `action` and it can have many values. In the code snippet above, we are making decisions based on the value of `action`.

