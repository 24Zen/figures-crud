package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var figures = []Figure{
	{
		1,
		"Ore no Imouto ga Konnani Kawaii Wake ga Nai (OreImo)",
		"Kousaka Kirino",
		[]string{
			"http://localhost:8080/images/Ore no Imouto ga Konnani Kawaii Wake ga Nai/Kousaka-Kirino/Kousaka-Kirino/00Kousaka-Kirino.jpg",
		},
		"Little Sister",
	},

	{
		2,
		"Ore no Imouto ga Konnani Kawaii Wake ga Nai (OreImo)",
		"Gokou Ruri",
		[]string{
			"http://localhost:8080/images/Ore no Imouto ga Konnani Kawaii Wake ga Nai/Gokou-Ruri/Riru.jpg",
		},
		"Little Sister's friend",
	},

	{
		3,
		"The Melancholy of Haruhi Suzumiya",
		"Suzumiya Haruhi",
		[]string{
			"http://localhost:8080/images/The Melancholy of Haruhi Suzumiya/Suzumiya Haruhi/Suzumiya Haruhi/00Suzumiya-Haruhi.jpeg",
		},
		"Leading actress",
	},

	{
		4,
		"The Melancholy of Haruhi Suzumiya",
		"Nagato Yuki",
		[]string{
			"http://localhost:8080/images/The Melancholy of Haruhi Suzumiya/Nagato Yuki/Nagato Yuki/00Nagato-Yuki.jpeg",
		},
		"Leading actress's friend",
	},

	{
		5,
		"The Melancholy of Haruhi Suzumiya",
		"Asahina Mikuru",
		[]string{
			"http://localhost:8080/images/The Melancholy of Haruhi Suzumiya/Asahina Mikuru/Asahina Mikuru/00Asahina-Mikuru.jpeg",
		},
		"Leading actress's friend",
	},

	{
		6,
		"Denpa Onna to Seishun Otoko",
		"Mifune Ryuuko",
		[]string{
			"http://localhost:8080/images/Denpa Onna to Seishun Otoko/Mifune-Ryuuko/00-Mifune-Ryuuko.jpeg",
			"http://localhost:8080/images/Denpa Onna to Seishun Otoko/Mifune-Ryuuko/01-Mifune-Ryuuko.jpeg",
			"http://localhost:8080/images/Denpa Onna to Seishun Otoko/Mifune-Ryuuko/02-Mifune-Ryuuko.jpeg",
			"http://localhost:8080/images/Denpa Onna to Seishun Otoko/Mifune-Ryuuko/03-Mifune-Ryuuko.jpeg",
		},

		"Leading actress's friend",
	},
	{
		7,
		"Vocaloid",
		"Hatsune Miku",
		[]string{
			"http://localhost:8080/images/Vocaloid/Hatsune-Miku/Miku.jpeg",
		},
		"Singer",
	},

	{
		8,
		"Working!!",
		"Yamada Aoi",
		[]string{
			"http://localhost:8080/images/Working!!/Yamada Aoi/Yamada Aoi/Aoi.jpg",
		},
		"Leading actress's friend",
	},

	{
		9,
		"To Aru Majutsu No Index",
		"Index",
		[]string{
			"http://localhost:8080/images/To Aru Majutsu No Index/Index/Index/00Index.jpeg",
		},
		"Leading actress",
	},
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// serve static images
	r.Static("/images", "./images")

	// READ ALL
	r.GET("/figures", func(c *gin.Context) {
		c.JSON(http.StatusOK, figures)
	})

	// READ ONE
	r.GET("/figures/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, f := range figures {
			if f.ID == id {
				c.JSON(http.StatusOK, f)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	// CREATE
	r.POST("/figures", func(c *gin.Context) {
		var newFigure Figure
		if err := c.BindJSON(&newFigure); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newFigure.ID = len(figures) + 1
		figures = append(figures, newFigure)
		c.JSON(http.StatusCreated, newFigure)
	})

	// UPDATE
	r.PUT("/figures/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i := range figures {
			if figures[i].ID == id {
				c.BindJSON(&figures[i])
				figures[i].ID = id
				c.JSON(http.StatusOK, figures[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	// DELETE
	r.DELETE("/figures/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, f := range figures {
			if f.ID == id {
				figures = append(figures[:i], figures[i+1:]...)
				c.Status(http.StatusNoContent)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	r.Run(":8080")
}
