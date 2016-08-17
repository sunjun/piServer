package models

import (
	//	"errors"
	"github.com/astaxie/beego/orm"
	//	"strconv"
	//	"time"
	"fmt"
)

type Picture struct {
	Id          int
	DeviceId	string
	PhotoDate	string `json:"photo_date"`
	Path	string `json:"path"`
}

func GetLastFive(deviceID string) []*Picture {
	var pictureList []*Picture
	o := orm.NewOrm()
	o.QueryTable("picture").Filter("DeviceId", deviceID).OrderBy("-Id").Limit(5).All(&pictureList)
	for _, v := range pictureList {
		v.Path = fmt.Sprintf("/static/img/%s", v.Path)
	}
	return pictureList
}

func init() {
	orm.RegisterModel(new(Picture))
}
