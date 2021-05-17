package curr

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"zhou/tools/statecode"

	"github.com/gin-gonic/gin"
	. "github.com/klauspost/cpuid/v2"
)

type Access struct {
	Url      string `json:"url"`
	Method   string `json:"method"`
	Number   string `json:"number"`
	PostBody string `json:"post_body"`
}

type feature struct {
	Label string `json:"label"`
	Value string `json:"value"`
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
			"content": gin.H{"result": &result},
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
			"content": gin.H{"result": &result},
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
	resp, err := client.Post(a.Url, "application/json", strings.NewReader(a.PostBody))
	//resp, err := client.Post(a.Url, "application/json", bytes.NewBuffer([]byte(a.PostBody)))
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

func PcFeature(c *gin.Context) {
	s := make([]feature, 0, 12)
	s = append(s, feature{
		Label: "brand",
		Value: CPU.BrandName,
	})
	s = append(s, feature{
		Label: "PhysicalCores",
		Value: strconv.Itoa(CPU.PhysicalCores),
	})

	s = append(s, feature{
		Label: "ThreadsPerCore",
		Value: strconv.Itoa(CPU.ThreadsPerCore),
	})

	s = append(s, feature{
		Label: "LogicalCores",
		Value: strconv.Itoa(CPU.LogicalCores),
	})

	s = append(s, feature{
		Label: "Family",
		Value: strconv.Itoa(CPU.Family),
	})

	s = append(s, feature{
		Label: "CacheLine",
		Value: strconv.Itoa(CPU.CacheLine),
	})

	s = append(s, feature{
		Label: "Model",
		Value: strconv.Itoa(CPU.Model),
	})

	c.JSON(http.StatusOK, gin.H{
		"code":    statecode.Success,
		"content": gin.H{"list": s},
		"message": "ok",
	})
}
