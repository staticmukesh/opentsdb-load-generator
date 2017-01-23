package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type Conf struct {
	Conn   int
	Server string
	Rate   int
}

func main() {
	conn := flag.Int("conn", 1, "Number of connection to Opentsdb")
	server := flag.String("tsdb", "localhost:4242", "Opentsdb server address")
	rate := flag.Int("rate", 1000, "Number of data points per second to be send")

	flag.Parse()

	conf := &Conf{
		Conn:   *conn,
		Server: *server,
		Rate:   *rate,
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
					req := fmt.Sprintf("put %s %d %d %s", "test.metric", timeStamp, value, "host=a")
					data <- req
					break
				}
			}
		}(data, conf)
	}
}

func pushData(data <-chan string, conf *Conf) {
	for i := 0; i < conf.Conn; i++ {
		go func(data <-chan string, conf *Conf) {
			conn, err := net.Dial("tcp", conf.Server)
			if err != nil {
				log.Fatalln(err.Error())
			}
			defer conn.Close()

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

			for {
				select {
				case req := <-data:
					conn.Write([]byte(req))
					break
				}
			}
		}(data, conf)
	}
}
