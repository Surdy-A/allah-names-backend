package main

import (
	"database/sql"
	"net/http"

	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Name struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Meaning string `json:"meaning"`
	Image   string `json:"image"`
	Usage   string `json:"usage"`
}

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", GetNames)

	r.GET("/name/:id", GetName)
	r.POST("/", CreateName)
	r.PUT("/", EditName)


	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

// getAlbums responds with the list of all albums as JSON.
func GetNames(c *gin.Context) {
	db, err := sql.Open("postgres", "user=postgres password=Goodman8349** host=127.0.0.1 port=5432 dbname=allah_names sslmode=disable")

	if err != nil {
		log.Fatal("Error ", err)
	}

	rows, err := db.Query("SELECT * FROM names")

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}

	defer rows.Close()

	var names []Name

	for rows.Next() {
		nm := Name{}
		err := rows.Scan(&nm.ID, &nm.Name, &nm.Meaning, &nm.Image, &nm.Usage)

		if err != nil {
			log.Fatal(err)
		}

		names = append(names, nm)
	}
	c.JSON(http.StatusOK, names)
}

// getname responds with the list of all albums as JSON.
func GetName(c *gin.Context) {
	db, err := sql.Open("postgres", "user=postgres password=Goodman8349** host=127.0.0.1 port=5432 dbname=allah_names sslmode=disable")

	if err != nil {
		log.Fatal("Error ", err)
	}

	id := c.Param("id")

	rows, err := db.Query("SELECT * FROM names WHERE id=$1", id)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}

	defer rows.Close()

	var names []Name

	for rows.Next() {
		nm := Name{}
		err := rows.Scan(&nm.ID, &nm.Name, &nm.Meaning, &nm.Image, &nm.Usage)

		if err != nil {
			log.Fatal(err)
		}

		names = append(names, nm)
	}
	c.JSON(http.StatusOK, names)
}

//postName adds a name from JSON received in the request body.
func CreateName(c *gin.Context) {
	db, err := sql.Open("postgres", "user=postgres password=Goodman8349** host=127.0.0.1 port=5432 dbname=allah_names sslmode=disable")

	if err != nil {
		log.Fatal("Error: ", err)
	}

	var newName Name
	 names := []Name{}

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newName); err != nil {
		return
	}
	fmt.Println(newName)

    sqlStatement := `INSERT INTO names (id, name, meaning, image, usage) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	
	query, err := db.Prepare(sqlStatement)

	if err != nil {
		log.Fatal("Error:--- ", err)
	}

	_, err = query.Exec(newName.ID, newName.Name, newName.Meaning, newName.Image, newName.Usage)
	if err != nil {
		panic(err)
	} else{
		fmt.Println("The value was successfully inserted!")
	}

	defer db.Close()
	// Add the new album to the slice.
	names = append(names, newName)
	c.IndentedJSON(http.StatusCreated, names)
}



//postName adds a name from JSON received in the request body.
func EditName(c *gin.Context) {
	db, err := sql.Open("postgres", "user=postgres password=Goodman8349** host=127.0.0.1 port=5432 dbname=allah_names sslmode=disable")

	if err != nil {
		log.Fatal("Error: ", err)
	}

	var newName Name
	 names := []Name{}

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newName); err != nil {
		return
	}
	fmt.Println(newName)

    sqlStatement := `UPDATE names SET id, name, meaning, image, usage WHHERE id =$1;`
	
	query, err := db.Prepare(sqlStatement)

	if err != nil {
		log.Fatal("Error:--- ", err)
	}

	_, err = query.Exec(newName.ID, newName.Name, newName.Meaning, newName.Image, newName.Usage, newName.ID)
	if err != nil {
		panic(err)
	} else{
		fmt.Println("The value was successfully inserted!")
	}

	defer db.Close()
	// Add the new album to the slice.
	names = append(names, newName)
	c.IndentedJSON(http.StatusCreated, names)
}