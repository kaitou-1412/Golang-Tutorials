package main

import (
	"fmt"
	"net/url"
)

const rawUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to Handling URLs in golang")
	fmt.Println(rawUrl)

	// parsing the url
	result, err := url.Parse(rawUrl)
	checkNilError(err)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("-------------------------Query Params---------------------------")
	qparams := result.Query()
	fmt.Printf("Type of query params is %T and Value is %v\n", qparams, qparams)
	fmt.Println(qparams["coursename"])
	for key, value := range qparams {
		fmt.Println(key, " -> ", value[0])
		fmt.Printf("Type of value is %T\n", value)
	}
	fmt.Println("-------------------------Create a URL out of chunks of info---------------------------")
	// always pass the reference in this case
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=ruban",
	}
	createdUrl := partsOfUrl.String()
	fmt.Println(createdUrl)
}

func checkNilError(err error) {
	if err != nil {
		//shuts down the execution of the program and shows the error
		panic(err)
	}
}
