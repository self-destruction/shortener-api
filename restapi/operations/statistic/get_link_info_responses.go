// Code generated by go-swagger; DO NOT EDIT.

package statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "shortener-api/models"
)

// GetLinkInfoOKCode is the HTTP code returned for type GetLinkInfoOK
const GetLinkInfoOKCode int = 200

/*GetLinkInfoOK Transitions for link

swagger:response getLinkInfoOK
*/
type GetLinkInfoOK struct {

	/*
	  In: Body
	*/
	Payload *GetLinkInfoOKBody `json:"body,omitempty"`
}

// NewGetLinkInfoOK creates GetLinkInfoOK with default headers values
func NewGetLinkInfoOK() *GetLinkInfoOK {

	return &GetLinkInfoOK{}
}

// WithPayload adds the payload to the get link info o k response
func (o *GetLinkInfoOK) WithPayload(payload *GetLinkInfoOKBody) *GetLinkInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get link info o k response
func (o *GetLinkInfoOK) SetPayload(payload *GetLinkInfoOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLinkInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLinkInfoBadRequestCode is the HTTP code returned for type GetLinkInfoBadRequest
const GetLinkInfoBadRequestCode int = 400

/*GetLinkInfoBadRequest Bad Request

swagger:response getLinkInfoBadRequest
*/
type GetLinkInfoBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLinkInfoBadRequest creates GetLinkInfoBadRequest with default headers values
func NewGetLinkInfoBadRequest() *GetLinkInfoBadRequest {

	return &GetLinkInfoBadRequest{}
}

// WithPayload adds the payload to the get link info bad request response
func (o *GetLinkInfoBadRequest) WithPayload(payload *models.Error) *GetLinkInfoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get link info bad request response
func (o *GetLinkInfoBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLinkInfoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLinkInfoNotFoundCode is the HTTP code returned for type GetLinkInfoNotFound
const GetLinkInfoNotFoundCode int = 404

/*GetLinkInfoNotFound The specified resource was not found

swagger:response getLinkInfoNotFound
*/
type GetLinkInfoNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLinkInfoNotFound creates GetLinkInfoNotFound with default headers values
func NewGetLinkInfoNotFound() *GetLinkInfoNotFound {

	return &GetLinkInfoNotFound{}
}

// WithPayload adds the payload to the get link info not found response
func (o *GetLinkInfoNotFound) WithPayload(payload *models.Error) *GetLinkInfoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get link info not found response
func (o *GetLinkInfoNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLinkInfoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}