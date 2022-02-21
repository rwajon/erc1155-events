package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rwajon/erc1155-events/db"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()

type NewAddressInWatch struct {
	Address string `json:"address"`
}

// Ping godoc
// @Summary add an address in watch list
// @Description add an address in watch list
// @Tags watch-list
// @Accept */*
// @Produce json
// @Param address body NewAddressInWatch true "new address to watch"
// @Success 201 {object} models.WatchList
// @Failure 400 {object} models.Error
// @Router /watch-list [post]
func AddAddressInWatchList(c *gin.Context) {
	var data models.WatchList

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if validationErr := validate.Struct(&data); validationErr != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Code:    http.StatusBadRequest,
			Message: validationErr.Error(),
		})
		return
	}

	result, err := db.WatchList.Save(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

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

	c.JSON(http.StatusCreated, models.Response{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    map[string]interface{}{"id": result.InsertedID, "address": data.Address},
	})
}

// Ping godoc
// @Summary get watch list
// @Description get watch list list
// @Tags watch-list
// @Accept */*
// @Produce json
// @Success 200 {object} []models.WatchList
// @Failure 404 {object} models.Error "no address in watch list found"
// @Router /watch-list [get]
func GetWatchList(c *gin.Context) {
	filter := bson.M{}
	var page, perPage int64 = 1, 100

	if c.Request != nil {
		for key, value := range c.Request.URL.Query() {
			if len(value) > 0 {
				switch key {
				case "page":
					page = utils.StringToInt(value[0])
				case "limit", "perPage":
					perPage = utils.StringToInt(value[0])
				default:
					filter[strings.ToLower(key)] = bson.M{"$regex": value[0], "$options": "im"}
				}
			}
		}
	}

	result, err := db.WatchList.GetManyAndCount(filter, &options.FindOptions{
		Skip: func() *int64 {
			p := page - 1
			if p < 0 {
				p = 1
			}
			return &p
		}(),
		Limit: &perPage,
	})

	var data []models.WatchList

	if err == nil {
		err = json.Unmarshal(utils.Jsonify(result.Data), &data)
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
			Message: "no address in watch list found",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
		Meta:    map[string]interface{}{"page": page, "perPage": perPage, "total": result.Count},
	})
}

// Ping godoc
// @Summary get single address in watch list
// @Description get single address in watch list
// @Tags watch-list
// @Accept */*
// @Produce json
// @Param   address     path    string     true        "address in watch list"
// @Success 200 {object} models.WatchList
// @Failure 404 {object} models.Error "address {address} not found in watch list"
// @Router /watch-list/{address} [get]
func GetOneAddressWatchList(c *gin.Context) {
	var data models.WatchList

	result, err := db.WatchList.GetOne(bson.M{
		"address": bson.M{"$regex": c.Param("address"), "$options": "im"},
	})

	if result == nil && err == nil {
		c.JSON(http.StatusNotFound, models.Error{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("address \"%s\" not found in watch list", c.Param("address")),
		})
		return
	}

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
	})
}

// Ping godoc
// @Summary add an address in watch list
// @Description add an address in watch list
// @Tags watch-list
// @Accept */*
// @Produce json
// @Param   addressId     path    string     true        "address ID to update"
// @Param address body NewAddressInWatch true "new address to watch"
// @Success 200 {object} models.WatchList
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error "address {addressId} not found in watch list"
// @Router /watch-list/{addressId} [put]
func UpdateAddressInWatchList(c *gin.Context) {
	var data models.WatchList

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if validationErr := validate.Struct(&data); validationErr != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Code:    http.StatusBadRequest,
			Message: validationErr.Error(),
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(c.Param("addressId"))
	result, err := db.WatchList.UpdateOne(bson.M{"_id": id},
		bson.D{{"$set", bson.D{{"address", data.Address}}}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    result.ModifiedCount,
	})
}

// Ping godoc
// @Summary delete an address in watch list
// @Description delete an address in watch list
// @Tags watch-list
// @Accept */*
// @Produce json
// @Param   addressId     path    string     true        "address ID to delete"
// @Success 200 {object} nil
// @Failure 404 {object} models.Error "address {addressId} not found in watch list"
// @Router /watch-list/{addressId} [delete]
func DeleteAddressInWatchList(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("addressId"))
	result, err := db.WatchList.DeleteOne(bson.M{"_id": id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    result.DeletedCount,
	})
}
