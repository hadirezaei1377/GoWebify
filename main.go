package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// normal function
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
