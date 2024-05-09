import {createRequest, createResponse} from "node-mocks-http";
import updateHandler from "../../src/handlers/update-handler";
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

            await updateHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "ID url parameter and text or completed body field are required"
            });
        });
    });

    describe("text and completion field is missing from the body", () => {
        it("returns bad request and error message", async () => {
            const mockRequest = createRequest({
                method: "PUT",
                url: "/todos",
                params: {
                    id: "6637e890ada4275ce9ec5a26"
                }
            });

            const mockResponse = createResponse();

            await updateHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "ID url parameter and text or completed body field are required"
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

            await updateHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.BAD_REQUEST);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "ID url parameter and text or completed body field are required"
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

            await updateHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.CONFLICT);
            expect(mockResponse._getJSONData()).toStrictEqual({
                message: "could not make update"
            });

            TodoMock.restore();
        });
    });

    describe("successfully updates todo item when text is updated", () => {
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

            await updateHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.OK);
            expect(mockResponse._getJSONData()).toStrictEqual(doc);

            TodoMock.restore();
        });
    });

    describe("successfully updates todo item when completed is updated", () => {
        it("returns status ok with updated todo item", async () => {
            const doc = {
                _id: "6637e890ada4275ce9ec5a26",
                text: "testing",
                completed: false
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
                    completed: false
                }
            });

            const mockResponse = createResponse();

            await updateHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.OK);
            expect(mockResponse._getJSONData()).toStrictEqual(doc);

            TodoMock.restore();
        });
    });

    describe("successfully updates todo item when both fields are updated", () => {
        it("returns stats ok with updated todo item", async () => {
            const doc = {
                _id: "6637e890ada4275ce9ec5a26",
                text: "testing",
                completed: false
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
                    text: "testing",
                    completed: false
                }
            });

            const mockResponse = createResponse();

            await updateHandler(mockRequest, mockResponse);
            expect(mockResponse.statusCode).toBe(statusCode.OK);
            expect(mockResponse._getJSONData()).toStrictEqual(doc);

            TodoMock.restore();
        });
    });
});