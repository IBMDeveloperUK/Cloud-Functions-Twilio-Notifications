const twilio = require('twilio');

function main(params) {
    if (params.action == "assigned") {
    var accountSid = params.accountsid;
    var authToken = params.accountauthtoken;
    const client = twilio(accountSid, authToken);
    return client.messages
      .create({
         body: 'New pull request assignee',
         from: '+447723452145',
         to: '+447710653583'
       })
      .then(function(message) {
        return { "message": "SMS sent" };
      });
    } else {
        return { "message": "SMS not sent" };
    }
}