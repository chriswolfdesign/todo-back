import { mock } from "sinon";
import Todo from "../../src/model/todo";
import { createRequest, createResponse } from "node-mocks-http";
import deleteTodoHandler from "../../src/handlers/delete-todo-handler";
import statusCode from "http-status-codes";

describe("delete todo handler tests", () => {
  describe("todo item could not be deleted", () => {
    it("returns a bad request with an error message", async () => {
      let TodoMock = mock(Todo);
      TodoMock.expects("findByIdAndDelete").returns(null);

      const mockRequest = createRequest({
        method: "DELETE",
        url: "todos",
        params: {
          id: "6637e890ada4275ce9ec5a26",
        },
      });

      const mockResponse = createResponse();

      await deleteTodoHandler(mockRequest, mockResponse);
      expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
      expect(mockResponse._getJSONData()).toStrictEqual({
        message: "could not delete todo item",
      });

      TodoMock.restore();
    });
  });

  describe("todo item was successfully deleted", () => {
    it("returns status ok with a copy of the deleted todo item", async () => {
      const doc = {
        _id: "6637e890ada4275ce9ec5a26",
        text: "delete me",
        completed: true,
      };

      let TodoMock = mock(Todo);
      TodoMock.expects("findByIdAndDelete").returns(doc);

      const mockRequest = createRequest({
        method: "DELETE",
        url: "todos",
        params: {
          id: "6637e890ada4275ce9ec5a26",
        },
      });

      const mockResponse = createResponse();

      await deleteTodoHandler(mockRequest, mockResponse);
      expect(mockResponse.statusCode).toBe(statusCode.OK);
      expect(mockResponse._getJSONData()).toStrictEqual({
        message: "todo item successfully deleted",
      });

      TodoMock.restore();
    });
  });
});

