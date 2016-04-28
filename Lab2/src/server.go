package main

// Lab 2
// By Steve Stilson

import (
	"github.com/drone/routes"
	"log"
	"net/http"
	"fmt"
	"strconv"
	"os"
	"strings"
)

// Create a hashmap to store values.
var myMap [10]map[int]string
var startPort string

type pair struct {
	key int
	value string
}

var ports []string 

func main() {
	fmt.Println("Now in server.go main.")

// Assume no more than 10.
	myMap = [10]map[int]string{}
	myMap[0] = make(map[int]string)
	myMap[1] = make(map[int]string)
	myMap[2] = make(map[int]string)
	myMap[3] = make(map[int]string)
	myMap[4] = make(map[int]string)
	myMap[5] = make(map[int]string)
	myMap[6] = make(map[int]string)
	myMap[7] = make(map[int]string)
	myMap[8] = make(map[int]string)
	myMap[9] = make(map[int]string)
	
//	fmt.Print("All values are: ")
//	fmt.Println(myMap)	// show an empty map for debugging purposes

    ports = make([]string, 10)
    params := os.Args[1:]  // params is an array of arguments on the command line.
    fmt.Print("Incoming parameters are: ")
    for u:=0; u < len(params); u++ {
        fmt.Printf("params[%d]: ", u)
        fmt.Println(params[u])
    }
 	var startAndEnd []string
	startAndEnd = make([]string, 2)
    if strings.Index(params[0], "-") == -1 {
        // Dash isn't in the second-to-last parameter, so it must be just a number like 3000.
        // So, in that case, make startAndEnd = [3000, 3000].
		startPort = params[0]
		startAndEnd[0] = startPort
		startAndEnd[1] = startPort
    } else {
         // dash is in the param, so split it on the dash to find the start and end.
        startAndEnd = strings.Split(params[0], "-")
    } // end dash check
 
    // startAndEnd would contain [3000, 3005] now.
    startPort = startAndEnd[0]
    var counter int
    start, err1 := strconv.Atoi(startAndEnd[0])
    if err1 != nil { fmt.Println(err1) }
    end, err2 := strconv.Atoi(startAndEnd[1])
    if err2 != nil { fmt.Println(err2) }
  
	mux := routes.New()
	mux.Get("/:key", getValue)	// Goes to the Getvalue() method
	mux.Put("/:key/:value", putValue)	// updates existing value
	// No post or delete

	mux.Get("/", getAllValues)
	http.Handle("/", mux)

    i := 0
    for counter = start; counter <= end; counter++ {
        ports[i] = strconv.Itoa(counter)
        fmt.Printf("Ports[%d] = %s.\n", i, ports[i])
        go listenOnPort(ports[i])
        i++
    } // end for
    fmt.Println("Hit a key to end the program: ")
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
} // end main

func listenOnPort(portParam string) {
	
//	fmt.Print("Now in listenOnPort with param ")
//	fmt.Println(portParam)
	var portString string
	portString = ":" + portParam
	
	log.Printf("Listening on port %s....", portString)
	http.ListenAndServe(portString, nil)
} // end listenOnPort

func getValue(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	keyToFind, err := strconv.Atoi(params.Get(":key"))
    if err != nil {
        fmt.Println(err)
    }
	//	w.Header().Set("Content-Type", "application/json")
	fmt.Print("Now looking up the value for ")
	fmt.Println(keyToFind)
	var result string
	var found bool
	
	myPort := whichMap(r)
	
	result, found = myMap[myPort][keyToFind] 
	if found == false {
		w.Write([]byte(strconv.Itoa(keyToFind)))
		w.Write([]byte(" not found."))
	} else {
		w.Write([]byte(result))
	}
	return
}

func putValue(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	keyToPut, err2 := strconv.Atoi(params.Get(":key"))
	if err2 != nil { panic(err2) }
	valueToPut := params.Get(":value")
	fmt.Print("Now in putValue with key ")
	fmt.Print(keyToPut)
	fmt.Print(" and valueToPut = ")
	fmt.Println(valueToPut)
	
	myPort := whichMap(r)
	myMap[myPort][keyToPut] = valueToPut 

	w.WriteHeader(http.StatusNoContent)
	return

} // end Put

func getAllValues(w http.ResponseWriter, r *http.Request) {
	var item pair
	
	myPort := whichMap(r)
	//	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("\n{\n   [\n"))
	
	for newKey, newValue := range myMap[myPort] {
		item.key = newKey
		item.value = newValue
		w.Write([]byte(item.toString())) 
	}
	w.Write([]byte("   ]\n}\n"))
	return
}

func (p pair) toString() string {
	var output string
	output = "\t{\n\t\tkey:\t\t"
	output += strconv.Itoa(p.key)
	output += "\n\t\tvalue:\t\t"
	output += p.value
	output += "\n\t}\n"

	return output
}

func whichMap (r *http.Request) int {
//	fmt.Println("r.Host is " + r.Host)
//	fmt.Println("I'm going to split it on : now.")
	urlArray := strings.Split(r.Host,":")
	myPort, err := strconv.Atoi(urlArray[1])
	if err != nil { fmt.Println(err) }
	var startNum int
	startNum, err0 := strconv.Atoi(startPort)
	if err0 != nil { fmt.Println(err0) }

	myPort = myPort - startNum
	fmt.Print("whichMap is returning port number ")
	fmt.Println(myPort)
	return myPort
}
