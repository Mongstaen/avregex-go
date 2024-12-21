package mvthandler

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Message struct {
	messageType          string `json:"messageType"`
	flightNumber         string `json:"flightNumber"`
	flightDate           string `json:"flightDate"`
	aircraftRegistration string `json:"aircraftRegistration"`
	flightOrigin         string `json:"fligthOrigin"`
	mvtType              string `json:"mvtType"`
	pushTime             string `json:"pushTime"`
	takeoffTime          string `json:"takeoffTime"`
	etaTime              string `json:"etaTime"`
	flightDestination    string `json:"flightDestination"`
	messageSi            string `json:"messageSi"`
}

func handleMVTType(mvtType string) {
	switch mvtType {
	case "AD":
		log.Println("Handled AD Movement Type")
	case "ED":
		log.Println("Handled ED Movement Type")
	case "AA":
		log.Println("Handled AA Movement Type")
	case "EA":
		log.Println("Handled EA Movement Type")
	case "EO":
		log.Println("Handled EO Movement Type")
	default:
		log.Println("Unknown Movement Type")
	}

	// TODO make functionality to handle AD, EA, AA messages
}

func findMVTType(text string) string {
	text = strings.TrimSpace(text)
	pattern := `(?m)^(?:.*\n){2}([A-Z]{2})`
	regex := regexp.MustCompile(pattern)
	mvtType := regex.FindStringSubmatch(text)[1]
	return mvtType
}

func newHandleMVTType(mvtType string) string {
	pattern := ""
	switch mvtType {
	case "AD":
		log.Println("Handled AD Movement Type")
		pattern = `(?i)([A-Z]{3})\n([A-Z0-9]{3,})\/([0-9]{2})\.([A-Z0-9]{5,}).([A-Z]{3})\n([A-Z]{2})([0-9]{4,6})\/([0-9]{4,6}) [A-Z]{2}([0-9]{4}) ([A-Z]{3})\n((?s).*)`
	case "ED":
		log.Println("Handled ED Movement Type")
		pattern = "ED"
	case "AA":
		log.Println("Handled AA Movement Type")
		pattern = `(?i)([A-Z]{3})\n([A-Z0-9]{3,})\/([0-9]{2})\.([A-Z0-9]{5,}).([A-Z]{3})\n([A-Z]{2})([0-9]{4,6})\/([0-9]{4,6})\n((?s).*)`
	case "EA":
		log.Println("Handled EA Movement Type")
		pattern = "EA"
	case "EO":
		log.Println("Handled EO Movement Type")
		pattern = "EO"
	default:
		log.Println("Unknown Movement Type")
	}

	return pattern
	// TODO make functionality to handle AD, EA, AA messages
}

func ProcessMVTMessage(text string) {
	text = strings.TrimSpace(text)
	pattern := `(?i)([A-Z]{3})\n([A-Z0-9]{3,})\/([0-9]{2})\.([A-Z0-9]{5,}).([A-Z]{3})\n([A-Z]{2})([0-9]{4,6})\/([0-9]{4,6}) [A-Z]{2}([0-9]{4}) ([A-Z]{3})\n((?s).*)`
	regex := regexp.MustCompile(pattern)
	result := regex.FindStringSubmatch(text)

	message := Message{
		messageType:          result[1],
		flightNumber:         result[2],
		flightDate:           result[3],
		aircraftRegistration: result[4],
		flightOrigin:         result[5],
		mvtType:              result[6],
		pushTime:             result[7],
		takeoffTime:          result[8],
		etaTime:              result[9],
		flightDestination:    result[10],
		messageSi:            result[11],
	}

	handleMVTType(message.mvtType)

	fmt.Println("Message Type: \t\t", message.messageType)
	fmt.Println("Flight Number: \t\t", message.flightNumber)
	fmt.Println("Flight Date: \t\t", message.flightDate)
	fmt.Println("Aircraft Registration: \t", message.aircraftRegistration)
	fmt.Println("Flight Origin: \t\t", message.flightOrigin)
	fmt.Println("Flight Destination: \t", message.flightDestination)
	fmt.Println("MVT Type: \t\t", message.mvtType)
	fmt.Println("Push Time: \t\t", message.pushTime)
	fmt.Println("Takeoff Time: \t\t", message.takeoffTime)
	fmt.Println("ETA Time: \t\t", message.etaTime)
	fmt.Println("SI: \n", message.messageSi)

	// TODO: Make UUID for message, and store entire message in DB
	// TODO: IE: UUID, Unique Flight ID, Message

	// TODO: Once flights DB, find flight and update required fields
}

func NewProcessMVTMessage(text string) {
	mvtType := findMVTType(text)
	pattern := newHandleMVTType(mvtType)
	regex := regexp.MustCompile(pattern)
	result := regex.FindStringSubmatch(text)

	fmt.Println(result[4])
	/* 	message := Message{
	   		messageType:          result[1],
	   		flightNumber:         result[2],
	   		flightDate:           result[3],
	   		aircraftRegistration: result[4],
	   		flightOrigin:         result[5],
	   		mvtType:              result[6],
	   		pushTime:             result[7],
	   		takeoffTime:          result[8],
	   		etaTime:              result[9],
	   		flightDestination:    result[10],
	   		messageSi:            result[11],
	   	}

	   	fmt.Println("Message Type: \t\t", message.messageType)
	   	fmt.Println("Flight Number: \t\t", message.flightNumber)
	   	fmt.Println("Flight Date: \t\t", message.flightDate)
	   	fmt.Println("Aircraft Registration: \t", message.aircraftRegistration)
	   	fmt.Println("Flight Origin: \t\t", message.flightOrigin)
	   	fmt.Println("Flight Destination: \t", message.flightDestination)
	   	fmt.Println("MVT Type: \t\t", message.mvtType)
	   	fmt.Println("Push Time: \t\t", message.pushTime)
	   	fmt.Println("Takeoff Time: \t\t", message.takeoffTime)
	   	fmt.Println("ETA Time: \t\t", message.etaTime)
	   	fmt.Println("SI: \n", message.messageSi) */

	// TODO: Make UUID for message, and store entire message in DB
	// TODO: IE: UUID, Unique Flight ID, Message

	// TODO: Once flights DB, find flight and update required fields
}
