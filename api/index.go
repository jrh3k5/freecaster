package api

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

}

type freeStuffWebhookRequest struct {
	Event  string   `json:"event"`
	Secret string   `json:"secret"`
	Data   []string `json:"data"`
}
