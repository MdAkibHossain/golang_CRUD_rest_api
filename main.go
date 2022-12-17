package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	courses = append(courses, Course{CourseId: "1", CourseName: "Course1", CoursePrice: 333, Author: &Author{FullName: "Akib", Website: "akib.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Course2", CoursePrice: 444, Author: &Author{FullName: "Hossain", Website: "hossain.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/all_course", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/create_course", createCourse).Methods("POST")
	r.HandleFunc("/update_course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/delete_course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))

}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome</h1>"))
}
func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}
func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No course found")
}
func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Send some data!")
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Enter the course name")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}
func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	params := mux.Vars(r)

	for index, c := range courses {
		if c.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}
	}
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, c := range courses {
		if c.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Id " + params["id"] + " " + "Delete successfull!")
			break
		}
	}
}
