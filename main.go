package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	mvthandler "github.com/mongstaen/avregex-go/messagehandlers/mvthandler"
)

func processNewMessage(text string) {
	text = strings.TrimSpace(text)
	messageTypePattern := `^\s*([A-Z]{3})\n`
	messageTypeRegex := regexp.MustCompile(messageTypePattern)
	messageType := strings.TrimSpace(messageTypeRegex.FindString(text))

	switch messageType {
	case "PDM":
		processPDMMessage(text)
	case "MVT":
		mvthandler.ProcessMVTMessage(text)
	case "LDM":
		fmt.Println("LDM")
	case "COR":
		processCORMessage(text)
	default:
		fmt.Println("Unknown Message Type")
	}
}

func processPDMMessage(text string) {
	// TODO add feature to the object "Is Potential Duplicate"
	newText := strings.TrimSpace(strings.ReplaceAll(text, "PDM", ""))
	log.Println("Processed PDM Message")
	processNewMessage(newText)
}

func processCORMessage(text string) {
	// TODO add feature to the object "Is Corrected"
	newText := strings.TrimSpace(strings.ReplaceAll(text, "COR", ""))
	log.Println("Processed COR Message")
	processNewMessage(newText)
}
func main() {
	text := `
COR
MVT
AA1203/07.LNREG.BGO
AD1402/1413 EA1524 SVG
PX123
SI MAP1405
`
	//fmt.Println(text)
	processNewMessage(text)
}
