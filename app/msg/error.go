package msg

import (
	"log"
)

func Check(err error) {

	if err == nil {
		return
	}

	log.Printf("%s\n", err)
}
