package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"strings"
	"time"
)

type Motivational_quotes struct {
	Id                int64  `orm:"auto"`
	MotivationalQuote string `orm:"size(255)"`
	Active            int
	DateCreated       time.Time `orm:"type(datetime)"`
	DateModified      time.Time `orm:"type(datetime)"`
	CreatedBy         int
	ModifiedBy        int
}

func init() {
	orm.RegisterModel(new(Motivational_quotes))
}

// AddMotivational_quotes insert a new Motivational_quotes into database and returns
// last inserted Id on success.
func AddMotivational_quotes(m *Motivational_quotes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMotivational_quotesById retrieves Motivational_quotes by Id. Returns error if
// Id doesn't exist
func GetMotivational_quotesById(id int64) (v *Motivational_quotes, err error) {
	o := orm.NewOrm()
	v = &Motivational_quotes{Id: id}
	if err = o.QueryTable(new(Motivational_quotes)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMotivational_quotes retrieves all Motivational_quotes matches certain condition. Returns empty list if
// no records exist
func GetAllMotivational_quotes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Motivational_quotes))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []Motivational_quotes
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateMotivational_quotes updates Motivational_quotes by Id and returns error if
// the record to be updated doesn't exist
func UpdateMotivational_quotesById(m *Motivational_quotes) (err error) {
	o := orm.NewOrm()
	v := Motivational_quotes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMotivational_quotes deletes Motivational_quotes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMotivational_quotes(id int64) (err error) {
	o := orm.NewOrm()
	v := Motivational_quotes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Motivational_quotes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
