package main

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/illfate2/search-system/search"
)

func main() {
	var docs []search.DocumentArg
	filepath.Walk("./testdata", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		words := strings.Split(string(file), " ")
		docs = append(docs, search.DocumentArg{
			Words: words,
			Name:  info.Name(),
		})
		return nil
	})
	service := search.BuildDocumentService(docs)
	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("query")
		queryWords := strings.Split(query, " ")
		results := service.Search(queryWords)
		c.JSON(http.StatusOK, results)
	})
	r.Run(":3333")
}
