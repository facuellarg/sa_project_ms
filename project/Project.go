package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Project struct {
	ProjectID     int      `json:"Proyecto_Id"`
	PlanningID    []int    `json:"Planeacion_Id"`
	Status        string   `json:"Status"`
	Members       []string `json:"Miembros"`
	ProjectLeader string   `json:"Lider_de_proyecto"`
	Title         string   `json:"Titulo"`
	StudyArea     []string `json:"Areas_de_estudio"`
	Description   string   `json:"Descripcion"`
}

func AllProjects(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var projects []Project
	results, err := db.Query("SELECT * FROM projects")
	log.Print("Empezo el get")
	for results.Next() {
		var members, studyAreas, planningID string
		var project Project
		err = results.Scan(&project.ProjectID,
			&planningID, &project.Status, &members, &project.ProjectLeader,
			&project.Title, &studyAreas, &project.Description)
		log.Print("esta en el for")
		if err != nil {
			log.Print("HAY ERROR")
			log.Print(err.Error()) // proper error handling instead of panic in your app
			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select project from database"})
			return
		}
		project.Members = strings.Split(members, ", ")
		project.StudyArea = strings.Split(studyAreas, ", ")
		arrp := strings.Split(planningID, ", ")
		for i := 0; i < len(arrp); i++ {
			num, _ := strconv.Atoi(arrp[i])
			project.PlanningID = append(project.PlanningID, num)
		}

		projects = append(projects, project)
	}

	json.NewEncoder(w).Encode(projects)
}
func GetProjectbyCode(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	projectID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
		return
	}

	var project Project
	var members, studyAreas, planningID string
	// Execute the query
	err = db.QueryRow("SELECT * FROM projects where Project_Id = ?", projectID).Scan(
		&project.ProjectID,
		&planningID, &project.Status, &members, &project.ProjectLeader,
		&project.Title, &studyAreas, &project.Description)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to select project from database"})
		return
	}
	project.Members = strings.Split(members, ", ")
	project.StudyArea = strings.Split(studyAreas, ", ")
	arrp := strings.Split(planningID, ", ")
	for i := 0; i < len(arrp); i++ {
		num, _ := strconv.Atoi(arrp[i])
		project.PlanningID = append(project.PlanningID, num)
	}

	json.NewEncoder(w).Encode(project)
}

func GetProjectbLeader(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	vars := mux.Vars(r)
	l := vars["id"]
	var project Project
	var members, studyAreas, planningID string
	// Execute the query
	err := db.QueryRow("SELECT * FROM projects where ProjectLeader = ?", l).Scan(
		&project.ProjectID,
		&planningID, &project.Status, &members, &project.ProjectLeader,
		&project.Title, &studyAreas, &project.Description)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to select project from database"})
		return
	}
	project.Members = strings.Split(members, ", ")
	project.StudyArea = strings.Split(studyAreas, ", ")
	arrp := strings.Split(planningID, ", ")
	for i := 0; i < len(arrp); i++ {
		num, _ := strconv.Atoi(arrp[i])
		project.PlanningID = append(project.PlanningID, num)
	}

	json.NewEncoder(w).Encode(project)
}

func InsertProject(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	decoder := json.NewDecoder(r.Body)
	var project Project

	err := decoder.Decode(&project)

	if err != nil {
		fmt.Println(project)
		log.Print(err.Error())
		return
	}
	if project.ProjectLeader == "" {
		json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Requerido: Lider_de_proyecto:string"})
		return
	}
	planningID := (arrayToString(project.PlanningID, ", "))
	members := (strings.Join(project.Members, ", "))
	studyAreas := strings.Join(project.StudyArea, ", ")

	stmt, _ := db.Prepare("INSERT INTO projects values (?,?,?,?,?,?,?,?)")
	res, err := stmt.Exec(0,
		planningID,
		project.Status,
		members,
		project.ProjectLeader,
		project.Title,
		studyAreas,
		project.Description)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(project)
		return
	}
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to get last insert id"})
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to get last insert id"})
	}
	project.ProjectID = int(id)
	json.NewEncoder(w).Encode(project)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	projectID := vars["id"]
	ID, err := strconv.Atoi(projectID)
	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
		return
	}

	stmt, err := db.Prepare("DELETE FROM projects where Project_Id = ?")
	if err != nil {
		log.Print(err.Error())
		return
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to delete project from database"})
		return
	}
	json.NewEncoder(w).Encode(ID)

}
func EditProjectMembers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var newProject Project
	err := decoder.Decode(&newProject)
	planningID := (arrayToString(newProject.PlanningID, ", "))
	members := (strings.Join(newProject.Members, ", "))
	studyAreas := strings.Join(newProject.StudyArea, ", ")

	vars := mux.Vars(r)
	projectID := vars["id"]
	ID, err := strconv.Atoi(projectID)
	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
		return
	}

	stmt, _ := db.Prepare("UPDATE projects SET Planning_Id=?,Estado=?,Members = ?,ProjectLeader=?,Title=?,StudyArea=?, Description=? WHERE Project_Id = ?")

	_, err = stmt.Exec(
		planningID,
		newProject.Status,
		members,
		newProject.ProjectLeader,
		newProject.Title,
		studyAreas,
		newProject.Description,
		ID)
	if err != nil {
		log.Print(err.Error())
		return
	}
	newProject.ProjectID = ID
	json.NewEncoder(w).Encode(newProject)

}
