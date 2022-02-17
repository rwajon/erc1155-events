package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/helpers"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var transactionCollection = config.GetCollection("transactions")

func getDateInterval(d string) (time.Time, time.Time) {
	d1, _ := dateparse.ParseAny(d)
	d2, _ := dateparse.ParseAny(fmt.Sprintf("%d-%d-%d 23:59:59", d1.Year(), d1.Month(), d1.Day()))
	return d1, d2
}

func GetTransactions(c *gin.Context) {
	filters := bson.M{}
	var page int64 = 0
	var perPage int64 = 100

	for key, value := range c.Request.URL.Query() {
		if len(value) > 0 {
			switch key {
			case "date":
				d1, d2 := getDateInterval(value[0])
				filters[strings.ToLower(key)] = bson.M{"$gte": d1, "$lte": d2}
			case "page":
				page = utils.StringToInt(value[0]) - 1
				if page < 0 {
					page = 0
				}
			case "limit", "perPage":
				perPage = utils.StringToInt(value[0])
			default:
				filters[strings.ToLower(key)] = bson.M{"$regex": value[0], "$options": "im"}
			}
		}
	}
	result, err := helpers.DBFindMany(transactionCollection, filters, &options.FindOptions{
		Skip:  &page,
		Limit: &perPage,
	})

	var data []models.Transaction

	if err == nil {
		err = json.Unmarshal(utils.Jsonify(result), &data)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Error:   err,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
		Meta:    map[string]interface{}{"page": page, "perPage": perPage, "total": len(data)},
	})
}
