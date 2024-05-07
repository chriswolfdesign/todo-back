import Todo from "../../src/model/todo";
import {createRequest, createResponse} from "node-mocks-http";
import createTodoHandler from "../../src/handlers/create-todo-handler";
import statusCode from "http-status-codes";
import {mock} from "sinon";

describe("create todo handler tests", () => {
    describe("it could not create the todo item", () => {
        it("returns bad request with an error message", async () => {
            let TodoMock = mock(Todo);
            TodoMock.expects("create").returns(null);

            const mockRequest = createRequest({
                method: "POST",
                url: "/todos",
                body: {
                    text: "did not create",
                    completed: false
                }
            });

            const mockResponse = createResponse();

            await createTodoHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "failed to create todo"
            });

            TodoMock.restore();
        });
    });

    describe("it could create the todo item", () => {
        it("returns status ok with the created todo", async () => {
            let doc = {
                _id: "6637eaa9a2b84ebd57b55028",
                text: "testing",
                completed: true
            }
            let TodoMock = mock(Todo);
            TodoMock.expects("create").returns(doc);

            const mockRequest = createRequest({
                method: "POST",
                url: "/todos",
                body: {
                    text: "testing",
                    completed: true
                }
            });

            const mockResponse = createResponse();

            await createTodoHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.OK);
            expect(mockResponse._getJSONData()).toStrictEqual(doc);

            TodoMock.restore();
        });
    });
});