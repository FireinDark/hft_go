package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ShopSalesRecords struct {
	Id           int       `orm:"column(id);auto" description:"主键id"`
	Count        int       `orm:"column(count)"`
	CreateUid    int       `orm:"column(create_uid);null"`
	OperateTime  time.Time `orm:"column(operate_time);type(timestamp);auto_now"`
	ProductId    string    `orm:"column(product_id);size(125);null" description:"产品id"`
	OrderId      string    `orm:"column(order_id);size(225)" description:"订单id"`
	SalesStatus  string    `orm:"column(sales_status);size(12);null"`
	WriteUid     int       `orm:"column(write_uid);null"`
	CreateDate   time.Time `orm:"column(create_date);type(timestamp)"`
	WriteDate    time.Time `orm:"column(write_date);type(timestamp)"`
	Operator     string    `orm:"column(operator);size(125);null"`
	PricePer     float32   `orm:"column(price_per);null"`
	PayType      string    `orm:"column(pay_type);size(125);null" description:"支付方式"`
	Name         string    `orm:"column(name);size(125);null"`
	OperatorName string    `orm:"column(operator_name);size(125);null"`
	StoreId      int       `orm:"column(store_id);null"`
}

func (t *ShopSalesRecords) TableName() string {
	return "shop_sales_records"
}

func init() {
	orm.RegisterModel(new(ShopSalesRecords))
}

// AddShopSalesRecords insert a new ShopSalesRecords into database and returns
// last inserted Id on success.
func AddShopSalesRecords(m *ShopSalesRecords) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetShopSalesRecordsById retrieves ShopSalesRecords by Id. Returns error if
// Id doesn't exist
func GetShopSalesRecordsById(id int) (v *ShopSalesRecords, err error) {
	o := orm.NewOrm()
	v = &ShopSalesRecords{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllShopSalesRecords retrieves all ShopSalesRecords matches certain condition. Returns empty list if
// no records exist
func GetAllShopSalesRecords(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ShopSalesRecords))
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

	var l []ShopSalesRecords
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

// UpdateShopSalesRecords updates ShopSalesRecords by Id and returns error if
// the record to be updated doesn't exist
func UpdateShopSalesRecordsById(m *ShopSalesRecords) (err error) {
	o := orm.NewOrm()
	v := ShopSalesRecords{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteShopSalesRecords deletes ShopSalesRecords by Id and returns error if
// the record to be deleted doesn't exist
func DeleteShopSalesRecords(id int) (err error) {
	o := orm.NewOrm()
	v := ShopSalesRecords{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ShopSalesRecords{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
