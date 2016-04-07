package main

<<<<<<< HEAD
// Assignment 2
=======
// Assignment 1
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
// By Steve Stilson

import (
	"github.com/drone/routes"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
<<<<<<< HEAD
	"io"
	"io/ioutil"
	"github.com/mkilling/goejdb"
    "labix.org/v2/mgo/bson"
    "os"
	"net"
 	"net/rpc"
 	"strconv"
    "github.com/pelletier/go-toml"
=======
	// "strconv"
	"io"
	"io/ioutil"
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
)

type foodstruct struct {
	Type string
	Drink_alcohol string
}

type musicstruct struct {
	Spotify_user_id string
}

type flightstruct struct {
	Seat string
}

type travelstruct struct {
  Flight flightstruct
}

type moviestruct struct {
	Tv_shows []string
	Movies []string
}

type profile struct {
    Email string
    Zip string
    Country string
    Profession string
    Favorite_color string
    Is_smoking string
    Favorite_sport string
    Food foodstruct
    Music musicstruct
    Movie moviestruct
    Travel travelstruct
}

func (p profile) toString() string {
	var output string
	output = "{\n\temail:\t" + p.Email + ",\n"
	output += "\tzip:\t" + p.Zip + ",\n"
	output += "\tcountry:\t" + p.Country + ",\n"
	output += "\tprofession:\t" + p.Profession + ",\n"
	output += "\tfavorite_color:\t" + p.Favorite_color + ",\n"
	output += "\tis_smoking:\t" + p.Is_smoking + ",\n"
	output += "\tfavorite_sport:\t" + p.Favorite_sport + ",\n"
	output += "\tfood:\t{\n"
	output += "\t\ttype:\t" + p.Food.Type + ",\n"
	output += "\t\tdrink_alcohol:\t" + p.Food.Drink_alcohol + "\n"
	output += "\t},\n\tmusic:\t{\n"

	output += "\t\tspotify_user_id:\t" + p.Music.Spotify_user_id + "\n"
	output += "\tmovie:\t{\n"

	output += "\t\tTv_shows:\t[ "
	for _, show := range p.Movie.Tv_shows { output += show + ", " }
	output += "],\n\t\tMovies: \t[ "
	for _, show := range p.Movie.Movies { output += show + ", " }
<<<<<<< HEAD
    output += "],\n\ttravel:\t{\n"
	output += "\t\tflight:\t{\n\t\tseat: " + p.Travel.Flight.Seat
    /*	if p.Travel.Flight.Seat == true { output += "aisle"
=======
  output += "],\n\ttravel:\t{\n"
	output += "\t\tflight:\t{\n\t\tseat: " + p.Travel.Flight.Seat
/*	if p.Travel.Flight.Seat == true { output += "aisle"
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
		} else { output += "window"
		}  */
	output += "\n\t\t}\n\t}\n}\n"

	return output
}

// Create an global array of profiles to store here.
var profiles []profile
<<<<<<< HEAD
type Listener int
=======
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86

func main() {
	fmt.Println("Now in MyRestApp main.")

 	// start with one element, so the array is not empty.
<<<<<<< HEAD
    // Insert one record:
=======
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
	var anotherProfile profile
	var food foodstruct
	var music musicstruct
	var flight flightstruct
	var travel travelstruct
	var movie moviestruct
	anotherProfile.Email = "s@g.com"
	anotherProfile.Zip = "12345"
	anotherProfile.Country = "USA"
	anotherProfile.Profession = "techie"
	anotherProfile.Favorite_color = "blue"
	anotherProfile.Is_smoking = "no"
	anotherProfile.Favorite_sport = "football"
	food.Type = "vegetarian"
	food.Drink_alcohol = "no"
	anotherProfile.Food = food
	anotherProfile.Music = music
	flight.Seat = "aisle"
	travel.Flight = flight
	anotherProfile.Travel = travel
	anotherProfile.Movie = movie

<<<<<<< HEAD
	tomlFilename := os.Args[1]
	fmt.Println("Opening toml file: " + tomlFilename)
	config, err9 := toml.LoadFile(tomlFilename)
	if err9 != nil {
		fmt.Println(err9)
	}
	// fmt.Println(config, err9)

	dbfile := config.Get("database.file_name").(string)
	fmt.Println("dbfile is " + dbfile)
	db9, err8 := goejdb.Open(dbfile, goejdb.JBOWRITER | goejdb.JBOCREAT | goejdb.JBOTRUNC)
	if err8 != nil {
		fmt.Println("Error opening " + dbfile)
	  	panic(err8)
	}

	profiles = append(profiles, anotherProfile)

	coll, _ := db9.CreateColl("profileList", nil)
	bsrec, _ := bson.Marshal(anotherProfile)
    coll.SaveBson(bsrec)
    db9.Close()

    fmt.Printf("\nSaved the first record.\n")
    var rpcchannel string
    var rpcnum int64
    rpcnum = config.Get("replication.rpc_server_port_num").(int64)
    rpcchannel = ":" + strconv.FormatInt(rpcnum, 10)
	go initializeRPCServer(rpcchannel)
=======
	profiles = append(profiles, anotherProfile)

  fmt.Print("All profiles are: ")
	fmt.Println(profiles)	// for debugging purposes
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86

	mux := routes.New()

	mux.Get("/profile/:email", GetProfile)
	mux.Put("/profile/:email", PutProfile)	// updates existing profile

	mux.Del("/profile/:email", DeleteProfile)
	mux.Post("/profile", PostProfile)	// creates a new profile

	http.Handle("/", mux)
<<<<<<< HEAD
	var uiport string
	var uiportnum int64
	uiportnum = config.Get("database.port_num").(int64)
	uiport = ":" + strconv.FormatInt(uiportnum, 10)
	log.Println("Listening on port " + uiport)
//	fmt.Printf("%s:%s\n", os.Getenv("IP"), os.Getenv("PORT"))

 	http.ListenAndServe(uiport, nil)

} // end main


