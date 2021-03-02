const twilio = require('twilio');

function main(params) {
    
var accountSid = params.accountsid;
var authToken = params.accountauthtoken;
var recipientNumber = params.recipientNumber;
var twilioNumber = params.twilioNumber;

const client = twilio(accountSid, authToken);
return client.messages
  .create({
     body: 'This is a node test',
     from: twilioNumber,
     to: recipientNumber
   })
  .then(function(message) {
    return { "messagesid": message.sid };
  });
}