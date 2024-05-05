import Todo from "../../src/model/todo";

describe("Get all todo handler tests", async () => {
    describe("there are no todos", async () => {
        Todo.find = async () => {
            return {};
        }
    });
});