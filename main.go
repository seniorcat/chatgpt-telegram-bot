package main

import (
	"context"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	gpt "github.com/sashabaranov/go-openai"
)

func main() {
	// Set up the Telegram bot API client
	bot, err := tgbotapi.NewBotAPI("enter bot api key here")
	if err != nil {
		log.Panic(err)
	}

	// Set up the OpenAI GPT-3 API client
	client := gpt.NewClient("++++++++++++++++++++++++++++++++++++")

	// Set up a ticker to run the bot's main loop once an hour
	ticker := time.NewTicker(time.Hour * 3)
	defer ticker.Stop()

	for range ticker.C {
		// Generate a joke about ChatGPT using the OpenAI GPT-3 API
		resp, err := client.CreateChatCompletion(
			context.Background(),
			gpt.ChatCompletionRequest{
				Model: gpt.GPT3Dot5Turbo,
				Messages: []gpt.ChatCompletionMessage{
					{
						Role:    gpt.ChatMessageRoleUser,
						Content: "Придумай одну шутку про сферу IT минимум на 50 слов",
					},
				},
			},
		)
		if err != nil {
			log.Panic(err)
		}

		// Create a message to post to the Telegram channel
		msg := tgbotapi.NewMessage(-11111111111111111, resp.Choices[0].Message.Content)

		// Send the message using the Telegram bot API
		_, err = bot.Send(msg)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Joke posted to Telegram channel")
		}
	}
}
