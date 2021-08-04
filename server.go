package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	// "os/user"
	"net/http"
	"os"

	"time"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

var welcome = Welcome{"Anonymous", time.Now().Format(time.Stamp)}

//Go application entrypoint
func main() {
	//set up css for html 
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static")))) //Go looks in the relative static directory first, then matches it to a

	http.HandleFunc("/", Serverfunc) // set router

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost

	http.HandleFunc("/api/check", check_api) 

	// chat room server
	http.HandleFunc("/chatroom", chat_func)

	http.HandleFunc("/1", redirect)

	//Print any errors from starting the webserver using fmt
	go fmt.Println(http.ListenAndServe(":8080", nil))

}

// initial router
func Serverfunc(w http.ResponseWriter, r *http.Request) {
	template, errors := template.ParseFiles("login.html")

	if errors != nil {
		fmt.Println("template error parsefile : ", errors)
	}
	errors = template.Execute(w, nil)
	if errors != nil {
		fmt.Println("template error execution  : ", errors)
	}

}

//funcion to call chat room platform
func chat_func(w http.ResponseWriter, r *http.Request) {

	template, errors := template.ParseFiles("chat_room.html")

	if errors != nil {
		fmt.Println("template error parsefile : ", errors)
	}
	errors = template.Execute(w, "user")
	if errors != nil {
		fmt.Println("template error execution  : ", errors)
	}

	bs, errors := ioutil.ReadAll(r.Body)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(string(bs))


}

//check if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// 
func check_api(w http.ResponseWriter, r *http.Request) {

	bs, errors := ioutil.ReadAll(r.Body)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusInternalServerError)
		return
	}

	//------------- check the requirement of the pseudo name -----------
	// get list of name from pseudo name file
	used_pseudo, err := readLines("pseudoname.txt")

	if err != nil {
        log.Fatalf("readLines: %s", err)
    }

	fmt.Println(used_pseudo)
	//get pseudo entered
	pseudo := string(bs)

	// check the requirement
	if len(pseudo) < 20 {

		if !contains(used_pseudo, pseudo) { // if the name is not in the list
			
			// http.Redirect(w, r, "/chatroom", http.StatusFound)
			used_pseudo=append(used_pseudo, pseudo)//add new pseudo in the list 
				
			err := writeLines(used_pseudo, "pseudoname.txt") // update file name 

			if err != nil {
				log.Fatalf("writeLines: %s", err)
			}
			go io.WriteString(w," Pseudo name can be used ")
			fmt.Println(used_pseudo)

		} else {
			io.WriteString(w, " Pseudo name have been used !! Please choose another one")
			return
		}
	} else {
		io.WriteString(w, "Pseudo name is too long (limited at 20 characters)")
		return
	}

}


// readLines reads a whole file into memory and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    for _, line := range lines {
        _, errors:=file.WriteString(line + "\n" )
		if errors != nil {
            log.Fatal(err)
        }
    }
    return nil
}

func redirect(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "http://www.google.com", http.StatusFound)
}