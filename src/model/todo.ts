import * as mongoose from "mongoose";

const TodoSchema = new mongoose.Schema({
    text: {
        type: String,
        required: true,
    },
    completed: {
        type: Boolean,
        required: true
    }
});

const Todo = mongoose.model("Todo", TodoSchema, "todos");

export default Todo;