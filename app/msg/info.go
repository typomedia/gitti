package msg

import (
	"fmt"
	"log"
)

func Info(format string, args ...interface{}) {
	log.Printf("%s\n", fmt.Sprintf(format, args...))
}
