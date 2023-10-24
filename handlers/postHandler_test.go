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

var _ = Describe("PostHandler", func() {
	var (
		e   *echo.Echo
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
		dm  dbfakes.FakeDatabaseManagerInterface
	)

	BeforeEach(func() {
		e = echo.New()

		req = httptest.NewRequest(http.MethodPost, "/todo", strings.NewReader(`{"text": "unit tests are fun", "completed": true}`))

		req.Header.Add("Content-Type", "application/json")

		rec = httptest.NewRecorder()
		ctx = e.NewContext(req, rec)

		dm = dbfakes.FakeDatabaseManagerInterface{}
	})

	When("DB handler successfully adds todo to database", func() {
		BeforeEach(func() {
			todo := model.Todo{
				ID:        5,
				Text:      "unit tests are fun",
				Completed: true,
			}

			dm.CreateTodoReturns(&todo, nil)
		})

		It("returns the todo created with a 200", func() {
			handler := handlers.CreateHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(rec.Body.String()).To(ContainSubstring("unit tests are fun"))
		})
	})

	When("Db is unable to add todo to database", func() {
		BeforeEach(func() {
			dm.CreateTodoReturns(nil, fmt.Errorf("unable to parse todo response"))
		})

		It("returns an error message with a 400", func() {
			handler := handlers.CreateHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).To(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			Expect(rec.Body.String()).To(ContainSubstring("unable to parse todo response"))
		})
	})
})
