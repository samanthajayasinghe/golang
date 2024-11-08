package pagerduty

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	pdApi "github.com/PagerDuty/go-pagerduty"
)

// Alert struct represents the data contained in an alert.
type Alert struct {
	ID          string
	Name        string
	IncidentID  string
	Severity    string
	Status      string
	CreatedAt   time.Time
	WebURL      string
	ClusterID   string
	ClusterName string
}

const (
	// PagerDuty Incident Statuses
	StatusTriggered    = "triggered"
	StatusAcknowledged = "acknowledged"
	StatusHigh         = "high"
	StatusLow          = "low"
)

type PagerDuty struct {
	client PagerDutyClient
}

func NewPagerDuty(client PagerDutyClient) *PagerDuty {
	return &PagerDuty{
		client: client,
	}
}

// NewWithToken initializes a new PDClient.
// The token can be created using the docs https://support.pagerduty.com/docs/api-access-keys#section-generate-a-user-token-rest-api-key.
func NewWithToken(authToken string, options ...pdApi.ClientOptions) (*PagerDuty, error) {

	pd := NewClient()
	err := pd.Connect(authToken, options...)

	if err != nil {
		return nil, err
	}

	return &PagerDuty{
		client: pd,
	}, nil

}

// GetIncidentAlerts returns all the alerts belonging to a particular incident.
func (pd *PagerDuty) GetIncidentAlerts(incidentID string) ([]Alert, error) {
	var alerts []Alert

	// Fetch alerts related to an incident via pagerduty API
	incidentAlerts, err := pd.client.ListIncidentAlerts(incidentID)

	if err != nil {
		var aerr pdApi.APIError

		if errors.As(err, &aerr) {
			if aerr.RateLimited() {
				return nil, fmt.Errorf("API rate limited")
			}

			return nil, fmt.Errorf("status code: %d, error: %s", aerr.StatusCode, err)
		}
	}

	for _, alert := range incidentAlerts.Alerts {
		currentAlert := alert
		formatAlert, err := pd.formatAlert(&currentAlert)

		if err != nil {
			return nil, err
		}
		alerts = append(alerts, formatAlert)
	}
	return alerts, nil

}

// GetClusterName interacts with the PD service endpoint and returns the cluster name string.
func (pd *PagerDuty) GetClusterName(serviceID string) (string, error) {
	service, err := pd.client.GetServiceWithContext(context.TODO(), serviceID, &pdApi.GetServiceOptions{})

	if err != nil {
		return "", err
	}

	clusterName := strings.Split(service.Description, " ")[0]

	return clusterName, nil
}

func (pd *PagerDuty) formatAlert(alert *pdApi.IncidentAlert) (formatAlert Alert, err error) {
	formatAlert.IncidentID = alert.Incident.ID

	formatAlert.Name = alert.Summary
	formatAlert.Status = alert.Status
	formatAlert.WebURL = alert.HTMLURL

	formatAlert.ClusterID = fmt.Sprint(alert.Body["details"].(map[string]interface{})["cluster_id"])
	formatAlert.ClusterName, err = pd.GetClusterName(alert.Service.ID)

	// If the service mapped to the current incident is not available (404)
	if err != nil {
		formatAlert.ClusterName = "N/A"
	}

	// If there's no cluster ID related to the given alert
	if formatAlert.ClusterID == "" {
		formatAlert.ClusterID = "N/A"
	}

	return formatAlert, nil
}
