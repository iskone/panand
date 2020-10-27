package main

import (
	"encoding/json"
	"github.com/iskone/panand/lib"
	"github.com/kataras/iris/v12"
	"os"
)

type AppInfo struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Key         string `json:"key"`
	RedirectUri string `json:"redirect_uri"`
}

type Config struct {
	App  AppInfo `json:"app"`
	Home string
	User *lib.AccessToken `json:"user"`
}
type Server struct {
	PanAnd  *lib.Panand
	isLogin bool
	app     *iris.Application
}

func LoadConfig() (*Config, error) {
	f, e := os.Open("./config.json")
	if e != nil {
		return nil, e
	}
	defer f.Close()
	var c Config
	e = json.NewDecoder(f).Decode(&c)
	if e != nil {
		return nil, e
	}
	return &c, nil
}

func SaveConfig(c Config) error {
	f, e := os.OpenFile("./config.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if e != nil {
		return e
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(c)
}
func main() {
	c, e := LoadConfig()
	if e != nil {
		panic(e)
	}

	authd := lib.NewClientOauth(c.App.Name, c.App.Id, c.App.Key, c.App.RedirectUri)
	var s = &Server{
		PanAnd: &lib.Panand{
			authd,
			c.User,
		},
		isLogin: false,
		app:     nil,
	}
	newHttp(s)
	_ = SaveConfig(*c)
}
