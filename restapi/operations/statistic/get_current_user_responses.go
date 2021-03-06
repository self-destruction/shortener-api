// Code generated by go-swagger; DO NOT EDIT.

package statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "shortener-api/models"
)

// GetCurrentUserOKCode is the HTTP code returned for type GetCurrentUserOK
const GetCurrentUserOKCode int = 200

/*GetCurrentUserOK Get user

swagger:response getCurrentUserOK
*/
type GetCurrentUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetCurrentUserOK creates GetCurrentUserOK with default headers values
func NewGetCurrentUserOK() *GetCurrentUserOK {

	return &GetCurrentUserOK{}
}

// WithPayload adds the payload to the get current user o k response
func (o *GetCurrentUserOK) WithPayload(payload *models.User) *GetCurrentUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user o k response
func (o *GetCurrentUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCurrentUserUnauthorizedCode is the HTTP code returned for type GetCurrentUserUnauthorized
const GetCurrentUserUnauthorizedCode int = 401

/*GetCurrentUserUnauthorized Authentication information is missing or invalid

swagger:response getCurrentUserUnauthorized
*/
type GetCurrentUserUnauthorized struct {
	/*

	 */
	WWWAuthenticate string `json:"WWW_Authenticate"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCurrentUserUnauthorized creates GetCurrentUserUnauthorized with default headers values
func NewGetCurrentUserUnauthorized() *GetCurrentUserUnauthorized {

	return &GetCurrentUserUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the get current user unauthorized response
func (o *GetCurrentUserUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *GetCurrentUserUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the get current user unauthorized response
func (o *GetCurrentUserUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WithPayload adds the payload to the get current user unauthorized response
func (o *GetCurrentUserUnauthorized) WithPayload(payload *models.Error) *GetCurrentUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user unauthorized response
func (o *GetCurrentUserUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
