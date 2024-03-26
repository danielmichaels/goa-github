package design

import . "goa.design/goa/v3/dsl"

// errors function is used to define and register error types.
func errors() {
	Error("not-found", NotFound, "not found")
	Error("bad-request", BadRequest, "bad request")
	Error("server-error", ServerError, "server error")
}

var NotFound = Type("not-found", func() {
	Description("not-found indicates the resource matching the id does not exist.")
	Attribute("user", String, "Github username", func() {
		Example("rs")
	})
	Attribute("name", String, "Name of error", func() { Example("not found") })
	Attribute("message", String, "Error message", func() {
		Example("not found")
	})
	Attribute("detail", String, "Error details", func() {
		Example("Failed to determine machine information. Cannot continue.")
	})
	Required("message", "detail", "name")

})
var BadRequest = Type("bad-request", func() {
	Description("bad-request indicates the values provided are invalid")
	Attribute("name", String, "Name of error", func() { Example("bad request") })
	Attribute("message", String, "Error message", func() {
		Example("bad request")
	})
	Attribute("detail", String, "Error details", func() {
		Example("Failed to determine machine information. Cannot continue.")
	})
	Required("message", "detail", "name")
})
var ServerError = Type("server-error", func() {
	Description("server-error indicates the server encountered an error.")
	Attribute("name", String, "Name of error", func() { Example("internal server error") })
	Attribute("message", String, "Error message", func() {
		Example("bad request")
	})
	Attribute("detail", String, "Error details", func() {
		Example("Failed to determine machine information. Cannot continue.")
	})
	Required("message", "detail", "name")
})
