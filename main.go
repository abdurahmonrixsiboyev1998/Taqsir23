// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// type Book struct {
// 	ID       int    `json:"id"`
// 	Title    string `json:"title"`
// 	Author   string `json:"author"`
// 	Year     int    `json:"year"`
// 	Pages    int    `json:"pages"`
// 	Language string `json:"language"`
// }

// var books []Book
// var jsonFilePath = "books.json"

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/books", getAllBooks).Methods("GET")
// 	r.HandleFunc("/books/{id}", getBookByID).Methods("GET")
// 	r.HandleFunc("/books", createBook).Methods("POST")
// 	r.HandleFunc("/books/{id}", updateBookByID).Methods("PUT")
// 	r.HandleFunc("/books/{id}", deleteBookByID).Methods("DELETE")

// 	loadBooksFromJSON()
// 	port := ":8080"
// 	fmt.Printf("Server is running on port %s...\n", port)
// 	log.Fatal(http.ListenAndServe(port, r))
// }

// func loadBooksFromJSON() {
// 	f, err := os.Open(jsonFilePath)
// 	if err != nil {
// 		books = []Book{}
// 		return
// 	}
// 	defer f.Close()

// 	decoder := json.NewDecoder(f)
// 	if err := decoder.Decode(&books); err != nil {
// 		log.Fatal("Error JSON decoding:", err)
// 	}
// }

// func saveBooksToJSON() {
// 	f2, err := os.Create(jsonFilePath)
// 	if err != nil {
// 		log.Fatal("Error creating JSON file:", err)
// 	}
// 	defer f2.Close()

// 	encoder := json.NewEncoder(f2)
// 	if err := encoder.Encode(books); err != nil {
// 		log.Fatal("Error encoding JSON:", err)
// 	}
// }

// func getAllBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// func getBookByID(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	num := mux.Vars(r)
// 	id, err := strconv.Atoi(num["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid book ID", http.StatusBadRequest)
// 		return
// 	}

// 	for _, book := range books {
// 		if book.ID == id {
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}

// 	http.Error(w, "Book not found", http.StatusNotFound)
// }

// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var newBook Book
// 	res, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Error reading request body", http.StatusBadRequest)
// 		return
// 	}

// 	if err := json.Unmarshal(res, &newBook); err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}

// 	if len(books) == 0 {
// 		newBook.ID = 1
// 	} else {
// 		last := books[len(books)-1]
// 		newBook.ID = last.ID + 1
// 	}

// 	books = append(books, newBook)
// 	saveBooksToJSON()
// 	json.NewEncoder(w).Encode(newBook)
// }

// func updateBookByID(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	num2 := mux.Vars(r)
// 	id, err := strconv.Atoi(num2["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid book ID", http.StatusBadRequest)
// 		return
// 	}

// 	var updatedBook Book
// 	nums, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}

// 	if err := json.Unmarshal(nums, &updatedBook); err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}

// 	for i, book := range books {
// 		if book.ID == id {
// 			books[i] = updatedBook
// 			saveBooksToJSON()
// 			json.NewEncoder(w).Encode(updatedBook)
// 			return
// 		}
// 	}

// 	http.Error(w, "Book not found", http.StatusNotFound)
// }

// func deleteBookByID(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	num3 := mux.Vars(r)
// 	id, err := strconv.Atoi(num3["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid book ID", http.StatusBadRequest)
// 		return
// 	}

// 	for i, book := range books {
// 		if book.ID == id {
// 			books = append(books[:i], books[i+1:]...)
// 			saveBooksToJSON()
// 			w.Write([]byte("Book deleted successfully"))
// 			return
// 		}
// 	}

// 	http.Error(w, "Book not found", http.StatusNotFound)
// }

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Pages    int    `json:"pages"`
	Language string `json:"language"`
}

var books []Book

var jsonFilePath = "books.json"

func loadBooksFromJSON() {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		books = []Book{}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&books); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
}

func saveBooksToJSON() {
	file, err := os.Create(jsonFilePath)
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(books); err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for _, book := range books {
		if book.ID == bookID {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &newBook); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if len(books) == 0 {
		newBook.ID = 1
	} else {
		lastBook := books[len(books)-1]
		newBook.ID = lastBook.ID + 1
	}

	books = append(books, newBook)
	saveBooksToJSON()
	json.NewEncoder(w).Encode(newBook)
}

func updateBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook Book
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &updatedBook); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == bookID {
			books[i] = updatedBook
			saveBooksToJSON()
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == bookID {
			books = append(books[:i], books[i+1:]...)
			saveBooksToJSON()
			w.Write([]byte("Book deleted successfully"))
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {
	num := mux.NewRouter()

	loadBooksFromJSON()

	num.HandleFunc("/books", getAllBooks).Methods("GET")
	num.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	num.HandleFunc("/books", createBook).Methods("POST")
	num.HandleFunc("/books/{id}", updateBookByID).Methods("PUT")
	num.HandleFunc("/books/{id}", deleteBookByID).Methods("DELETE")

	port := ":8080"

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, num))
}
