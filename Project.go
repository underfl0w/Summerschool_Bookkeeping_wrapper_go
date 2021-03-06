package aliceREST

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"math/rand"
)

func Testunitary(server string) {

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

	var tempo int

	var id int


	url = "http://" + server + "/api/post/entry/data/"

	_, err := http.Get(url)

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	}

start:

	fmt.Println("Enter a number\n")

	tempo, _ = fmt.Scan(&id)

	if tempo == 0 {

		fmt.Println(" \nThis is not a number \n ")

		goto start

	}

	for i := 0; i < id; i++ {

		subsystem += string(rand.Intn(122-97) + 97)

		class += string(rand.Intn(122-97) + 97)

		typelog += string(rand.Intn(122-97) + 97)

		run += string(rand.Intn(122-97) + 97)

		author += string(rand.Intn(122-97) + 97)

		title += string(rand.Intn(122-97) + 97)

		text += string(rand.Intn(122-97) + 97)

		followsup += string(rand.Intn(122-97) + 97)

		interventiontype += string(rand.Intn(122-97) + 97)

	}

	date = strconv.Itoa(rand.Intn(20)) + strconv.Itoa(rand.Intn(100)) + "-" + strconv.Itoa(rand.Intn(13)) + "-" + strconv.Itoa(rand.Intn(32)) + " " + strconv.Itoa(rand.Intn(24)) + ":" + strconv.Itoa(rand.Intn(24)) + ":" + strconv.Itoa(rand.Intn(24))

	interruptionduration = strconv.Itoa(rand.Intn(20)) + strconv.Itoa(rand.Intn(100)) + "-" + strconv.Itoa(rand.Intn(13)) + "-" + strconv.Itoa(rand.Intn(32)) + " " + strconv.Itoa(rand.Intn(24)) + ":" + strconv.Itoa(rand.Intn(24)) + ":" + strconv.Itoa(rand.Intn(24))
/*
	fmt.Println("Created :")

	fmt.Println("\n")

	fmt.Println(date)

	fmt.Println("\nSubsystem :")

	fmt.Println("\n")

	fmt.Println(subsystem)

	fmt.Println("\nClass :")

	fmt.Println("\n")

	fmt.Println(class)

	fmt.Println("\nType :")

	fmt.Println("\n")

	fmt.Println(typelog)

	fmt.Println("\nRun Number :")

	fmt.Println("\n")

	fmt.Println(run)

	fmt.Println("\nAuthor :")

	fmt.Println("\n")

	fmt.Println(author)

	fmt.Println("\nTitle :")

	fmt.Println("\n")

	fmt.Println(title)

	fmt.Println("\nlog_entry_text :")

	fmt.Println("\n")

	fmt.Println(text)

	fmt.Println("\nfollowsup:")

	fmt.Println("\n")

	fmt.Println(followsup)

	fmt.Println("\nInterruption_duration :")

	fmt.Println("\n")

	fmt.Println(interruptionduration)

	fmt.Println("\nIntervention_type :")

	fmt.Println("\n")

	fmt.Println(interventiontype)

	fmt.Println("\n")
*/
	jsonData := map[string]string{"created": date, "subsystem": subsystem, "class": class, "type": typelog, "run": run, "author": author, "title": title, "log_entry_text": text, "follow_ups": followsup, "interruption_duration": interruptionduration, "intervention_type": interventiontype}

	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		panic(err)

	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	//alllog()

}

func Singlelog(logIdentifier int, server string) string{

	//var	logIdentifier int

	//var	tempo int

	var	url string

	var	err error

//start:

	//fmt.Println("Enter a number")

	//tempo, _ = fmt.Scan(&logIdentifier)

//	if tempo == 0 {

//		fmt.Println(" \nThis is not a number \n ")

//		goto start

//	}


	url = "http://" + server + "/api/single/entry/"

	url += strconv.Itoa(logIdentifier)

	//url += "?token=" + token

	response, err := http.Get(url)

	if err != nil {
		return "The HTTP request failed with error" + err.Error()

	} else {

		contents, _ := ioutil.ReadAll(response.Body)

		Stringcontents := string(contents)

		number_0 := strings.Count(Stringcontents, ",")

		number_1 := strings.Count(Stringcontents, "{")

		number_2 := strings.Count(Stringcontents, "}")

		data := strings.Replace(Stringcontents, ",", ",\n", number_0)

		data = strings.Replace(data, "{", "\n{\n", number_1)

		data = strings.Replace(data, "}", "\n}\n", number_2)

		return data

	}
	return "nill"
}

