package yyui

import (
	"pddApp/common"
	"strings"
)

type Goods struct {
	CatId           int64    `json:"CatId"`          //叶子类目ID
	CostTemplateId  int64    `json:"CostTemplateId"` //物流运费模板ID
	CountryId       int      `json:"CountryId"`      //地区/国家ID
	GoodsName       string   `json:"GoodsName"`      // 商品标题
	GoodsDesc       string   `json:"GoodsDesc"`      //商品描述
	GoodsType       int      `json:"GoodsType"`      //1-国内普通商品
	IsLowPrice      bool     `json:"IsLowPrice"`     //
	IsOnline        bool     `json:"IsOnline"`
	DetailGallery   []string `json:"DetailGallery"`
	CarouselGallery []string `json:"CarouselGallery"`
}

func (s *ShowInput) GetGoods() (goods []Goods, err error) {
	goodsMap, err := common.GetGoodsMap(s.ShopExcel.Text, s.ShopSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: " + err.Error())
		return
	}
	goodsImageMap, err := common.GetGoodsComparison(s.ModelExcel.Text, s.ModelSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: " + err.Error())
		return
	}
	goodsConfig, err := common.GetGoodsConfig(s.SkuExcel.Text, s.SkuSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: " + err.Error())
		return
	}
	pubDir := s.PubFileDir.Text
	picDir := s.PicKitDir.Text
	for k, v := range goodsMap {
		var modelList []string
		var skuList []string
		var isLowPriceList []string
		var isOnlineList []string
		var isLowPrice bool
		var isOnline bool
		var imageDir string
		for _, l := range v {
			key := strings.ToLower(l.Model)
			val, ok := goodsImageMap[key] // 从map查找图片目录是否存在
			if ok {
				is, _ := common.IsPathExists(s.PicKitDir.Text + "/" + *val.PicDir)
				if is {
					imageDir = *val.PicDir
					break
				}
			}
		}
		for _, l := range v {
			modelList = append(modelList, l.Model)
			skuList = append(skuList, l.SkuDisplay)
			isLowPriceList = append(isLowPriceList, l.IsLowPrice)
			isOnlineList = append(isOnlineList, l.IsOnline)

		}
		if common.IsEleExistsSlice("低价", isLowPriceList) {
			isLowPrice = true
		}
		if common.IsEleExistsSlice("是", isOnlineList) {
			isOnline = true
		}
		goodsDetailGallery := make([]string, len(goodsConfig.DetailGalleryConfigList))
		goodsCarouselGallery := make([]string, len(goodsConfig.CarouselGalleryConfigList))
		for _, i := range goodsConfig.DetailGalleryConfigList {
			if i.IsPublic {
				goodsDetailGallery[i.Num-1] = pubDir + "/" + i.FileName + ".jpg"
			} else {
				goodsDetailGallery[i.Num-1] = picDir + "/" + imageDir + "/" + i.FileName + ".jpg"
			}
		}
		for _, i := range goodsConfig.CarouselGalleryConfigList {
			goodsCarouselGallery[i.Num-1] = picDir + "/" + imageDir + "/" + i.FileName + ".jpg"
		}
		goods = append(goods, Goods{
			CatId:           1234,
			CostTemplateId:  1234,
			CountryId:       1234,
			GoodsName:       k,
			GoodsDesc:       k,
			GoodsType:       1,
			IsOnline:        isOnline,
			IsLowPrice:      isLowPrice,
			DetailGallery:   goodsDetailGallery,
			CarouselGallery: goodsCarouselGallery,
		})
	}
	return goods, nil
}
