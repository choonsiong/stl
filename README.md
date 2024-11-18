# STL

A simple demo project that provides web API to read STL file from a directory and save the STL data into a database.

Endpoints:

- `GET /`
- `GET /ping`
- `POST /get-stl-list`
  - result: `{"stl_list":["Demo.STL"]}`
- `POST /save-stl-mongo`
  - body: `{"name":"Demo.STL"}`
- `POST /query-stl-mongo`