package sys

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wailsapp/wails"

	"github.com/peterhellberg/acr122u"
)

// Card .
type Card struct {
	log *wails.CustomLogger
}

// RFID .
type RFID struct {
	ID string `json:"id"`
}

// WailsInit .
func (c *Card) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("Card")

	go func() {

		ctx, err := acr122u.EstablishContext()
		if err != nil {
			fmt.Println("device acr122u not connected ! ")
			runtime.Events.Emit("rfid", c.GetNotifyCard())

		} else {
			fmt.Println("device acr122u connected ! ", ctx.Readers())
			ctx.ServeFunc(func(card acr122u.Card) {
				runtime.Events.Emit("rfid", c.GetRFID(card))
				// runtime.Events.Emit("rfidid", c.GetRFID(c))
			})
		}
	}()

	return nil
}

// GetNotifyCard
func (c *Card) GetNotifyCard() *RFID {

	return &RFID{
		ID: "device acr122u not connected !",
	}
}

// GetRFID .
func (c *Card) GetRFID(card acr122u.Card) *RFID {

	return &RFID{
		ID: convertUID(card.UID()),
	}
}

func convertUID(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, "")
}
