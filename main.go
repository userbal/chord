package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

const (
	defaultHost = "localhost"
	defaultPort = "3410"
)

type Node struct{}

var Port int = 3401

func readCommands(command string, arg string) {
	switch command {
	case "help":
		doHelp()
	case "port":
		doPort(arg)
	case "create":
		doCreate()
	case "join":
		doJoin(arg)
	case "quit":
		doQuit()
	case "put":
		doPut(arg)
	case "putRandom":
		doPutRandom(arg)
	case "get":
		doGet(arg)
	case "delete":
		doDelete(arg)
	case "dump":
		doDump()
	case "dumpkey":
		doDumpKey(arg)
	case "dumpaddr":
		doDumpAddr(arg)
	case "dumpall":
		doDumpAll()
	}
}

func doHelp() {
	fmt.Printf("commands:\nport <port number>\ncreate\njoin <address>\nput <key> <value>\nputrandom <number of keys>\nget <key>\ndeletekey <key>\ndump\ndumpaddr <address>\ndumpkey <key>\ndumpall\nquit\n")
}

func doPort(portNumber string) {
	port, err := strconv.Atoi(portNumber)
	if err != nil {
		fmt.Printf("invalid port")
		return
	}
	Port = port
}

func doCreate() {
	node := new(Node)
	address := defaultHost + ":" + defaultPort
	go server(address)
}

func server(address string) {
	actor := startActor()
	rpc.Register(actor)
	rpc.HandleHTTP()
	fmt.Printf("Listening")

	l, e := net.Listen("tcp", address)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	if err := http.Serve(l, nil); err != nil {
		log.Fatalf("http.Server: %v", err)
	}
}

func (s server) ping() error {
	finished := make(chan struct{})
	s <- func(f *Feed) {
		if len(f.Messages) < count {
			count = len(f.Messages)

		}
		*reply = make([]string, count)
		copy(*reply, f.Messages[len(f.Messages)-count:])
		finished <- struct{}{}
	}
	<-finished
	return nil
}

func startActor() Server {
	ch := make(chan handler)
	state := new(Feed)
	go func() {

		for f := range ch {
			f(state)
		}
	}()
	return ch
}
func doJoin(address string) {}
func doQuit() {
	os.Exit(0)
}
func doPut(keyandValue string)   {}
func doPutRandom(address string) {}
func doGet(key string)           {}
func doDelete(key string)        {}
func doDump()                    {}
func doDumpAddr(address string)  {}
func doDumpKey(address string)   {}
func doDumpAll()                 {}

//func connection(restart string) {
//	ln, err := net.listen("tcp", ListeningToPort)
//	if err != nil {
//		log.Printf("listen failed: %v", err)
//	}
//
//	for {
//		conn, err := ln.Accept()
//		if restart == true {
//			return
//		}
//		if err != nil {
//			log.Printf("accept error: %v", err)
//			continue
//		}
//
//	}
//}

func main() {

	fmt.Printf("type help for a list of commands\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stringline := scanner.Text()
		line := strings.Fields(stringline)

		if len(line) > 0 {
			command := line[0]
			line = line[1:]
			stringl := strings.Join(line, " ")
			readCommands(command, stringl)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}
}
