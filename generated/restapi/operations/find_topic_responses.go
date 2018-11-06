// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sul-dlss/rialto-entity-resolver/generated/models"
)

// FindTopicOKCode is the HTTP code returned for type FindTopicOK
const FindTopicOKCode int = 200

/*FindTopicOK returns a URI for the topic in RIALTO

swagger:response findTopicOK
*/
type FindTopicOK struct {

	/*contains the actual URI
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewFindTopicOK creates FindTopicOK with default headers values
func NewFindTopicOK() *FindTopicOK {

	return &FindTopicOK{}
}

// WithPayload adds the payload to the find topic o k response
func (o *FindTopicOK) WithPayload(payload string) *FindTopicOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find topic o k response
func (o *FindTopicOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindTopicOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// FindTopicNotFoundCode is the HTTP code returned for type FindTopicNotFound
const FindTopicNotFoundCode int = 404

/*FindTopicNotFound Topic not found

swagger:response findTopicNotFound
*/
type FindTopicNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFindTopicNotFound creates FindTopicNotFound with default headers values
func NewFindTopicNotFound() *FindTopicNotFound {

	return &FindTopicNotFound{}
}

// WithPayload adds the payload to the find topic not found response
func (o *FindTopicNotFound) WithPayload(payload *models.ErrorResponse) *FindTopicNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find topic not found response
func (o *FindTopicNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindTopicNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindTopicInternalServerErrorCode is the HTTP code returned for type FindTopicInternalServerError
const FindTopicInternalServerErrorCode int = 500

/*FindTopicInternalServerError Server error

swagger:response findTopicInternalServerError
*/
type FindTopicInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFindTopicInternalServerError creates FindTopicInternalServerError with default headers values
func NewFindTopicInternalServerError() *FindTopicInternalServerError {

	return &FindTopicInternalServerError{}
}

// WithPayload adds the payload to the find topic internal server error response
func (o *FindTopicInternalServerError) WithPayload(payload *models.ErrorResponse) *FindTopicInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find topic internal server error response
func (o *FindTopicInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindTopicInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
