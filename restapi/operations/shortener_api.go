// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"shortener-api/restapi/operations/link"
	"shortener-api/restapi/operations/statistic"
	"shortener-api/restapi/operations/user"
)

// NewShortenerAPI creates a new Shortener instance
func NewShortenerAPI(spec *loads.Document) *ShortenerAPI {
	return &ShortenerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		LinkCreateShortLinkHandler: link.CreateShortLinkHandlerFunc(func(params link.CreateShortLinkParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LinkCreateShortLink has not yet been implemented")
		}),
		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserCreateUser has not yet been implemented")
		}),
		StatisticGetCurrentUserHandler: statistic.GetCurrentUserHandlerFunc(func(params statistic.GetCurrentUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation StatisticGetCurrentUser has not yet been implemented")
		}),
		LinkGetLinkHandler: link.GetLinkHandlerFunc(func(params link.GetLinkParams) middleware.Responder {
			return middleware.NotImplemented("operation LinkGetLink has not yet been implemented")
		}),
		StatisticGetLinkInfoHandler: statistic.GetLinkInfoHandlerFunc(func(params statistic.GetLinkInfoParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation StatisticGetLinkInfo has not yet been implemented")
		}),
		LinkGetLinksHandler: link.GetLinksHandlerFunc(func(params link.GetLinksParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LinkGetLinks has not yet been implemented")
		}),
		StatisticGetReferersHandler: statistic.GetReferersHandlerFunc(func(params statistic.GetReferersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation StatisticGetReferers has not yet been implemented")
		}),
		UserLoginUserHandler: user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserLoginUser has not yet been implemented")
		}),
		UserLogoutUserHandler: user.LogoutUserHandlerFunc(func(params user.LogoutUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation UserLogoutUser has not yet been implemented")
		}),
		LinkRemoveLinkHandler: link.RemoveLinkHandlerFunc(func(params link.RemoveLinkParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LinkRemoveLink has not yet been implemented")
		}),

		// Applies when the Authorization header is set with the Basic scheme
		BasicAuthAuth: func(user string, pass string) (interface{}, error) {
			return nil, errors.NotImplemented("basic auth  (BasicAuth) has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ShortenerAPI This is a simple API for Shortener URL */
type ShortenerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// BasicAuthAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	BasicAuthAuth func(string, string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// LinkCreateShortLinkHandler sets the operation handler for the create short link operation
	LinkCreateShortLinkHandler link.CreateShortLinkHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// StatisticGetCurrentUserHandler sets the operation handler for the get current user operation
	StatisticGetCurrentUserHandler statistic.GetCurrentUserHandler
	// LinkGetLinkHandler sets the operation handler for the get link operation
	LinkGetLinkHandler link.GetLinkHandler
	// StatisticGetLinkInfoHandler sets the operation handler for the get link info operation
	StatisticGetLinkInfoHandler statistic.GetLinkInfoHandler
	// LinkGetLinksHandler sets the operation handler for the get links operation
	LinkGetLinksHandler link.GetLinksHandler
	// StatisticGetReferersHandler sets the operation handler for the get referers operation
	StatisticGetReferersHandler statistic.GetReferersHandler
	// UserLoginUserHandler sets the operation handler for the login user operation
	UserLoginUserHandler user.LoginUserHandler
	// UserLogoutUserHandler sets the operation handler for the logout user operation
	UserLogoutUserHandler user.LogoutUserHandler
	// LinkRemoveLinkHandler sets the operation handler for the remove link operation
	LinkRemoveLinkHandler link.RemoveLinkHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ShortenerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ShortenerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ShortenerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ShortenerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ShortenerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ShortenerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ShortenerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ShortenerAPI
func (o *ShortenerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BasicAuthAuth == nil {
		unregistered = append(unregistered, "BasicAuthAuth")
	}

	if o.LinkCreateShortLinkHandler == nil {
		unregistered = append(unregistered, "link.CreateShortLinkHandler")
	}

	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}

	if o.StatisticGetCurrentUserHandler == nil {
		unregistered = append(unregistered, "statistic.GetCurrentUserHandler")
	}

	if o.LinkGetLinkHandler == nil {
		unregistered = append(unregistered, "link.GetLinkHandler")
	}

	if o.StatisticGetLinkInfoHandler == nil {
		unregistered = append(unregistered, "statistic.GetLinkInfoHandler")
	}

	if o.LinkGetLinksHandler == nil {
		unregistered = append(unregistered, "link.GetLinksHandler")
	}

	if o.StatisticGetReferersHandler == nil {
		unregistered = append(unregistered, "statistic.GetReferersHandler")
	}

	if o.UserLoginUserHandler == nil {
		unregistered = append(unregistered, "user.LoginUserHandler")
	}

	if o.UserLogoutUserHandler == nil {
		unregistered = append(unregistered, "user.LogoutUserHandler")
	}

	if o.LinkRemoveLinkHandler == nil {
		unregistered = append(unregistered, "link.RemoveLinkHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ShortenerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ShortenerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "BasicAuth":
			_ = scheme
			result[name] = o.BasicAuthenticator(o.BasicAuthAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *ShortenerAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *ShortenerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ShortenerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ShortenerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the shortener API
func (o *ShortenerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ShortenerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/me/shorten_urls"] = link.NewCreateShortLink(o.context, o.LinkCreateShortLinkHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me"] = statistic.NewGetCurrentUser(o.context, o.StatisticGetCurrentUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/shorten_urls/{hash}"] = link.NewGetLink(o.context, o.LinkGetLinkHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/shorten_urls/{hash}"] = statistic.NewGetLinkInfo(o.context, o.StatisticGetLinkInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/shorten_urls"] = link.NewGetLinks(o.context, o.LinkGetLinksHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/shorten_urls/{hash}/referers"] = statistic.NewGetReferers(o.context, o.StatisticGetReferersHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/login"] = user.NewLoginUser(o.context, o.UserLoginUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/logout"] = user.NewLogoutUser(o.context, o.UserLogoutUserHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/me/shorten_urls/{hash}"] = link.NewRemoveLink(o.context, o.LinkRemoveLinkHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ShortenerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ShortenerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ShortenerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ShortenerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
