import dotenv from "dotenv";
import express, {Express, Request, Response} from "express";
import indexHandler from "./handlers/index-handler";
import healthCheckHandler from "./handlers/health-check-handler";
import * as mongoose from "mongoose";
import Todo from "./model/todo";
import statusCode from "http-status-codes";
import getAllTodoHandler from "./handlers/get-all-todo-handler";
import getSingleTodoHandler from "./handlers/get-single-todo-handler";

dotenv.config();

const app: Express = express();
const port = process.env.PORT || 2000;

mongoose.connect("mongodb://127.0.0.1:27017/local").then(r => console.log("Connected to database"));

app.get("/", indexHandler);
app.get("/check", healthCheckHandler);
app.get("/todos", getAllTodoHandler);
app.get("/todos/:id", getSingleTodoHandler);

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});