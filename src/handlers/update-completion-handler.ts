import {Request, Response} from "express";
import statusCode from "http-status-codes";
import Todo from "../model/todo";

export default async function updateCompletionHandler(req: Request, res: Response) {
    const id = req.params.id;
    const completed = req.body.completed;

    res.setHeader("Content-Type", "application/json");

    if (id === undefined || completed === undefined) {
        res.status(statusCode.BAD_REQUEST).json({
            message: "ID url parameter and completed body field are required"
        });
        return;
    }

    const filter = {
        _id: id
    };

    const update = {
        completed
    };

    let result = await Todo.findOneAndUpdate(filter, update, {new: true});

    if (result === null) {
        res.status(statusCode.CONFLICT).json({
            message: "could not make update"
        });
        return;
    }

    res.status(statusCode.OK).json(result);
}