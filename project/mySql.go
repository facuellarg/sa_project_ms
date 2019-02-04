package main

import (
	"database/sql"

	"log"
)

func connect() *sql.DB {
	//local 127.0.0.1:4003
	//ms 192.168.99.102:4003
	//cloud 35.227.50.158:4003
	db, err := sql.Open("mysql", "root:1234@tcp(35.227.50.158:4003)/project")

	if err != nil {
		log.Fatal("Could not connect to database")
		log.Fatal("No entro base de datos")

	}

	return db
}

func initDB() {
	db := connect()
	defer db.Close()
	var statemen string
	statemen = "CREATE TABLE IF NOT EXISTS`projects` (`Project_Id` int(11) NOT NULL AUTO_INCREMENT,`Planning_Id` varchar(400) DEFAULT '',`Estado` varchar(400) DEFAULT '',`Members` varchar(500) DEFAULT '',`ProjectLeader` varchar(100) NOT NULL,`Title` varchar(50) DEFAULT '',`StudyArea` varchar(400) DEFAULT '',`Description` varchar(200) DEFAULT '',PRIMARY KEY (`Project_Id`));"
	stmt, err := db.Prepare(statemen)
	if err != nil {
		log.Fatal("Error al crear la tabla projects")
	} else {
		log.Printf("La tabla projects fue creada con exito")
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("Error al crear la tabla projects")
	} else {
		log.Printf("La tabla projects fue creada con exito")
	}

}
