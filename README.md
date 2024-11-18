# STL

A simple demo project that provides web API to read [STL format](https://en.wikipedia.org/wiki/STL_(file_format)) file from a directory and save the STL data into a database.

Endpoints:

- `GET /`
- `GET /ping`
- `POST /get-stl-list`
  - result: `{"stl_list":["Demo.STL"]}`
- `POST /save-stl-mongo`
  - body: `{"name":"Demo.STL"}`
- `POST /query-stl-mongo`
  - body: `{"name":"Demo.STL"}`

Use below command to start up a MongoDB container:

`docker run --name mongodb -v "$DOCKER_MONGODB_DATA:/data/db" --publish 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=password -d mongo:latest`

Above command assume the environment variable `DOCKER_MONGODB_DATA` exists, else create a directory somewhere and add the environment variable (e.g., `DOCKER_MONGODB_DATA=/path/to/your/local/directory`) to your shell startup script (e.g., `.bashrc`, `.zshrc`).

To create the executable, run `make all`:
```
$ make all
rm -f stl-server
go build -o stl-server .
$ 
```