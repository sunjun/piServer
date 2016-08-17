package mywebsocket

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/websocket"
)

type device struct {
	Id       string `json:"id"`
	WsID     string `json:"wsID"`
	IsOnline bool   `json:"isOnline"`
	Ws       *websocket.Conn
}

const (
	UPADTE = iota
	UPLOAD_ID
	ALL_DEVICES
	HEART_BEAT
	TAKE_PHOTO
	STATUS
)

type Command struct {
	CommandCode    int
	DeviceID       string
	CommandMessage string
}

var deviceList []*device
var tempWs *websocket.Conn

var port *int = flag.Int("p", 23456, "Port to listen.")

func initDeviceList() {
	deviceList = make([]*device, 0, 100)
}

// sendRecvServer echoes back text messages sent from client
// using websocket.Message.
func sendRecvServer(ws *websocket.Conn) {
	fmt.Printf("sendRecvServer %#v\n", ws)
	for {
		var buf string
		// Receive receives a text message from client, since buf is string.
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf)
		// Send sends a text message to client, since buf is string.
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%q\n", buf)
	}
	fmt.Println("sendRecvServer finished")
}

type T struct {
	Msg  string
	Path string
}

// jsonServer echoes back json string sent from client using websocket.JSON.
func clientUpdateServer(ws *websocket.Conn) {
	fmt.Printf("clientUpdateServer %#v\n", ws.Config())
	for {
		var command Command
		// Receive receives a text message serialized T as JSON.
		err := websocket.JSON.Receive(ws, &command)
		if err != nil {
			fmt.Println("clientUpdateServer:" + err.Error())
			break
		}

		switch command.CommandCode {
		case UPLOAD_ID:
			newDevice := &device{Id: command.DeviceID, IsOnline: true, Ws: ws}
			deviceList = append(deviceList, newDevice)
			fmt.Println("clientUpdateServer:%v", deviceList)
		case HEART_BEAT:
			fmt.Println(command.CommandMessage)

		}

		//		fmt.Printf("recv:%#v\n", msg)
		//		// Send send a text message serialized T as JSON.
		//		err = websocket.JSON.Send(ws, msg)
		//		if err != nil {
		//			fmt.Println(err)
		//			break
		//		}
		//		fmt.Printf("send:%#v\n", msg)
	}
}

// jsonServer echoes back json string sent from client using websocket.JSON.
func clientMainServer(ws *websocket.Conn) {
	fmt.Printf("clientMainServer %#v\n", ws.Config())
	for {
		var command Command
		// Receive receives a text message serialized T as JSON.
		err := websocket.JSON.Receive(ws, &command)
		if err != nil {
			fmt.Println("clientMainServer:" + err.Error())
			break
		}

		switch command.CommandCode {
		case UPLOAD_ID:
			newDevice := &device{Id: command.DeviceID, IsOnline: true, Ws: ws}
			deviceList = append(deviceList, newDevice)
			fmt.Println("clientMainServer:%v", deviceList)
		case HEART_BEAT:
			fmt.Println(command.CommandMessage)

		case TAKE_PHOTO:
			b, err := json.Marshal(command)

			command.CommandMessage = "http://192.168.1.103:23456/static/img/image.jpg"
			err = websocket.Message.Send(tempWs, string(b))
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(command.CommandMessage)
		}

		//		fmt.Printf("recv:%#v\n", msg)
		//		// Send send a text message serialized T as JSON.
		//		err = websocket.JSON.Send(ws, msg)
		//		if err != nil {
		//			fmt.Println(err)
		//			break
		//		}
		//		fmt.Printf("send:%#v\n", msg)
	}
}

// jsonServer echoes back json string sent from client using websocket.JSON.
func webControlServer(ws *websocket.Conn) {
	fmt.Printf("webControlServer %#v\n", ws.Config())
	for {
		var command Command
		// Receive receives a text message serialized T as JSON.
		err := websocket.JSON.Receive(ws, &command)
		fmt.Println("webControlServer got message")
		fmt.Printf("webControlServer %#v\n", command)
		if err != nil {
			fmt.Println("webControlServer:" + err.Error())
			break
		}

		switch command.CommandCode {
		case ALL_DEVICES:
			// Send send a text message serialized T as JSON.
			b, err := json.Marshal(deviceList)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("stringb:" + string(b))
			err = websocket.Message.Send(ws, string(b))
			if err != nil {
				fmt.Println(err)
				break
			}

		case STATUS:
			var sendCommand Command
			var selectDevice *device
			selectDevice = getDeviceFromDeviceList(command.DeviceID)
			fmt.Printf("webControlServer %#v\n", deviceList)
			sendCommand.CommandCode = STATUS
			sendCommand.DeviceID = command.DeviceID
			sendCommand.CommandMessage = "offline"
			if (selectDevice != nil) && (selectDevice.IsOnline == true) {
				sendCommand.CommandMessage = "online"
			}
			b, err := json.Marshal(sendCommand)
			err = websocket.Message.Send(ws, string(b))
			if err != nil {
				fmt.Println(err)
				return
			}
		case TAKE_PHOTO:
			//send command to client to take photo and upload
			var selectDevice *device
			selectDevice = getDeviceFromDeviceList(command.DeviceID)
			var clientTakePhotoCommand Command
			clientTakePhotoCommand.CommandCode = TAKE_PHOTO
			clientTakePhotoCommand.DeviceID = selectDevice.Id
			clientTakePhotoCommand.CommandMessage = "take photo"
			b, err := json.Marshal(clientTakePhotoCommand)
			err = websocket.Message.Send(selectDevice.Ws, string(b))
			tempWs = ws
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func getDeviceFromDeviceList(deviceID string) *device {
	var selectDevice *device
	for i, v := range deviceList {
		if v.Id == deviceID {
			selectDevice = deviceList[i]
			return selectDevice
		}
	}
	return selectDevice
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	fileName := path.Base(handler.Filename)

	f, err := os.OpenFile("./static/img/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func RunWebsocket() {
	flag.Parse()
	http.Handle("/sendRecvText", websocket.Handler(sendRecvServer))
	http.Handle("/clientUpdateServer", websocket.Handler(clientUpdateServer))
	http.Handle("/clientMainServer", websocket.Handler(clientMainServer))
	http.Handle("/webControlServer", websocket.Handler(webControlServer))
	http.HandleFunc("/upload", upload)
	fmt.Printf("http://localhost:%d/\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic("ListenANdServe: " + err.Error())
	}
}
