package controller

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Port struct {
	Host      string `json:"host"`
	StartPort string `json:"start_port"`
	EndPort   string `json:"end_port"`
}

var (
	MsgChan = make(chan string, 10)
)

func (p *Port) Scan(c *gin.Context) {
	if "" == p.Host {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"content": "",
			"message": "请输入域名或IP",
			"err":     "",
		})
		return
	}

	startPort, err := strconv.Atoi(p.StartPort)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"content": "",
			"message": "起始端口请输入数字",
			"err":     err.Error(),
		})
		return
	}

	if startPort <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"content": "",
			"message": "起始端口必须大于0",
			"err":     "",
		})
		return
	}

	endPort, err := strconv.Atoi(p.EndPort)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"content": "",
			"message": "截止端口请输入数字",
			"err":     err.Error(),
		})
		return
	}

	if endPort <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"content": "",
			"message": "截止端口必须大于0",
			"err":     "",
		})
		return
	}

	if startPort > endPort {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"content": "",
			"message": "截止端口必须大于起始端口",
			"err":     "",
		})
		return
	}

	go func() {
		if startPort == endPort {
			MsgChan <- scanOne(p.Host, startPort)
		} else {
			scanRange(p.Host, startPort, endPort, MsgChan)
		}
		MsgChan <- "扫描完毕"
	}()

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"content": "",
		"message": "ok",
	})
}

func scanOne(host string, port int) string {
	_, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), time.Duration(time.Second))
	if err != nil {
		return err.Error() + " port: " + strconv.Itoa(port)
	}

	return "端口" + strconv.Itoa(port) + "已开放"
}

func scanRange(host string, startPort, endPort int, msg chan string) {
	for i := startPort; i <= endPort; i++ {

		_, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(i), time.Duration(time.Second))
		if err != nil {
			msg <- err.Error() + " port: " + strconv.Itoa(i)
			continue
		}
		msg <- "端口" + strconv.Itoa(i) + "已开放"
	}

}
