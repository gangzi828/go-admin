package heroine

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"net/http"
)

// @Summary 添加女星
// @Description 获取JSON
// @Tags 女星
// @Accept  application/json
// @Product application/json
// @Param data body models.Heroine true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/heroine [post]
// @Security Bearer
func InsertHeroine(c *gin.Context) {
	var data models.Heroine
	err := c.Bind(&data)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c,result,"")
}

// @Summary 女星列表
// @Description 获取JSON
// @Tags 女星
// @Param era query int false "年代"
// @Param name query string false "姓名"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/heroinelist [get]
// @Security
func GetHeroineLogList(c *gin.Context) {
	var data models.Heroine
	var err error
	var pageSize = 10
	var pageIndex = 1

	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	index := c.Request.FormValue("pageIndex")
	if index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Name = c.Request.FormValue("name")

	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize

	var res app.Response
	res.Data = mp
	c.JSON(http.StatusOK, res.ReturnOK())
}
