package handlers

import (
	"bechallenge/interfaces"
	"bechallenge/types"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

var userActionCountPattern = regexp.MustCompile(`^/users/(\d+)/actions/count$`)

type userActionCount struct {
	processingService interfaces.ProcessingService
}

// NewUserActionCountHandler creates a new instance of the UserActionCountHandler
func NewUserActionCountHandler(processingService interfaces.ProcessingService) userActionCount {
	return userActionCount{
		processingService: processingService,
	}
}

type userActionCountResp struct {
	Count int `json:"count"`
}

func (uac userActionCount) Handle(w http.ResponseWriter, r *http.Request) {
	matches := userActionCountPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	idString := matches[1] // Extract the first captured group (the ID)
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	count, err := uac.processingService.UserActionCount(id)
	if err != nil && errors.Is(err, types.ErrUserNotPresent) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		errorMsg := fmt.Sprintf("error getting action count for user with ID: %d", id)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	userActionCount := userActionCountResp{
		Count: count,
	}

	resp, _ := json.Marshal(userActionCount)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
