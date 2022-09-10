package requests

import (
	"encoding/json"
	"net/http"

	"blob-base/resources"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateAccountRequest struct {
	Data resources.TransactionEnvelope
}

func NewCreateAccountRequest(r *http.Request) (CreateAccountRequest, error) {
	var request CreateAccountRequest

	if err := json.NewDecoder(r.Body).Decode(&request.Data); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateAccountRequest) validate() error {
	return mergeErrors(validation.Errors{
		"attributes": validation.Validate(&r.Data.Attributes, validation.Required),
	}).Filter()
}
