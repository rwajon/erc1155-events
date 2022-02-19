package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

// Ping godoc
// @Summary get transactions
// @Description get transactions list
// @Tags transactions
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Transaction
// @Router /transactions [get]
func GetTransactions(c *gin.Context) {
	filters := bson.M{}
	var page, perPage int64 = 1, 100

	if c.Request != nil {
		for key, value := range c.Request.URL.Query() {
			if len(value) > 0 {
				switch key {
				case "date":
					d1, _ := dateparse.ParseAny(value[0])
					d2, _ := dateparse.ParseAny(fmt.Sprintf("%d-%d-%d 23:59:59", d1.Year(), d1.Month(), d1.Day()))
					filters[strings.ToLower(key)] = bson.M{"$gte": d1, "$lte": d2}
				case "page":
					page = utils.StringToInt(value[0])
				case "limit", "perPage":
					perPage = utils.StringToInt(value[0])
				default:
					filters[strings.ToLower(key)] = bson.M{"$regex": value[0], "$options": "im"}
				}
			}
		}
	}

	result, err := helpers.DBFindManyAndCount(transactionCollection, filters, &options.FindOptions{
		Skip: func() *int64 {
			p := page - 1
			if p < 0 {
				p = 1
			}
			return &p
		}(),
		Limit: &perPage,
	})

	var data []models.Transaction

	if err == nil {
		err = json.Unmarshal(utils.Jsonify(result["data"]), &data)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Error:   err,
		})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusNotFound, models.Error{
			Code:    http.StatusNotFound,
			Message: "no transactions found",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
		Meta:    map[string]interface{}{"page": page, "perPage": perPage, "total": result["count"]},
	})
}

// Ping godoc
// @Summary get single transaction
// @Description get single transaction
// @Tags transactions
// @Accept */*
// @Produce json
// @Param   hash     path    string     true        "transaction hash"
// @Success 200 {object} models.Transaction
// @Failure 404 {object} models.Error "no transaction with hash: {hash} found"
// @Router /transactions/{hash} [get]
func GetOneTransaction(c *gin.Context) {
	result, err := helpers.DBFindOne(transactionCollection, bson.M{
		"hash": bson.M{"$regex": c.Param("hash"), "$options": "im"},
	})

	var data models.Transaction

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

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, models.Error{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("no transaction with hash: \"%s\" found", c.Param("hash")),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
