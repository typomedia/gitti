package cfg

import (
	"github.com/spf13/viper"
	"github.com/typomedia/gitti/app"
	"github.com/typomedia/gitti/app/msg"
	"os"
)

func Read() {
	viper.SetConfigName(app.App.Name)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	msg.Check(err)

	// if not found, create a new cfg file
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		Write()
		Read()
		msg.Info("New Config File: " + viper.ConfigFileUsed())
		msg.Info("Please edit the Config File and restart the Server.")
		os.Exit(0)
	}
}
