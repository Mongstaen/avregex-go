package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
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
		processMVTMessage(text)
	case "LDM":
		fmt.Println("LDM")
	case "COR":
		processCORMessage(text)
	default:
		fmt.Println("Unknown Message Type")
	}
}

func processMVTMessage(text string) {
	text = strings.TrimSpace(text)
	lines := strings.Split(text, "\n")
	fmt.Println(text)

	// Handle line 3
	parts := strings.Split(lines[2], " ")

	for _, part := range parts {
		//fmt.Println(part)
		if strings.HasPrefix(part, "ED") {
			fmt.Println("Part starts with ED:", part)
		}
		if strings.HasPrefix(part, "AD") {
			fmt.Println("Part starts with AD:", part)
		}
		if strings.HasPrefix(part, "EA") {
			fmt.Println("Part starts with EA:", part)
		}
		if strings.HasPrefix(part, "AA") {
			fmt.Println("Part starts with AA:", part)
		}
		if strings.HasPrefix(part, "EO") {
			fmt.Println("Part starts with EO:", part)
		}
		if strings.HasPrefix(part, "NI") {
			fmt.Println("Part starts with NI:", part)
		}
		if len(part) == 3 {
			fmt.Println("Flight Destination:", part)
		}
	}

	for _, line := range lines[3:] {
		if strings.HasPrefix(line, "PX") {
			numberOfPassengers := strings.TrimPrefix(line, "PX")
			fmt.Println("PX is present:", numberOfPassengers)
		}
	}
	if strings.HasPrefix(lines[3], "DL") {
		fmt.Println("DL is present")
		// TODO Handle DL..
	}

	fmt.Println("\n")

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
PDM
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
	dly := `
MVT
SD200/22.PMDFG.CDG
ED221125
DL72/0025
`
	ni := `
MVT
SD200/22.PMDFG.CDG
NI221150
SI ENGINE TROUBLE
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
	processNewMessage(dly)
	processNewMessage(ni)
	processNewMessage(text)
	processNewMessage(mvt)
	//processNewMessage(ldm)
}
