package api

import (
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NavJson struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (n NavJson) TableName() string {
	return "nav"
}

type AController struct {
}

func (a *AController) Index(c *gin.Context) {
	// 所有数据
	navList := []models.Nav{}
	//models.DB.Find(&navList)

	// id > 3 的数据
	//models.DB.Where("id > 3").Find(&navList)

	// id > 3 && id < 9
	//var a1 = 3
	//var b = 9
	//models.DB.Where("id > ? AND id < ?", a1, b).Find(&navList)

	// 使用 in 查询 id 在 3, 5, 6中的数据
	//models.DB.Where("id in ?", []int{3, 4, 6}).Find(&navList)

	// 使用 like 查询标题里含有 title 的内容
	//models.DB.Where("title like ?", "%title%").Find(&navList)

	// 使用 id 在 3和9 之间的数据, between and 包括 3和9
	//models.DB.Where("id between ? and ?", 3, 9).Find(&navList)

	// 查询 id = 2 或者 3 的数据 or
	//models.DB.Where("id = ? or id = ?", 2, 3).Find(&navList)

	// 使用 Select 指定返回的字段
	//models.DB.Select("id, title, sort").Find(&navList) // 这种其它字段返回默认值
	//navList1 := []NavJson{}
	//models.DB.Select("id, title").Find(&navList1)

	// Order 排序
	//models.DB.Order("id desc").Order("sort asc").Find(&navList)

	// Limit 限制查询数量 Offset 跳过几条

	// Count 统计总数
	//var num int64
	//models.DB.Find(&navList).Count(&num)

	// 使用原生 sql 删除 nav 的一条数据
	// 执行一个语句
	//models.DB.Exec("delete from nav where id = ?", 5)

	// 使用 sql 修改
	//models.DB.Exec("update nav set title='大明星' where id=?", 5)

	c.JSON(http.StatusOK, gin.H{
		"result": navList,
	})
}

func (a *AController) List(c *gin.Context) {
	c.String(http.StatusOK, "我是一个api接口-list")
}
