## What is this?
Waiter is a simple http server that serves file. This is for learning purposes.

## Why this?
Built this to understand more about how requests are parsed and response formed in HTTP.

## Limitations
This serves just txt and html files and currently can return only 3 status codes. `200`, `404` and `500`.

## How to run the code
1. Clone this repository
2. From the root directory, run `go run main.go localhost:8080` to start the server
3. Open your web browser and make a request to `localhost:8080/file.txt` or `localhost:8080/file.html` 

