package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

type Conf struct {
	Conn   int
	Server string
	Rate   int
	Metric string
	Host   string
}

func main() {
	conn := flag.Int("conn", 1, "Number of connection to Opentsdb")
	server := flag.String("tsdb", "localhost:4242", "Opentsdb server address")
	rate := flag.Int("rate", 1000, "Number of data points per second to be send")
	metric := flag.String("metric", "test.metric", "Metric name to be send.")

	flag.Parse()

	host, err := os.Hostname()
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}

	conf := &Conf{
		Conn:   *conn,
		Server: *server,
		Rate:   *rate,
		Metric: *metric,
		Host:   host,
	}

	data := make(chan string)

	generateLoad(data, conf)
	pushData(data, conf)

	var exit chan string
	<-exit
}

func generateLoad(data chan<- string, conf *Conf) {
	for i := 0; i < conf.Rate; i++ {
		go func(data chan<- string, conf *Conf) {
			rand.Seed(time.Now().Unix())
			ticker := time.NewTicker(time.Second)

			for {
				select {
				case <-ticker.C:
					timeStamp := time.Now().Unix()
					value := rand.Intn(100)
					req := fmt.Sprintf("put %s %d %d host=%s", conf.Metric, timeStamp, value, conf.Host)
					data <- req
					break
				}
			}
		}(data, conf)
	}
}

func pushData(data <-chan string, conf *Conf) {
	for i := 0; i < conf.Conn; i++ {
		go func(data <-chan string, conf *Conf, connId int) {
			conn, err := net.Dial("tcp", conf.Server)
			if err != nil {
				log.Fatalln(err.Error())
			}
			defer conn.Close()

			log.Printf("Conn: %d, Connected to %s", connId, conf.Server)
			go func(conn net.Conn) {
				for {
					resp, err := bufio.NewReader(conn).ReadString('\n')
					if err != nil {
						log.Fatalln(err.Error())
						break
					}

					log.Println(resp)
				}
			}(conn)

			ticker := time.NewTicker(time.Second)
			count := 0
			for {
				select {
				case req := <-data:
					conn.Write([]byte(req))
					break
				case <-ticker.C:
					log.Printf("Pushed %d data points in last 1 second on Conn: %d\n", count, i)
					break
				}
			}
		}(data, conf, i)
	}
}
