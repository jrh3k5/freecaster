package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jrh3k5/freecaster/config"
	"github.com/jrh3k5/freecaster/logging"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := logging.GetSlogLogger(ctx)

	freeStuffConfig, err := config.GetFreeStuffConfig(ctx)
	if err != nil {
		logger.Error("Failed to resolve FreeStuff config", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		_ = r.Body.Close()
	}()

	var requestBody *freeStuffWebhookRequest
	if decodeErr := json.NewDecoder(r.Body).Decode(requestBody); decodeErr != nil {
		logger.Debug("Failed to decode request body", "err", decodeErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if requestBody.Secret != freeStuffConfig.WebhookSecret {
		logger.Debug("Webhook secret is not correct")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gameIDs := make([]uint64, 0, len(requestBody.Data))
	for _, datum := range requestBody.Data {
		parsedGameID, parseErr := strconv.ParseUint(datum, 10, 64)
		if parseErr != nil {
			logger.Debug("Failed to parse game ID from data: "+datum, "err", parseErr)
			continue
		}

		gameIDs = append(gameIDs, parsedGameID)
	}

	// TODO: do something with the game IDs

	if _, writeErr := w.Write([]byte(`{ "status": "ok" }`)); writeErr != nil {
		logger.Error("failed to write successful response body", "err", writeErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type freeStuffWebhookRequest struct {
	Event  string   `json:"event"`
	Secret string   `json:"secret"`
	Data   []string `json:"data"`
}
