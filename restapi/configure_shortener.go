// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/dre1080/recover"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ivahaev/go-logger"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"shortener-api/api/handlers"
	"shortener-api/restapi/operations"
	"shortener-api/restapi/operations/link"
	"shortener-api/restapi/operations/statistic"
	"shortener-api/restapi/operations/user"
	"strings"
)

//go:generate swagger generate server --target ..\..\shortener-api --name Shortener --spec ..\swagger.yml

func configureFlags(api *operations.ShortenerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

var h handlers.Handler

func configureAPI(api *operations.ShortenerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = h.Authorize

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.LinkCreateShortLinkHandler = link.CreateShortLinkHandlerFunc(h.CreateShortLink)
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(h.CreateUser)
	api.StatisticGetCurrentUserHandler = statistic.GetCurrentUserHandlerFunc(h.GetCurrentUser)
	api.LinkGetLinkHandler = link.GetLinkHandlerFunc(h.GetLink)
	api.StatisticGetLinkInfoHandler = statistic.GetLinkInfoHandlerFunc(h.GetLinkInfo)
	api.LinkGetLinksHandler = link.GetLinksHandlerFunc(h.GetLinks)
	api.StatisticGetReferersHandler = statistic.GetReferersHandlerFunc(func(params statistic.GetReferersParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation statistic.GetReferers has not yet been implemented")
	})
	api.UserLoginUserHandler = user.LoginUserHandlerFunc(h.LoginUser)
	api.UserLogoutUserHandler = user.LogoutUserHandlerFunc(func(params user.LogoutUserParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
	})
	api.LinkRemoveLinkHandler = link.RemoveLinkHandlerFunc(func(params link.RemoveLinkParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation link.RemoveLink has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	// reading config from file
	viper.SetConfigName("database")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// initial config string
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("host"),
		viper.GetString("port"),
		viper.GetString("database"),
	)
	db, err := sqlx.Connect(viper.GetString("adapter"), connectionString)
	if err != nil {
		log.Fatalf(err.Error())
	}
	//defer db.Close()

	// https://github.com/jmoiron/sqlx/issues/322
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	h = handlers.Handler{Connect: db}

	err = logger.SetLevel("DEBUG")
	if err != nil {
		log.Fatalf(err.Error())
	}
	recovery := recover.New(&recover.Options{
		Log: logger.Crit,
	})
	return recovery(handler)
}
