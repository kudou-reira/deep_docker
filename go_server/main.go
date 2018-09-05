package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	// routes consist of a path and a handler function.
	// r.HandleFunc("/xdccAnilist", xdccTempSearch).Methods("GET")

	r.HandleFunc("/", xdccRoot).Methods("GET")

	// bind to a port and pass our router in
	http.Handle("/", &middleWareServer{r})

	log.Fatal(http.ListenAndServe(":8080", nil))
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

// func xdccAnilist(w http.ResponseWriter, r *http.Request) {
// 	if r != nil {
// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		query := r.Form

// 		tempSeason := query["season"][0]
// 		tempStartDate := query["year"][0]

// 		fmt.Println("this is the tempSeason", tempSeason)
// 		fmt.Println("this is the tempStartDate", tempStartDate)

// 		a := anilistMain(tempSeason, tempStartDate)

// 		m := dataMedia{Media: a}

// 		mediaJSON, err := json.Marshal(m)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(mediaJSON)

// 		// temp := []byte("this is xdccAnilist")
// 		// w.Write(temp)
// 	}
// }

func xdccRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	temp := []byte("welcome to the server")
	w.Write(temp)
}

type middleWareServer struct {
	r *mux.Router
}

func (s *middleWareServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}