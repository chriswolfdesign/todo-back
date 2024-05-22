import dotenv from "dotenv";
import express, { Express } from "express";
import indexHandler from "./handlers/index-handler";
import healthCheckHandler from "./handlers/health-check-handler";
import * as mongoose from "mongoose";
import getAllTodoHandler from "./handlers/get-all-todo-handler";
import getSingleTodoHandler from "./handlers/get-single-todo-handler";
import createTodoHandler from "./handlers/create-todo-handler";
import updateHandler from "./handlers/update-handler";
import deleteTodoHandler from "./handlers/delete-todo-handler";
import cors from "cors";

dotenv.config();

const app: Express = express();
const port = process.env.PORT || 8080;
const mongoUrl = process.env.MONGO_URL || "mongodb://127.0.0.1:27017/local";

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cors());

mongoose.connect(mongoUrl).then((_) => console.log("Connected to database"));

app.get("/", indexHandler);
app.get("/check", healthCheckHandler);
app.get("/todos", getAllTodoHandler);
app.get("/todos/:id", getSingleTodoHandler);

app.post("/todos", createTodoHandler);

app.put("/todos/:id", updateHandler);

app.delete("/todos/:id", deleteTodoHandler);

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
