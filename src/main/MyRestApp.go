package main

// Assignment 1
// By Steve Stilson

import (
	"github.com/drone/routes"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	// "strconv"
	"io"
	"io/ioutil"
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
  output += "],\n\ttravel:\t{\n"
	output += "\t\tflight:\t{\n\t\tseat: " + p.Travel.Flight.Seat
/*	if p.Travel.Flight.Seat == true { output += "aisle"
		} else { output += "window"
		}  */
	output += "\n\t\t}\n\t}\n}\n"

	return output
}

// Create an global array of profiles to store here.
var profiles []profile

func main() {
	fmt.Println("Now in MyRestApp main.")

 	// start with one element, so the array is not empty.
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

	profiles = append(profiles, anotherProfile)

  fmt.Print("All profiles are: ")
	fmt.Println(profiles)	// for debugging purposes

	mux := routes.New()

	mux.Get("/profile/:email", GetProfile)
	mux.Put("/profile/:email", PutProfile)	// updates existing profile

	mux.Del("/profile/:email", DeleteProfile)
	mux.Post("/profile", PostProfile)	// creates a new profile

	http.Handle("/", mux)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
} // end main

func GetProfile(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	emailToFind := params.Get(":email")
	w.Header().Set("Content-Type", "application/json")
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

func PostProfile(w http.ResponseWriter, r *http.Request) {
	// creates a new profile
	fmt.Println("Now in PostProfile.")
//	decoder := json.NewDecoder(r.Body)
// I tried decoder := json.NewDecoder(strings.NewReader(r.Body)). Didn't work.
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
		panic("Bad body")
  }
	var newProfile profile
      	// err := decoder.Decode(&newProfile)
	err = json.Unmarshal(body, &newProfile)
       	// body, err := ioutil.ReadAll(r.Body)
	if  err == io.EOF {
	      fmt.Println("end of file.")
	  } else if err != nil {
	      fmt.Println("Error in decoding: ")
				fmt.Println(err)
	  }

	fmt.Println("Now creating the profile for " + newProfile.Email + ":\n\n")
	fmt.Print("newProfile = ")
	fmt.Println(newProfile)	// for debugging purposes
	if newProfile.Is_smoking != "yes" && newProfile.Is_smoking != "no" {
		fmt.Println("Bad value for Is_smoking: " + newProfile.Is_smoking)   }

  if newProfile.Food.Drink_alcohol != "yes" && newProfile.Food.Drink_alcohol != "no" {
			fmt.Println("Bad value for Drink_alcohol: " + newProfile.Food.Drink_alcohol)   }

  if newProfile.Travel.Flight.Seat != "aisle" && newProfile.Travel.Flight.Seat != "window" {
			fmt.Println("Bad value for Flight.Seat: " + newProfile.Travel.Flight.Seat)   }

	profiles = append(profiles, newProfile)
	fmt.Println(profiles)	// for debugging purposes
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	return
}

func PutProfile(w http.ResponseWriter, r *http.Request) {
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
    if err != nil {
        panic("Bad body")
    }
//    fmt.Println("Found this in the body: ")
//		fmt.Println(body)
    err = json.Unmarshal(body, &newProfile)
				// Unmarshal will only update the fields of the keys sent to it, not all fields.
		if  err == io.EOF {
        fmt.Println("end of file.")
      } else if err != nil {
          fmt.Println("Bad marshal: ")
					fmt.Println(err)
      }

		fmt.Print("newProfile = ")
		fmt.Println(newProfile)	// for debugging purposes
		if newProfile.Is_smoking != "yes" && newProfile.Is_smoking != "no" {
			fmt.Println("Bad value for Is_smoking: " + newProfile.Is_smoking)   }

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
