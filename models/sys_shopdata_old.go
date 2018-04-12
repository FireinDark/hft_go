package models

import "time"

type SysShopdataOld struct {
	ShopId            string    `orm:"column(shop_id);size(20)"`
	Code              string    `orm:"column(code);size(50)"`
	Name              string    `orm:"column(name);size(200);null"`
	Spec              string    `orm:"column(spec);size(200);null"`
	Trademark         string    `orm:"column(trademark);size(100);null"`
	Units             string    `orm:"column(units);size(50);null"`
	FactoryName       string    `orm:"column(factory_name);size(200);null"`
	Addr              string    `orm:"column(addr);size(250);null"`
	Img               string    `orm:"column(img);size(500);null"`
	TradePrice        string    `orm:"column(trade_price);size(14);null"`
	RetailPrice       string    `orm:"column(retail_price);size(14);null"`
	StatusId          int       `orm:"column(status_id);null" description:"0: 处理中， 1： 处理完毕 "`
	UpdateAt          time.Time `orm:"column(updateAt);type(timestamp);auto_now"`
	Repertory         uint      `orm:"column(repertory);null" description:"库存数量"`
	IsStandardBarcode int       `orm:"column(is_standard_barcode);null" description:"是否标准条码 1是 0否"`
	AffirmPrice       int       `orm:"column(affirm_price);null"`
	UpDownShelves     int       `orm:"column(up_down_shelves);null" description:"上下架，0表示下架，1表示上架"`
	ExpiryDate        int       `orm:"column(expiry_date);null" description:"保质期"`
	Proportion        float32   `orm:"column(proportion);null" description:"换算比例"`
	IsChild           int       `orm:"column(is_child);null" description:"大小包装状态，0表示大包装，1表示小包装"`
	IsConfig          int       `orm:"column(is_config);null" description:"商品是否配置，0表示未配置，1表示已配置"`
	UnitId            int       `orm:"column(unit_id);null" description:"单位"`
	Brevity           string    `orm:"column(brevity);size(50);null" description:"简码"`
}
