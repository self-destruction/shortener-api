// Code generated by go-swagger; DO NOT EDIT.

package link

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "shortener-api/models"
)

// GetLinkMovedPermanentlyCode is the HTTP code returned for type GetLinkMovedPermanently
const GetLinkMovedPermanentlyCode int = 301

/*GetLinkMovedPermanently Link redirect

swagger:response getLinkMovedPermanently
*/
type GetLinkMovedPermanently struct {
}

// NewGetLinkMovedPermanently creates GetLinkMovedPermanently with default headers values
func NewGetLinkMovedPermanently() *GetLinkMovedPermanently {

	return &GetLinkMovedPermanently{}
}

// WriteResponse to the client
func (o *GetLinkMovedPermanently) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(301)
}

// GetLinkNotFoundCode is the HTTP code returned for type GetLinkNotFound
const GetLinkNotFoundCode int = 404

/*GetLinkNotFound The specified resource was not found

swagger:response getLinkNotFound
*/
type GetLinkNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLinkNotFound creates GetLinkNotFound with default headers values
func NewGetLinkNotFound() *GetLinkNotFound {

	return &GetLinkNotFound{}
}

// WithPayload adds the payload to the get link not found response
func (o *GetLinkNotFound) WithPayload(payload *models.Error) *GetLinkNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get link not found response
func (o *GetLinkNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLinkNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
