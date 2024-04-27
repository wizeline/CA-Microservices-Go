package controller

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/wizeline/CA-Microservices-Go/internal/repository"
	"github.com/wizeline/CA-Microservices-Go/internal/service"

	"github.com/go-chi/render"
)

const (
	repoErrStatus        errStatus = "RepositoryError"
	svcErrStatus         errStatus = "ServiceError"
	ctrlPayloadErrStatus errStatus = "ControllerPayloadError"
	ctrlParamErrStatus   errStatus = "ControllerParameterError"
)

var _ fmt.Stringer = errStatus("")

// errHTTP represents the default http error responses.
type errHTTP struct {
	Code    int       `json:"code"`
	Status  errStatus `json:"status"`
	Message string    `json:"message"`
}

// errStatus represents the error classification or status types.
type errStatus string

func (e errStatus) String() string {
	return string(e)
}

type PayloadErr struct {
	Err error
}

func (e PayloadErr) Error() string {
	return fmt.Sprintf("invalid payload: %s", e.Err)
}

type ParameterErr struct {
	Param string
	Err   string
}

func (e ParameterErr) Error() string {
	return fmt.Sprintf("invalid %v parameter: %v", e.Param, e.Err)
}

// newErrHTTP returns a new HTTP error message based on error types
func newErrHTTP(err error) errHTTP {
	var (
		repoErr        *repository.Err
		svcErr         *service.Err
		ctrlPayloadErr *PayloadErr
		ctrlParamErr   *ParameterErr
	)

	switch {

	// ########### REPOSITORY ERRORS ###########

	case errors.As(err, &repoErr):

		// TODO: evaluate the rest of the repository errors

		return errHTTP{
			Code:    http.StatusInternalServerError,
			Status:  repoErrStatus,
			Message: err.Error(),
		}

	// ########### SERVICE ERRORS ###########

	case errors.As(err, &svcErr):

		// TODO: evaluate the rest of the service errors

		return errHTTP{
			Code:    http.StatusBadRequest,
			Status:  svcErrStatus,
			Message: err.Error(),
		}

	// ########### CONTROLLER ERRORS ###########

	case errors.As(err, &ctrlParamErr):
		return errHTTP{
			Code:    http.StatusBadRequest,
			Status:  ctrlParamErrStatus,
			Message: err.Error(),
		}

	case errors.As(err, &ctrlPayloadErr):
		return errHTTP{
			Code:    http.StatusUnsupportedMediaType,
			Status:  ctrlPayloadErrStatus,
			Message: err.Error(),
		}

	// ########### DEFAULT ERRORS ###########

	default:
		return errHTTP{
			Code:    http.StatusBadRequest,
			Status:  errStatus(reflect.TypeOf(err).String()),
			Message: err.Error(),
		}
	}
}

// errJSON returns an error JSON response
func errJSON(w http.ResponseWriter, r *http.Request, err error) {
	errHttp := newErrHTTP(err)
	render.Status(r, errHttp.Code)
	render.JSON(w, r, errHttp)
}
