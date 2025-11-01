# setup basic cli (x)
> enable user to choose csv to consume
lookForSecrets -I inventory.csv

# setup csv interfacing (x)
> create an example csv interfacing structs
> iteratively parse fields through csv token field
> secrets = (type,provider,creator,value)

# integrate an api to get simulated code (x)

> load the recursively generated and loaded codebase{we assume only code,config,etc exist within the codebase}
> for each iteration of the token field {check for the codebase for match of the pattern(use fuzzy?(has confidence))}

## documentation ()

## alerting mechanism through mail ()
> generate gmail creds (pat)
> generic way to build smtp req
> https://www.youtube.com/watch?v=H0HZc4FgX7E

## alerting mechanism through Slack () {similar to discord's setup}
> generate creds 
> use below req to send data
` 
POST https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
    Content-type: application/json
{
    "text": "{alert text here}"
}
`
> https://docs.slack.dev/messaging/sending-messages-using-incoming-webhooks/

### add location details through metadata ()
