import Todo from "../../src/model/todo";
import { createRequest, createResponse } from "node-mocks-http";
import getSingleTodoHandler from "../../src/handlers/get-single-todo-handler";
import statusCode from "http-status-codes";
import { mock } from "sinon";

describe("Get single todo handler tests", () => {
  describe("there is no todo with that id", () => {
    it("should return bad request with error message", async () => {
      let TodoMock = mock(Todo);
      TodoMock.expects("findOne").returns(null);

      const mockRequest = createRequest({
        method: "GET",
        url: "/todos",
        params: [
          {
            id: "1",
          },
        ],
      });

      const mockResponse = createResponse();

      await getSingleTodoHandler(mockRequest, mockResponse);
      expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
      expect(mockResponse._getJSONData()).toStrictEqual({
        message: "could not find todo with that id",
      });

      TodoMock.restore();
    });
  });

  describe("there is a todo with that id", () => {
    it("should return status ok with that todo", async () => {
      const doc = {
        _id: "6637e890ada4275ce9ec5a26",
        text: "testing",
        completed: true,
      };

      let TodoMock = mock(Todo);
      TodoMock.expects("findOne").returns(doc);

      const mockRequest = createRequest({
        method: "GET",
        url: "/todos",
        params: {
          id: "6637e890ada4275ce9ec5a26",
        },
      });

      const mockResponse = createResponse();

      await getSingleTodoHandler(mockRequest, mockResponse);

      expect(mockResponse.statusCode).toBe(statusCode.OK);
      expect(mockResponse._getJSONData()).toStrictEqual(doc);

      TodoMock.restore();
    });
  });
});

