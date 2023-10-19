package main

import (
	"encoding/json"
	"github.com/labi-le/hyprland-ipc-client"
	"log"
	"os"
	"os/exec"
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

func ReadFirstLayout() {
	devices := map[string][]map[string]any{}
	out, err := exec.Command("hyprctl", "devices", "-j").Output()
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(out, &devices); err != nil {
		log.Fatal(err)
	}

	for _, device := range devices["keyboards"] {
		if device["main"] == true {
			text, ok := device["active_keymap"].(string)
			if !ok {
				log.Fatal("Could not read active keymap")
			} else {
				ToStdOut(Output{ 
					Text:    text,
					Tooltip: "Current keyboard layout",
					Class:   "keyboard-layout",
				})
			}

			
		}
	}
}

func main() {
	ReadFirstLayout()
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
