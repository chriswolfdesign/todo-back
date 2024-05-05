import {Request, Response} from "express";

export default function indexHandler(req: Request, res: Response) {
    res.send("Hello world!")
}