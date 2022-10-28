package main

import (
	"encoding/json"
	"github.com/labi-le/hyprland-ipc-client"
	"log"
	"os"
)

type Output struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Class   string `json:"class"`
}

func ToStdOut(s any) {
	if err := json.NewEncoder(os.Stdout).Encode(s); err != nil {
		log.Fatal(err)
	}
}

type ed struct {
	client.DummyEvHandler
}

func main() {
	c := client.NewClient(os.Getenv("HYPRLAND_INSTANCE_SIGNATURE"))
	e := &ed{}
	client.Subscribe(c, e, client.EventActiveLayout)
}

func (e *ed) ActiveLayout(layout client.ActiveLayout) {
	ToStdOut(Output{
		Text:    layout.Name,
		Tooltip: "Current keyboard layout",
		Class:   "keyboard-layout",
	})
}
