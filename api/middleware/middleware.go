package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// Queue
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Que declaration
var Que []string
var QueItems []*Queue
//var Que []Queue

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	//var queue = []Queue
	var queue []*Queue
	//qq := Queue{}
	for scanner.Scan() {
		//newQueue := Queue{}
		if scanner.Text() == "" {
			continue
		}

		var nodeQueue = &Queue{}
		nodeQueue.Domain = scanner.Text()
		scanner.Scan()
		//nodeQueue.Weight = 1

		//nodeQueue.Weight = scanner.Text()
		nodeQueue.Weight, _ = strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
		scanner.Scan()
		nodeQueue.Weight, _ = strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
		//fmt.Println("IN", scanner.Text())

		queue = append(queue, nodeQueue)

		//fmt.Println("IN", nodeQueue)
	}
	return queue
}


// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}
	//fmt.Println("FROM HEADER", domain)

	nodeItem := &Queue{}
	for _, row := range repo.Read() {
		//fmt.Println("FROM SOURCE", row.Domain)

		//  ALGORITHM HERE...
		//  USE QUE

		if domain == row.Domain {
			nodeItem = &Queue{
				Priority: row.Priority,
				Weight:   row.Weight,
				Domain:   row.Domain,
			}

			if nodeItem.Weight >= 5 && nodeItem.Priority >= 5 {
				fmt.Println("High")

			} else if nodeItem.Weight < 5 && nodeItem.Priority < 5 {
				fmt.Println("Low")
			} else {
				fmt.Println("Medium")
			}
		}
	}


	Que = append(Que, domain)

	c.Next()
}