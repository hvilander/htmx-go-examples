// All of the stuff for dealing with people in and out of the db
package people 


// probably useful 
// https://go.dev/src/database/sql/sql_test.go?h=TestQuery#L252

import (
	"fmt"
  "log"


  "mywebsite.tv/name/cmd/database"
)


type Person struct {
  FName string
  LName string
}

type People = []Person


func GetAll() People {
 rows, err := database.DBCon.Query(`SELECT fname, lname FROM people`)
 defer rows.Close()

 people := []Person{}
 for rows.Next() {
   var p Person 
   err = rows.Scan(&p.FName, &p.LName)
   if err != nil {
     log.Printf("Scan: %v", err)
   }
   people = append(people, p)
 }

  log.Print("error: ", err)
  log.Print("rows : ", rows)
  return people 
}

func Create(p Person) (int64, error){
  log.Print(p)
  result, err := database.DBCon.Exec(
    "INSERT INTO people (fName, lName) VALUES(?, ?)",
    p.FName, p.LName,
  )
  id, err := result.LastInsertId()
  if err != nil {
    return 0, fmt.Errorf("Create person failed: %v", err)
  }

  return id, nil 
}
  

