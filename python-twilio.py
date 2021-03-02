import sys
import requests
from requests.auth import HTTPBasicAuth

def main(dict):
    if dict["action"] == "assigned":
        accountSid = dict["accountSid"]
        msgData = {
            "To": dict["recipientNumber"],
            "From": dict["twilioNumber"],
            "Body": "New pull request assignee"
        }
        urlStr = "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
        
        req = requests.post(urlStr, data=msgData, auth=HTTPBasicAuth(dict["accountSid"], dict["authToken"]), headers={"Accept":"application/json", "Content-Type":"application/x-www-form-urlencoded"})
        
        print(req.status_code)
        return { 'message': 'SMS sent'}
    else:
        return { 'message': 'SMS not sent'}