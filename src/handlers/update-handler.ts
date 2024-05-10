import { Request, Response } from "express";
import statusCode from "http-status-codes";
import Todo from "../model/todo";

export default async function updateHandler(req: Request, res: Response) {
  res.setHeader("Content-Type", "application/json");

  if (!isValidRequest(req)) {
    res.status(statusCode.BAD_REQUEST).json({
      message: "ID url parameter and text or completed body field are required",
    });
    return;
  }

  const filter = {
    _id: req.params.id,
  };

  let result = await Todo.findOneAndUpdate(filter, req.body, { new: true });

  if (result === null) {
    res.status(statusCode.CONFLICT).json({
      message: "could not make update",
    });
    return;
  }

  res.status(statusCode.OK).json(result);
}

function isValidRequest(req: any): boolean {
  return (
    req.params.id !== undefined &&
    (req.body.text !== undefined || req.body.completed !== undefined)
  );
}

