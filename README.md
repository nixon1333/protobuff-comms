# Idea
Trying to write an application based on protobuff communication. 
The main idea is there will be micro-service and client written in 
- python service
- golang

A common protobuff data structure will be shared among them. They will use that proto libs to comminucate between them using http protocol.

# Folder structure
```
/protos
├──/libs
│   ├──/go
│   ├──/python
├──/python-service
└──/go-client 
```

`go-client` is holding go implementation for protobuf request and response parsing

`python-service` is serving and parsing data using protobuf over http.

`protos` is holding original `.proto` file and generate libs.


# Commands
## proto lib generate

To run library file run from root directory following command 

### python lib
`protoc --go_out=protos/libs/go protos/users.proto`

### go lib
`protoc --go_out=protos/libs/go protos/users.proto`

## python project run 
- go to python project directory `cd python-service`
- setup virtual environment `virtualenv venv` (first time only)
- source it `source venv/bin/activate`
- install requirements `pip install -r requirements.txt` (first time only)
- to run the basic flask app run `export FLASK_APP=main.py && python -m flask run`

Boom! Server is running in `http://127.0.0.1:5000/` :) 

## golang project run
- go to go project directory `cd go-client`
- run `go run *.go`

Yeahhh! Client is now calling `http://127.0.0.1:5000/` to posting proto file, reading proto  :) 

# Contact
Ashraful Islam Nixon - @nixon1333

Project Link: https://github.com/nixon1333/protobuff-comms
