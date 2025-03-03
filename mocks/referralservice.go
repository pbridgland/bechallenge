package mocks

type ReferralService struct {
	NextReferralIndexesResult map[int]int
	NextReferralIndexesErr    error
}

// ReferralIndexes implements interfaces.ReferralService.
func (r *ReferralService) ReferralIndexes() (map[int]int, error) {
	return r.NextReferralIndexesResult, r.NextReferralIndexesErr
}
