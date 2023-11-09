package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Assessment_options struct {
	Id               int64 `orm:"auto"`
	AssessmentId     int
	AssessmentOption string `orm:"size(255)"`
	Active           int
	DateCreated      time.Time `orm:"type(datetime)"`
	DateModified     time.Time `orm:"type(datetime)"`
	CreatedBy        int
	ModifiedBy       int
}

func init() {
	orm.RegisterModel(new(Assessment_options))
}

// AddAssessment_options insert a new Assessment_options into database and returns
// last inserted Id on success.
func AddAssessment_options(m *Assessment_options) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAssessment_optionsById retrieves Assessment_options by Id. Returns error if
// Id doesn't exist
func GetAssessment_optionsById(id int64) (v *Assessment_options, err error) {
	o := orm.NewOrm()
	v = &Assessment_options{Id: id}
	if err = o.QueryTable(new(Assessment_options)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAssessment_options retrieves all Assessment_options matches certain condition. Returns empty list if
// no records exist
func GetAllAssessment_options(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Assessment_options))
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

	var l []Assessment_options
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

// UpdateAssessment_options updates Assessment_options by Id and returns error if
// the record to be updated doesn't exist
func UpdateAssessment_optionsById(m *Assessment_options) (err error) {
	o := orm.NewOrm()
	v := Assessment_options{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAssessment_options deletes Assessment_options by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAssessment_options(id int64) (err error) {
	o := orm.NewOrm()
	v := Assessment_options{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Assessment_options{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
