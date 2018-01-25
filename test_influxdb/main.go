package main

import (
"log"
"time"

"github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB = "mydb"
	username = ""
	password = ""
)

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://10.33.126.76:9202",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	theTime := time.Now()
	pt1, err := client.NewPoint("cpu_usage", tags, fields, theTime)
	theTime = theTime.Add(time.Second * 1)
	pt2, err := client.NewPoint("cpu_usage", tags, fields, theTime)
	theTime = theTime.Add(time.Second * 1)
	pt3, err := client.NewPoint("cpu_usage", tags, fields, theTime)


	tags1 := map[string]string{"cpu1": "cpu-total1"}
	fields1 := map[string]interface{}{
		"idle":   1.1,
		"system": 5.3,
		"user":   4.6,
	}
	theTime = theTime.Add(time.Second * 1)
	pt4, err := client.NewPoint("cpu_usage", tags1, fields1, theTime)
	theTime = theTime.Add(time.Second * 1)
	pt5, err := client.NewPoint("cpu_usage", tags1, fields1, theTime)
	theTime = theTime.Add(time.Second * 1)
	pt6, err := client.NewPoint("cpu_usage", tags1, fields1, theTime)

	bp.AddPoint(pt1)
	bp.AddPoint(pt2)
	bp.AddPoint(pt3)

	bp.AddPoint(pt4)
	bp.AddPoint(pt5)
	bp.AddPoint(pt6)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
}
