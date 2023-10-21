package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"todo-back/handlers"

	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetHandler", func() {
	var (
		e   *echo.Echo
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
	)

	BeforeEach(func() {
		e = echo.New()
		req = httptest.NewRequest(http.MethodGet, "/test", nil)
		rec = httptest.NewRecorder()
		ctx = e.NewContext(req, rec)
	})

	When("Testing the test handler", func() {
		It("returns the expected string with 200", func() {
			err := handlers.GetTest(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Result().StatusCode).To(Equal(http.StatusOK))
			Expect(rec.Body.String()).To(Equal("Echo server running correctly"))
		})
	})
})
