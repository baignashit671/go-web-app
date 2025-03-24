package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Home page accessed") // Debug Log
    http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Courses page accessed") // Debug Log
    http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("About page accessed") // Debug Log
    http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Contact page accessed") // Debug Log
    http.ServeFile(w, r, "static/contact.html")
}

func main() {
    fmt.Println("Starting server on port 8080...")  // Debug Log
    http.HandleFunc("/home", homePage)
    http.HandleFunc("/courses", coursePage)
    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/contact", contactPage)

    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)  // Improved Error Log
    }
}

