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
	"todo-back/model"
)

var _ = Describe("GetHandler", func() {
	var (
		e   *echo.Echo
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
		dm  dbfakes.FakeDatabaseManagerInterface
	)

	BeforeEach(func() {
		e = echo.New()

		req = httptest.NewRequest(http.MethodGet, "/todo/1", nil)

		rec = httptest.NewRecorder()
		ctx = e.NewContext(req, rec)

		ctx.SetPath("/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		dm = dbfakes.FakeDatabaseManagerInterface{}
	})

	When("DB handler correctly returns the todo item", func() {
		BeforeEach(func() {
			todo := model.Todo{
				ID:        1,
				Text:      "first item",
				Completed: false,
			}

			dm.GetTodoReturns(&todo, nil)
		})

		It("sends the todo with a positive status", func() {
			handler := handlers.GetHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(rec.Body.String()).To(ContainSubstring("first item"))
		})
	})

	When("an ID is requested that does not exist", func() {
		BeforeEach(func() {

			todo := model.Todo{}

			dm.GetTodoReturns(&todo, fmt.Errorf("some error"))
		})

		It("sends a string informing the id was not found", func() {
			handler := handlers.GetHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).To(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusForbidden))
		})
	})
})
