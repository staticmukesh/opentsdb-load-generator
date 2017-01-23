# opentsdb-load-generator
A simple tool for generating time series load for benchmarking Opentsdb.

## Compile from source

```go
go get github.com/staticmukesh/opentsdb-load-generator
```
The `opentsdb-load-generator` binary should now be available at $GOPATH/bin/opentsdb-load-generator

## Usage
```go
user@localhost:~/$ ./opentsdb-load-generator --help
Usage of ./opentsdb-load-generator:
  -conn int
    	Number of connection to Opentsdb (default 1)
  -rate int
    	Number of data points per second to be send (default 1000)
  -tsdb string
    	Opentsdb server address (default "localhost:4242")

```

## License
Copyright (c) 2017 Mukesh Sharma. Licensed under the MIT License.
