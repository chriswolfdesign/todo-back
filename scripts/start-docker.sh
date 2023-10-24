#! /bin/bash

docker run --network todo --name todo-back -d -p 3000:3000 todo-back
