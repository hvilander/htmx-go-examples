package main

import (
  "html/template"
  "io"
  "os"
  "log"
  "database/sql"

  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"

  "mywebsite.tv/name/cmd/database"
)

type Templates struct {
  templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
  return &Templates{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
}

type Count struct {
  Count int
}

type Contact struct {
  Name string
  Email string
}

type Person struct {
  FName string
  LName string
}

func newContact(name, email string) Contact {
  return Contact{
    Name: name,
    Email: email,
  }
}

type Contacts = []Contact

type Data struct {
  Contacts Contacts
}

func newData() Data {
  return Data{
    Contacts: []Contact{
      newContact("John", "jd@gmail.com"),
      newContact("Clara", "cd@gmail.com"),
    },
  }
}


func main() {
  // Process env file
  log.Print("Processing env vars...")
  err := godotenv.Load()

  if err != nil {
    log.Fatalf("error loading env file: %", err)
  }

  dbConnectionString := os.Getenv("DB_CONN_STR")

  // Setup connection to postgresSQL db
  log.Print("setting up db connection...")
  database.DBCon, err = sql.Open("postgres", dbConnectionString)

 
  if err != nil {
    panic(err)
  }

  defer database.DBCon.Close()
  var version string
  
  if err := database.DBCon.QueryRow("select version()").Scan(&version); err != nil {
    panic(err)
  }

  // Print db version
  log.Printf("db connected with version=%s\n", version)

  // Start the server
  e := echo.New()
  e.Use(middleware.Logger())


  data := newData()
  e.Renderer = newTemplate()


  // handlers TODO move these to their own files at some point

  e.GET("/", func(c echo.Context) error {
    // TODO Get people and add them to data



    //allPeople := people.getAll(db)
    


    //people.Test()

    return c.Render(200, "index", data)
  })


  e.POST("/contacts", func(c echo.Context) error {
    name := c.FormValue("name")
    email:= c.FormValue("email")
    data.Contacts = append(data.Contacts, newContact(name, email))

    return c.Render(200, "index", data)
  })

  e.Logger.Fatal(e.Start(":42069"))
}
