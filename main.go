
package main

import (
  "fmt"
  "strings"
  "errors"
  "context"
  "log"
  "os"
  "time"
  "github.com/solywsh/chatgpt"

  "github.com/joho/godotenv"
  "github.com/slack-go/slack"
  "github.com/slack-go/slack/slackevents"
  "github.com/slack-go/slack/socketmode"
)

func main() {
  godotenv.Load(".env")

  token := os.Getenv("SLACK_AUTH_TOKEN")
  appToken := os.Getenv("SLACK_APP_TOKEN")


  client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))


  socketClient := socketmode.New(
    client,
    socketmode.OptionDebug(true),
    socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
  )

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
    for {
      select {
      case <-ctx.Done():
        log.Println("Shutting down socketmode listener")
        return
      case event := <-socketClient.Events:
        switch event.Type {
        case socketmode.EventTypeEventsAPI:
          eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
          if !ok {
            log.Printf("Could not type cast the event to the EventsAPIEvent: %v\n", event)
            continue
          }

          socketClient.Ack(*event.Request)
          err := handleEventMessage(eventsAPIEvent, client)
          if err != nil {
            log.Fatal(err)
          }
        }

      }
    }
  }(ctx, client, socketClient)

  socketClient.Run()
}

func handleEventMessage(event slackevents.EventsAPIEvent, client *slack.Client) error {
	switch event.Type {
	case slackevents.CallbackEvent:

		innerEvent := event.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := handleAppMentionEvent(ev, client)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("unsupported event type")
	}
	return nil
}

func handleAppMentionEvent(event *slackevents.AppMentionEvent, client *slack.Client) error {

	user, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}

	text := strings.ToLower(event.Text)

	attachment := slack.Attachment{}

  fmt.Println("DarkBrian was asked a question by, ", user.Name)
  attachment.Pretext = getChatAnswer(text)

	_, _, err = client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}

//func getPrompt(prompt string) (result string) {
//	re := regexp.MustCompile(`\[(.*?)\]`)
//	submatchall := re.FindAllString(prompt, -1)
//	for _, element := range submatchall {
//		element = strings.Trim(element, "[")
//		element = strings.Trim(element, "]")
//    return element
//	}
//  return ""
//}

func getChatAnswer(prompt string) string {
  chatGPTToken := os.Getenv("CHAT_GPT_TOKEN")
  chat := chatgpt.New(chatGPTToken, "user_id(not required)", 30*time.Second)
	defer chat.Close()

	answer, err := chat.Chat(prompt)
	if err != nil {
		fmt.Println(err)
	}
	return answer
}
