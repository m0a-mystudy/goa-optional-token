package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/m0a-mystudy/goa-optional-token/app"
)

// NewBasicAuthMiddleware creates a middleware that checks for the presence of a basic auth header
// and validates its content.
func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			// A real app would do something more interesting here
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("missing auth")
			}

			// Proceed
			goa.LogInfo(ctx, "basic", "user", user, "pass", pass)
			return h(ctx, rw, req)
		}
	}
}

type OptionalBasicAuth struct {
	User string
	Pass string
}

type optionalBasicAuthKeyType int

const (
	optionalBasicAuthKey optionalBasicAuthKeyType = iota + 1
)

func NewOptinalBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			// A real app would do something more interesting here
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return h(ctx, rw, req)
			}

			// Proceed
			goa.LogInfo(ctx, "basic", "user", user, "pass", pass)
			ctx = context.WithValue(ctx, optionalBasicAuthKey, &OptionalBasicAuth{
				User: user,
				Pass: pass,
			})
			return h(ctx, rw, req)
		}
	}
}

func ContextOptionalBasicAuth(ctx context.Context) *OptionalBasicAuth {
	if v := ctx.Value(optionalBasicAuthKey); v != nil {
		return v.(*OptionalBasicAuth)
	}
	return nil
}

// BasicController implements the BasicAuth resource.
type BasicController struct {
	*goa.Controller
}

// NewBasicController creates a BasicAuth controller.
func NewBasicController(service *goa.Service) *BasicController {
	return &BasicController{Controller: service.NewController("BasicController")}
}

// Secure runs the secure action.
func (c *BasicController) Secure(ctx *app.SecureBasicContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Optional runs the secure action.
func (c *BasicController) Optional(ctx *app.OptionalBasicContext) error {
	fmt.Println("in Optional", ContextOptionalBasicAuth(ctx))
	res := &app.Success{OK: true}
	return ctx.OK(res)
}

// Unsecure runs the unsecure action.
func (c *BasicController) Unsecure(ctx *app.UnsecureBasicContext) error {
	res := &app.Success{OK: true}
	return ctx.OK(res)
}
