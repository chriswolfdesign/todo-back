import dotenv from "dotenv";
import express, {Express, Request, Response} from "express";

dotenv.config();

const app: Express = express();
const port = process.env.PORT || 2000;

app.get("/", (req: Request, res: Response) => {
    res.send("Hello world!")
});

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});