package handlers

import (
	"bechallenge/interfaces"
	"encoding/json"
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
		http.Error(w, "error getting referralindexes", http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(referralIndexes)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
