import {Request, Response} from "express";
import statusCode from "http-status-codes";

export default function indexHandler(req: Request, res: Response) {
    res.statusCode = statusCode.OK;
    res.send("Hello world!")
}