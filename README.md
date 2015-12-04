RPC is a way of connecting two separate services via a raw TCP socket

> Note: SOAP, Thrift, REST APIs, message queues such as RabbitMQ, and key value stores such as Etcd are examples of other tools and protocols

## Credit

Code modified only slightly from https://github.com/haisum/rpcexample

## Basic outline

The fundamental principle is, you define an RPC service:

- Write a function
- Add some RPC configuration
- Register our function as a RPC service
- Start the service and have it listen for messages on a specific port

From here we have a client service that calls the RPC service:

- Write code which calls RPC function
- Call the function via a specific ip/port
- The 'message' is passed over as valid JSON

## JSON-RPC

The client JSON could look like:

- `method`: name of method/service
- `params`: Array of arguments to be passed
- `id`: usually an integer; makes it easier for client to know which request it got a response to (if the RPC calls are done asynchroneously)
 
```json
{"method": "Arith.Multiply", "params": [{"A": 2, "B": 3}], "id": 1}
```

The RPC server JSON response could look like:

- `result`: contains return value of method called (`null` if error ocurred)
- `error`: if error occurred, indicates error code or error message, otherwise `null`
- `id`: the id of the request it is responding to

```json
{"result": 6, "error": null, "id": 1}
```

To run the example code in this repo you'll need both:

- https://github.com/Integralist/rpcremote
- https://github.com/Integralist/rpcclient

In the remote repo run:

```bash
go run rpcremote/server.go
2015/12/04 10:25:29 Serving RPC handler
```

In the client repo run:

```bash
go run rpcclient/client.go
2015/12/04 10:27:22 2*3=6
```

This will trigger a log in the remote:

```bash
2015/12/04 10:27:22 Multiplying 2 with 3
```
