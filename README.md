# Super Super Simple Mock Server based on File and Docker

## TL;DR
You can set up the mock server simply by placing a JSON file in the `/responses` directory below.

## How To Use
Clone the repository and start Docker with the command below.
Replace `PORT` with the port number you want to use, or specify it using an environment variable.
```bash
# cd mock-api
# PORT=8080
docker build -t mock-api .
docker run -d -p ${PORT}:8080 -v $(pwd)/responses:/responses mock-api
```

if you use PowerShell on Windows,
```ps1
# cd mock-api
# $PORT=8080
$PWD="$((pwd).Path.Replace('\', '/'))"
docker build -t mock-api .
docker run -d -p ${PORT}:8080 -v ${PWD}/responses:/responses mock-api
```

Since the files under `/responses` are mounted in the container, you can continue testing by modifying the JSON files even while the container is running.

## Example
- if you want to test `http://localhost:8080/v2/index`  
create a `v2` directory under `/responses` and place `index.json` inside it.
The file path should be /responses/v2/index.json.

- if you want to test `http://localhost:8080/v2` 
After that, place `/responses/v2.json` 

## Notes
- It only recognizes `*.json` files.
