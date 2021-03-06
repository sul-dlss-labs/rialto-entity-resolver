// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindGrantParams creates a new FindGrantParams object
// no default values defined in spec.
func NewFindGrantParams() FindGrantParams {

	return FindGrantParams{}
}

// FindGrantParams contains all the bound params for the find grant operation
// typically these are obtained from a http.Request
//
// swagger:parameters FindGrant
type FindGrantParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A grant identifier
	  Required: true
	  In: query
	*/
	Identifier string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindGrantParams() beforehand.
func (o *FindGrantParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qIdentifier, qhkIdentifier, _ := qs.GetOK("identifier")
	if err := o.bindIdentifier(qIdentifier, qhkIdentifier, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIdentifier binds and validates parameter Identifier from query.
func (o *FindGrantParams) bindIdentifier(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("identifier", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("identifier", "query", raw); err != nil {
		return err
	}

	o.Identifier = raw

	return nil
}
