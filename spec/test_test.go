package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type AppConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type User struct {
	Username string
	Password string
	Request  string
}

var app = AppConfig{
	Host:     "localhost",
	Port:     5432,
	Username: "postgres",
	Password: "password",
}

var _ = Describe("Reading code systems from the application", func() {
	BeforeEach(func() {
		FetchApi()
	})
	Context("GIVEN the service is started", func() {
		var appStart *app // Assuming App is a type defined in phinvads
		Context("GIVEN the database loads successfully", func() {
			BeforeEach(func() {
				// Initialize app here if needed
			})
			Context("GIVEN a valid user with access to the API", func() {
				var user *User
				BeforeAll(func() {
					user = &User{
						Username: "testuser",
						Password: "testpassword",
					}
				})
				Context("WHEN the user requests all code systems", func() {
					user.Request = "/api/code-systems"
					It("THEN return a list of all code systems", func() {
						Expect(codestystemslist).To(Equal("codestystemslist"))
					})
				})
				Context("WHEN the user requests a specific code system by OID", func() {
					user.Request = "/api/code-systems/{oid}"
					It("THEN return the desired code system", func() {
						Expect(codesystem).To(Equal("codesystem"))
					})
				})
				Context("WHEN the user requests a code system using an INVALID OID", func() {
					user.Request = "/api/code-systems/zero"
					It("THEN return a 400 error", func() {
						Expect(statuscode).To(Equal(400))
					})
				})
				Context("WHEN the user requests a non-existant endpoint", func() {
					user.Request = "/api/invalid-endpoint"
					It("THEN return a 404 error", func() {
						Expect(statuscode).To(Equal(404))
					})
				})
			})
		})
	})
})
