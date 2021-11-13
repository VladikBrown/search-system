package main

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/illfate2/search-system/mq"
	"github.com/illfate2/search-system/search"
)

type MQMessage struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func main() {
	var (
		m sync.RWMutex
		docs []search.DocumentArg
	)
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

	conn, ch, q, err := mq.SetUpMQConnection("amqp://guest:guest@localhost:5672/", "hello")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	defer ch.Close()
	go mq.ReceiveFromQueue(ch, q, func(data []byte) {
		var message MQMessage
		err := json.Unmarshal(data, &message)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Received message: ", message.Text)

		m.Lock()
		docs = append(docs, search.DocumentArg{
			Words: strings.Split(message.Text, " "),
			Name:  message.Name,
		})
		service = search.BuildDocumentService(docs)
		m.Unlock()
	})

	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("query")
		queryWords := strings.Split(query, " ")
		m.RLock()
		results := service.Search(queryWords)
		log.Println(len(docs))
		m.RUnlock()
		c.JSON(http.StatusOK, results)
	})
	r.Run(":3333")
}
