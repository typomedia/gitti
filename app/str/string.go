package str

import (
	"bytes"
	"strings"
)

func Trim(byte []byte) string {
	// remove empty lines from byte
	byte = bytes.Trim(byte, "\x00")
	byte = bytes.Trim(byte, "\n")
	byte = bytes.Trim(byte, "\r")

	return string(byte)
}

func Split(str string, sep string) []string {
	return strings.Split(str, sep)
}

func After(str string, sep string) string {
	return Split(str, sep)[1]
}

func Before(str string, sep string) string {
	return Split(str, sep)[0]
}

func Starts(str string, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}
