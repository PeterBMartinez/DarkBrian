package main

import (
  "github.com/joho/godotenv"
  "github.com/PeterBMartinez/DarkBrian/slackmanager"
)

func main() {
  godotenv.Load(".env")
  slackmanager.Start()
}

