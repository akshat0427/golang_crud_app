package main

import (
	"fmt"
	"net/http"

	"example.com/home-1/app"
	"github.com/gin-gonic/gin"
)

var globalVar string

func main() {
	router := gin.Default()
	// router.LoadHTMLGlob("/templates")
	router.LoadHTMLGlob("templates/*.html")

	// tmpl, err := template.ParseFiles("index.html")

	// if err != nil {
	// 	c.String(http.StatusInternalServerError, "Error parsing HTML template")
	// 	return
	// }

	// Define route to serve the HTML file
	router.GET("/", func(c *gin.Context) {
		

		// Send the HTML file as the response
		c.HTML(200, "index.html", gin.H{})
		// c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/read", func(c *gin.Context) {

		// items := []string{"apple", "banana", "orange", "grape"}

		v, v2 := app.PrintCollection()

		// router.LoadHTMLFiles("read.html")

		// fmt.Println(v2)

		// Render a template passing the list
		c.HTML(http.StatusOK, "read.html", gin.H{
			"Name": v,
			"Id":   v2,

			// "Items": items,
		})

		// ,

		// Example list
		// items := []string{"apple", "banana", "orange", "grape"}

		// v, v2 :=app.PrintCollection()
		// // Render a template passing the list
		// c.HTML(http.StatusOK, "read.html", gin.H{
		// "Name": v,
		// "Data":v2,
		// })

		// c.String(http.StatusOK, "this fine")

		// fmt.Println(v[0], v2[0])

	})

	router.POST("/sub", func(c *gin.Context) {

		title := c.PostForm("title")

		content := c.PostForm("content")

		app.Ins(title, content)

		fmt.Println(title, content)

		c.Redirect(http.StatusSeeOther, "/submit")

	})

	router.GET("/submit", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")

		c.File("submit.html")

	})

	router.GET("/del/:stringValue", func(c *gin.Context) {

		stringValuestr := c.Param("stringValue")

		app.DropDocument(stringValuestr)

		// fmt.Println(stringValuestr)
		c.Redirect(http.StatusSeeOther, "/read")
		// c.String(http.StatusOK, "this fine")

	})

	router.GET("/update/:stv", func(c *gin.Context) {
		globalVar = c.Param("stv")

		// commonv :=

		strvstr := app.Find(c.Param("stv"))
		// router.LoadHTMLFiles("update.html")
		// fmt.Println(strvstr)

		// c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "update.html", gin.H{
			"N": strvstr,
		})
		// c.File("update.html")
		// c.String(http.StatusOK, "this fine")

	})

	router.POST("/edit", func(c *gin.Context) {

		v := c.PostForm("updatedtext")

		app.Update(globalVar, v)

		c.Redirect(http.StatusSeeOther, "/read")

	})
	router.GET("/value/:intValue", func(c *gin.Context) {
		intValueStr := c.Param("intValue")
		fmt.Println(intValueStr)
		c.Redirect(http.StatusSeeOther, "/")
	})

	// Run the server
	router.Run(":8080")
}
