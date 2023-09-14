package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"os"

	"solver/dpll"
)

var PORT = os.Getenv("SERVER_PORT")
var CLIENT = os.Getenv("CLIENT_PORT")
var URL = os.Getenv("URL")

func parseRequest(request string) [][]string {
	var lines [][]string

	for _, clause := range strings.Split(request, "\n") {
		clause = strings.TrimSpace(clause)
		if clause != "" {
			re := regexp.MustCompile(`\s+`)
			literalsList := re.Split(clause, -1)
			lines = append(lines, literalsList)
		}
	}
	return lines
}

func writeResponse(w http.ResponseWriter, jsonResponse []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("%s:%s", URL, CLIENT))
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(jsonResponse)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	bodyStr := string(body)

	lines := parseRequest(bodyStr)
	solution := dpll.RunDPLL(lines)
	jsonResponse, err := json.Marshal(solution)

	if err != nil {
		panic(err.Error())
	}

	writeResponse(w, jsonResponse)

}

func main() {
	http.HandleFunc("/solve", solveHandler)
	fmt.Printf("reading on: %s:%s\n", URL, CLIENT)
	if PORT == "" {
		println(fmt.Errorf("SERVER_PORT env variable is not defined"))
	}

	fmt.Println("Starting server on port", PORT)
	if err := http.ListenAndServe(fmt.Sprint(":", PORT), nil); err != nil {
		println(err.Error())
	}
}
