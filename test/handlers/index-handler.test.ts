import { createRequest, createResponse } from "node-mocks-http";
import indexHandler from "../../src/handlers/index-handler";
import statusCode from "http-status-codes";

describe("Index handler tests", () => {
  it("should return 'Hello world!' and 200 status", () => {
    const mockRequest = createRequest({
      method: "GET",
      url: "/",
    });

    const mockResponse = createResponse();

    indexHandler(mockRequest, mockResponse);
    expect(mockResponse.statusCode).toBe(statusCode.OK);
    expect(mockResponse._getData()).toBe("Hello world!");
  });
});

