package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"todo-back/db/dbfakes"
	"todo-back/handlers"
	"todo-back/model"

	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetAllHandler", func() {
	var (
		e   *echo.Echo
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
		dm  dbfakes.FakeDatabaseManagerInterface
	)

	BeforeEach(func() {
		e = echo.New()
		req = httptest.NewRequest(http.MethodGet, "/todos", nil)
		rec = httptest.NewRecorder()
		ctx = e.NewContext(req, rec)
	})

	When("DB handler returns an error", func() {
		BeforeEach(func() {
			dm = dbfakes.FakeDatabaseManagerInterface{}
			dm.GetAllTodosReturns(nil, fmt.Errorf("some error"))
		})

		It("sends the user an error status code and message", func() {
			handler := handlers.GetAllHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("some error"))
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			Expect(rec.Body.String()).To(Equal("Unable to retrieve list of todo items"))
		})
	})

	When("DB handler returns valid todo items", func() {
		BeforeEach(func() {
			todos := []model.Todo{
				{
					ID:        1,
					Text:      "first item",
					Completed: true,
				},
				{
					ID:        2,
					Text:      "second item",
					Completed: false,
				},
			}

			dm = dbfakes.FakeDatabaseManagerInterface{}
			dm.GetAllTodosReturns(todos, nil)
		})

		It("sends the user the expected todo items with a positive status", func() {
			handler := handlers.GetAllHandler(ctx, &dm)
			err := handler(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(rec.Body.String()).To(ContainSubstring("first item"))
			Expect(rec.Body.String()).To(ContainSubstring("second item"))
		})
	})
})
