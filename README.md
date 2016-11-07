# helloworld
Helloworld is a trivial demo web application written in Go.
It serves the following endpoints (all using the ```GET``` verb): 
* ```/hello```. Returns a hello message, including a count of hello's.
* ```/env```. Returns all environment variables with the prefix ```NOMAD```.
* ```/fs/<path>```. Serves the file or folder specified by ```<path>```
* ```/health```. Returns the health status of the application.
* ```/fail```. Sets the health status to unhealthy.
* ```/ok```. Sets the health status to healthy.

## running
```
LISTEN_ADDRESS=127.0.0.1:8081 helloworld "message to include in hello"
```
