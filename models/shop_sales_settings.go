package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ShopSalesSettings struct {
	Id           int       `orm:"column(id);auto" description:"主键id"`
	CreateUid    int       `orm:"column(create_uid);null"`
	StoreId      int       `orm:"column(store_id);null" description:"店铺id"`
	PrintReceive int8      `orm:"column(print_receive);null"`
	EraseCent    int8      `orm:"column(erase_cent);null"`
	UseAlipay    int8      `orm:"column(use_alipay);null"`
	UseWechat    int8      `orm:"column(use_wechat);null"`
	WriteUid     int       `orm:"column(write_uid);null"`
	WriteDate    time.Time `orm:"column(write_date);type(timestamp);auto_now"`
	CreateDate   time.Time `orm:"column(create_date);type(timestamp)"`
}

func (t *ShopSalesSettings) TableName() string {
	return "shop_sales_settings"
}

func init() {
	orm.RegisterModel(new(ShopSalesSettings))
}

// AddShopSalesSettings insert a new ShopSalesSettings into database and returns
// last inserted Id on success.
func AddShopSalesSettings(m *ShopSalesSettings) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetShopSalesSettingsById retrieves ShopSalesSettings by Id. Returns error if
// Id doesn't exist
func GetShopSalesSettingsById(id int) (v *ShopSalesSettings, err error) {
	o := orm.NewOrm()
	v = &ShopSalesSettings{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllShopSalesSettings retrieves all ShopSalesSettings matches certain condition. Returns empty list if
// no records exist
func GetAllShopSalesSettings(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ShopSalesSettings))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ShopSalesSettings
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateShopSalesSettings updates ShopSalesSettings by Id and returns error if
// the record to be updated doesn't exist
func UpdateShopSalesSettingsById(m *ShopSalesSettings) (err error) {
	o := orm.NewOrm()
	v := ShopSalesSettings{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteShopSalesSettings deletes ShopSalesSettings by Id and returns error if
// the record to be deleted doesn't exist
func DeleteShopSalesSettings(id int) (err error) {
	o := orm.NewOrm()
	v := ShopSalesSettings{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ShopSalesSettings{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
