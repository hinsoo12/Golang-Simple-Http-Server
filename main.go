package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func welcomeHandler(rw http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/welcome" {
		http.Error(rw, "Ooops the page you are requesting is not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(rw, "Ooops unsupported method!", 404)
		return
	}
	fmt.Fprintf(rw, "Welcome 2F-Capital PLC")
}

func registerHandler(rw http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(rw, "Unable to parse the form :%v", err)
		return
	}

	// must have to be validated before registeration

	fullname := req.FormValue("fullname")
	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")
	address := req.FormValue("address")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	//valid := validator.New()

	fmt.Fprintf(rw, "Successfully registered\n")

	fmt.Fprintf(rw, "Fullname : %s\n", fullname)
	fmt.Fprintf(rw, "Username : %s\n", username)
	fmt.Fprintf(rw, "Email : %s\n", email)
	fmt.Fprintln(rw, "Password :", hashedPassword)
	fmt.Fprintf(rw, "Address : %s\n", address)

}

func main() {
	fmt.Println("^---- 2F-Capital plc --> Simple HTTP Server ----^")

	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("Server is started running")

	log.Fatal(http.ListenAndServe(":8000", nil))

}
