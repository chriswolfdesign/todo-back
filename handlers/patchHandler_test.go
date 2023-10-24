package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"todo-back/db/dbfakes"
	"todo-back/handlers"
	"todo-back/model"
)

var _ = Describe("PatchHandler", func() {
	var (
		e   *echo.Echo
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
		dm  dbfakes.FakeDatabaseManagerInterface
	)

	BeforeEach(func() {
		e = echo.New()

		req = httptest.NewRequest(http.MethodPatch, "/todo/1", strings.NewReader(`{"text": "new text", "completed": false}`))

		req.Header.Add("Content-Type", "application/json")

		rec = httptest.NewRecorder()
		ctx = e.NewContext(req, rec)

		ctx.SetPath("/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		dm = dbfakes.FakeDatabaseManagerInterface{}
	})

	When("DB handler correctly updates the todo item", func() {
		BeforeEach(func() {
			todo := model.Todo{
				ID:        1,
				Text:      "new text",
				Completed: false,
			}

			dm.UpdateTodoReturns(&todo, nil)
		})

		It("sends the todo with a 200", func() {
			handler := handlers.PatchHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(rec.Body.String()).To(ContainSubstring("new text"))
			Expect(rec.Body.String()).To(ContainSubstring("false"))
		})
	})

	When("DB is uanble to update todo in database", func() {
		BeforeEach(func() {
			dm.UpdateTodoReturns(nil, fmt.Errorf("some error"))
		})

		It("returns an error message with a 400", func() {
			handler := handlers.PatchHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).To(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			Expect(rec.Body.String()).To(ContainSubstring("unable to update todo"))
		})
	})
})
