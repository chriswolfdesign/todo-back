import Todo from "../../src/model/todo";
import { createRequest, createResponse } from "node-mocks-http";
import getAllTodoHandler from "../../src/handlers/get-all-todo-handler";
import statusCode from "http-status-codes";
import { mock } from "sinon";

describe("Get all todo handler tests", () => {
  describe("there are no todos", () => {
    it("should return an empty array", async () => {
      let TodoMock = mock(Todo);
      TodoMock.expects("find").returns([]);

      const mockRequest = createRequest({
        method: "GET",
        url: "/todos",
      });

      const mockResponse = createResponse();

      await getAllTodoHandler(mockRequest, mockResponse);
      expect(mockResponse.statusCode).toBe(statusCode.OK);
      expect(mockResponse._getJSONData()).toStrictEqual({
        todos: [],
      });

      TodoMock.restore();
    });
  });

  describe("there is one todo", () => {
    it("should have one todo in payload", async () => {
      const doc = [
        {
          _id: "6637e890ada4275ce9ec5a26",
          text: "bubble",
          completed: false,
        },
      ];

      let TodoMock = mock(Todo);

      TodoMock.expects("find").returns(doc);

      const mockRequest = createRequest({
        method: "GET",
        url: "/todos",
      });

      const mockResponse = createResponse();

      await getAllTodoHandler(mockRequest, mockResponse);
      expect(mockResponse.statusCode).toBe(statusCode.OK);
      expect(mockResponse._getJSONData()).toStrictEqual({ todos: doc });

      TodoMock.restore();
    });
  });

  describe("there are multiple todos", () => {
    it("should have all todos in the payload", async () => {
      const doc = [
        {
          _id: "6637eaa9a2b84ebd57b55027",
          text: "fizz",
          completed: false,
        },
        {
          _id: "6637eaa9a2b84ebd57b55028",
          text: "buzz",
          completed: true,
        },
      ];

      let TodoMock = mock(Todo);
      TodoMock.expects("find").returns(doc);

      const mockRequest = createRequest({
        method: "GET",
        url: "/todos",
      });

      const mockResponse = createResponse();

      await getAllTodoHandler(mockRequest, mockResponse);
      expect(mockResponse.statusCode).toBe(statusCode.OK);
      expect(mockResponse._getJSONData()).toStrictEqual({ todos: doc });

      TodoMock.restore();
    });
  });
});

