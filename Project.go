package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func singlelog() {

	var id int

	var tempo int

	var url string

	var err error

start:

	fmt.Println("Enter a number")

	tempo, _ = fmt.Scan(&id)

	if tempo == 0 {

		fmt.Println(" \nThis is not a number \n ")

		goto start

	}

	fmt.Println("\nAwaiting reponse\n ")

	url = "http://localhost/api/single/entry/"

	url += strconv.Itoa(id)

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))

	}
}

func retrievefile() {

	var id int

	var tempo int

	var url string

	var err error

start:

	fmt.Println("Enter a number")

	tempo, _ = fmt.Scan(&id)

	if tempo == 0 {

		fmt.Println("\nThis is not a number \n ")

		goto start

	}

	fmt.Println("\nAwaiting reponse\n ")

	url = "http://localhost/api/single/entry/file/"

	url += strconv.Itoa(id)

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))

	}
}

func uploadfile() {

	fmt.Println("Upload file")

}

func alllog() {

	var url string

	var err error

	fmt.Println("\nAwaiting reponse\n ")

	url = "http://localhost/api/all/entries/"

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))

	}
}

func userinfo() {

	fmt.Println("Info")

}

func creatlog() {

	fmt.Println("\nAwaiting reponse\n ")

	response, err := http.Get("http://localhost")

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	}

	jsonData := map[string]string{"created": "Orange Juice", "subsystem": "French Fries", "class": "3A", "type": "???", "run": "8", "author": "Sangoku", "title": "Pressing orange", "log_entry_text": "A orange explode", "follow_ups": "Where does it come from ?", "interruption_duration": "2018-07-07 21:21:21", "intervention_type": "Emergency"}

	jsonValue, _ := json.Marshal(jsonData)

	response, err = http.Post("http://localhost/api/post/entry/data/", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))

	}

	fmt.Println("End")
}

func printMenu() {

	fmt.Println("\n|-------------------------------------------------------|")

	fmt.Println("|Choose 1 to retrieve a single log entry                |")

	fmt.Println("|Choose 2 to retrieve a file from the single log entry  |")

	fmt.Println("|Choose 3 to retrieves all log entries                  |")

	fmt.Println("|Choose 4 to create an log entry                        |")

	fmt.Println("|Choose 5 to uploads a file to the log entry            |")

	fmt.Println("|Choose 6 to find a user                                |")

	fmt.Println("|Choose 7 to exit                                       |")

	fmt.Println("|-------------------------------------------------------|\n ")

}

func main() {

	var choice int

	for choice != 7 {

		printMenu()

		fmt.Scan(&choice)

		switch choice {

		case 1:

			singlelog()

		case 2:

			retrievefile()

		case 3:

			alllog()

		case 4:

			creatlog()

		case 5:

			uploadfile()

		case 6:

			userinfo()

		case 7:

			fmt.Println("Bye !")

		default:

			fmt.Printf("Wrong choice !")
		}
	}
}
