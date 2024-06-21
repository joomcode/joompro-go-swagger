// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/joomcode/joompro-go-swagger/examples/composed-auth/models"
)

// NewMultiAuthExampleAPI creates a new MultiAuthExample instance
func NewMultiAuthExampleAPI(spec *loads.Document) *MultiAuthExampleAPI {
	return &MultiAuthExampleAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		AddOrderHandler: AddOrderHandlerFunc(func(params AddOrderParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation AddOrder has not yet been implemented")
		}),
		GetAccountHandler: GetAccountHandlerFunc(func(params GetAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetAccount has not yet been implemented")
		}),
		GetItemsHandler: GetItemsHandlerFunc(func(params GetItemsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetItems has not yet been implemented")
		}),
		GetOrderHandler: GetOrderHandlerFunc(func(params GetOrderParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetOrder has not yet been implemented")
		}),
		GetOrdersForItemHandler: GetOrdersForItemHandlerFunc(func(params GetOrdersForItemParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetOrdersForItem has not yet been implemented")
		}),

		HasRoleAuth: func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (hasRole) has not yet been implemented")
		},
		// Applies when the Authorization header is set with the Basic scheme
		IsRegisteredAuth: func(user string, pass string) (*models.Principal, error) {
			return nil, errors.NotImplemented("basic auth  (isRegistered) has not yet been implemented")
		},
		// Applies when the "X-Custom-Key" header is set
		IsResellerAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (isReseller) X-Custom-Key from header param [X-Custom-Key] has not yet been implemented")
		},
		// Applies when the "CustomKeyAsQuery" query is set
		IsResellerQueryAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (isResellerQuery) CustomKeyAsQuery from query param [CustomKeyAsQuery] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*
MultiAuthExampleAPI This sample API demonstrates how to compose several authentication schemes
and configure complex security requirements for your operations.

This API simulates a very simple market place with customers and resellers
of items.

Personas:
  - as a first time user, I want to see all items on sales
  - as a registered customer, I want to post orders for items and
    consult my past orders
  - as a registered reseller, I want to see all pending orders on the items
    I am selling on the market place
  - as a reseller managing my own inventories, I want to post replenishment orders for the items I provide
  - as a register user, I want to consult my personal account infos

The situation we defined on the authentication side is as follows:
  - every known user is authenticated using a basic token
  - resellers are authenticated using API keys - we let the option to authenticate using a header or query param
  - any registered user (customer or reseller) will add a signed JWT to access more API endpoints

Obviously, there are several ways to achieve the same result. We just wanted to demonstrate here how
security requirements may compose several schemes.

Note that we used the "OAuth2" declaration here but don't implement a real
OAuth2 workflow: our intend here is just to be able to extract scopes from a passed JWT token (the
only way to manipulate scoped authorizers with Swagger 2.0 is to declare them with type "oauth2").
*/
type MultiAuthExampleAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// HasRoleAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	HasRoleAuth func(string, []string) (*models.Principal, error)

	// IsRegisteredAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	IsRegisteredAuth func(string, string) (*models.Principal, error)

	// IsResellerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Custom-Key provided in the header
	IsResellerAuth func(string) (*models.Principal, error)

	// IsResellerQueryAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key CustomKeyAsQuery provided in the query
	IsResellerQueryAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// AddOrderHandler sets the operation handler for the add order operation
	AddOrderHandler AddOrderHandler
	// GetAccountHandler sets the operation handler for the get account operation
	GetAccountHandler GetAccountHandler
	// GetItemsHandler sets the operation handler for the get items operation
	GetItemsHandler GetItemsHandler
	// GetOrderHandler sets the operation handler for the get order operation
	GetOrderHandler GetOrderHandler
	// GetOrdersForItemHandler sets the operation handler for the get orders for item operation
	GetOrdersForItemHandler GetOrdersForItemHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *MultiAuthExampleAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *MultiAuthExampleAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *MultiAuthExampleAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MultiAuthExampleAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MultiAuthExampleAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MultiAuthExampleAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MultiAuthExampleAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MultiAuthExampleAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MultiAuthExampleAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MultiAuthExampleAPI
func (o *MultiAuthExampleAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.HasRoleAuth == nil {
		unregistered = append(unregistered, "HasRoleAuth")
	}
	if o.IsRegisteredAuth == nil {
		unregistered = append(unregistered, "IsRegisteredAuth")
	}
	if o.IsResellerAuth == nil {
		unregistered = append(unregistered, "XCustomKeyAuth")
	}
	if o.IsResellerQueryAuth == nil {
		unregistered = append(unregistered, "CustomKeyAsQueryAuth")
	}

	if o.AddOrderHandler == nil {
		unregistered = append(unregistered, "AddOrderHandler")
	}
	if o.GetAccountHandler == nil {
		unregistered = append(unregistered, "GetAccountHandler")
	}
	if o.GetItemsHandler == nil {
		unregistered = append(unregistered, "GetItemsHandler")
	}
	if o.GetOrderHandler == nil {
		unregistered = append(unregistered, "GetOrderHandler")
	}
	if o.GetOrdersForItemHandler == nil {
		unregistered = append(unregistered, "GetOrdersForItemHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MultiAuthExampleAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MultiAuthExampleAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "hasRole":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.HasRoleAuth(token, scopes)
			})

		case "isRegistered":
			result[name] = o.BasicAuthenticator(func(username, password string) (interface{}, error) {
				return o.IsRegisteredAuth(username, password)
			})

		case "isReseller":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.IsResellerAuth(token)
			})

		case "isResellerQuery":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.IsResellerQueryAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *MultiAuthExampleAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MultiAuthExampleAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
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

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *MultiAuthExampleAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
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
func (o *MultiAuthExampleAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the multi auth example API
func (o *MultiAuthExampleAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MultiAuthExampleAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/order/add"] = NewAddOrder(o.context, o.AddOrderHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/account"] = NewGetAccount(o.context, o.GetAccountHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/items"] = NewGetItems(o.context, o.GetItemsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/order/{orderID}"] = NewGetOrder(o.context, o.GetOrderHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/orders/{itemID}"] = NewGetOrdersForItem(o.context, o.GetOrdersForItemHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MultiAuthExampleAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MultiAuthExampleAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MultiAuthExampleAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MultiAuthExampleAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MultiAuthExampleAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
