import dotenv from "dotenv";
import express, {Express, Request, Response} from "express";
import indexHandler from "./handlers/index-handler";

dotenv.config();

const app: Express = express();
const port = process.env.PORT || 2000;

app.get("/", indexHandler);

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});