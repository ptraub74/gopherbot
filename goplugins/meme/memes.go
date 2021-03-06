package memes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/parsley42/gopherbot/bot"
)

var (
	gobot   bot.Robot
	botName string
)

type MemeConfig struct {
	Username string
	Password string
}

var config MemeConfig
var configured bool = false

func memegen(r bot.Robot, command string, args ...string) {
	switch command {
	case "init":
		gobot = r
		botName = r.User
		err := r.GetPluginConfig(&config)
		if err == nil {
			configured = true
		}
	case "simply":
		sendMeme(r, "61579", "ONE DOES NOT SIMPLY", args[0])

	case "prepare":
		sendMeme(r, "47779539", "You "+args[0], "PREPARE TO DIE")

	case "prettymuch":
		sendMeme(r, "8070362", args[0]+" pretty much", "the "+args[1]+" ever "+args[2])

	case "gosh":
		sendMeme(r, "18304105", args[0], "Gosh!")

	case "skills":
		sendMeme(r, "20509936", args[0]+" "+args[1], args[2])

	}
}

func sendMeme(r bot.Robot, templateId, topText, bottomText string) {
	url, err := createMeme(templateId, topText, bottomText)
	if err == nil {
		r.Say(url)
	} else {
		r.Reply("Sorry, something went wrong. Check the logs?")
		r.Log(bot.Error, fmt.Errorf("Generating a meme: %v", err))
	}
}

// Compose imgflip meme - thanks to Adam Georgeson for this function
func createMeme(templateId, topText, bottomText string) (string, error) {
	values := url.Values{}
	values.Set("template_id", templateId)
	values.Set("username", config.Username)
	values.Set("password", config.Password)
	values.Set("text0", topText)
	values.Set("text1", bottomText)
	resp, err := http.PostForm("https://api.imgflip.com/caption_image", values)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if !data["success"].(bool) {
		return "", errors.New(data["error_message"].(string))
	}

	url := data["data"].(map[string]interface{})["url"].(string)

	return url, nil
}

func init() {
	bot.RegisterPlugin("memes", memegen)
}
