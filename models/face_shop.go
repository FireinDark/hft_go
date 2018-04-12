package models

import "time"

type FaceShop struct {
	ShopId        string    `orm:"column(shop_id);size(40)" description:"店铺id"`
	PersongroupId string    `orm:"column(persongroup_id);size(60)" description:"店铺人脸识别唯一id"`
	GroupName     string    `orm:"column(group_name);size(255);null" description:"group信息"`
	CreateTime    time.Time `orm:"column(create_time);type(datetime);null"`
	IsDeleted     int8      `orm:"column(is_deleted);null" description:"是否被删除，默认0为未删除， 1为已删除"`
	UpdateTime    time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
}
