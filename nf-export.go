package main

import (
	influxdb "github.com/nerdalert/nfexport/Godeps/_workspace/src/github.com/influxdb/influxdb/client"
)

func main() {
	client := NewInfluxClient(false)
	err := GetTables(client)
	if err != nil {
		log.Error("Error retrieving influxdb tables: ", err)
	}

	col := []float32{0, 1, 2, 3}
	err = WritePoints(client, col)
	if err != nil {
		log.Error("Error writing a point tables: ", err)
	}
	log.Info("Post-Write!")
}

func NewInfluxClient(compression bool) influxdb.Client {
	c, err := influxdb.NewClient(&influxdb.ClientConfig{
		Host:     "192.168.59.103:8086",
		Username: "root",
		Password: "root",
		Database: "data",
	})
	if err != nil {
		panic(err)
	}
	return *c
}

func WritePoints(c influxdb.Client, cols []float32) error {
	series := &influxdb.Series{Name: "data",
		Columns: []string{"col1", "col2", "col3", "col4"},
		Points: [][]interface{}{
			[]interface{}{cols[0], cols[1], cols[2], cols[3]},
		}}
	return c.WriteSeries([]*influxdb.Series{series})
}

func GetTables(c influxdb.Client) error {
	tables, err := c.GetDatabaseList()
	if err != nil {
		log.Errorf("Oh noes --> %s", err)
		return err
	}
	prettyPrint(tables, "json")
	return nil
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
