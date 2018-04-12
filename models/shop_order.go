package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ShopOrder struct {
	Id              int       `orm:"column(id);auto" description:"主键id"`
	Receivable      float32   `orm:"column(receivable)" description:"应收款"`
	PayType         string    `orm:"column(pay_type);size(12)" description:"支付方式"`
	StoreId         int       `orm:"column(store_id)" description:"店铺id"`
	DayOrderSeq     int       `orm:"column(day_order_seq)" description:"day_order_seq"`
	TotalNoDiscount float32   `orm:"column(total_no_discount)"`
	OrderId         string    `orm:"column(order_id);size(225)" description:"订单id"`
	GoodsCatagory   int       `orm:"column(goods_catagory)"`
	SmallChange     float32   `orm:"column(small_change)"`
	WriteUid        int       `orm:"column(write_uid)" description:"Last Updated by"`
	Discount        float32   `orm:"column(discount)" description:"折扣"`
	Status          int       `orm:"column(status);null" description:"status"`
	PayedTotal      float32   `orm:"column(payed_total)"`
	WriteDate       time.Time `orm:"column(write_date);type(timestamp);auto_now_add"`
	CreateUid       int       `orm:"column(create_uid);null"`
	Operator        string    `orm:"column(operator);size(125);null"`
	CreateDate      time.Time `orm:"column(create_date);type(timestamp)"`
	MemberId        int64     `orm:"column(member_id);null" description:"会员id"`
}

func (t *ShopOrder) TableName() string {
	return "shop_order"
}

func init() {
	orm.RegisterModel(new(ShopOrder))
}

// AddShopOrder insert a new ShopOrder into database and returns
// last inserted Id on success.
func AddShopOrder(m *ShopOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetShopOrderById retrieves ShopOrder by Id. Returns error if
// Id doesn't exist
func GetShopOrderById(id int) (v *ShopOrder, err error) {
	o := orm.NewOrm()
	v = &ShopOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllShopOrder retrieves all ShopOrder matches certain condition. Returns empty list if
// no records exist
func GetAllShopOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ShopOrder))
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

	var l []ShopOrder
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

// UpdateShopOrder updates ShopOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateShopOrderById(m *ShopOrder) (err error) {
	o := orm.NewOrm()
	v := ShopOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteShopOrder deletes ShopOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteShopOrder(id int) (err error) {
	o := orm.NewOrm()
	v := ShopOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ShopOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
