package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/modules"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
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
	now := time.Unix(time.Now().Unix()-86400*14, 0).Format("2006-01-02")

	rows, err := modules.DbConn.Table("api_login_log").Select("count(*) as ips,sum(`count`) as num,login_date").Where("login_date >= ?", now).Group("login_date").Rows()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}
	defer rows.Close()

	var ipsTotal, numTotal, ips, num int
	ipsArr := make([]int, 0)
	numArr := make([]int, 0)
	title := make([]string, 0)
	var loginDate time.Time
	for {
		if !rows.Next() {
			break
		}
		err = rows.Scan(&ips, &num, &loginDate)
		if err != nil {
			logrus.Error(err)
			break
		}
		ipsArr = append(ipsArr, ips)
		ipsTotal += ips

		numArr = append(numArr, num)
		numTotal += num

		title = append(title, loginDate.Format("2006-01-02"))
	}

	// 计算平均

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": []modules.ChartData{{
			Name: fmt.Sprintf("ip(av: %d)", ipsTotal/len(ipsArr)),
			Data: ipsArr,
		}, {
			Name: fmt.Sprintf("pages(av: %d)", numTotal/len(numArr)),
			Data: numArr,
		}},
		"title": title,
	})
}
