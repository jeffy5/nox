package skus

import (
	"ditigram_iot_backend/app/models"
	"ditigram_iot_backend/app/utils/uid"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ListSkuCustomCate 获取商品自定义分类列表
func ListSkuCustomCate(page, pageSize int, step string, filter models.SkuCustomCateListFilter) ([]models.SkuCustomCateListRes, int64, error) {
	list := make([]models.SkuCustomCateListRes, 0)
	var count int64 = 0

	if step == "start" {
		err := setSkuCustomCateListFilter(models.DB.Model(&models.SkuCustomCate{}), filter).Count(&count).Error
		if err != nil {
			return list, count, fmt.Errorf("统计总数失败 -> %w", err)
		}

		if count <= 0 {
			return list, 0, nil
		}
	}

	fields := []string{
		"sku_custom_cate_no", "sku_custom_cate_name", "sku_custom_cate_level", 
	}
	offset := (page - 1) * pageSize
	err := setSkuCustomCateListFilter(
		models.DB.Model(&models.SkuCustomCate{}).Select(fields), filter,
	).Limit(pageSize).Offset(offset).Scan(&list).Error
	if err != nil {
		return list, count, fmt.Errorf("查询失败 -> %w", err)
	}
	return list, count, nil
}

// CreateSkuCustomCate 创建商品自定义分类
func CreateSkuCustomCate(req models.SkuCustomCateCreateReq) error {
	appCategory := models.SkuCustomCate{
		SkuCustomCateNO:       uid.GetDBUidByTableName(models.SkuCustomCateTableName),
        SkuCustomCateName: req.SkuCustomCateName,
        SkuCustomCateLevel: req.SkuCustomCateLevel,
        
	}
	appCategory.SetSkuCustomCateNameI18n(req.SkuCustomCateNameI18n)
	err := models.DB.Model(&models.SkuCustomCate{}).Create(&appCategory).Error
	if err != nil {
		return fmt.Errorf("插入失败 -> %w", err)
	}
	return nil
}

// UpdateSkuCustomCate 更新商品自定义分类
func UpdateSkuCustomCate(req models.SkuCustomCateUpdateReq) error {
	var appCategory models.SkuCustomCate
	appCategory.SetSkuCustomCateNameI18n(req.SkuCustomCateNameI18n)
	updateMap := map[string]interface{}{
        "sku_custom_cate_name": req.SkuCustomCateName,
        "sku_custom_cate_level": req.SkuCustomCateLevel,
        
	}
	err := models.DB.Model(&models.SkuCustomCate{}).Where(
		"sku_custom_cate_no = ? AND sku_custom_cate_status = ?",
		req.SkuCustomCateNO, models.SkuCustomCateStatusNormal,
	).Updates(updateMap).Error
	if err != nil {
		return fmt.Errorf("更新失败 -> %w", err)
	}
	return nil
}

// DeleteSkuCustomCate 删除商品自定义分类
func DeleteSkuCustomCate(req models.SkuCustomCateDeleteReq) error {
	updateMap := map[string]interface{}{
		"sku_custom_cate_status": models.SkuCustomCateStatusDelete,
	}
	err := models.DB.Model(&models.SkuCustomCate{}).Where(
		"sku_custom_cate_no = ? AND sku_custom_cate_status = ?",
		req.SkuCustomCateNO, models.SkuCustomCateStatusNormal,
	).Updates(updateMap).Error
	if err != nil {
		return fmt.Errorf("更新失败 -> %w", err)
	}
	return nil
}

// setSkuCustomCateListFilter 设置获取商品自定义分类列表过滤参数
func setSkuCustomCateListFilter(db *gorm.DB, filter models.SkuCustomCateListFilter) *gorm.DB {
	db = db.Where("sku_custom_cate_status = ?", models.SkuCustomCateStatusNormal)
	return db
}
