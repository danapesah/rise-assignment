package api

import (
	"github.com/gin-gonic/gin"
	"riseAssignment/db"
)

var metrics Metrics

type Metrics struct {
	ID                  int `bson:"_id"`
	dbConnectionMaxTime int `bson:"db_connection_max_time"`
	dbConnectionMinTime int `bson:"db_connection_min_time"`
	dbConnectionAvgTime int `bson:"db_connection_avg_time"`
	dbConnectionCount   int `bson:"dbConnectionCount"`
	dbReadMaxTime       int `bson:"db_read_max_time"`
	dbReadMinTime       int `bson:"db_read_min_time"`
	dbReadAvgTime       int `bson:"db_read_avg_time"`
	dbReadAvgCount      int `bson:"db_read_avg_count"`
	dbSaveMaxTime       int `bson:"db_save_max_time"`
	dbSaveMinTime       int `bson:"db_save_min_time"`
	dbSaveAvgTime       int `bson:"db_save_avg_time"`
	dbSaveAvgCount      int `bson:"db_save_avg_count"`
}

func loadMetrics(database db.MongoDB) (metrics Metrics) {
	return database.LoadByID(0, "metrics", metrics).(Metrics)
}

func (metrics Metrics) updateConnectionAVGTime(database db.MongoDB, connectionTime int) {
	averageTime := metrics.dbConnectionAvgTime
	count := metrics.dbConnectionCount

	metrics.dbConnectionAvgTime = averageTime + connectionTime/(count+1)
	metrics.dbConnectionCount = count + 1

	if connectionTime > metrics.dbConnectionMaxTime {
		metrics.dbConnectionMaxTime = connectionTime
	}

	if connectionTime < metrics.dbConnectionMinTime {
		metrics.dbConnectionMinTime = connectionTime
	}
}

func (metrics Metrics) updateReadMetrics(database db.MongoDB, connectionTime int) {
	averageTime := metrics.dbReadAvgTime
	count := metrics.dbReadAvgCount

	metrics.dbReadAvgTime = averageTime + connectionTime/(count+1)
	metrics.dbReadAvgCount = count + 1

	if connectionTime > metrics.dbReadMaxTime {
		metrics.dbReadMaxTime = connectionTime
	}

	if connectionTime < metrics.dbReadMinTime {
		metrics.dbReadMinTime = connectionTime
	}
}

func (metrics Metrics) updateSaveAVGTime(database db.MongoDB, connectionTime int) {
	averageTime := metrics.dbSaveAvgTime
	count := metrics.dbSaveAvgCount

	metrics.dbSaveAvgTime = averageTime + connectionTime/(count+1)
	metrics.dbSaveAvgCount = count + 1

	if connectionTime > metrics.dbSaveMaxTime {
		metrics.dbSaveMaxTime = connectionTime
	}

	if connectionTime < metrics.dbSaveMinTime {
		metrics.dbSaveMinTime = connectionTime
	}
}

func GetMetrics(c *gin.Context) {
	database := db.GetDatabase()
	defer database.Disconnect()

	database.LoadByID(0, "metrics", metrics)

	c.JSON(200, gin.H{"metrics": metrics})
}
