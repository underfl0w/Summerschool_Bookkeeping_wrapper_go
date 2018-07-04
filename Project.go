package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func singlelog(token []byte) {

	tokensize := bytes.IndexByte(token, 0)

	var (
		id int

		tempo int

		url string

		err error
	)

start:

	fmt.Println("Enter a number")

	tempo, _ = fmt.Scan(&id)

	if tempo == 0 {

		fmt.Println(" \nThis is not a number \n ")

		goto start

	}

	fmt.Println("\nAwaiting reponse\n ")

	url = "http://localhost:8080/api/single/entry/"

	url += strconv.Itoa(id)

	url += "?token=" + string(token[:tokensize])

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))

	}
}

func retrievefile(token []byte) {

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

	url = "http://localhost:8080/api/single/entry/file/"

	url += strconv.Itoa(id)

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))

	}
}

func uploadfile(token []byte) {

	var path string

	var name string

	var url string

	var id int

	url = "http://localhost:8080/api/upload/"

	fmt.Println("Enter the path of the file with the name")

	fmt.Scan(&path)

	if _, err := os.Stat(path); os.IsNotExist(err) {

		fmt.Println("File does not exist")

	} else {

		fmt.Println("File exists")

		file, err := os.Open(path)

		if err != nil {

			panic(err)

		}

		defer file.Close()

		if err != nil {

			panic(err)

		}

		buff := make([]byte, 512)

		_, err = file.Read(buff)

		if err != nil {

			fmt.Println(err)

			os.Exit(1)

		}

		filetype := http.DetectContentType(buff)

		fmt.Println(filetype)

		switch filetype {

		case "image/JPG", "image/jpg":

			fmt.Println(filetype)

		case "image/gif", "image/GIF":

			fmt.Println(filetype)

		case "application/msword":

			fmt.Println(filetype)

		case "application/pdf", "application/PDF":

			fmt.Println(filetype)

		default:

			fmt.Println("unknown file type uploaded")
		}

		fmt.Println("Enter the name of the file")

		fmt.Scan(&name)

		fmt.Println("Enter the ID of the log")

		fmt.Scan(&id)

		url += strconv.Itoa(id)

		res, err := http.Post(url, filetype, file)

		if err != nil {

			panic(err)

		}

		defer res.Body.Close()

		message, _ := ioutil.ReadAll(res.Body)

		fmt.Printf(string(message))

	}

}

func alllog(token []byte) {

	var url string

	var err error

	fmt.Println("\nAwaiting reponse\n ")

	url = "http://localhost:8080/api/all/entries/"

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))

	}
}

func userinfo(token []byte) {

	fmt.Println("Info")

}

func requesttoken() []byte {
	/*
		Fetch the authentication token that is needed for making requests.
	*/
	var url string

	url = "http://localhost/"

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

		return nil

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		return data

	}
}

func createlog() {

	fmt.Println("\nAwaiting reponse\n ")

	response, err := http.Get("http://localhost:8080/api/post/entry/data/")

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
	var token []byte

	var choice int

	token = requesttoken()

	for choice != 7 {

		printMenu()

		fmt.Scan(&choice)

		switch choice {

		case 1:

			singlelog(token)

		case 2:

			retrievefile(token)

		case 3:

			alllog(token)

		case 4:

			createlog()

		case 5:

			uploadfile(token)

		case 6:

			userinfo(token)

		case 7:

			fmt.Println("Bye !")

		default:

			fmt.Printf("Wrong choice !")
		}
	}
}
