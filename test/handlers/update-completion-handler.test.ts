import {createRequest, createResponse} from "node-mocks-http";
import updateCompletionHandler from "../../src/handlers/update-completion-handler";
import statusCode from "http-status-codes";
import {mock} from "sinon";
import Todo from "../../src/model/todo";

describe("update completion handler tests", () => {
    describe("id is missing from the url", () => {
        it("returns bad request and error message", async () => {
            const mockRequest = createRequest({
                method: "PUT",
                url: "/todos",
                body: {
                    completed: false
                }
            });

            const mockResponse = createResponse();

            await updateCompletionHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "ID url parameter and completed body field are required"
            });
        });
    });

    describe("completion field is missing from the body", () => {
        it("returns bad request and error message", async () => {
            const mockRequest = createRequest({
                method: "PUT",
                url: "/todos",
                params: {
                    id: "6637e890ada4275ce9ec5a26"
                }
            });

            const mockResponse = createResponse();

            await updateCompletionHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "ID url parameter and completed body field are required"
            });
        });
    });

    describe("both id and completion are missing", () => {
        it("returns bad request and error message", async () => {
            const mockRequest = createRequest({
                method: "PUT",
                url: "/todos"
            });

            const mockResponse = createResponse();

            await updateCompletionHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "ID url parameter and completed body field are required"
            });
        });
    });

    describe("could not update todo item", () => {
        it("returns conflict status and error message", async () => {
            let TodoMock = mock(Todo);
            TodoMock.expects("findOneAndUpdate").returns(null);

            const mockRequest = createRequest({
                method: "PUT",
                url: "/todos",
                params: {
                    id: "6637e890ada4275ce9ec5a26"
                },
                body: {
                    completed: false
                }
            });

            const mockResponse = createResponse();

            await updateCompletionHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.CONFLICT);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "could not make update"
            });

            TodoMock.restore();
        });
    });

    describe("successfully updates todo item", () => {
        it("returns status ok with updated todo item", async () => {
            const doc = {
                _id: "6637e890ada4275ce9ec5a26",
                text: "testing",
                completed: true,
            };

            let TodoMock = mock(Todo);
            TodoMock.expects("findOneAndUpdate").returns(doc);

            const mockRequest = createRequest({
                method: "PUT",
                url: "/todos",
                params: {
                    id: "6637e890ada4275ce9ec5a26"
                },
                body: {
                    completed: true
                }
            });

            const mockResponse = createResponse();

            await updateCompletionHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.OK);
            expect(mockResponse._getJSONData()).toStrictEqual(doc);
        });
    });
});