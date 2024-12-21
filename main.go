package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/mongstaen/avregex-go/messagehandlers/mvthandler"
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
		mvthandler.NewProcessMVTMessage(text)
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
	dep := `
MVT
SD200/21.PMDFG.CDG
AD1100/1115 EA1500 FRA
DL72/0015
PX112
SI DEICING
`
	arr := `
MVT
SD200/21.PMDFG.CDG
AA1515/1520
FLD22
`
	text := `
MVT
AA1203/07.LNREG.BGO
AD1402/1413 EA1524 SVG
PX123
SI MAP1405
`
	mvt := `
MVT
AZ1074/08.HBXXX.ZRH
AD1454/1458 EA1555 FRA
DL46/02/0005/0004
PX023
DLA/02B//
`
	/* 	ldm := `
	LDM
	AZ464/08.HBXXX.C10Y92.2/3
	-LCY.28/14/0/0.0.T262.1/262.PAX/1/41.JMP/0.CRW/0.PAD/0/0
	SI LCY BAG 262 POS 0 FRE 0
	PAX WEIGHTS USED M88 F70 C35 I0
	SERVICE WEIGHT ADJUSTMENT WEIGHT/INDEX
	ADD
	NIL
	DEDUCTIONS
	NIL
	LCY C 0 M 0 B 18/ 262 O 0 T 0
	` */
	//fmt.Println(text)
	processNewMessage(dep)
	processNewMessage(arr)
	processNewMessage(text)
	processNewMessage(mvt)
	//processNewMessage(ldm)
}
