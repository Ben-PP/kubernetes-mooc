# crate-todo.sh
Simple script that calls the todo-backend running at `TODO_URL` with post request and creates todo with `TODO_TEXT` as the content. If the `TODO_TEXT` is empty, it fetches random wikipedia article url and uses that as content.

## Usage
To run this script, you can use docker.
```bash
docker run --rm -e TODO_TEXT="Remember to push this image" -e TODO_URL="http://localhos:8000" benpp/create-todo:latest
```