=======
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
} // end main

>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
func GetProfile(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	emailToFind := params.Get(":email")
	w.Header().Set("Content-Type", "application/json")
<<<<<<< HEAD
	var profileReturned profile
	profileReturned = GetOneRecord(emailToFind, false, false)
	if profileReturned.Email == "No such email" {
		fmt.Println("No such email in the database.")
	}
    js, err := json.Marshal(profileReturned)	// have to convert it again to json to send it to the client.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.Write(js)

	printAllRecordsInDB()

} // end get
=======
  fmt.Println("Looking up the profile for " + emailToFind + ":\n\n")
	var newProfile profile
  var found bool
	found = false

  // Go through the slice until you find the one that matches that email.  Set newProfile to it.
	for _,element := range profiles {
		if element.Email == emailToFind {
			newProfile = element
		  found = true }
}
	if found == false { w.Write([]byte(emailToFind + " not found."))
	} else {
		fmt.Println("Found " + newProfile.toString())
		js, err := json.Marshal(newProfile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	}
	return
}
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86

func PostProfile(w http.ResponseWriter, r *http.Request) {
	// creates a new profile
	fmt.Println("Now in PostProfile.")
//	decoder := json.NewDecoder(r.Body)
// I tried decoder := json.NewDecoder(strings.NewReader(r.Body)). Didn't work.
<<<<<<< HEAD
    body, err := ioutil.ReadAll(r.Body)
    if err != nil { panic("Bad body")  }
=======
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
		panic("Bad body")
  }
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
	var newProfile profile
      	// err := decoder.Decode(&newProfile)
	err = json.Unmarshal(body, &newProfile)
       	// body, err := ioutil.ReadAll(r.Body)
<<<<<<< HEAD
	if  err == io.EOF { fmt.Println("end of file.")
	  } else if err != nil {
	      fmt.Println("Error in decoding: ")
		  fmt.Println(err)
	  }

	fmt.Println("Now creating the profile for " + newProfile.Email + ":\n")
=======
	if  err == io.EOF {
	      fmt.Println("end of file.")
	  } else if err != nil {
	      fmt.Println("Error in decoding: ")
				fmt.Println(err)
	  }

	fmt.Println("Now creating the profile for " + newProfile.Email + ":\n\n")
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
	fmt.Print("newProfile = ")
	fmt.Println(newProfile)	// for debugging purposes
	if newProfile.Is_smoking != "yes" && newProfile.Is_smoking != "no" {
		fmt.Println("Bad value for Is_smoking: " + newProfile.Is_smoking)   }

