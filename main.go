package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showCommands()

	command := readCommand()

	switch command {

	case 1:
		fmt.Println("Monitoring...")
		startMonitoring()

	case 2:
		fmt.Println("Showing logs")

	case 3:
		fmt.Println("Exiting...")
		exit()

	default:
		fmt.Println("Command not found")
		unexpectedExit()
	}
}

func showCommands() {
	fmt.Println("1 - Begin monitorization")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")
}

func readCommand() int {
	var command int

	fmt.Scan(&command)

	return command
}

func exit() {
	os.Exit(0)
}

func unexpectedExit() {
	os.Exit(-1)
}

func startMonitoring() {
	url := "https://www.clouflare.com"
	response, _ := http.Get(url)
	fmt.Println(response)
}
