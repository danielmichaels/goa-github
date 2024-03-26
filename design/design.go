package design

import . "goa.design/goa/v3/dsl"

var _ = API("goa-github", func() {
	Title("Goa API Demo")
	Description("Goa API demonstration server that fetches and stores Github user data.")
	Server("server", func() {
		Description("development server")
		Host("localhost", func() { URI("http://localhost:9090") })
	})
})

var _ = Service("users_v1", func() {
	Description("Github users endpoints v1")
	HTTP(func() {
		Path("/v1/users")
	})
	errors()
	Method("listUsers", func() {
		Payload(func() {
			Attribute("view", String, "view to render", func() {
				Enum("default", "detailed")
			})
		})
		HTTP(func() {
			GET("/")
			Param("view")
			Response(StatusOK)
		})
		Result(CollectionOf(Users), func() { View("default") })
	})
	Method("retrieveUser", func() {
		Payload(func() {
			Attribute("username", String, "username of Github user to retrieve", func() {
				Example("rs")
				Default("rs")
			})
			Required("username")
		})
		HTTP(func() {
			GET("/{username}")
			Response(StatusOK)
		})
		Result(Users, func() { View("detailed") })
	})
	Method("createUser", func() {
		Payload(func() {
			Attribute("username", String, "username of Github user to create in database", func() {
				Example("rs")
				Default("rs")
			})
			Required("username")
		})
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		Result(User)
	})
	Method("deleteUser", func() {
		Payload(func() {
			Attribute("username", String, "username of Github user to delete from database", func() {
				Example("rs")
				Default("rs")
			})
			Required("username")
		})
		HTTP(func() {
			DELETE("/{username}")
			Response(StatusOK)
		})
		Result(Empty)
	})
})

var User = Type("User", func() {
	Description("Github user object")
	Attribute("username", String, "Username of Github user", func() {
		Example("rs")
	})
	Attribute("name", String, "Name of Github user", func() {
		Example("Olivier Poitrey")
	})
	Attribute("followers", Int, "Number of Github users following", func() {
		Example(4467)
	})
	Attribute("repos", Int, "Number of public Github repositories this user owns", func() {
		Example(116)
	})
	Attribute("company", String, "User supplied employer name", func() {
		Example("Netflix")
	})
	Attribute("location", String, "User supplied location", func() {
		Example("Silicon Valley, California, USA")
	})
	Attribute("blog", String, "User supplied website/blog", func() {
		Example("https://nextdns.io")
	})
})

var Users = ResultType("application/vnd.goa-github.users", func() {
	TypeName("Users")
	Description("A list of Users")
	Reference(User)

	Attributes(func() {
		Attribute("username")
		Attribute("name")
		Attribute("followers")
		Attribute("repos")
		Attribute("company")
		Attribute("location")
		Attribute("blog")
	})

	// default is returned when listing users or when retrieving an individual user
	View("default", func() {
		Attribute("username")
		Attribute("name")
		Attribute("followers")
		Attribute("repos")
	})

	// detailed can be returned when retrieving an individual user
	View("detailed", func() {
		Attribute("username")
		Attribute("name")
		Attribute("followers")
		Attribute("repos")
		Attribute("company")
		Attribute("location")
		Attribute("blog")
	})
	Required(
		"username",
		"name",
		"followers",
		"repos",
	)
})
