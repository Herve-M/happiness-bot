package cmd

import (
	"fmt"

  "github.com/Herve-M/happiness-bot/core"
)

func TestCmd(flags *core.SlackTestFlags) {
  fmt.Printf("TestCmd with %v", flags)
}
