package main

import (
	"encoding/json"
	"github.com/labi-le/hyprland-ipc-client/v3"
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

func ReadFirstLayout(ipc client.IPC, evDispatcher client.EventHandler) {
	devices, err := ipc.Devices()
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices.Keyboards {
		if device.Main {
			evDispatcher.ActiveLayout(client.ActiveLayout{
				Type: "keyboard",
				Name: device.ActiveKeymap,
			})
			break
		}
	}
}

func main() {
	var (
		ipc          = client.MustClient(os.Getenv("HYPRLAND_INSTANCE_SIGNATURE"))
		evDispatcher = &ed{}
	)

	ReadFirstLayout(ipc, evDispatcher)

	if err := client.Subscribe(ipc, evDispatcher, client.EventActiveLayout); err != nil {
		log.Fatal(err)
	}
}

func (e *ed) ActiveLayout(layout client.ActiveLayout) {
	ToStdOut(Output{
		Text:    layout.Name,
		Tooltip: "Current keyboard layout",
		Class:   "keyboard-layout",
	})
}
