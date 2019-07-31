package scaleft

import (
	"encoding/json"
	"gopkg.in/resty.v1"
	"log"
)

type Config struct {
	key    string
	secret string
	team   string
}

const url string = "https://app.scaleft.com/v1"

var teamName = ""

type Bearer struct {
	BearerToken string `json:"bearer_token"`
}

// returns interface that contains Bearer token.
func (c *Config) Authorization() (interface{}, error) {

	// get bearer token
	resp := GetToken(c.team, c.key, c.secret)

	teamName = c.team

	token := Bearer{}
	jsonErr := json.Unmarshal(resp, &token)
	if jsonErr != nil {
		log.Printf("[DEBUG] Error getting bearer token:%s", jsonErr)
	}

	return token, nil
}

func GetToken(team, key, secret string) []byte {
	log.Printf("[DEBUG] Getting bearer token from Config.")

	composedUrl := url + "/teams/" + team + "/service_token"

	credentials := map[string]interface{}{"key_id": key, "key_secret": secret}

	resp, _ := resty.R().
		SetBody(credentials).
		SetHeaders(map[string]string{
			"Accept":       "application/json",
			"Content-Type": "Application/json"}).
		Post(composedUrl)

	return resp.Body()

}
