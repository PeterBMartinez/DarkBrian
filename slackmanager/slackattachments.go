package slackmanager

import (
  "github.com/PeterBMartinez/DarkBrian/chatgptmanager"
  "github.com/slack-go/slack"
)

func chatGPTResponseAttachment(text string) slack.Attachment {
  return slack.Attachment{
    Pretext: chatgptmanager.GetChatAnswer(text),
  }
}