func Retrievefile(id int, server string) {

	//var id int

	//var tempo int

	var url string

	var err error

//start:

	//fmt.Println("Enter a number")

	//tempo, _ = fmt.Scan(&id)

	//if tempo == 0 {

	//	fmt.Println("\nThis is not a number \n ")

	//	goto start

	//}


	url = "http://"+server+"/api/single/entry/file/"

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

func Uploadfile(path string,name string,id int, server string) {

	//var path string

	//var name string

	var url string

	//var id int

	url = "http://" + server + "/api/upload/"

	//fmt.Println("Enter the path of the file with the name")

	//fmt.Scan(&path)

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

	//	fmt.Println("Enter the name of the file")

	//	fmt.Scan(&name)

	//	fmt.Println("Enter the ID of the log")

	//	fmt.Scan(&id)

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

func Alllog(server string) string{

	var url string

	var err error

	fmt.Println("\nAwaiting reponse")

	url = "http://" + server + "/api/all/entries/"

	//url += "?token=" + token

	response, err := http.Get(url)

	if err != nil {
		return "The HTTP request failed with error" + err.Error()

	} else {

		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {

			panic(err)

		}

		Stringcontents := string(contents)

		number_0 := strings.Count(Stringcontents, ",")

		number_1 := strings.Count(Stringcontents, "{")

		number_2 := strings.Count(Stringcontents, "}")

		data := strings.Replace(Stringcontents, ",", ",\n", number_0)

		data = strings.Replace(data, "{", "\n{\n", number_1)

		data = strings.Replace(data, "}", "\n}\n", number_2)

		return data

	}
}

func Userinfo(server string) {

	fmt.Println("Info")

}

// This function is not needed anymore. The need for tokens has been disabled by Frederick.
//func Requesttoken() string {
//	/*
//		Fetch the authentication token that is needed for making requests.
//	*/
//	var url string
//
//	url = server
//
//	response, err := http.Get(url)
//
//	fmt.Println("\n", response, "\n")
//
//	if err != nil {
//
//		fmt.Printf("The HTTP request failed with error %s\n", err)
//
//		return ""
//
//	} else {
//
//		return string(response.Request.URL.RawQuery)
//
//	}
//}

func Createlog(date string,subsystem string,class string,typelog string,run string,author string,title string,text string,followsup string,interruptionduration string,interventiontype string, server string) []byte{

	var url string

	url = "http://" + server + "/api/post/entry/data/"

	_, err := http.Get(url)

	if err != nil {
		return nil

	}

	//created := time.Now()

	//date = created.Format("2006-01-02 15:04:05")

	//fmt.Println(date, "\n")

	//fmt.Println("Enter the susbsystem : \n")

	//reader := bufio.NewReader(os.Stdin)

	//subsystem, _ = reader.ReadString('\n')
	//
	//subsystem = strings.Replace(subsystem, "\n", "", -1)
	//
	//fmt.Println("\nEnter the class : \n ")
	//
	//reader1 := bufio.NewReader(os.Stdin)
	//
	//class, _ = reader1.ReadString('\n')
	//
	//class = strings.Replace(class, "\n", "", -1)
	//
	//fmt.Println("\nEnter the type of run : \n ")
	//
	//reader2 := bufio.NewReader(os.Stdin)
	//
	//typelog, _ = reader2.ReadString('\n')
	//
	//typelog = strings.Replace(typelog, "\n", "", -1)
	//
	//fmt.Println("\nEnter the run number : \n ")
	//
	//reader3 := bufio.NewReader(os.Stdin)
	//
	//run, _ = reader3.ReadString('\n')
	//
	//run = strings.Replace(run, "\n", "", -1)
	//
	//fmt.Println("\nEnter the author : \n ")
	//
	//reader4 := bufio.NewReader(os.Stdin)
	//
	//author, _ = reader4.ReadString('\n')
	//
	//author = strings.Replace(author, "\n", "", -1)
	//
	//fmt.Println("\nEnter the title : \n ")
	//
	//reader5 := bufio.NewReader(os.Stdin)
	//
	//title, _ = reader5.ReadString('\n')
	//
	//title = strings.Replace(title, "\n", "", -1)
	//
	//fmt.Println("\nEnter the description : \n ")
	//
	//reader6 := bufio.NewReader(os.Stdin)
	//
	//text, _ = reader6.ReadString('\n')
	//
	//text = strings.Replace(text, "\n", "", -1)
	//
	//fmt.Println("\nEnter the follow up : \n ")
	//
	//reader7 := bufio.NewReader(os.Stdin)
	//
	//followsup, _ = reader7.ReadString('\n')
	//
	//followsup = strings.Replace(followsup, "\n", "", -1)
	//
	//fmt.Println("\nEnter the interruption duration : \n ")
	//
	//reader8 := bufio.NewReader(os.Stdin)
	//
	//interruptionduration, _ = reader8.ReadString('\n')
	//
	//interruptionduration = strings.Replace(interruptionduration, "\n", "", -1)
	//
	//fmt.Println("\nEnter the intervention type : \n ")
	//
	//reader9 := bufio.NewReader(os.Stdin)
	//
	//interventiontype, _ = reader9.ReadString('\n')
	//
	//interventiontype = strings.Replace(interventiontype, "\n", "", -1)

	jsonData := map[string]string{"created": date, "subsystem": subsystem, "class": class, "type": typelog, "run": run, "author": author, "title": title, "log_entry_text": text, "follow_ups": followsup, "interruption_duration": interruptionduration, "intervention_type": interventiontype}

	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		panic(err)

	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

//func printMenu() {
//
//	fmt.Println("\n|-------------------------------------------------------|")
//
//	fmt.Println("|Choose 1 to retrieve a single log entry                |")
//
//	fmt.Println("|Choose 2 to retrieve a file from the single log entry  |")
//
//	fmt.Println("|Choose 3 to retrieves all log entries                  |")
//
//	fmt.Println("|Choose 4 to create an log entry                        |")
//
//	fmt.Println("|Choose 5 to uploads a file to the log entry            |")
//
//	fmt.Println("|Choose 6 to find a user                                |")
//
//	fmt.Println("|Choose 7 to exit                                       |")
//
//	fmt.Println("|-------------------------------------------------------|\n ")
//
//}

//func main() {
//	//var token string
//
//	var choice int
//
//	//token = requesttoken()
//
//	for choice != 7 {
//
//		printMenu()
//
//		fmt.Scan(&choice)
//
//		switch choice {
//
//		case 1:
//
//			//singlelog()
//
//		case 2:
//
//			//retrievefile()
//
//		case 3:
//
//			Alllog()
//
//		case 4:
//
//			//createlog()
//
//		case 5:
//
//			//uploadfile()
//
//		case 6:
//
//			Userinfo()
//
//		case 7:
//
//			fmt.Println("Bye !")
//
//		case 8:
//
//			Testunitary()
//
//		default:
//
//			fmt.Printf("Wrong choice !")
//		}
//	}
//}

func printMenu() {
	//
	//	fmt.Println("\n|-------------------------------------------------------|")
	//
	//	fmt.Println("|Choose 1 to retrieve a single log entry                |")
	//
	//	fmt.Println("|Choose 2 to retrieve a file from the single log entry  |")
	//
	//	fmt.Println("|Choose 3 to retrieves all log entries                  |")
	//
	//	fmt.Println("|Choose 4 to create an log entry                        |")
	//
	//	fmt.Println("|Choose 5 to uploads a file to the log entry            |")
	//
	//	fmt.Println("|Choose 6 to find a user                                |")
	//
	//	fmt.Println("|Choose 7 to exit                                       |")
	//
	//	fmt.Println("|-------------------------------------------------------|\n ")
	//
	//}

	//func main() {
	//	//var token string
	//
	//	var choice int
	//
	//	//token = Requesttoken()
	//
	//	for choice != 7 {
	//
	//		printMenu()
	//
	//		fmt.Scan(&choice)
	//
	//		switch choice {
	//
	//		case 1:
	//
	//			Singlelog()
	//
	//		case 2:
	//
	//			Retrievefile()
	//
	//		case 3:
	//
	//			alllog()
	//
	//		case 4:
	//
	//			Createlog()
	//
	//		case 5:
	//
	//			Uploadfile()
	//
	//		case 6:
	//
	//			Userinfo()
	//
	//		case 7:
	//
	//			fmt.Println("Bye !")
	//
	//		case 8:
	//
	//			test_unitary()
	//
	//		default:
	//
	//			fmt.Printf("Wrong choice !")
	//		}
	//	}
	//}
}
