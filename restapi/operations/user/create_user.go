// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// CreateUserHandlerFunc turns a function with the right signature into a create user handler
type CreateUserHandlerFunc func(CreateUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateUserHandlerFunc) Handle(params CreateUserParams) middleware.Responder {
	return fn(params)
}

// CreateUserHandler interface for that can handle valid create user params
type CreateUserHandler interface {
	Handle(CreateUserParams) middleware.Responder
}

// NewCreateUser creates a new http.Handler for the create user operation
func NewCreateUser(ctx *middleware.Context, handler CreateUserHandler) *CreateUser {
	return &CreateUser{Context: ctx, Handler: handler}
}

/*CreateUser swagger:route POST /users User createUser

User registrations

*/
type CreateUser struct {
	Context *middleware.Context
	Handler CreateUserHandler
}

func (o *CreateUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateUserBody create user body
// swagger:model CreateUserBody
type CreateUserBody struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// language
	Language string `json:"language,omitempty"`

	// password
	// Required: true
	// Format: password
	Password *strfmt.Password `json:"password"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this create user body
func (o *CreateUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateUserBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *CreateUserBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	if err := validate.FormatOf("body"+"."+"password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateUserBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateUserBody) UnmarshalBinary(b []byte) error {
	var res CreateUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}