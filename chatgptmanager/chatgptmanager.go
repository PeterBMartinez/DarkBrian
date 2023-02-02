package chatgptmanager

import (
  "os"
  "github.com/solywsh/chatgpt"
  "fmt"
  "time"
)

func GetChatAnswer(prompt string) string {
  chatGPTToken := os.Getenv("CHAT_GPT_TOKEN")
  chat := chatgpt.New(chatGPTToken, "user_id(not required)", 30*time.Second)
	defer chat.Close()

	answer, err := chat.Chat(prompt)
	if err != nil {
		fmt.Println(err)
	}
	return answer
}
