package pagerduty

import (
	"context"

	pdApi "github.com/PagerDuty/go-pagerduty"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	pdMock "samanthajayasinghe.github.io/golang/features/testing/mocking/mock"
)

var _ = Describe("Pagerduty", func() {
	var (
		mockCtrl        *gomock.Controller
		mockPdClient    *pdMock.MockPagerDutyClient
		pagerDuty       *PagerDuty
		testIncidentID  string
		testClusterID   string
		testServiceID   string
		testClusterName string
		testAlertName   string
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockPdClient = pdMock.NewMockPagerDutyClient(mockCtrl)
		pagerDuty = NewPagerDuty(mockPdClient)
		testIncidentID = "incident-id-000"
		testClusterID = "cluster-id-000"
		testServiceID = "service-id-000"
		testClusterName = "test-cluster-name"
		testAlertName = "alert-error-budget"
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("When pagerduty client executes", func() {
		It("Should return incidents", func() {

			// Mock alert response
			alertResponse := &pdApi.ListAlertsResponse{
				Alerts: []pdApi.IncidentAlert{
					alert(
						testIncidentID,
						testServiceID,
						testAlertName,
						testClusterID,
						StatusTriggered,
					),
				},
			}

			// Mock the service response
			serviceResponse := &pdApi.Service{
				Description: testClusterName,
			}

			mockPdClient.EXPECT().ListIncidentAlerts(testIncidentID).Return(alertResponse, nil).Times(1)
			mockPdClient.EXPECT().GetServiceWithContext(context.TODO(), testServiceID, gomock.Any()).Return(serviceResponse, nil).Times(1)

			alerts, err := pagerDuty.GetIncidentAlerts(testIncidentID)
			Expect(err).To(BeNil())
			Expect(len(alerts)).To(Equal(1))

		})

		It("Should returns a format alert", func() {

			nonFormatAlert := alert(
				testIncidentID,
				testServiceID,
				testAlertName,
				testClusterID,
				StatusTriggered,
			)

			// Mock the service response
			serviceResponse := &pdApi.Service{
				Description: testClusterName,
			}

			mockPdClient.EXPECT().GetServiceWithContext(context.TODO(), testServiceID, gomock.Any()).Return(serviceResponse, nil).Times(1)

			formatAlert, err := pagerDuty.formatAlert(&nonFormatAlert)
			Expect(err).To(BeNil())
			Expect(formatAlert.ClusterID).To(Equal(testClusterID))
			Expect(formatAlert.ClusterName).To(Equal(testClusterName))
			Expect(formatAlert.Name).To(Equal(testAlertName))
			Expect(formatAlert.IncidentID).To(Equal(testIncidentID))

		})

	})
})

// alert retuns a pagerduty alert object with pre-configured data.
func alert(
	incidentID string,
	serviceID string,
	name string,
	clusterID string,
	status string,
) pdApi.IncidentAlert {

	body := map[string]interface{}{
		"details": map[string]interface{}{
			"cluster_id": clusterID,
		},
	}

	if name == "CHGM" {
		body = map[string]interface{}{
			"details": map[string]interface{}{
				"notes": [...]string{
					"cluster_id: " + clusterID,
				},
			},
		}
	}
	return pdApi.IncidentAlert{

		Incident: pdApi.APIReference{
			ID: incidentID,
		},

		Service: pdApi.APIObject{
			ID: serviceID,
		},

		APIObject: pdApi.APIObject{
			Summary: name,
		},

		Body: body,

		Status: status,
	}
}
