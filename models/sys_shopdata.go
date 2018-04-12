package models

import "time"

type SysShopdata struct {
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
	UpDownShelves     int       `orm:"column(up_down_shelves);null" description:"���¼ܣ�0��ʾ�¼ܣ�1��ʾ�ϼ�"`
	ExpiryDate        int       `orm:"column(expiry_date);null" description:"������"`
	Proportion        float32   `orm:"column(proportion);null" description:"�������"`
	IsChild           int       `orm:"column(is_child);null" description:"��С��װ״̬��0��ʾ���װ��1��ʾС��װ"`
	IsConfig          int       `orm:"column(is_config);null" description:"��Ʒ�Ƿ����ã�0��ʾδ���ã�1��ʾ������"`
	UnitId            int       `orm:"column(unit_id);null" description:"��λ"`
	Brevity           string    `orm:"column(brevity);size(50);null" description:"����"`
}
