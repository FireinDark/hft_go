package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysBarcodeOld struct {
	Id                 int       `orm:"column(code);pk"`
	Name               string    `orm:"column(name);size(200);null"`
	Spec               string    `orm:"column(spec);size(200);null"`
	Specnum            int       `orm:"column(specnum);null" description:"规格数量"`
	Trademark          string    `orm:"column(trademark);size(100);null"`
	Addr               string    `orm:"column(addr);size(200);null"`
	Units              string    `orm:"column(units);size(50);null"`
	FactoryName        string    `orm:"column(factory_name);size(200);null"`
	TradePrice         string    `orm:"column(trade_price);size(20);null"`
	RetailPrice        string    `orm:"column(retail_price);size(20);null"`
	UpdateAt           time.Time `orm:"column(updateAt);type(timestamp);null;auto_now"`
	Wholeunit          string    `orm:"column(wholeunit);size(50);null"`
	Wholenum           int       `orm:"column(wholenum);null"`
	Img                string    `orm:"column(img);size(500);null"`
	Src                string    `orm:"column(src);size(20);null"`
	UnitId             int       `orm:"column(unit_id)" description:"单位id 可对于该条码对应单位个数"`
	Code2              string    `orm:"column(code2);size(50);null"`
	Name2              string    `orm:"column(name2);size(200);null"`
	FactoryName2       string    `orm:"column(factory_name2);size(100);null"`
	Trademark2         string    `orm:"column(trademark2);size(100);null"`
	Spec2              string    `orm:"column(spec2);size(200);null"`
	Addr2              string    `orm:"column(addr2);size(200);null"`
	ExpirationDate     int       `orm:"column(expiration_date);null" description:"保质期"`
	ExpirationDateUnit int16     `orm:"column(expiration_date_unit);null" description:"保质期单位（1 月 2 天 3小时）"`
	Status             int16     `orm:"column(status);null" description:"1默认 2 已修改 3已复查"`
	Issame             int16     `orm:"column(issame);null" description:"是否同一产品 1是 2否 3 没找到"`
	Batch              int       `orm:"column(batch);null" description:"数据批次"`
	IsPack             int16     `orm:"column(is_pack)" description:"是否为包装 0默认 1否 2是"`
	IsHigh             int16     `orm:"column(is_high);null"`
}

func (t *SysBarcodeOld) TableName() string {
	return "sys_barcode_old"
}

func init() {
	orm.RegisterModel(new(SysBarcodeOld))
}

// AddSysBarcodeOld insert a new SysBarcodeOld into database and returns
// last inserted Id on success.
func AddSysBarcodeOld(m *SysBarcodeOld) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysBarcodeOldById retrieves SysBarcodeOld by Id. Returns error if
// Id doesn't exist
func GetSysBarcodeOldById(id int) (v *SysBarcodeOld, err error) {
	o := orm.NewOrm()
	v = &SysBarcodeOld{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysBarcodeOld retrieves all SysBarcodeOld matches certain condition. Returns empty list if
// no records exist
func GetAllSysBarcodeOld(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysBarcodeOld))
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

	var l []SysBarcodeOld
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

// UpdateSysBarcodeOld updates SysBarcodeOld by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysBarcodeOldById(m *SysBarcodeOld) (err error) {
	o := orm.NewOrm()
	v := SysBarcodeOld{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysBarcodeOld deletes SysBarcodeOld by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysBarcodeOld(id int) (err error) {
	o := orm.NewOrm()
	v := SysBarcodeOld{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysBarcodeOld{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
