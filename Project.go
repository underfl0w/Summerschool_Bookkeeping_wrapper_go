package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func singlelog() {

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

	url = "http://heikovm.hihva.nl/api/single/entry/"

	url += strconv.Itoa(id)

	//url += "?token=" + token

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

	url = "http://heikovm.hihva.nl/api/single/entry/file/"

	url += strconv.Itoa(id)

	//url += "?token="

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))

	}
}

func uploadfile() {

	var path string

	var name string

	var url string

	var id int

	url = "http://heikovm.hihva.nl/api/upload/"

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

func alllog() {

	var url string

	var err error

	fmt.Println("\nAwaiting reponse\n ")

	url = "http://heikovm.hihva.nl/api/all/entries/"

	//url += "?token=" + token

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

// This function is not needed anymore. The need for tokens has been disabled by Frederick.
func requesttoken() string {
	/*
		Fetch the authentication token that is needed for making requests.
	*/
	var url string

	url = "http://heikovm.hihva.nl/"

	response, err := http.Get(url)

	fmt.Println("\n", response, "\n")

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

		return ""

	} else {

		return string(response.Request.URL.RawQuery)

	}
}

func createlog() {

	var url string

	var subsystem string

	var class string

	var typelog string

	var run string

	var author string

	var title string

	var text string

	var followsup string

	var interruptionduration string

	var interventiontype string

	var date string

	fmt.Println("\nAwaiting reponse\n ")

	url = "http://heikovm.hihva.nl/api/post/entry/data/"

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	}

	created := time.Now()

	date = created.Format("2006-01-02 15:04:05")

	fmt.Println(date, "\n")

	fmt.Println("Enter the susbsystem : \n")

	fmt.Scan(&subsystem)

	fmt.Println("\nEnter the class : \n ")

	fmt.Scan(&class)

	fmt.Println("\nEnter the type of run : \n ")

	fmt.Scan(&typelog)

	fmt.Println("\nEnter the run number : \n ")

	fmt.Scan(&run)

	fmt.Println("\nEnter the author : \n ")

	fmt.Scan(&author)

	fmt.Println("\nEnter the title : \n ")

	fmt.Scan(&title)

	fmt.Println("\nEnter the description : \n ")

	fmt.Scan(&text)

	fmt.Println("\nEnter the follow up : \n ")

	fmt.Scan(&followsup)

	fmt.Println("\nEnter the interruption duration : \n ")

	fmt.Scanf("%s", interruptionduration)

	fmt.Println("\nEnter the intervention type : \n ")

	fmt.Scan(&interventiontype)

	jsonData := map[string]string{"created": date, "subsystem": subsystem, "class": class, "type": typelog, "run": run, "author": author, "title": title, "log_entry_text": text, "follow_ups": followsup, "interruption_duration": interruptionduration, "intervention_type": interventiontype}

	jsonValue, _ := json.Marshal(jsonData)

	response, err = http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

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
	//var token string

	var choice int

	//token = requesttoken()

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

			createlog()

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
