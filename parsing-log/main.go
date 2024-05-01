package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("/home/shivam/Desktop/30-April-2024 17-Response Processor.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	jsonStr := "[{}"

	fmt.Println("scanning file: ", file.Name())

	start := time.Now()

	words := ""

	count := 0

	for scanner.Scan() {

		line := scanner.Text()

		words = words + line

		if count == 150 {

			if strings.Contains(line, "metric poll response: {") {

				jsonStr += ",{"

			} else {

				jsonStr += line
			}

			count = 0

		}

		count++

	}

	jsonStr += "]"

	fmt.Println("End :", time.Since(start))

	var dat []interface{}

	if err := json.Unmarshal([]byte(jsonStr), &dat); err != nil {
		panic(err)
	}

	objectTarget := os.Args[1]

	metricType := os.Args[2]

	for _, d := range dat {
		map1 := d.(map[string]interface{})

		if len(map1) == 0 {
			continue
		}

		if map1["object.target"].(string) == objectTarget && map1["metric.name"].(string) == metricType {

			timestamp := map1["event.timestamp"].(float64)
			unixSeconds := int64(timestamp)
			unixNanoseconds := int64((timestamp - float64(unixSeconds)) * 1e9)
			t := time.Unix(unixSeconds, unixNanoseconds)
			humanReadableTime := t.Format("02-January-2006 03:04:05.000 pm")

			result, _ := json.Marshal(map1["result"])

			fmt.Printf("%v \t %v\t%v\n", humanReadableTime, map1["status"], string(result))
		}

	}

}
