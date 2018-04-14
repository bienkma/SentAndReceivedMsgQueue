package main

import (
	"fmt"
	"bytes"
	"os"
	"log"
	"time"
	"bufio"
	"io/ioutil"
	"encoding/json"

	"net/http"
	url2 "net/url"
)

// Sender UI
func sender(url string, topic string, msg string) {
	form := url2.Values{
		"topic": {topic},
		"msg":   {msg},
	}
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Println(err)
	}
	defer rsp.Body.Close()
}

// Receiver UI
func receiver(url string)  {
	rsp, err := http.Get(url)
	if err != nil{
		log.Println(err)
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	var content map[string]interface{}
	json.Unmarshal(body, &content)
	if content["Status"] == 200.0{
		fmt.Println("$Receiver->",content["Msg"])
	}
	time.Sleep(3)
}

// Banner function display
func banner() {
	fmt.Println("\033[H\033[2J") // Clear screen
	fmt.Println("+---------------------------------------------------------------+")
	fmt.Println("|\t\t\t CHAT CLIENT \t\t\t\t|")
	fmt.Println("|Current System Time: " + time.Now().Format(time.RFC850) + "\t\t|")
	fmt.Println("+---------------------------------------------------------------+")
}

func menu() {
	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1. Choice s for sender mode")
	fmt.Println("2. Choice r for receiver mode")
	fmt.Println("3. Choice q for exit program")
}

func main() {
	var topic string = ""
	var apiSenderUrl string = ""
	var apiReceiverUrl string = ""

	// Get Variable from environment OS, if nil set variable from default value
	if topic = os.Getenv("TOPIC"); topic == "" {
		topic = "roomchat"
	}
	if apiSenderUrl = os.Getenv("API_SENDER_URL"); apiSenderUrl == "" {
		apiSenderUrl = "http://localhost:8080/apis/sender"
	}

	if apiReceiverUrl = os.Getenv("API_RECEIVER_URL"); apiReceiverUrl == "" {
		apiReceiverUrl = "http://localhost:8080/apis/receiver?topic="+topic
	}

	banner()
	menu()

	for {
	Start:
		fmt.Print("->")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		switch scan.Text() {
		case "s":
			for {
				fmt.Print("$sender->")
				scan := bufio.NewScanner(os.Stdin)
				scan.Scan()
				sender(apiSenderUrl, topic, scan.Text())
			}
		case "r":
			fmt.Print("$receiver->")
			for {
				receiver(apiReceiverUrl)
			}
		case "q":
			fmt.Println("Good bye!")
			os.Exit(0)
		default:
			goto Start
		}
	}
}
