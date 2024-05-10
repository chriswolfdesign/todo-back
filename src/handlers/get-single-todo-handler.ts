import { Request, Response } from "express";
import Todo from "../model/todo";
import statusCode from "http-status-codes";

export default async function getSingleTodoHandler(
  req: Request,
  res: Response,
) {
  const id = req.params.id;

  const todo = await Todo.findOne({ _id: id });

  if (todo === null) {
    res.setHeader("Content-Type", "application/json");
    res.status(statusCode.BAD_REQUEST).json({
      message: "could not find todo with that id",
    });
    return;
  }

  res.setHeader("Content-Type", "application/json");
  res.status(statusCode.OK).json(todo);
}

