package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CloudpayStatementDetail struct {
	Id            int       `orm:"column(id);auto" description:"主键id"`
	Chcd          string    `orm:"column(chcd);size(0);null" description:"交易渠道"`
	TransTime     time.Time `orm:"column(trans_time);type(timestamp);auto_now" description:"交易时间"`
	CreateUid     int       `orm:"column(create_uid);null" description:"Created by"`
	Clearing      float32   `orm:"column(clearing);null" description:"商户清算金额"`
	Terminalid    string    `orm:"column(terminalid);size(0);null"`
	WriteUid      int       `orm:"column(write_uid);null" description:"Last Updated by"`
	Amount        float32   `orm:"column(amount);null" description:"商户交易金额"`
	WriteDate     time.Time `orm:"column(write_date);type(timestamp)" description:"Last Updated on"`
	Poundage      float32   `orm:"column(poundage);null" description:"商户手续费"`
	OrderNum      string    `orm:"column(order_num);size(0);null" description:"订单号"`
	TransOrderNum string    `orm:"column(trans_order_num);size(0);null"`
	CreateDate    time.Time `orm:"column(create_date);type(timestamp)" description:"Created on"`
	Type          string    `orm:"column(type);size(0);null" description:"交易类型"`
	Mchntid       string    `orm:"column(mchntid);size(0);null" description:"服务商户号"`
}

func (t *CloudpayStatementDetail) TableName() string {
	return "cloudpay_statement_detail"
}

func init() {
	orm.RegisterModel(new(CloudpayStatementDetail))
}

// AddCloudpayStatementDetail insert a new CloudpayStatementDetail into database and returns
// last inserted Id on success.
func AddCloudpayStatementDetail(m *CloudpayStatementDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCloudpayStatementDetailById retrieves CloudpayStatementDetail by Id. Returns error if
// Id doesn't exist
func GetCloudpayStatementDetailById(id int) (v *CloudpayStatementDetail, err error) {
	o := orm.NewOrm()
	v = &CloudpayStatementDetail{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCloudpayStatementDetail retrieves all CloudpayStatementDetail matches certain condition. Returns empty list if
// no records exist
func GetAllCloudpayStatementDetail(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CloudpayStatementDetail))
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

	var l []CloudpayStatementDetail
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

// UpdateCloudpayStatementDetail updates CloudpayStatementDetail by Id and returns error if
// the record to be updated doesn't exist
func UpdateCloudpayStatementDetailById(m *CloudpayStatementDetail) (err error) {
	o := orm.NewOrm()
	v := CloudpayStatementDetail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCloudpayStatementDetail deletes CloudpayStatementDetail by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCloudpayStatementDetail(id int) (err error) {
	o := orm.NewOrm()
	v := CloudpayStatementDetail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CloudpayStatementDetail{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
