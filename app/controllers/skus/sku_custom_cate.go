package skus

import (
	"ditigram_iot_backend/app/controllers"
	"ditigram_iot_backend/app/models"
	"ditigram_iot_backend/app/services/skus"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// SkuCustomCateIndex 获取商品自定义分类列表
func SkuCustomCateIndex(ctx *gin.Context) {
	page, pageSize, step, condition := controllers.GetPage(ctx)
	var filter models.SkuCustomCateListFilter
	err := json.Unmarshal(condition, &filter)
	if err != nil {
		controllers.ResponseError(ctx, controllers.ErrParamsInvalid, err.Error())
		return
	}

	list, count, err := skus.ListSkuCustomCate(page, pageSize, step, filter)
	if err != nil {
		controllers.ResponseError(ctx, fmt.Errorf("【获取商品自定义分类列表】获取列表失败 -> %w", err))
		return
	}
	controllers.ResponseSuccess(ctx, map[string]interface{}{
		"list":  list,
		"count": count,
	})
}

// SkuCustomCateCreate 添加商品自定义分类
func SkuCustomCateCreate(ctx *gin.Context) {
	var req models.SkuCustomCateCreateReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		controllers.ResponseError(ctx, controllers.ErrParamsInvalid, err.Error())
		return
	}

	err = skus.CreateSkuCustomCate(req)
	if err != nil {
		controllers.ResponseError(ctx, fmt.Errorf("【添加商品自定义分类】添加商品自定义分类失败 -> %w", err))
		return
	}
	controllers.ResponseSuccess(ctx)
}

// SkuCustomCateUpdate 添加商品自定义分类
func SkuCustomCateUpdate(ctx *gin.Context) {
	var req models.SkuCustomCateUpdateReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		controllers.ResponseError(ctx, controllers.ErrParamsInvalid, err.Error())
		return
	}

	err = skus.UpdateSkuCustomCate(req)
	if err != nil {
		controllers.ResponseError(ctx, fmt.Errorf("【修改商品自定义分类】修改商品自定义分类失败 -> %w", err))
		return
	}
	controllers.ResponseSuccess(ctx)
}

// SkuCustomCateDelete 删除商品自定义分类
func SkuCustomCateDelete(ctx *gin.Context) {
	var req models.SkuCustomCateDeleteReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		controllers.ResponseError(ctx, controllers.ErrParamsInvalid, err.Error())
		return
	}

	err = skus.DeleteSkuCustomCate(req)
	if err != nil {
		controllers.ResponseError(ctx, fmt.Errorf("【删除商品自定义分类】删除商品自定义分类失败 -> %w", err))
		return
	}
	controllers.ResponseSuccess(ctx)
}
