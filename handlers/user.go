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

var userIDPattern = regexp.MustCompile(`^/users/(\d+)$`)

type user struct {
	processingService interfaces.ProcessingService
}

func NewUserHandler(processingService interfaces.ProcessingService) user {
	return user{
		processingService: processingService,
	}
}

func (u user) Handle(w http.ResponseWriter, r *http.Request) {
	matches := userIDPattern.FindStringSubmatch(r.URL.Path)
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

	user, err := u.processingService.User(id)
	if err != nil && errors.Is(err, types.ErrUserNotPresent) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		errorMsg := fmt.Sprintf("error getting user with ID: %d", id)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
