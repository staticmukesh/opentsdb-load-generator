# opentsdb-load-generator

### Overview
opentsdb-load-generator is a simple tool for load-testing Opentsdb. It can generate and push large amount of data to the Opentsdb. Currently, it uses `telnet` based put request

### Compile from source

```go
go get github.com/staticmukesh/opentsdb-load-generator
```
The `opentsdb-load-generator` binary should now be available at `$GOPATH/bin`/opentsdb-load-generator

### Usage
```bash
$ ./opentsdb-load-generator --help
Usage of ./opentsdb-load-generator:
  -conn int
    	Number of connection to Opentsdb (default 1)
  -metric string
    	Metric name to be send. (default "test.metric")
  -rate int
    	Number of data points per second to be send (default 1000)
  -tsdb string
    	Opentsdb server address (default "localhost:4242")
```

e.g. to push data at the rate of 10k datapoints per second on 10 connections, use the following command:
```bash
$ ./opentsdb-load-generator -conn=5 -rate=10000
2017/01/24 21:57:46 Conn No: 0, connected to localhost:4242
2017/01/24 21:57:46 Conn No: 2, connected to localhost:4242
2017/01/24 21:57:46 Conn No: 1, connected to localhost:4242
2017/01/24 21:57:46 Conn No: 4, connected to localhost:4242
2017/01/24 21:57:46 Conn No: 3, connected to localhost:4242
2017/01/24 21:57:47 Pushed 2000 data points in last 1 second on Conn: 1
2017/01/24 21:57:47 Pushed 2000 data points in last 1 second on Conn: 3
2017/01/24 21:57:47 Pushed 2000 data points in last 1 second on Conn: 0
2017/01/24 21:57:47 Pushed 2000 data points in last 1 second on Conn: 2
2017/01/24 21:57:47 Pushed 2000 data points in last 1 second on Conn: 4
```

### Contributing
Feel free to raise PR for any feature improvement or issue.

### License
Copyright (c) 2017 Mukesh Sharma. Licensed under the MIT License.
