package service

import (
	"log"
	"net/http"
	"httpserve/natsstreaming"
	"github.com/gin-gonic/gin"
	
)

type Request struct {
	Channel, Message, ClientId string
}

func PushMsg(c *gin.Context) {
	
	var param Request

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":          1,
				"error_message": err.Error(),
			},
		)
		return
	}

	if param.Channel == "" || param.Message == "" || param.ClientId == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"code":          1,
				"error_message": "params should not empty!",
			},
		)
		return
	}
	log.Println(natsstreaming.Conn)
	natsstreaming.Conn.Publish(param.Channel, []byte(param.Message))

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":          0,
			"error_message": "",
			"data":          param,
		},
	)
}
