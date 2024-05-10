import { Request, Response } from "express";
import Todo from "../model/todo";
import statusCode from "http-status-codes";

export default async function deleteTodoHandler(req: Request, res: Response) {
  const id = req.params.id;

  const result = await Todo.findByIdAndDelete(id);

  if (result === null) {
    res.status(statusCode.BAD_REQUEST).json({
      message: "could not delete todo item",
    });
    return;
  }

  res.status(statusCode.OK).json({
    message: "todo item successfully deleted",
  });
}

