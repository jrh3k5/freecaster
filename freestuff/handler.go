package freestuff

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/jrh3k5/freecaster/config"
	v1 "github.com/jrh3k5/freestuff-api-go/pkg/client/v1"
	freestuffhttp "github.com/jrh3k5/freestuff-api-go/pkg/client/v1/http"
)

var (
	handlerOnce     sync.Once
	handlerInitErr  error
	handlerInstance *Handler
)

// Handler is a handler for FreeStuff webhook requests.
type Handler struct {
	apiClient v1.Client
}

// GetHandler gets a handler for FreeStuff webhook requests.
func GetHandler(ctx context.Context) (*Handler, error) {
	handlerOnce.Do(func() {
		freeStuffConfig, err := config.GetFreeStuffConfig(ctx)
		if err != nil {
			handlerInitErr = fmt.Errorf("failed to initialize FreeStuff configuration: %w", err)
			return
		}

		apiClient := freestuffhttp.NewHTTPClient(freeStuffConfig.APIKey, http.DefaultClient)
		handlerInstance = NewHandler(apiClient)
	})

	if handlerInitErr != nil {
		return nil, fmt.Errorf("failed to initialize handler: %w", handlerInitErr)
	}

	return handlerInstance, nil
}

func NewHandler(apiClient v1.Client) *Handler {
	return &Handler{
		apiClient: apiClient,
	}
}

// HandleFreeGames handles the announcement of free games.
func (h *Handler) HandleFreeGames(ctx context.Context, gameIDs []int64) error {
	_, err := h.apiClient.GetGameInfo(ctx, gameIDs)
	if err != nil {
		gameIDStrings := make([]string, len(gameIDs))
		for idIndex, gameID := range gameIDs {
			gameIDStrings[idIndex] = strconv.FormatInt(gameID, 10)
		}
		return fmt.Errorf("failed to get game infos for game IDs: [%s]: %w", strings.Join(gameIDStrings, ", "), err)
	}
	return nil
}
