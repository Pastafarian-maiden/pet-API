package internal

import "github.com/gin-gonic/gin"

type name struct {
	id        string `json:",omitempty"`
	name      string `json:",omitempty"`
	addressId string `json:",omitempty"`
}

type address struct {
	id   string `json:",omitempty"`
	name string `json:",omitempty"`
}

func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func GetNames(c *gin.Context)     {}
func PostNames(c *gin.Context)    {}
func GetAddresses(c *gin.Context) {}
func PostAdresses(c *gin.Context) {}
