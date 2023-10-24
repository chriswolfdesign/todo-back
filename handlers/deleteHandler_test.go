package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"todo-back/db/dbfakes"
	"todo-back/handlers"
)

var _ = Describe("DeleteHandler", func() {
	var (
		e   *echo.Echo
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
		dm  dbfakes.FakeDatabaseManagerInterface
	)

	BeforeEach(func() {
		e = echo.New()

		req = httptest.NewRequest(http.MethodDelete, "/todo/1", nil)

		rec = httptest.NewRecorder()
		ctx = e.NewContext(req, rec)

		ctx.SetPath("/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

	})

	When("DB handler correctly deletes the todo item", func() {
		BeforeEach(func() {
			dm.DeleteTodoReturns(nil)
		})

		It("sends a 200", func() {
			handler := handlers.DeleteHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(rec.Body.String()).To(ContainSubstring("successfully"))
		})
	})

	When("DB cannot delete the todo item", func() {
		BeforeEach(func() {
			dm.DeleteTodoReturns(fmt.Errorf("some error"))
		})

		It("sends a 400", func() {
			handler := handlers.DeleteHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).To(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusForbidden))
			Expect(rec.Body.String()).To(ContainSubstring("unable to delete"))
		})
	})
})
