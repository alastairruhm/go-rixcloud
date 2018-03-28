package rixcloud

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ProfileService provides a method for credential verification.
type ProfileService struct {
	sling *sling.Sling
}

// Profile the user profile
type Profile struct {
	UserID int `json:"userid,omitempty"`
}

// Service the service which user subscribe
type Service struct {
	Name      string `json:"name,omitempty"`
	IP        string `json:"ip,omitempty"`
	Provider  string `json:"provider,omitempty"`
	Bandwidth string `json:"bandwidth,omitempty"`
	Feature   string `json:"feature,omitempty"`
	ISP       string `json:"isp,omitempty"`
	Rate      string `json:"rate,omitempty"`
}

// NewProfileService returns a new AccountService.
func NewProfileService(sling *sling.Sling) *ProfileService {
	return &ProfileService{
		sling: sling.Path("profile/"),
	}
}

// VerifyCredentials ...
func (s *ProfileService) VerifyCredentials() (*Profile, *http.Response, error) {
	profile := new(Profile)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").Receive(profile, apiError)
	return profile, resp, firstError(err, apiError)
}

// func (s *ProfileService) ListServices() (*)
