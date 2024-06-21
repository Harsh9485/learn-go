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
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses/add", addCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	courses = append(courses, Course{
		ID:       "1",
		Name:     "python",
		Date: 	  time.Now(),
		Price:    100,
		Author:   &Author{
			Name: "harsh", 
			Email: "harsh@gmail.com",
			Website: "harsh.com",
		},
	})
	courses = append(courses, Course{
		ID:       "2",
		Name:     "java",
		Date: 	  time.Now(),
		Price:    100,
		Author:   &Author{
			Name:    "harsh", 
			Email:   "harsh@gmail.com",
			Website: "harsh.com",
		},
	})
	http.ListenAndServe(
		":8000",
		r,
	)
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
			return 
		}
	}
	json.NewEncoder(w).Encode("not found course this id")             
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add new course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		fmt.Println("empty body")
		json.NewEncoder(w).Encode("empty body plz send data")
		return
	}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	fmt.Println(course)
	if err!= nil {
		fmt.Println("error", err)
		json.NewEncoder(w).Encode("error")
	}
	fmt.Println("this course",course)
	if course.isEmpty() {
		fmt.Println("empty data")
		json.NewEncoder(w).Encode("empty data")
		return
	}
	rand.Seed(time.Now().UnixNano())
	
	course.ID = strconv.Itoa(rand.Intn(100) + 1)
	course.Date = time.Now()
	courses = append(courses, course)
	fmt.Println("data created")
	json.NewEncoder(w).Encode(course)
} 

func updateCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		fmt.Println("empty body")
		json.NewEncoder(w).Encode("empty body")
		return
	}
	
	for i, course := range courses {
		if course.ID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			if course.isEmpty() {
				fmt.Println("empty body")
				json.NewEncoder(w).Encode("empty body")
				return
			}
			course.ID = params["id"]
			course.Date = time.Now()
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for i, course := range courses {
		if course.ID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("not found course this id")
}
// Create struct
var courses []Course

type Course struct {
	ID     string  `json:"id"`
	Name   string  `json:"courseName"`
	Price  float64 `json:"price"`
	Date   time.Time  `json:"date"`
	Author *Author `json:"author"`
}

type Author struct {
	Name    string `json:"authorName"`
	Email   string `json:"email"`
	Website string `json:"website"`
}
func (c *Course) isEmpty() bool {
	if c.Name == "" {
		return true
	}
	return false
}



 