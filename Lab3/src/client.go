package main

// Lab 3
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
    "encoding/binary"
)

var ports []string 
var numPorts int
var weights []int
var sizes []int

func main() {

    fmt.Println("Now in client.go main.")

    // Set up the incoming number of nodes first.
    // The params are the ports where a server will start executing on.
    // If it contains a dash, then it's a range.
    
    sizes = []int{0, 0, 0, 0, 0, 0 }    // How many entries sent to each port so far
    ports = make([]string, 6)           // Assume never more than 6
    params := os.Args[1:]               // params is an array of arguments on the command line.
    fmt.Print("Incoming parameters are: ")
    for u:=0; u < len(params); u++ {
        fmt.Printf("params[%d]: ", u)
        fmt.Println(params[u])
    }
    if len(params) < 2 {
        fmt.Println("len params is less than 2.")
        fmt.Println("Usage: go run client.go startPort-endPort key->value,key->value,key->value,key->value, etc., with\nno spaces around the dash or the commas.")
        os.Exit(1)
    }
    if strings.Index(params[1], ",") == -1 {
        // comma isn't in the last parameter
        fmt.Println("Last parameter doesn't have a comma in it.")
        fmt.Println("Usage: go run client.go startPort-endPort key->value,key->value,key->value,key->value, etc., with\nno spaces around the dash or the commas.")
        os.Exit(1)
    }
    if strings.Index(params[0], "-") == -1 {
        // dash isn't in the second-to-last parameter
        fmt.Println("Second-to-last parameter doesn't have a dash in it.")
        fmt.Println("Usage: go run client.go startPort-endPort key->value,key->value,key->value,key->value, etc., with\nno spaces around the dash or the commas.")
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

        var weights []int
        var maxNode int
        // Use the size of each node to calculate the weight for each node.  Then choose the one with the maximum score.
        for count := 0; count < numPorts; count++ {
            w := weight(sizes[count], keyToAdd)
            fmt.Printf("Port %d's score is %d.\n", count, w)
            weights = append(weights, w)
            // At this point, weights[0] will contain the w (score) for node 0,
            // weights[1] will contain the weight for node 1, etc.
            // maxnode will start out pointing to node 0, then be compared to each node's 
            // weight, and the one with the largest weight will be pointed to by maxnode each time.
            // maxNode is an integer pointing to which node number, NOT the value itself.
            if count == 0 { 
                maxNode = 0
            } else {
                if weights[count] > weights[maxNode] { maxNode = count }
            }
        }
	fmt.Printf("So key %s will be sent to the maximum, which is server node %d.\n", keyToAdd, maxNode)
	// That shows which "slot" (server) to send that data to.
    sizes[maxNode]++
    
	var jsonStr []byte
	var url string
	url = "https://cmpe273-stilsons.c9users.io:" + ports[maxNode] + "/" + keyToAdd + "/" + valueToAdd
	fmt.Println("Doing PUT on url " + url)
    req, err4 := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
    if err4 != nil { fmt.Println(err4) }
    client := &http.Client{}
    resp, err5 := client.Do(req)
    if err5 != nil { panic(err5) }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)
    fmt.Println("Response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))

} // end add

func weight(nodeSize int, key string) int {
    var returnVal int
    var a int
    var b int
    a = 1103515245
    b = 12345
	h := md5.New()
	io.WriteString(h, key)
	hashValue := h.Sum(nil)
	buf := bytes.NewBuffer(hashValue) 
    hashInt, _ := binary.ReadVarint(buf)  // converts from []byte to int

    // This uses a slight variation of the preferred method shown on page 19
    // of the original Rendezvous HRW paper.
	returnVal = (a * ((a * nodeSize + b) + int(hashInt) + b) % (2^31))
	// golang uses ^ for bitwise XOR on integers, and % does integer modulus.
	// hashValue is a byte array.  2^31 is a 10-digit number.
	// The mod operation makes bigger-sized nodes return smaller scores.  Good!

    return returnVal
}
