package slackmanager

import (
  "fmt"
  "strings"
  "errors"
  "context"
  "log"
  "os"
  "github.com/slack-go/slack"
  "github.com/slack-go/slack/slackevents"
  "github.com/slack-go/slack/socketmode"
)


func Start() {
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

  go sendMessageBasedOnEvent(ctx, client, socketClient)

  socketClient.Run()
}

func sendMessageBasedOnEvent(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
  for {
    select {
    case <-ctx.Done():
      return
    case event := <-socketClient.Events:
      switch event.Type {
      case socketmode.EventTypeEventsAPI:
        eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
        if !ok {
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

	attachment := chatGPTResponseAttachment(text)
  fmt.Println("User %s has messaged", user.Name)
	_, _, err = client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}