<<<<<<< HEAD
    if newProfile.Food.Drink_alcohol != "yes" && newProfile.Food.Drink_alcohol != "no" {
			fmt.Println("Bad value for Drink_alcohol: " + newProfile.Food.Drink_alcohol)   }

    if newProfile.Travel.Flight.Seat != "aisle" && newProfile.Travel.Flight.Seat != "window" {
			fmt.Println("Bad value for Flight.Seat: " + newProfile.Travel.Flight.Seat)   }

	// Add it to the list
	profiles = append(profiles, newProfile)

	// And also add it to the database.  Need to use new variables each time, for some reason.
	tomlFilename := os.Args[1]
	fmt.Println("Opening toml file: " + tomlFilename)
	config, err6 := toml.LoadFile(tomlFilename)
	if err6 != nil {
		fmt.Println(err6)
	}
	dbfile := config.Get("database.file_name").(string)
	db1, err1 := goejdb.Open(dbfile, goejdb.JBOWRITER | goejdb.JBOCREAT )
	if err1 != nil {
		fmt.Println("Problem opening " + dbfile + " in POST:")
		panic(err1)
	}
    coll, _ := db1.CreateColl("profileList", nil)

	bsrec, _ := bson.Marshal(newProfile)
    coll.SaveBson(bsrec)
    db1.Close()

	printAllRecordsInDB()	// Print all profiles for debugging purposes
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	// Now send the POST to the other instance.
	tomlFilename = os.Args[1]
	fmt.Println("Opening toml file: " + tomlFilename)
	config1, err7 := toml.LoadFile(tomlFilename)
	if err7 != nil {
		fmt.Println(err7)
	}
    var rpcchannel string
    var rpcnum int64
    rpcnum = config1.Get("replication.rpc_server_port_num").(int64)
    rpcchannel = ":" + strconv.FormatInt(rpcnum, 10)
	fmt.Println("RPC channel is ")
	fmt.Println(rpcchannel)
	client, err := rpc.Dial("tcp", rpcchannel)
 	if err != nil {
 		fmt.Println("Error dialing: ")
 		log.Fatal(err)
 	}

 //	in := bufio.NewReader(os.Stdin)
	var reply int
 		// Then call the server.method.

 	err = client.Call("Listener.RPCPost", newProfile.Email, &reply)
 	if err != nil {
 		fmt.Println("Error calling RPCPost: ")
 		log.Fatal(err)
 	}
=======
  if newProfile.Food.Drink_alcohol != "yes" && newProfile.Food.Drink_alcohol != "no" {
			fmt.Println("Bad value for Drink_alcohol: " + newProfile.Food.Drink_alcohol)   }

  if newProfile.Travel.Flight.Seat != "aisle" && newProfile.Travel.Flight.Seat != "window" {
			fmt.Println("Bad value for Flight.Seat: " + newProfile.Travel.Flight.Seat)   }

	profiles = append(profiles, newProfile)
	fmt.Println(profiles)	// for debugging purposes
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
	return
}

