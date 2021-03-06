package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// BasicAuth defines a security scheme using basic authentication.
var BasicAuth = BasicAuthSecurity("basic_auth")

// OptionalBasicAuth defines is all user access
var OptionalBasicAuth = BasicAuthSecurity("optional_basic_auth")

var _ = Resource("basic", func() {
	Description("This resource uses basic auth to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(BasicAuth)

	Action("secure", func() {
		Description("This action is secure with the basic_auth scheme")
		Routing(GET("/basic"))
		Response(OK)
		Response(Unauthorized)
	})
	Action("optional", func() {
		Description("This action is optional secure with the basic_auth scheme")
		Security(OptionalBasicAuth)
		Routing(GET("/basic/optional"))
		Response(OK)
	})

	Action("unsecure", func() {
		Description("This action does not require auth")
		Routing(GET("/basic/unsecure"))
		NoSecurity()
		Response(OK)
	})
})
