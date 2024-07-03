package main

import(
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	// get visitor anem
	vName := r.URL.Query().Get("visitor_name")
	if vName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing vistor_name parameter")
	}

	// get client IP addres
	clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		clientIP = "Unkown"
	}

	// simulate location
	location := "Ebonyi"

	greeting  := fmt.Sprintf("Hello, %s! The temperature is 11 degress Celcius in %s", vName, location)
	responseData := map[string]string{
		"client_ip": clientIP,
		"location": location,
		"greeting": greeting,
	}

	// Encoding response data to JSON
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error creating responses: %v", err)
		return
	}
	// Write JSON response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main ()  {
	// define the handler func for th api/hello route
	http.HandleFunc("/api/hello", handler)
	
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the server on port 8080
	fmt.Println("Server listening on port %s \n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}