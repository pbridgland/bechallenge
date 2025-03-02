package handlers

import (
	"bechallenge/interfaces"
	"encoding/json"
	"fmt"
	"net/http"
)

type referalIndexes struct {
	referralService interfaces.ReferralService
}

// NewReferralIndexesHandler creates a new instance of the Referral Indexes Handler
func NewReferralIndexesHandler(referralService interfaces.ReferralService) referalIndexes {
	return referalIndexes{
		referralService: referralService,
	}
}

func (ri referalIndexes) Handle(w http.ResponseWriter, r *http.Request) {
	referralIndexes, err := ri.referralService.ReferralIndexes()
	if err != nil {
		errorMsg := fmt.Sprintf("error getting referralindexes")
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(referralIndexes)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
