package models

import (
	//	"errors"
	"github.com/astaxie/beego/orm"
	//	"strconv"
	//	"time"
)

type Device struct {
	Id          int
	DeviceId	string
	DeviceGroup	int
}

func init() {
	orm.RegisterModel(new(Device))
}

func GetAllDevices() ([]*Device, int, error) {
	var devices []*Device
	o := orm.NewOrm()
	num, err := o.QueryTable("device").All(&devices)

	return devices, int(num), err
}