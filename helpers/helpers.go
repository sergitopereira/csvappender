package helpers

import (
	"log"
)

func TerminalHelper(args []string) {
	if len(args) > 3 {
		log.Fatal("Program requires TEST DATA and OUTPUT fodler path")
	} else if len(args) <= 2 {
		log.Fatal("Program requires TEST DATA and OUTPUT folders path")
	} else {
		return
	}
}
