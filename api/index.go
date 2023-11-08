package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jrh3k5/freecaster/config"
	"github.com/jrh3k5/freecaster/logging"
	"go.uber.org/zap"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger, err := logging.GetZapLogger(ctx)
	if err != nil {
		log.Printf("Failed to initialize logger: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	freeStuffConfig, err := config.GetFreeStuffConfig(ctx)
	if err != nil {
		logger.Error("Failed to resolve FreeStuff config", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		_ = r.Body.Close()
	}()

	requestBody := &freeStuffWebhookRequest{}
	if decodeErr := json.NewDecoder(r.Body).Decode(requestBody); decodeErr != nil {
		logger.Debug("Failed to decode request body", zap.Error(decodeErr))
		w.WriteHeader(http.StatusBadRequest)
		if _, writeErr := w.Write([]byte(`{ "status": "notok", "error": "invalid JSON request body" }`)); writeErr != nil {
			logger.Error("unable to write bad JSON response", zap.Error(writeErr))
		}
		return
	}

	if requestBody.Secret != freeStuffConfig.WebhookSecret {
		logger.Debug("Webhook secret is not correct")
		w.WriteHeader(http.StatusBadRequest)
		if _, writeErr := w.Write([]byte(`{ "status": "notok", "error": "invalid webhook secret" }`)); writeErr != nil {
			logger.Error("unable to write bad webhook secret response", zap.Error(writeErr))
		}
		return
	}

	gameIDs := make([]uint64, 0, len(requestBody.Data))
	for _, datum := range requestBody.Data {
		datumString := datum.String()
		parsedGameID, parseErr := strconv.ParseUint(datumString, 10, 64)
		if parseErr != nil {
			logger.Debug("Failed to parse game ID from data: "+datumString, zap.Error(parseErr))
			continue
		}

		gameIDs = append(gameIDs, parsedGameID)
	}

	// TODO: do something with the game IDs

	if _, writeErr := w.Write([]byte(`{ "status": "ok" }`)); writeErr != nil {
		logger.Error("failed to write successful response body", zap.Error(writeErr))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type freeStuffWebhookRequest struct {
	Event  string        `json:"event"`
	Secret string        `json:"secret"`
	Data   []json.Number `json:"data"`
}
