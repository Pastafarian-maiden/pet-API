package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type name struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	AddressId string `json:"addressId,omitempty"`
}

var names = []name{
	{Id: "1", Name: "Pastafarian"},
	{Id: "2", Name: "Galiya", AddressId: "1"},
}

type address struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var addresses = []address{
	{Id: "3", Name: "Pastafarian"},
	{Id: "4", Name: "Galiya"},
}

func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func GetNames(c *gin.Context) {
	c.JSON(http.StatusOK, names)
}

func PostNames(c *gin.Context) {
	var newNames name

	// Call BindJSON to bind the received JSON to newNames.
	if err := c.BindJSON(&newNames); err != nil {
		return
	}
	names = append(names, newNames)
	c.JSON(http.StatusCreated, newNames)
}

func GetAddresses(c *gin.Context) {
	c.JSON(http.StatusOK, addresses)
}

func PostAdresses(c *gin.Context) {
	var newAddress address
	if err := c.BindJSON(&newAddress); err != nil {
		return
	}
	addresses = append(addresses, newAddress)
	c.JSON(http.StatusCreated, newAddress)
}
