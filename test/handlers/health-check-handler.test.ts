import {createRequest, createResponse} from "node-mocks-http";
import healthCheckHandler from "../../src/handlers/health-check-handler";
import {json} from "express";

describe("Health check handler tests", () => {
    it("should return 200 status with healthy message", () => {
        const mockRequest = createRequest({
            method: "GET",
            url: "/check",
        });

        const mockResponse = createResponse();

        healthCheckHandler(mockRequest, mockResponse);
        expect(mockResponse.statusCode).toBe(200);
        expect(mockResponse._getJSONData()).toStrictEqual({
            message: "healthy"
        });
    });
});