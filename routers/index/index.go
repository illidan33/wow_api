package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/global"
	"github.com/illidan33/wow_tools/modules"
	"github.com/sirupsen/logrus"
	"net/http"
	"sort"
	"time"
)

func Index(c *gin.Context) {
	html := "index.html"
	modules.CreateLoginLog(c, html, 1)
	c.HTML(http.StatusOK, html, gin.H{
		"apiPage": "home",
	})
}

// 中间件
func AuthMiddleware(c *gin.Context) {
	t, err := c.Cookie("token")
	if err != nil || t == "" {
		c.HTML(http.StatusUnauthorized, "unauth.html", nil)
		c.Abort()
		return
	} else {
		c.Next()
	}
}

func ChartIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "chart.html", gin.H{
		"apiPage": "chart",
	})
}

func GetChartData(c *gin.Context) {
	startDate := c.Query("sd")
	endDate := c.Query("ed")
	var startTime, endTime time.Time
	titleArr := make(map[string]string, 0)
	now := time.Now()

	if startDate == "" {
		u := now.Unix() - 86400*global.Config.ChartDay
		startTime = time.Unix(u, 0)
		startDate = startTime.Format("2006-01-02")
	} else {
		startTime, _ = time.Parse("2006-01-02", startDate)
	}
	if endDate == "" {
		endTime = now
		endDate = endTime.Format("2006-01-02")
	} else {
		endTime, _ = time.Parse(endDate, "2006-01-02")
	}
	for i := startTime.Unix(); i <= endTime.Unix(); i += 86400 {
		t := time.Unix(i, 0)
		d := t.Format("2006-01-02")
		titleArr[d] = d
	}

	htmls, err := modules.DbConn.Table("api_login_log").Select("count(DISTINCT(ip)) as ips,sum(`count`) as num,login_date").Where("login_date >= ? and login_date <= ? and type = ? ", startDate, endDate, 1).Group("login_date").Rows()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}
	defer htmls.Close()

	var ipsTotal, ips, num int
	ipsArr := make(map[string]int, 0)
	htmlNumArr := make(map[string]int, 0)
	apiNumArr := make(map[string]int, 0)
	var loginDate time.Time
	for {
		if !htmls.Next() {
			break
		}
		err = htmls.Scan(&ips, &num, &loginDate)
		if err != nil {
			logrus.Error(err)
			break
		}
		date := loginDate.Format("2006-01-02")

		ipsArr[date] = ips
		ipsTotal += ips

		// 访问数除10，方便查看
		num = num / 10
		htmlNumArr[date] = num
	}

	apis, err := modules.DbConn.Table("api_login_log").Select("count(DISTINCT(ip)) as ips,sum(`count`) as num,login_date").Where("login_date >= ? and login_date <= ? and type = ? ", startDate, endDate, 2).Group("login_date").Rows()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}
	defer apis.Close()
	for {
		if !apis.Next() {
			break
		}
		err = apis.Scan(&ips, &num, &loginDate)
		if err != nil {
			logrus.Error(err)
			break
		}
		date := loginDate.Format("2006-01-02")
		ipsTotal += ips

		if _, ok := ipsArr[date]; ok {
			ipsArr[date] += ips
		} else {
			ipsArr[date] = ips
		}

		// 访问数除10，方便查看
		num = num / 10
		apiNumArr[date] = num
	}

	for _, v := range titleArr {
		if _, ok := htmlNumArr[v]; !ok {
			htmlNumArr[v] = 0
		}
		if _, ok := apiNumArr[v]; !ok {
			apiNumArr[v] = 0
		}
	}
	dateArr := make([]string, 0)
	for _, v := range titleArr {
		dateArr = append(dateArr, v)
	}
	sort.Strings(dateArr);
	ipsA := make([]int, 0)
	htmlsA := make([]int, 0)
	apisA := make([]int, 0)
	for _, date := range dateArr {
		ipsA = append(ipsA, ipsArr[date])
		htmlsA = append(htmlsA, htmlNumArr[date])
		apisA = append(apisA, apiNumArr[date])
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": []modules.ChartData{{
			Name: fmt.Sprintf("IP"),
			Data: ipsA,
		}, {
			Name: fmt.Sprintf("Page/10"),
			Data: htmlsA,
		}, {
			Name: fmt.Sprintf("Api/10"),
			Data: apisA,
		}},
		"title": dateArr,
	})
}
