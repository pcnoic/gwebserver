package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"net/http"
	"syscall"
)

const serverPortNumber = ":9090"
const serverPath = "/"

func initMsg(){
	fmt.Printf("Starting the server at port 90... \n")
}

func runAndServe(port string){
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}else {
		log.Print("Server is listening to connections...")
	}
}

func closeHandler(){
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<- c
		fmt.Println("\r Ctrl+C pressed  | Exiting")
		os.Exit(0)
	}()
}

func reqHandler(){
	http.HandleFunc(serverPath, func(writer http.ResponseWriter, r *http.Request){
		fmt.Println(writer, "\r request received")
	})
}

func main()  {
	closeHandler()
	initMsg()
	reqHandler()
	runAndServe(serverPortNumber)
}