import {Response, Request} from "express";
import Todo from "../model/todo";
import statusCode from "http-status-codes";

export default async function createTodoHandler(req: Request, res: Response) {
    const result = await Todo.create({
        text: req.body.text,
        completed: req.body.completed,
    });

    res.setHeader("Content-Type", "application/json");

    if (result === null) {
        res.status(statusCode.BAD_REQUEST).json({
            message: "failed to create todo"
        });
        return;
    }

    res.status(statusCode.OK).json(result);
};