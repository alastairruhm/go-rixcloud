package rixcloud

import (
	"fmt"
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

// ActivedService ...
type ActivedService struct {
	ServiceID    int    `json:"serviceid,omitempty"`
	Name         string `json:"name,omitempty"`
	GroupName    string `json:"groupname,omitempty"`
	GroupID      int    `json:"groupid,omitempty"`
	NextDueDate  string `json:"nextduedate,omitempty"`
	BillingCycle string `json:"billingcycle,omitempty"`
}

// DeactivedServie ...
type DeactivedServie struct {
	ServiceID int    `json:"serviceid,omitempty"`
	Status    string `json:"status,omitempty"`
}

// ServiceAPIResponse ...
type ServiceAPIResponse struct {
	Data *ServiceOverview `json:"data,omitempty"`
}

// ServiceOverview ...
type ServiceOverview struct {
	ActiveCount  int               `json:"activecount,omitempty"`
	TotalService int               `json:"totalserveice,omitempty"`
	Actived      []ActivedService  `json:"actived,omitempty"`
	Deactived    []DeactivedServie `json:"deactived,omitempty"`
}

// ServiceNode the node which service provide
type ServiceNode struct {
	Name      string `json:"name,omitempty"`
	IP        string `json:"ip,omitempty"`
	Provider  string `json:"provider,omitempty"`
	Bandwidth string `json:"bandwidth,omitempty"`
	Feature   string `json:"feature,omitempty"`
	ISP       string `json:"isp,omitempty"`
	Rate      string `json:"rate,omitempty"`
}

// TrafficAPIResponse ...
type TrafficAPIResponse struct {
	Data *Traffic `json:"data,omitempty"`
}

// Traffic data
type Traffic struct {
	Upload   string `json:"upload,omitempty"`
	Download string `json:"download,omitempty"`
	Total    string `json:"total,omitempty"`
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

// GetServiceOverview ...
func (s *ProfileService) GetServiceOverview() (*ServiceOverview, *http.Response, error) {
	res := new(ServiceAPIResponse)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("service/").Receive(res, apiError)
	return res.Data, resp, firstError(err, apiError)
}

// GetTrafficData ...
func (s *ProfileService) GetTrafficData(serviceid string) (*Traffic, *http.Response, error) {
	url := fmt.Sprintf("service/%s/traffic", serviceid)
	res := new(TrafficAPIResponse)
	apiError := new(APIError)

	resp, err := s.sling.New().Get(url).Receive(res, apiError)
	return res.Data, resp, firstError(err, apiError)
}
