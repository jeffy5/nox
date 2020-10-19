skuCustomCateNS := apiV1.Group("/sku_custom_cate")
{
    skuCustomCateNS.POST("/index", skus.SkuCustomCateIndex)
    skuCustomCateNS.POST("/create", skus.SkuCustomCateCreate)
    skuCustomCateNS.POST("/update", skus.SkuCustomCateUpdate)
    skuCustomCateNS.POST("/delete", skus.SkuCustomCateDelete)
}
