package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sunjun/piServer/models"
	// "fmt"
)

type GroupDevice struct {
	Text	string			`json:"text"`
	Nodes	[]*GroupDevice  `json:"nodes"`
}

type response struct {
	Lists  []*models.Device `json:"lists"`
}


type DeviceController struct {
	beego.Controller
}

// @router /status [get]
func (c *DeviceController) Status() {
	deviceId := c.GetString("device_id")

	c.Data["json"] = models.GetLastFive(deviceId)
	c.ServeJSON()
}

// @router /group [get]
func (c *DeviceController) Group() {
	groups, _, err := models.GetAllGroups()
	devices, _, err := models.GetAllDevices()


	groupDevices := make([]*GroupDevice, 0, 10)
	if (err == nil) {
		for _, v := range groups {
			groupDevice := &GroupDevice{Text:v.GroupName}
			nodes := make([]*GroupDevice, 0, 10)
			j := 0
			for _, d := range devices {
				if (d.DeviceGroup == v.Id) {
					node := &GroupDevice{Text:d.DeviceId}
					nodes = append(nodes, node)
					j++
				}
			}

			if (j > 0) {
				groupDevice.Nodes = nodes
			}
			groupDevices = append(groupDevices,groupDevice)
		}
	}

	c.Data["json"] = groupDevices
	c.ServeJSON()
}
