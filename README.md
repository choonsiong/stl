# STL

A simple demo project that provides web API to read [STL format](https://en.wikipedia.org/wiki/STL_(file_format)) file from a directory and save the STL data into a database.

Endpoints:

```
- GET /
- GET /ping
- POST /get-stl-list
  - result: {"stl_list":["Demo.STL"]}
- POST /save-stl-mongo
  - body: {"name":"Demo.STL"}
- POST /query-stl-mongo
  - body: {"name":"Demo.STL"}
```

Use below command to start up a MongoDB container:

```
docker run --name mongodb -v "$DOCKER_MONGODB_DATA:/data/db" --publish 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=password -d mongo:latest
```

Above command assume the environment variable `DOCKER_MONGODB_DATA` exists, else create a directory somewhere and add the environment variable (e.g., `DOCKER_MONGODB_DATA=/path/to/your/local/directory`) to your shell startup script (e.g., `.bashrc`, `.zshrc`).

To create the executable, run `make all`:
```
$ make all
rm -f stl-server
go build -o stl-server .
$ 
```

To access MongoDB shell:

```
$ docker exec -it mongodb bash
root@eb756328a645:/# 
root@eb756328a645:/# 
root@eb756328a645:/# mongosh -u mongoadmin
Enter password: ********
Current Mongosh Log ID:	673b53c53eecb762966b5ba4
Connecting to:		mongodb://<credentials>@127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.1.1
Using MongoDB:		7.0.5
Using Mongosh:		2.1.1

For mongosh info see: https://docs.mongodb.com/mongodb-shell/

------
   The server generated these startup warnings when booting
   2024-11-18T14:48:16.871+00:00: /sys/kernel/mm/transparent_hugepage/enabled is 'always'. We suggest setting it to 'never'
   2024-11-18T14:48:16.871+00:00: vm.max_map_count is too low
------

test> 

test> show dbs
STL       2.26 MiB
admin   148.00 KiB
config   48.00 KiB
local    72.00 KiB
test> use STL
switched to db STL
STL> show collections
Binary
STL> 
```