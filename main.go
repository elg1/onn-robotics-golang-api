package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "firebase.google.com/go"
    "firebase.google.com/go/db"
    "github.com/gorilla/mux"
    "google.golang.org/api/option"
)

//const (
//    firebaseConfigFile = "path/to/your/firebaseConfig.json"
//    firebaseDBURL      = "gs://onn-robotics-golang-api.appspot.com"
//)

// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
	apiKey: "AIzaSyACfliiKDfB7twq3IWhI_7TNsxV4IJ1gb8",
	authDomain: "onn-robotics-golang-api.firebaseapp.com",
	projectId: "onn-robotics-golang-api",
	storageBucket: "onn-robotics-golang-api.appspot.com",
	messagingSenderId: "566703221787",
	appId: "1:566703221787:web:4884f03a723c21d541dcbe",
	measurementId: "G-W8KK3X3LP0"
  };

var (
    ctx context.Context
    app *firebase.App
)

func main() {
    // Initialize Firebase
    ctx = context.Background()
    opt := option.WithCredentialsFile(firebaseConfigFile)
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("Firebase initialization error: %v\n", err)
    }

    // Initialize Firestore
    client, err := app.DatabaseWithURL(ctx, firebaseDBURL)
    if err != nil {
        log.Fatalf("Firestore initialization error: %v\n", err)
    }

    // Initialize the Router
    router := mux.NewRouter()

    // Define API routes (e.g., /api/books)
    router.HandleFunc("/api/books", getBooks).Methods("GET")
    router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    router.HandleFunc("/api/books", createBook).Methods("POST")
    router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
    router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

    // Start the API server
    port := ":8080"
    fmt.Printf("Server is running on port %s...\n", port)
    log.Fatal(http.ListenAndServe(port, router))
}

// Define API handlers (getBooks, getBook, createBook, updateBook, deleteBook)