package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)



func main() {
	fmt.Println("welcome to golang course. talk about how to create server...")

}

// controller file
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang course. talk about how to create server...</h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)                                                                   
}

func getCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("get one course by id ")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, course := range courses {
		if (course.ID == params["id"]){
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("not found course this id")
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add new course")
	if r.Body == nil {
		fmt.Println("empty body")
		json.NewEncoder(w).Encode("empty body")
		return
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		fmt.Println("empty body")
		json.NewEncoder(w).Encode("empty body")
		return
	}
	rand.Seed(time.Now().UnixNano())
	
	course.ID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
} 


// Create struct
var courses []Course

type Course struct {
	ID     string  `json:"id"`
	Name   string  `json:"courseName"`
	Price  float64 `json:"price"`
	Date   string  `json:"date"`
	Author *Author `json:"author"`
}

type Author struct {
	Name    string `json:"authorName"`
	Email   string `json:"email"`
	Website string `json:"website"`
}
func (c *Course) isEmpty() bool {
	return c.Name == ""
}