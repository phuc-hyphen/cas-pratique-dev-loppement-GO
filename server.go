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
	// "net/url"
	"os"

	"time"
)

//Go application entrypoint
func main() {
	//set up css for html
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static")))) //Go looks in the relative static directory first, then matches it to a

	http.HandleFunc("/", Serverfunc) // set router

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost

	http.HandleFunc("/api/check", check_api)// server used for checking condition 

	// chat room server
	http.HandleFunc("/chatroom", chat_func)

	//Print any errors from starting the webserver using fmt
	go fmt.Println(http.ListenAndServe(":8080", nil))

}

// initial router
func Serverfunc(w http.ResponseWriter, r *http.Request) {

	// redirection page
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Println("template error parsefile : ", err)
		}
		// now redirect to the url
		
		http.Redirect(w, r, "/chatroom", http.StatusFound)
		return
	}

	template, errors := template.ParseFiles("login.html")

	if errors != nil {
		fmt.Println("template error parsefile : ", errors)
	}
	errors = template.Execute(w, nil)
	if errors != nil {
		fmt.Println("template error execution  : ", errors)
	}

	// // add user name in url
	// u, _ := url.Parse("http://localhost:8080/chatroom")
	// fmt.Println("original:",u)
	// user := "tom.jones"
	// u.User = url.User(user)

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

	// if r.URL.Query().Get("step") == "2" {
	// 	// show the form / page described in (3) above.
	// 	// ...
	// 	return

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

// function response resquest
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
	if len(pseudo) < 20 && len(pseudo) > 0 {

		if !contains(used_pseudo, pseudo) { // if the name is not in the list

			//response resquest
			io.WriteString(w, " Pseudo name can be used ")
			fmt.Println(used_pseudo)

			used_pseudo = append(used_pseudo, pseudo) //add new pseudo in the list

			// update file name
			err := writeLines(used_pseudo, "pseudoname.txt")
			if err != nil {
				log.Fatalf("writeLines: %s", err)
			}

			// write in a log file txt
			errors := writeLogs(pseudo, "log.txt", "enter the room")

			if errors != nil {
				log.Fatalf("writeLines: %s", err)
			}

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
		_, errors := file.WriteString(line + "\n")
		if errors != nil {
			log.Fatal(err)
		}
	}
	return nil
}

// write
func writeLogs(pseudo string, path string, action string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	newline := pseudo + "-----" + string(time.Now().Format(time.Stamp)) + "---- " + action

	_, err = fmt.Fprintln(file, newline)
	if err != nil {
		return err
	}
	return nil
}
