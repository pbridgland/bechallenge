package handlers

import (
	"bechallenge/interfaces"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

var nextActionsPattern = regexp.MustCompile(`^/actions/(.+)/nextactions$`)

type nextActions struct {
	processingService interfaces.ProcessingService
}

// NewNextActionsHandler creates a new instance of the Next Actions Handler
func NewNextActionsHandler(processingService interfaces.ProcessingService) nextActions {
	return nextActions{
		processingService: processingService,
	}
}

func (aps nextActions) Handle(w http.ResponseWriter, r *http.Request) {
	matches := nextActionsPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		http.Error(w, "Invalid action type", http.StatusBadRequest)
		return
	}

	actionType := matches[1] // Extract the first captured group (the action type)

	probabilities, err := aps.processingService.NextActions(actionType)
	if err != nil {
		errorMsg := fmt.Sprintf("error getting next actions for action type: %s", actionType)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(probabilities)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