func PutProfile(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	//  In order to update the profile, search for it in the database, delete it and then append a new one with the new info.
	fmt.Println("Now in PutProfile.")

	params := r.URL.Query()
	emailToFind := params.Get(":email")
	var newProfile profile
	newProfile = GetOneRecord(emailToFind, true, false)
	fmt.Println("Deleted the profile for " + newProfile.Email + ".\n")

	body, err := ioutil.ReadAll(r.Body)
=======
	//  In order to update the profile, search for it, delete it and then append a new one with the new info.
	fmt.Println("Now in PutProfile.")
	params := r.URL.Query()
	emailToFind := params.Get(":email")

	var newProfile profile
	var found bool
	found = false

	// Go through the slice until you find the one that matches that email.  Set newProfile to it.
	for counter,element := range profiles {
		if element.Email == emailToFind {
			newProfile = element
			found = true
			profiles = append(profiles[:counter], profiles[counter+1:]...)	// deletes it
		}
  }
	if found == false { w.Write([]byte(emailToFind + " not found."))
	} else {
		fmt.Println("Now updating the profile for " + emailToFind + ":\n\n")
		body, err := ioutil.ReadAll(r.Body)
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
    if err != nil {
        panic("Bad body")
    }
//    fmt.Println("Found this in the body: ")
//		fmt.Println(body)
    err = json.Unmarshal(body, &newProfile)
				// Unmarshal will only update the fields of the keys sent to it, not all fields.
<<<<<<< HEAD
		if  err == io.EOF { fmt.Println("end of file.")
		} else if err != nil {
          fmt.Println("Bad marshal: ")
		  fmt.Println(err)
        }

		fmt.Print("Updated newProfile = ")
=======
		if  err == io.EOF {
        fmt.Println("end of file.")
      } else if err != nil {
          fmt.Println("Bad marshal: ")
					fmt.Println(err)
      }

		fmt.Print("newProfile = ")
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
		fmt.Println(newProfile)	// for debugging purposes
		if newProfile.Is_smoking != "yes" && newProfile.Is_smoking != "no" {
			fmt.Println("Bad value for Is_smoking: " + newProfile.Is_smoking)   }

<<<<<<< HEAD
	    if newProfile.Food.Drink_alcohol != "yes" && newProfile.Food.Drink_alcohol != "no" {
				fmt.Println("Bad value for Drink_alcohol: " + newProfile.Food.Drink_alcohol)   }

	    if newProfile.Travel.Flight.Seat != "aisle" && newProfile.Travel.Flight.Seat != "window" {
				fmt.Println("Bad value for Flight.Seat: " + newProfile.Travel.Flight.Seat)   }

		profiles = append(profiles, newProfile)

		tomlFilename := os.Args[1]
		fmt.Println("Opening toml file: " + tomlFilename)
		config, err6 := toml.LoadFile(tomlFilename)
		if err6 != nil {
			fmt.Println(err6)
		}
		dbfile := config.Get("database.file_name").(string)
    	fmt.Println("Looking up the profile for " + newProfile.Email + " in database: " + dbfile + "\n")

		db2, err2 := goejdb.Open(dbfile, goejdb.JBOWRITER | goejdb.JBOCREAT )
		if err2 != nil {
			fmt.Println("Problem creating into the database " + dbfile + " in PUT:")
			panic(err2)
		}
    	coll, _ := db2.CreateColl("profileList", nil)

		bsrec, _ := bson.Marshal(newProfile)
    	coll.SaveBson(bsrec)
    	db2.Close()

		printAllRecordsInDB()	// for debugging purposes
		w.WriteHeader(http.StatusNoContent)

		// Now send the PUT to the other instance.
	  	var rpcchannel string
	 	var rpcnum int64
    	rpcnum = config.Get("replication.rpc_server_port_num").(int64)
    	rpcchannel = ":" + strconv.FormatInt(rpcnum, 10)
		client, err := rpc.Dial("tcp", rpcchannel)
	 	if err != nil {
 			fmt.Println("Error dialing: ")
 			log.Fatal(err)
 		}

 	//	in := bufio.NewReader(os.Stdin)
		var reply int
 		// Then call the server.method.

 		err = client.Call("Listener.RPCPut", newProfile.Email, &reply)
 		if err != nil {
 			fmt.Println("Error calling RPCPut: ")
 			log.Fatal(err)
 		}
 		fmt.Println("client = ")
 		fmt.Println(client)
// 		fmt.Println("line = " + line)
 		fmt.Println("reply = ")
 		fmt.Println(reply)

		return
} // end Put

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	emailToFind := params.Get(":email")
	var profileReturned profile
	profileReturned = GetOneRecord(emailToFind, true, false)
	fmt.Println("Deleted the profile for " + profileReturned.Email + ".\n")

	printAllRecordsInDB()
	w.WriteHeader(http.StatusNoContent)

	// Now send the delete to the other instance via RPC.
	tomlFilename := os.Args[1]
	fmt.Println("Opening toml file: " + tomlFilename)
	config, err6 := toml.LoadFile(tomlFilename)
	if err6 != nil {
		fmt.Println(err6)
	}
    var rpcchannel string
    var rpcnum int64
    rpcnum = config.Get("replication.rpc_server_port_num").(int64)
    rpcchannel = strconv.FormatInt(rpcnum, 10)
	client, err := rpc.Dial("tcp", rpcchannel)
 	if err != nil {
 		fmt.Println("Error dialing: ")
 		log.Fatal(err)
 	}

 //	in := bufio.NewReader(os.Stdin)
	line := "Hello. This is a test."
	var reply int
 		// Then call the server.method.

 		err = client.Call("Listener.RPCDelete", line, &reply)
 		if err != nil {
 			fmt.Println("Error calling RPCDelete: ")
 			log.Fatal(err)
 		}
	return
}

func printAllRecordsInDB (){

	var newProfile2 profile

	tomlFilename := os.Args[1]
	fmt.Println("Opening toml file: " + tomlFilename)
	config, err6 := toml.LoadFile(tomlFilename)
	if err6 != nil {
		fmt.Println(err6)
	}
	dbfile := config.Get("database.file_name").(string)
	db3, err3 := goejdb.Open(dbfile, goejdb.JBOREADER )
	if err3 != nil {
		fmt.Println("Error in opening " + dbfile)
	  	panic(err3)
	}
	coll3, _ := db3.CreateColl("profileList", nil)

    fmt.Print("\nHere are all the records in the database " + dbfile + ": ")
    results3, _ := coll3.Find(``)
    fmt.Printf("Records found: %d\n", len(results3))
	for counter2,element2 := range results3 {
		 bson.Unmarshal(element2, &newProfile2)
		 fmt.Print("Record ")
		 fmt.Print(counter2)
		 fmt.Print(": ")
		 fmt.Println(newProfile2)
	}
	db3.Close()
}

