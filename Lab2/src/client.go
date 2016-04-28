package main

// Lab 2
// By Steve Stilson

import (
    "fmt"
    "os"
    "io"
    "strings"
    "strconv"
    "crypto/md5"
    "net/http"
    "bytes"
    "io/ioutil"
)

var nodes map[int]int
var ports []string 
var numPorts int

func main() {

    fmt.Println("Now in client.go main.")

    // Set up the incoming number of nodes first.
    // The params are the ports where a server will start executing on.
    // If it contains a dash, then it's a range.
    
    nodes = make(map[int]int)
    ports = make([]string, 6)
    params := os.Args[1:]  // params is an array of arguments on the command line.
    fmt.Print("Incoming parameters are: ")
    for u:=0; u < len(params); u++ {
        fmt.Printf("params[%d]: ", u)
        fmt.Println(params[u])
    }
    if len(params) < 2 {
        fmt.Println("len params is less than 2.")
        fmt.Println("Usage: go run client.go startPort-endPort key->value,key->value,key->value,key->value, etc.\nNo spaces around the dash or the commas.")
        os.Exit(1)
    }
    if strings.Index(params[1], ",") == -1 {
        // comma isn't in the last parameter
        fmt.Println("Last parameter doesn't have a comma in it.")
        fmt.Println("Usage: go run client.go startPort-endPort key->value,key->value,key->value,key->value, etc.\nNo spaces around the dash or the commas.")
        os.Exit(1)
    }
    if strings.Index(params[0], "-") == -1 {
        // dash isn't in the second-to-last parameter
        fmt.Println("Second-to-last parameter doesn't have a dash in it.")
        fmt.Println("Usage: go run client.go startPort-endPort key->value,key->value,key->value,key->value, etc.\nNo spaces around the dash or the commas.")
        os.Exit(1)
    }
        // dash is in the param, so split it on the dash to find the start and end.
        startAndEnd := strings.Split(params[0], "-")
        // startAndEnd would contain [1000, 1005] now.
        var counter int
        start, err1 := strconv.Atoi(startAndEnd[0])
        if err1 != nil { fmt.Println(err1) }
        end, err2 := strconv.Atoi(startAndEnd[1])
        if err2 != nil { fmt.Println(err2) }
        numPorts = 0

        // This loop prints out the ports and also serves to initialize numPorts.
        for counter = start; counter <= end; counter++ {
            ports[numPorts] = strconv.Itoa(counter)
            fmt.Printf("Ports[%d] = %s.\n", numPorts, ports[numPorts])
            numPorts++
        } // end for

    // Now parse the parameter for the data to add to the servers.
    
    dataArray := strings.Split(params[1], ",")
    for counter = 0; counter < len(dataArray); counter++ {
        pair := strings.Split(dataArray[counter], "->")
        add(pair[0], pair[1])
      }
} // end main

func add(keyToAdd string, valueToAdd string) {
      fmt.Printf("I'll insert new key: %s with new value: %s.\n", keyToAdd, valueToAdd)

	h := md5.New()
	io.WriteString(h, valueToAdd)
	hashValue := h.Sum(nil)
	fmt.Printf("Hash for value %s is %x.\n", valueToAdd, hashValue)

	// Md5 always returns 32 hex bytes. theFirstInt is the first byte, in the range 0-255.
	// 256 / number of ports gives slotwidths.  Divide the hashvalue / slotwidths to 
	// See which "slot" (server) to send that data to.
	
	var temp string
	temp = string(hashValue[0])
    theFirstInt := int(temp[0])
    
	slotWidths := 256 / numPorts
	fmt.Print("theFirstInt = ")
	fmt.Println(theFirstInt)
    fmt.Printf("slotWidths = %d ", slotWidths)
    
	whichServer := theFirstInt / slotWidths
	fmt.Print("So I'm sending this data to server on port ")
	fmt.Println(ports[whichServer])
	
	var jsonStr []byte
	var url string
	url = "https://cmpe273-stilsons.c9users.io:" + ports[whichServer] + "/" + keyToAdd + "/" + valueToAdd
	fmt.Println("Doing PUT on url " + url)
    req, err4 := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
    if err4 != nil { fmt.Println(err4) }
    client := &http.Client{}
    resp, err5 := client.Do(req)
    if err5 != nil { panic(err5) }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

} // end add
