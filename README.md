# DarkBrian
<kbd> <img width="685" style="border-radius: 0%;" alt="Screenshot 2023-01-31 at 8 34 51 PM" src="https://user-images.githubusercontent.com/42549333/216329444-b935e37d-c991-4275-9866-1f1438c16caf.png"></kbd> 

## Local Environment Setup
Verify if you have go lang installed on your computer by running `go version` in your terminal, you should see something along the lines of `go version go1.19.5 darwin/amd64`.

If not installed for MacOs you can run the following `brew install go`

Once complete re run `go version` to verify

Pull down the repo by using the following command `git clone git@github.com:PeterBMartinez/DarkBrian.git`

Create a .env file in your main directory(`touch .env`) and place the following environment variables in the .env file
```
SLACK_AUTH_TOKEN="{{ Slack Authentication Token }}"
SLACK_CHANNEL_ID="{{ Channel Name }}"
SLACK_APP_TOKEN="{{ Slack Application Token }}"
CHAT_GPT_TOKEN="{{ ChatGPT Token }}"
```
If you don't have these tokens, here are some helpful links on how to retrieve them!
- Slack tokens: [api.slack.com](https://api.slack.com/tutorials/tracks/getting-a-token)
- Slack channel Id: [socialinterns.com](https://help.socialintents.com/article/148-how-to-find-your-slack-team-id-and-slack-channel-id#:~:text=the%20Team%20ID.-,Open%20any%20web%20browser%20and%20log%20in%20to%20your%20Slack,represents%20your%20Slack%20Channel%20ID.)
- ChatGPT token: [platform.openai.com](https://platform.openai.com/account/api-keys)

Run the following in your terminal or command prompt to install all the package dependencies `go mod tidy`
Then run `go run app/main.go`

## Running DarkBrian
Run the following in your terminal in the DarkBrian top level directory `go run app/main.go`

## Usage 
In the channel that you selected for your bot to run simply mention your bot and provide your chatgpt prompt

### Example: 
<kbd><img width="1261" alt="Screenshot 2023-01-31 at 8 59 37 PM" src="https://user-images.githubusercontent.com/42549333/216330200-0f4d7272-5a18-4bfa-8897-4453615d532a.png"></kdb>

## Todo List:
- Direct slack messaging event catching
- Additional integrations such as Telegram, and Discord which would be optional based on environment variables
- Support for adding your own learning models via a json file


### If you have requests please let me know! 

