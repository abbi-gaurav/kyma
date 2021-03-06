package broker

import (
	"github.com/kyma-project/kyma/components/remote-environment-broker/internal"
	"github.com/pkg/errors"
)

// DTOs for Open Service Broker v2.12 API

type contextDTO struct {
	Platform  string             `json:"platform"`
	Namespace internal.Namespace `json:"namespace"`
}

// ProvisionRequestDTO represents provision request
type ProvisionRequestDTO struct {
	ServiceID        internal.ServiceID     `json:"service_id"`
	PlanID           internal.ServicePlanID `json:"plan_id"`
	OrganizationGUID string                 `json:"organization_guid"`
	SpaceGUID        string                 `json:"space_guid"`
	Parameters       map[string]interface{} `json:"parameters,omitempty"`
	Context          contextDTO             `json:"context,omitempty"`
}

// ProvisionSuccessResponseDTO represents response after successful provisioning
type ProvisionSuccessResponseDTO struct {
	DashboardURL *string               `json:"dashboard_url"`
	Operation    *internal.OperationID `json:"operation,omitempty"`
}

// DeprovisionSuccessResponseDTO represents response after successful deprovisioning
type DeprovisionSuccessResponseDTO struct {
	Operation *internal.OperationID `json:"operation,omitempty"`
}

// LastOperationSuccessResponseDTO represents info response about last successful operation
type LastOperationSuccessResponseDTO struct {
	State       internal.OperationState `json:"state"`
	Description *string                 `json:"description,omitempty"`
}

// BindSuccessResponseDTO represents response with credentials for service instance after successful binding
type BindSuccessResponseDTO struct {
	// Credentials is a free-form hash of credentials that can be used by
	// applications or users to access the service.
	Credentials map[string]interface{} `json:"credentials,omitempty"`
}

// BindParametersDTO contains parameters sent by Service Catalog in the body of bind request.
type BindParametersDTO struct {
	ServiceID string `json:"service_id"`
	PlanID    string `json:"plan_id"`
}

// Validate checks if bind parameters aren't empty
func (params *BindParametersDTO) Validate() error {
	if params.PlanID == "" || params.ServiceID == "" {
		return errors.New("bind parameters cannot be empty")
	}
	return nil
}
