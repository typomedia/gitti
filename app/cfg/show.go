package cfg

import (
	"github.com/spf13/viper"
	"github.com/typomedia/gitti/app/msg"
	"os"
)

func Show() (string, string) {
	Read()
	file := viper.ConfigFileUsed()
	content, err := os.ReadFile(file)
	msg.Check(err)

	return file, string(content)
}
