package cfg

import (
	"github.com/spf13/viper"
	"github.com/typomedia/gitti/app"
	"github.com/typomedia/gitti/app/msg"
	"github.com/typomedia/gitti/app/str"
)

func Write() {
	viper.Set("app.name", app.App.Name)
	viper.Set("app.version", app.App.Version)
	viper.Set("auth.enabled", false)
	viper.Set("auth.token", str.Hex())
	viper.Set("repos", map[string]string{
		"example": "/path/to/example",
	})
	err := viper.WriteConfigAs(app.App.Name + ".toml")
	msg.Check(err)
}
