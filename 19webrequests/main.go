package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Handling web requests in golang")

	response, err := http.Get(url)
	checkNilError(err)
	defer closeHTTPConnection(response)
	fmt.Printf("Response is of type: %T\n", response)

	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	content := string(dataBytes)
	fmt.Println(content)
}

func closeHTTPConnection(response *http.Response) {
	// caller's responsibility to close the connection
	// since defer is specific to the scope it is bring used in...
	// so it will execute the code before returning to the main function
	defer response.Body.Close()
}

func checkNilError(err error) {
	if err != nil {
		//shuts down the execution of the program and shows the error
		panic(err)
	}
}
