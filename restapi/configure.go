// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"emptyApi/config"
	"emptyApi/controllers"
	"emptyApi/lib"
	"emptyApi/repositories"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"emptyApi/restapi/operations"
)

//go:generate swagger generate server --target .. --name api --spec ../api-schema.yaml --exclude-main

func configureFlags(api *operations.API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	config.Init()
	ah:= controllers.ApiHandler{repositories.NewTestRepository(lib.Connect)}

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	api.BearerAuth = ah.Authorization
	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	api.KeyHandler = operations.KeyHandlerFunc(ah.Key)
	api.GetHandler = operations.GetHandlerFunc(ah.Get)
	api.PostHandler = operations.PostHandlerFunc(ah.Post)
	api.HealthHandler = operations.HealthHandlerFunc(ah.Health)

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
	return lib.Middleware(handler)
}
