package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  // Print "Hello, World!" to the console
  fmt.Println("Hello, World!")

  // Create a telegram service. Ignoring error for demo simplicity.
  botApi := os.Getenv("TG_BOT_API")
telegramService, _ := telegram.New(botApi)

// Passing a telegram chat id as receiver for our messages.
// Basically where should our message be sent?
chatId := os.Getenv("TG_CHAT_ID")
// convert chatId to int64
chatIdInt, _ := strconv.ParseInt(chatId, 10, 64)
// Set the chat id as receiver for our messages.
fmt.Println("Adding receivers")
fmt.Println(chatIdInt)
telegramService.AddReceivers(chatIdInt)
// Tell our notifier to use the telegram service. You can repeat the above process
// for as many services as you like and just tell the notifier to use them.
// Inspired by http middlewares used in higher level libraries.
notify.UseServices(telegramService)

// Send a test message.
_ = notify.Send(
	context.Background(),
	"Hello, sandy bunda :)",
	"KIIII",
)
}