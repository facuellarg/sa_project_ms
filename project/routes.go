package main

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Api Microservice Arquitectura")
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},

	Route{"AllProjects", "GET", "/projects", AllProjects},
	Route{"PostProject", "POST", "/projects", InsertProject},
	Route{"ProjectbyCode", "GET", "/projects/{id}", GetProjectbyCode},
	Route{"ProjectbyLeader", "GET", "/projects/leader/{id}", GetProjectbLeader},
	Route{"DeleteTag", "DELETE", "/projects/{id}", DeleteProject},
	Route{"EditProjectMembers", "PUT", "/projects/{id}", EditProjectMembers},
}
