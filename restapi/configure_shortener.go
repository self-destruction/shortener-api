// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/dre1080/recover"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ivahaev/go-logger"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"log"
	"net/http"
	"shortener-api/models"
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

var connect *sqlx.DB

func hitError(err error, errorCode int) middleware.Responder {
	return user.NewCreateUserBadRequest().WithPayload(
		&models.Error{
			Code:    swag.Int64(int64(errorCode)),
			Message: err.Error(),
		},
	)
}

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
	api.BasicAuthAuth = func(user string, pass string) (interface{}, error) {
		return nil, errors.NotImplemented("basic auth  (BasicAuth) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.LinkCreateShortLinkHandler = link.CreateShortLinkHandlerFunc(func(params link.CreateShortLinkParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation link.CreateShortLink has not yet been implemented")
	})
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
		userRequest := params.Body

		stmt, err := connect.Prepare("INSERT `shortener`.`user` SET username=?, hash=?, email=?, timezone=?, language=?")
		if err != nil {
			logger.Debug(err)
			return hitError(err, user.CreateUserBadRequestCode)
		}

		// insert user into db
		res, err := stmt.Exec(userRequest.Username, userRequest.Password, userRequest.Email, userRequest.Timezone, userRequest.Language)
		if err != nil {
			logger.Debug(err)
			return hitError(err, user.CreateUserBadRequestCode)
		}

		insertId, err := res.LastInsertId()
		if err != nil {
			logger.Debug(err)
			return hitError(err, user.CreateUserBadRequestCode)
		}

		// select user by last insert id for response
		userDB := &models.User{}
		query := "SELECT id, username, hash, email, timezone, language, createdAt AS dateCreated FROM `shortener`.`user` WHERE id=? LIMIT 1"
		err = connect.Get(userDB, query, insertId)
		if err != nil {
			logger.Debug(err)
			return hitError(err, user.CreateUserBadRequestCode)
		}

		logger.JSON(userDB)
		return user.NewCreateUserOK().WithPayload(userDB)
	})
	api.StatisticGetCurrentUserHandler = statistic.GetCurrentUserHandlerFunc(func(params statistic.GetCurrentUserParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation statistic.GetCurrentUser has not yet been implemented")
	})
	api.LinkGetLinkHandler = link.GetLinkHandlerFunc(func(params link.GetLinkParams) middleware.Responder {
		return middleware.NotImplemented("operation link.GetLink has not yet been implemented")
	})
	api.StatisticGetLinkInfoHandler = statistic.GetLinkInfoHandlerFunc(func(params statistic.GetLinkInfoParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation statistic.GetLinkInfo has not yet been implemented")
	})
	api.LinkGetLinksHandler = link.GetLinksHandlerFunc(func(params link.GetLinksParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation link.GetLinks has not yet been implemented")
	})
	api.StatisticGetReferersHandler = statistic.GetReferersHandlerFunc(func(params statistic.GetReferersParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation statistic.GetReferers has not yet been implemented")
	})
	api.UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
	})
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
	db, err := sqlx.Open("mysql", "root:pwd@tcp(192.168.99.100:3307)/shortener?charset=utf8")
	if err != nil {
		log.Fatalf(err.Error())
	}
	//defer db.Close()

	// Ping to connection
	err = db.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// https://github.com/jmoiron/sqlx/issues/322
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	connect = db

	err = logger.SetLevel("DEBUG")
	if err != nil {
		log.Fatalf(err.Error())
	}

	recovery := recover.New(&recover.Options{
		Log: logger.Info,
	})

	return recovery(handler)
}
