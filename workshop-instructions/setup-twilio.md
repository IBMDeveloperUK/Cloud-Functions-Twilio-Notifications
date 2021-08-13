# Create a Twilio Account

In this workshop, the Twilio API is going to be called and used to send a text message to your phone. It doesn't require a credit card to sign up and all you need to do is sign up for a [Twilio trial account](https://www.twilio.com/try-twilio) that includes $15 credit.

## Step 1 - Sign up
Using the link above, head over to Twilio and sign up for the free trial account and enter your details into the form shown.

![twilio sign up](../workshop-assets/twilio/twilio-sign-up.png "Twilio Sign Up")

## Step 2 - Buy a number
Once you have logged into your account you will need to set it up.

Head over to the left side panel and click on the `#`. 

Next, click on `Buy a number`. 

Select your country and tick the box `SMS` as this number only needs to send a text message. 

Click `Search` and purchase a number.

> :rotating_light: This will use your trial credit so do not panic :rotating_light:

![twilio buy a number](../workshop-assets/twilio/twilio-buy-number.png "Buy a Twilio Number")

## Step 3 - Check your dashboard

Navigate back to your dashboard and it should look something similar to this:

> Ensure you have a trial number visible

![twilio dashboard](../workshop-assets/twilio/twilio-dashboard.png "Twilio Dashboard")

That is all you needed to do for Twilio but don't close it down just yet. You will need to use the `Account SID` & `Auth Token` shortly!

Let's get coding and create the serverless function! - [IBM Cloud Functions](./setup-ibm-cloud-function.md)