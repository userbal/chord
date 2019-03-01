package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func doCreate()             {}
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
