## TL;DR
You can set up the mock server simply by placing a JSON file in the `/responses` directory below.

## Limitations
Because it is a very simple configuration, the following cases are not currently supported.
Please consider using `JSON-server` or `postman`. Or, fork the repository and implement it.

- Multiple response patterns
- Status codes other than `200`
- Conditional branching according to body or header
- Responses other than JSON

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
```powershell
# cd mock-api
# $PORT=8080
$PWD="$((pwd).Path.Replace('\', '/'))"
docker build -t mock-api .
docker run -d -p ${PORT}:8080 -v ${PWD}/responses:/responses mock-api
```

Since the files under `responses` are mounted in the container, you can continue testing by modifying the JSON files even while the container is running.

## Example. 
- You want to test `/v2/index`.  
Create a `v2` directory under `/responses` and place `index.json` under it.
Place `/responses/v2/index.json`.  

- If you want to test `/v2`.  
There are two ways to implement this.
    - Place `/responses/v2.json`.  
    - Place `.json` under `/responses/v2`.

- If you want to test `/v2?id=123&name=json`.
Place `v2_id=123&name=json.json` under `/responses`.
> The query parameter should be `?` can be specified by replacing `?` (since `?` because it cannot be used).

- If you want to test `/v2` with POST requests.
Place `v2.post.json` under `/responses`.
> The method should be `(path).(method).json` (Only `GET` can be omitted.)
