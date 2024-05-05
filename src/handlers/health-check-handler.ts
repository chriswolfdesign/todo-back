import {Request, Response} from "express";
import statusCode from "http-status-codes";

export default function healthCheckHandler(req: Request, res: Response) {
    res.setHeader("Content-Type", "application/json");
    res.statusCode = statusCode.OK;
    res.json({
        message: "healthy"
    });
}
