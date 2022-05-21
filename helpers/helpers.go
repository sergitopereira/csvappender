package helpers

import (
	"log"
)

func TerminalHelper(args []string) {
	if len(args) > 3 {
		log.Fatal("Program requires TEST and OUPUT Folders")
	} else if len(args) <= 2 {
		log.Fatal("Program requires TEST and OUPUT  Folders")
	} else {
		return
	}
}
