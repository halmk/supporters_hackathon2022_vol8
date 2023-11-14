package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/temoto/robotstxt"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf(fmt.Sprintf("$%s must be set", strings.ToUpper("port")))
	}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("APP_ORIGIN"),
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/metadata", func(c *gin.Context) {
		recv_url := c.Query("url")
		log.Println(recv_url)
		if len(recv_url) == 0 {
			c.String(http.StatusBadRequest, "url param required")
			return
		}
		u, err := url.Parse(recv_url)
		if err != nil {
			log.Println("Error parsing received url:", err.Error())
			c.String(http.StatusInternalServerError, "failed to parse url")
			return
		}
		endpoint := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
		robots_url, err := url.Parse(endpoint)
		if err != nil {
			log.Println("Error parsing index url:", err.Error())
			c.String(http.StatusInternalServerError, "failed to parse url")
			return
		}
		robots_url.Path = path.Join(robots_url.Path, "robots.txt")
		resp, err := http.Get(robots_url.String())
		if err != nil {
			log.Println("Error getting robots.txt:", err.Error())
			c.String(http.StatusInternalServerError, "failed to get robots.txt")
			return
		}

		robots, err := robotstxt.FromResponse(resp)
		defer resp.Body.Close()
		if err != nil {
			log.Println("Error parsing robots.txt:", err.Error())
			c.String(http.StatusInternalServerError, "failed to load robots.txt")
			return
		}
		allow := robots.TestAgent(u.Path, "bot")
		if !allow {
			log.Println("Scraping is forbidden")
			c.String(http.StatusBadRequest, "scraping is forbidden")
			return
		}

		doc, err := goquery.NewDocument(recv_url)
		if err != nil {
			log.Println("Error getting document", err.Error())
			c.String(http.StatusInternalServerError, "failed to get document")
			return
		}

		var title string
		var description string
		var image string

		title = doc.Find("title").Text()
		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			property, _ := s.Attr("property")
			name, _ := s.Attr("name")
			content, _ := s.Attr("content")
			if strings.Contains(property, "description") && len(description) == 0 {
				description = content
			} else if strings.Contains(property, "image") && len(image) == 0 {
				image = content
			} else if strings.Contains(name, "description") && len(description) == 0 {
				description = content
			} else if strings.Contains(name, "image") && len(image) == 0 {
				image = content
			}
		})

		fmt.Printf("Title: %s\n", title)
		fmt.Printf("Description: %s\n", description)
		fmt.Printf("Image: %s\n", image)

		c.JSON(http.StatusOK, gin.H{
			"title":       title,
			"description": description,
			"image":       image,
		})
	})

	r.Run(":" + port)
}

