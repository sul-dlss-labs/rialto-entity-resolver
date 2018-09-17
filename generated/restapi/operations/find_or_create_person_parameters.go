// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindOrCreatePersonParams creates a new FindOrCreatePersonParams object
// no default values defined in spec.
func NewFindOrCreatePersonParams() FindOrCreatePersonParams {

	return FindOrCreatePersonParams{}
}

// FindOrCreatePersonParams contains all the bound params for the find or create person operation
// typically these are obtained from a http.Request
//
// swagger:parameters findOrCreatePerson
type FindOrCreatePersonParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*First name of the person
	  In: query
	*/
	FirstName *string
	/*Full name of the person
	  In: query
	*/
	FullName *string
	/*Last name of the person
	  In: query
	*/
	LastName *string
	/*ORCID of the person
	  In: query
	*/
	Orcid *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindOrCreatePersonParams() beforehand.
func (o *FindOrCreatePersonParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFirstName, qhkFirstName, _ := qs.GetOK("first_name")
	if err := o.bindFirstName(qFirstName, qhkFirstName, route.Formats); err != nil {
		res = append(res, err)
	}

	qFullName, qhkFullName, _ := qs.GetOK("full_name")
	if err := o.bindFullName(qFullName, qhkFullName, route.Formats); err != nil {
		res = append(res, err)
	}

	qLastName, qhkLastName, _ := qs.GetOK("last_name")
	if err := o.bindLastName(qLastName, qhkLastName, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrcid, qhkOrcid, _ := qs.GetOK("orcid")
	if err := o.bindOrcid(qOrcid, qhkOrcid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFirstName binds and validates parameter FirstName from query.
func (o *FindOrCreatePersonParams) bindFirstName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FirstName = &raw

	return nil
}

// bindFullName binds and validates parameter FullName from query.
func (o *FindOrCreatePersonParams) bindFullName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FullName = &raw

	return nil
}

// bindLastName binds and validates parameter LastName from query.
func (o *FindOrCreatePersonParams) bindLastName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.LastName = &raw

	return nil
}

// bindOrcid binds and validates parameter Orcid from query.
func (o *FindOrCreatePersonParams) bindOrcid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Orcid = &raw

	return nil
}
