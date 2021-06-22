package civ

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	discordChan = "741356800988807218"
)

var (
	discordSess *discordgo.Session
	playerMap   = map[string]string{
		"sleazegull":          "222905440302923776",
		"cereal killer":       "195001919528370176",
		"i love my large son": "285306283525931009",
		"ragingpandafury":     "96662773823455232",
		"hard luck hero":      "346374307443638276",
		"shogunaut":           "438447596705677332",
		"congee":              "691145093721358377",
		"natschley":           "663587852470190111",
	}
)

// JSON ...
type JSON struct {
	GameName string `json:"value1"`
	UserName string `json:"value2"`
	Turn     string `json:"value3"`
}

// ListenAndServe ...
func ListenAndServe(sess *discordgo.Session) error {
	discordSess = sess

	r := mux.NewRouter()
	r.HandleFunc("/civ", WebhookHandler)

	s := http.Server{
		Addr:    ":80",
		Handler: r,
	}

	return s.ListenAndServe()
}

// WebhookHandler ...
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("error reading civ webhook body: %v", err)
		return
	}

	fmt.Println("webhook incoming")
	fmt.Println(string(b))

	var j JSON
	err = json.Unmarshal(b, &j)
	if err != nil {
		log.Errorf("error unmarhsalling civ JSON: %v", err)
		return
	}

	user := strings.ToLower(playerMap[j.UserName])
	msg := fmt.Sprintf("It's your turn <@%s>", user)
	if user == "" {
		user = j.UserName
		msg = fmt.Sprintf("It's your turn %s", user)
	}

	_, err = discordSess.ChannelMessageSend(discordChan, msg)
	if err != nil {
		log.Errorf("error sending to discord: %v", err)
	}
}
