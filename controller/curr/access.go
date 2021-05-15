package curr

import (
	"bytes"
	"net/http"
	"strconv"
	"sync"
	"time"
	"zhou/tools/statecode"

	"github.com/gin-gonic/gin"
)

type Access struct {
	Url      string `json:"url"`
	Method   string `json:"method"`
	Number   string `json:"number"`
	PostBody string `json:"post_body"`
}

func (a *Access) HttpRewrite(c *gin.Context) {

	// 参数校验
	if a.Url == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    statecode.ErrorParamEmpty,
			"content": "",
			"message": "请求URL不能为空",
		})
		return
	}

	number, err := strconv.Atoi(a.Number)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    statecode.ErrorParamNotAllow,
			"content": "",
			"message": err,
		})
		return
	}

	if number < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    statecode.ErrorParamNotAllow,
			"content": "",
			"message": "number 值不合法",
		})
		return
	}

	if a.Method == "GET" {
		result := make([]string, 0, number)
		wg := sync.WaitGroup{}
		wg.Add(number)
		for i := 0; i < number; i++ {
			go httpGet(a, &result, &wg)
		}

		wg.Wait()
		c.JSON(http.StatusOK, gin.H{
			"code":    statecode.Success,
			"content": &result,
			"message": "ok",
		})
		return

	}

	if a.Method == "POST" {

		result := make([]string, 0, number)
		wg := sync.WaitGroup{}
		wg.Add(number)
		for i := 0; i < number; i++ {
			go httpPost(a, &result, &wg)
		}

		wg.Wait()
		c.JSON(http.StatusOK, gin.H{
			"code":    statecode.Success,
			"content": &result,
			"message": "ok",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    statecode.Success,
		"content": "",
		"message": "ok",
	})
}

func httpGet(a *Access, result *[]string, wg *sync.WaitGroup) {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(a.Url)
	if err != nil {
		*result = append(*result, err.Error())

		wg.Done()
		return
	}

	defer resp.Body.Close()
	*result = append(*result, time.Now().Format("2006-01-02 15:04:05.000000"))
	// rs, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(rs))
	// *result = append(*result, string(rs))
	wg.Done()
}

func httpPost(a *Access, result *[]string, wg *sync.WaitGroup) {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Post(a.Url, "application/json", bytes.NewBuffer([]byte(a.PostBody)))
	if err != nil {
		*result = append(*result, err.Error())

		wg.Done()
		return
	}

	defer resp.Body.Close()
	*result = append(*result, time.Now().Format("2006-01-02 15:04:05.000000"))
	// rs, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(rs))
	// *result = append(*result, string(rs))
	wg.Done()
}
