import {Request, Response} from "express";

export default function healthCheckHandler(req: Request, res: Response) {
    res.setHeader("Content-Type", "application/json");
    res.statusCode = 200;
    res.json({
        message: "healthy"
    });
}
