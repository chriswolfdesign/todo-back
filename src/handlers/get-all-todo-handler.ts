import Todo from "../model/todo";
import statusCode from "http-status-codes";
import {Request, Response} from "express";

export default async function getAllTodoHandler(req: Request, res: Response) {
    const allTodos = await Todo.find();

    return res.status(statusCode.OK).json({
        todos: allTodos
    });
}