package lxcapp



import (
	"net/http"
	"fmt"
)

type Appstore struct {

}

func LxcappContrinerContrineridConfigGet(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprintf(w, "Hello World! Get")
		w.WriteHeader(http.StatusOK)
}

func LxcappContrinerContrineridConfigPost(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func LxcappContrinerContrineridDelete(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func LxcappContrinerListGet(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func LxcappContrinerListPost(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

