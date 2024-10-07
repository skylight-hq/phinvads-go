package main_test

import {
	"github/onsi/ginkgo/v2"
}

ENDPOINTS = [
	"endpoint/1",

]

var _ = Describe("[PHIN VADS API]", func() {	
	Context("Given a running phinvads-go environment that has been properly populated", func() {
		BeforeAll(func() {
			// spin up environment if it's not already running
			// populate database
		})


	})
})

// package main_test

// import (
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// )

// type AppConfig struct {
// 	Host     string
// 	Port     int
// 	Username string
// 	Password string
// }

// type User struct {
// 	Username string
// 	Password string
// 	Request  string
// }

// var app = AppConfig{
// 	Host:     "localhost",
// 	Port:     5432,
// 	Username: "postgres",
// 	Password: "password",
// }

// var _ = Describe("Reading code systems from the application", func() {
// 	BeforeEach(func() {
// 		FetchApi()
// 	})
// 	Context("GIVEN the service is started", func() {
// 		var appStart *app // Assuming App is a type defined in phinvads
// 		Context("GIVEN the database loads successfully", func() {
// 			BeforeEach(func() {
// 				// Initialize app here if needed
// 			})
// 			Context("GIVEN a valid user with access to the API", func() {
// 				var user *User
// 				BeforeAll(func() {
// 					user = &User{
// 						Username: "testuser",
// 						Password: "testpassword",
// 					}
// 				})
// 				Context("WHEN the user requests all code systems", func() {
// 					user.Request = "/api/code-systems"
// 					It("THEN return a list of all code systems", func() {
// 						Expect(codestystemslist).To(Equal("codestystemslist"))
// 					})
// 				})
// 				Context("WHEN the user requests a specific code system by OID", func() {
// 					user.Request = "/api/code-systems/{oid}"
// 					It("THEN return the desired code system", func() {
// 						Expect(codesystem).To(Equal("codesystem"))
// 					})
// 				})
// 				Context("WHEN the user requests a code system using an INVALID OID", func() {
// 					user.Request = "/api/code-systems/zero"
// 					It("THEN return a 400 error", func() {
// 						Expect(statuscode).To(Equal(400))
// 					})
// 				})
// 				Context("WHEN the user requests a non-existant endpoint", func() {
// 					user.Request = "/api/invalid-endpoint"
// 					It("THEN return a 404 error", func() {
// 						Expect(statuscode).To(Equal(404))
// 					})
// 				})
// 			})
// 		})
// 	})
// })

package main_test

import {
	"github/onsi/ginkgo/v2"
}

// These endpoints should respond to simple GET requests with 200 and a response body
const GETTABLE_ENDPOINTS = [
	"/",
	"/assets/",
	"/api",
	"/api/code-systems",
	"/api/code-system-concepts",
	"/api/value-sets",
	"/api/views",
	"/load-hot-topics",
	"/search"
]

// These endpoints have show-type response bodies
// The ID to be passed in to these specs is acquired from the corresponding index action
const GETTABLE_SHOW_ENDPOINTS_WITH_PREDICTABLE_IDS = [
	"code-systems", "code-system-concepts", "value-sets", "views"
]

func IsValidGetRequest(response) {
	return (
		response.code == '200' &&
		response.body.length != 0
		JSON.parse(response.body).errors.length == 0
	)
}

var _ = Describe("[PHIN VADS API]", func() {	
	Context("Given an API environment that has been populated with data", func() {
		BeforeEach(func {
			// load data into the environment
			// dependent upon sample dataset, may require a BeforeAll
		})

		Context("Given that I do not have an API key", func() {
			Context("When I GET the 'create API key' endpoint", func() {
				// pending
				// describe steps for getting an API key
			})

			Context("When I GET any other valid endpoint", func() {
				// generate specs for each possible endpoint
				// I should get a 400
				// response body should say "You need an API key idiot"
			})
		})

		Context("Given that I have a valid API key", func() {
			BeforeEach(func() {
				// generate an API key
			})

			Context("When I attempt to GET an invalid endpoint", func() {
				BeforeEach(func() {
					response = doTheThing("/known/invalid/endpoint");
				})

				It("Then the request should fail", func() {
					Expect(response.code).To(Equal('404'));
					Expect(response.body).To(BeZero()));
				})
			})

			Context("When I GET /" func() {
				BeforeEach(func() {
					response = doTheThing("/")
				})



			})

			Context("When I GET /assets/", func() {
				BeforeEach(func() {
					response = doTheThing("/assets/")
				})

				It("Then the response should succeed")
				It("Then the response should be compressed")
				It("Then the response should be UTF-8 encoded")
			})

		})
	})
	
})

// package main_test

// import (
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// )

// type AppConfig struct {
// 	Host     string
// 	Port     int
// 	Username string
// 	Password string
// }

// type User struct {
// 	Username string
// 	Password string
// 	Request  string
// }

// var app = AppConfig{
// 	Host:     "localhost",
// 	Port:     5432,
// 	Username: "postgres",
// 	Password: "password",
// }

// var _ = Describe("Reading code systems from the application", func() {
// 	BeforeEach(func() {
// 		FetchApi()
// 	})
// 	Context("GIVEN the service is started", func() {
// 		var appStart *app // Assuming App is a type defined in phinvads
// 		Context("GIVEN the database loads successfully", func() {
// 			BeforeEach(func() {
// 				// Initialize app here if needed
// 			})
// 			Context("GIVEN a valid user with access to the API", func() {
// 				var user *User
// 				BeforeAll(func() {
// 					user = &User{
// 						Username: "testuser",
// 						Password: "testpassword",
// 					}
// 				})
// 				Context("WHEN the user requests all code systems", func() {
// 					user.Request = "/api/code-systems"
// 					It("THEN return a list of all code systems", func() {
// 						Expect(codestystemslist).To(Equal("codestystemslist"))
// 					})
// 				})
// 				Context("WHEN the user requests a specific code system by OID", func() {
// 					user.Request = "/api/code-systems/{oid}"
// 					It("THEN return the desired code system", func() {
// 						Expect(codesystem).To(Equal("codesystem"))
// 					})
// 				})
// 				Context("WHEN the user requests a code system using an INVALID OID", func() {
// 					user.Request = "/api/code-systems/zero"
// 					It("THEN return a 400 error", func() {
// 						Expect(statuscode).To(Equal(400))
// 					})
// 				})
// 				Context("WHEN the user requests a non-existant endpoint", func() {
// 					user.Request = "/api/invalid-endpoint"
// 					It("THEN return a 404 error", func() {
// 						Expect(statuscode).To(Equal(404))
// 					})
// 				})
// 			})
// 		})
// 	})
// })
