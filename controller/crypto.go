package controller

import (
	"net/http"
	"testGin/httputil"
	"testGin/micro"
	"testGin/model"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary 查询sql语句
// @Description 根据sql语句查询对应结果
// @Tags 数据
// @Accept json
// @Produce json
// @Param searchSql body model.SearchSql true "根据sql查询结果"
// @Success 200 {object} model.SearchResponse "查询结果集，包括flowID、txID"
// @Failure 400 {object} httputil.HTTPError
// @Router /querySql [post]
func (c *Controller) SearchResults(ctx *gin.Context) {
	searchSql := new(model.SearchSql)
	if err := ctx.ShouldBindJSON(searchSql); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	// 校验请求的入参
	if err := searchSql.Validation(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	// 调用micro client，进行实际逻辑处理
	res, err := micro.GetResults(searchSql)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
