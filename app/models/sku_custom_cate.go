// SkuCustomCateListFilter 获取商品自定义分类过滤参数
type SkuCustomCateListFilter struct {
    SkuCustomCateName string `json:"sku_custom_cate_name"`
    SkuCustomCateLevel int `json:"sku_custom_cate_level"`
    
}

// SkuCustomCateListRes 获取商品自定义分类响应结果
type SkuCustomCateListRes struct {
    SkuCustomCateName string `json:"sku_custom_cate_name"`
    SkuCustomCateLevel int `json:"sku_custom_cate_level"`
    
}

// SkuCustomCateCreateReq 添加商品自定义分类请求参数
type SkuCustomCateCreateReq struct {
    SkuCustomCateName string `json:"sku_custom_cate_name"`
    SkuCustomCateLevel int `json:"sku_custom_cate_level"`
    
}

// SkuCustomCateUpdateReq 更新商品自定义分类请求参数
type SkuCustomCateUpdateReq struct {
    SkuCustomCateName string `json:"sku_custom_cate_name"`
    SkuCustomCateLevel int `json:"sku_custom_cate_level"`
    
}

// SkuCustomCateDeleteReq 删除商品自定义分类请求参数
type SkuCustomCateDeleteReq struct {
    SkuCustomCateName string `json:"sku_custom_cate_name"`
    SkuCustomCateLevel int `json:"sku_custom_cate_level"`
    
}
