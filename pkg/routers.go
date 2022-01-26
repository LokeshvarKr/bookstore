package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters(r *mux.Router) {
	// r.HandleFunc("/book", CreateBook)
	r.HandleFunc("/home", HomePage).Methods("GET")
	r.HandleFunc("/book", GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", GetBookById).Methods("GET")
}

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// }

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Book Store Home Page"))

}
func GetBook(w http.ResponseWriter, r *http.Request) {
	books := ReadBooksFromDB()
	data, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(501)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	book := ReadBookByIdFromDB(bookId)
	if book == nil {
		w.WriteHeader(501)
		w.Write([]byte("Error! Wrong Book Id"))
		return
	}
	data, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(501)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
