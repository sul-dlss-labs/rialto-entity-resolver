// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sul-dlss/rialto-entity-resolver/generated/models"
)

// FindPersonOKCode is the HTTP code returned for type FindPersonOK
const FindPersonOKCode int = 200

/*FindPersonOK returns a URI for the person in Rialto

swagger:response findPersonOK
*/
type FindPersonOK struct {

	/*contains the actual URI
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewFindPersonOK creates FindPersonOK with default headers values
func NewFindPersonOK() *FindPersonOK {

	return &FindPersonOK{}
}

// WithPayload adds the payload to the find person o k response
func (o *FindPersonOK) WithPayload(payload string) *FindPersonOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find person o k response
func (o *FindPersonOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPersonOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// FindPersonNotFoundCode is the HTTP code returned for type FindPersonNotFound
const FindPersonNotFoundCode int = 404

/*FindPersonNotFound Person not found

swagger:response findPersonNotFound
*/
type FindPersonNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFindPersonNotFound creates FindPersonNotFound with default headers values
func NewFindPersonNotFound() *FindPersonNotFound {

	return &FindPersonNotFound{}
}

// WithPayload adds the payload to the find person not found response
func (o *FindPersonNotFound) WithPayload(payload *models.ErrorResponse) *FindPersonNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find person not found response
func (o *FindPersonNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPersonNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindPersonInternalServerErrorCode is the HTTP code returned for type FindPersonInternalServerError
const FindPersonInternalServerErrorCode int = 500

/*FindPersonInternalServerError Server error

swagger:response findPersonInternalServerError
*/
type FindPersonInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFindPersonInternalServerError creates FindPersonInternalServerError with default headers values
func NewFindPersonInternalServerError() *FindPersonInternalServerError {

	return &FindPersonInternalServerError{}
}

// WithPayload adds the payload to the find person internal server error response
func (o *FindPersonInternalServerError) WithPayload(payload *models.ErrorResponse) *FindPersonInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find person internal server error response
func (o *FindPersonInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPersonInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