func initializeRPCServer (serverport string){

	fmt.Println("Starting initializeRPCServer on " + serverport)
	addy, err := net.ResolveTCPAddr("tcp", serverport)
 	if err != nil {
 		log.Fatal(err)
 	}

	inbound, err := net.ListenTCP("tcp", addy)
 	if err != nil {
 		log.Fatal(err)
 	}

 	listener := new(Listener)
 	rpc.Register(listener)
 	rpc.Accept(inbound)
}

func (t *Listener) RPCDelete(line string, reply *int) error {
			fmt.Print("Now in RPCDelete with input ")
			fmt.Println(line)
			GetOneRecord(line, true, true)  // deletes
			return nil
		}

func (t *Listener) RPCPost(line string, reply *int) error {
			fmt.Print("Now in RPCPost with input ")
			fmt.Println(line)
			GetOneRecord(line, false, true)
			return nil
		}

func (t *Listener) RPCPut(line string, reply *int) error {
			fmt.Print("Now in RPCPut with input ")
			fmt.Println(line)
			GetOneRecord(line, false, true)
			return nil
		}

func GetOneRecord (emailLookUp string, deleteFlag bool, switchDB bool) profile {

	tomlFilename := os.Args[1]
	fmt.Println("Opening toml file: " + tomlFilename)
	config, err6 := toml.LoadFile(tomlFilename)
	if err6 != nil {
		fmt.Println(err6)
	}
	dbfile := config.Get("database.file_name").(string)
	if switchDB == true {
		if dbfile == "app1.db" {
			dbfile = "app2.db"
		} else {
			dbfile = "app1.db"
		}
	}
    fmt.Println("Looking up the profile for " + emailLookUp + " and database: " + dbfile + "\n")

	var newProfile profile
	newProfile.Email = "No such email"	// Default value for non-existent records in database.

	db, err := goejdb.Open(dbfile, goejdb.JBOWRITER | goejdb.JBOCREAT )	// JBOWRITER can do both reads and writes.
	if err != nil { panic(err) 	}
	coll, _ := db.CreateColl("profileList", nil)

	// Now execute query
	// For some reason, email has to be lower case, even though the field starts with a capital E. :-(

	query := fmt.Sprintf("{\"email\" : \"%s\"}", emailLookUp) // The default is just to retrieve, not delete.
	results, err4 := coll.Find(query)						 // These variables will be written over if delete is done.
    fmt.Printf("Records found: %d\n", len(results))
    if err4 != nil {
    	fmt.Println("Error in retrieval: ")
    	fmt.Println(err4)
    }
	if deleteFlag == true {
		fmt.Println("Delete flag set.")
 		query = fmt.Sprintf("{\"email\" : \"%s\" , \"$dropall\": true }", emailLookUp)
		_, err0 := coll.Update(query)				// .Update returns an int, but .Find returns a result set.
	    if err0 != nil {
    		fmt.Println("Error in deleting: ")
    		fmt.Println(err0)
    	}
	}
    db.Close()

  // Go through the slice of those that were found and show the records.  Should be only one.  Set newProfile to it.
	for _,element := range results {
		 bson.Unmarshal(element, &newProfile)	// Will write over the default "no such email".
		 fmt.Print("Found the following record: ")
		 fmt.Println(newProfile)

	} // end for
	return newProfile

} // end function
=======
	  if newProfile.Food.Drink_alcohol != "yes" && newProfile.Food.Drink_alcohol != "no" {
				fmt.Println("Bad value for Drink_alcohol: " + newProfile.Food.Drink_alcohol)   }

	  if newProfile.Travel.Flight.Seat != "aisle" && newProfile.Travel.Flight.Seat != "window" {
				fmt.Println("Bad value for Flight.Seat: " + newProfile.Travel.Flight.Seat)   }

		profiles = append(profiles, newProfile)
		fmt.Println(profiles)	// for debugging purposes
		w.WriteHeader(http.StatusNoContent)
		return
	} // end found
} // end Put

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	//  Returns an object, in this case a profile.
	params := r.URL.Query()
	emailToFind := params.Get(":email")
	fmt.Println("Now deleting the profile for " + emailToFind + ":\n\n")
	var found bool
	found = false

	for counter,element := range profiles {
		if element.Email == emailToFind {
		  found = true
			fmt.Println("Found " + emailToFind + ". Now deleting it.")
			fmt.Println(element.toString())
			profiles = append(profiles[:counter], profiles[counter+1:]...)
			}
		}
	if found == false { fmt.Println(emailToFind + " not found.") }
	w.WriteHeader(http.StatusNoContent)
	return
}
>>>>>>> 805bc29e00779c399a8c5652f88d300bcc78bb86
