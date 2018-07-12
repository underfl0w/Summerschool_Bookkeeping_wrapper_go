package main

import (
	"aliceREST"
)

func main() {
	// The server to contact and the port number that goes with it.
	var server = "localhost:8081"

	// Request a single log entry from the rest API
	var cool = aliceREST.Singlelog(1, server)

	// Submit an entry log into the server backend using the rest API
	aliceREST.Createlog("2018-07-06T13:45:12.000Z","partical",
		"taart","ERROR","222","Jurjen",
		"Test by Jurjen","crash of the universe", "fix universe",
		"2018-07-06T13:45:12.000Z","ERROR", server)

	// Request all entries from rest API
	aliceREST.Alllog(server)


	print(cool)

}