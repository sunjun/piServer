package models

import (
	//	"errors"
	"github.com/astaxie/beego/orm"
	//	"strconv"
	//	"time"
)

type Group struct {
	Id          int
	GroupName	string
}

func init() {
	orm.RegisterModel(new(Group))
}

func GetAllGroups() ([]*Group, int, error) {
	var groups []*Group
	o := orm.NewOrm()
	num, err := o.QueryTable("group").All(&groups)

	return groups, int(num), err
}