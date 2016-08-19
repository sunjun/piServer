var currentNode;
var conn = null;
var CommandCode = {
    UPADTE:0,
    UPLOAD_ID:1,
    ALL_DEVICES:2,
    HEART_BEAT:3,
    TAKE_PHOTO:4,
    STATUS:5
};
function mainFunction()
{
    startClock();
    addBaiduMap();
    updateTreeView();
    addDeviceTabListener();
    settingDialogAdd();
    takePhotoDialogAdd();
    startWebsocket();
    if (conn) {
        waitForSocketConnection(conn, function() {
         //   getDevices();
        }); 
    }

}
function waitForSocketConnection(socket, callback){
    setTimeout(
            function(){
                if (socket.readyState === 1) {
                    if(callback !== undefined){
                        callback();
                    }
                    return;
                } else {
                    waitForSocketConnection(socket,callback);
                }
            }, 5);
};

function getDevicesStatus(node)
{
    var getAll = {"CommandCode":CommandCode.STATUS, "DeviceID":node.text, "CommandMessage":"status"};
    conn.send(JSON.stringify(getAll));
}
function startWebsocket()
{
    conn = new WebSocket("ws://mindfulstart.net:23456/webControlServer");
    conn.onopen = function(e) {
        onOpen(e);
    }
    conn.onmessage = function(e) {
        onMessage(e);
   }
    conn.onclose = function(e) {
        onClose(e);
    }
}


function onOpen(e) { 
    console.log("Connection established!");
}  

function onClose(e) { 
    console.log(e.data);
    conn.close();
}

function onMessage(e) { 
    console.log(e.data);
    var data = JSON.parse(e.data);
    switch (data.CommandCode) {
        case CommandCode.STATUS:
        if (data.CommandMessage == "online") 
            $("#online").text("在线");
        else 
            $("#online").text("离线");
        break;
        case CommandCode.TAKE_PHOTO:
        $("#mdl-spinner").remove();
        var dialog = document.querySelector('#take-photo-dialog');
        var image ='<img width=100% height=100% src="'+data.CommandMessage+'" />';

        dialog.querySelector('.mdl-dialog__content').innerHTML = image;
        updateDeviceStatus(currentNode);
    }
    
}  
function onError(e) {
}  

function addBaiduMap()
{
    var map = new BMap.Map("baidu-device-map");

    var geolocation = new BMap.Geolocation();
    geolocation.getCurrentPosition(function(r){
        if(this.getStatus() == BMAP_STATUS_SUCCESS){
            var mk = new BMap.Marker(r.point);
            map.addOverlay(mk);
            map.centerAndZoom(r.point,12);
        } else {
            alert('failed'+this.getStatus());
        }        
    },{enableHighAccuracy: true})
}

function startClock() 
{
    var today=new Date();
    var  y=today.getFullYear();
    var  mon=today.getMonth()+1;
    mon=changnum(mon);
    var  d=today.getDay();
    d=changnum(d);
    var  h=today.getHours();
    var  m=today.getMinutes();
    m=changnum(m);
    var  s=today.getSeconds();
    s=changnum(s);
    document.getElementById("timeid").innerHTML=y+"年"+mon+"月"+d+"日"+h+":"+m+":"+s;
    t=setTimeout(function(){
        startClock();
    },500);
}

function changnum(i){      //月、日、秒如果小于10数字前加0
    if(i<10){
        return "0"+i;
    }
    return i;
}

function addDeviceTabListener() 
{
    $("li#device-map").click(function(){
        $(this).attr("class", "active");
        $("li#device-status").attr("class", "");
        $("li#device-info").attr("class", "");

        $("div#baidu-device-map").show();
        $("div#device-status").hide()
        $("div#device-info").hide();
    });

    $("li#device-status").click(function(){
        $(this).attr("class", "active");
        $("li#device-map").attr("class", "");
        $("li#device-info").attr("class", "");

        $("div#baidu-device-map").hide();
        $("div#device-status").show()
        $("div#device-info").hide();
        updateDeviceStatus(currentNode);
    });

    $("li#device-info").click(function(){
        $(this).attr("class", "active");
        $("li#device-map").attr("class", "");
        $("li#device-status").attr("class", "");

        $("div#baidu-device-map").hide();
        $("div#device-status").hide()
        $("div#device-info").show();
    });
}

function getCurActiveDeviceTab()
{
    var result = 0;
    if ($("li#device-map").attr("class") === "active")
        result = 0;
    if ($("li#device-status").attr("class") === "active")
        result = 1;
    if ($("li#device-info").attr("class") === "active")
        result = 2;

    return result;
}

function updateMap(node)
{
    $.get("device_info",{device_id: node.text}, function(data,status){
        var point = new BMap.Point(data[0].longitude, data[0].latitude);
        var map = new BMap.Map("baidu-device-map");    // 创建Map实例

        map.centerAndZoom(point, 15);  // 初始化地图,设置中心点坐标和地图级别

        map.addControl(new BMap.MapTypeControl());   //添加地图类型控件
        map.setCurrentCity("济南");          // 设置地图显示的城市 此项是必须设置的
        map.enableScrollWheelZoom(true);     //开启鼠标滚轮缩放
        var mk = new BMap.Marker(point);
        map.addOverlay(mk);
    });
}

function updateDeviceStatus(node)
{
    $.get("v1/device/status",{device_id: node.text}, function(data,status){
        //alert(data[0].device_photo_size);
        if (data.length < 5) {
            $("#picture-0").attr("src", "");
            $("#picture-1").attr("src", "");
            $("#picture-2").attr("src", "");
            $("#picture-3").attr("src", "");
            $("#picture-4").attr("src", "");

            $("#picture-date0").html("");
            $("#picture-date1").html("");
            $("#picture-date2").html("");
            $("#picture-date3").html("");
            $("#picture-date4").html("");

        } else {
            $("#picture-0").attr("src", data[0].path);
            $("#picture-1").attr("src", data[1].path);
            $("#picture-2").attr("src", data[2].path);
            $("#picture-3").attr("src", data[3].path);
            $("#picture-4").attr("src", data[4].path);

            $("#picture-date0").html(data[0].photo_date);
            $("#picture-date1").html(data[1].photo_date);
            $("#picture-date2").html(data[2].photo_date);
            $("#picture-date3").html(data[3].photo_date);
            $("#picture-date4").html(data[4].photo_date);    
        }
    });
    getDevicesStatus(node);
}

function updateTreeView()
{
    $.get("v1/device/group",function(data,status){
        $('#tree').treeview({
            data: data,
        color: "#428bca",

        onNodeSelected: function(event, node) {
            currentNode = node;
            if (getCurActiveDeviceTab() === 0) {
                updateMap(node);
            } else if (getCurActiveDeviceTab() === 1) {
                updateDeviceStatus(node);
            }
        }
        });
    });
}

function settingDialogAdd()
{
    var dialog = document.querySelector('#setting-dialog');
    // var dialog = $("#setting-dialog");
    var showDialogButton = document.querySelector('#device-setting-dialog');
    if (! dialog.showModal) {
        dialogPolyfill.registerDialog(dialog);
    }
    showDialogButton.addEventListener('click', function() {
        dialog.showModal();
        var device_id = currentNode.text;

        $.get("device_info", {device_id:device_id},function(data,status){
            $("#device").text(data[0].device_id);
            $("#device-boot").val(data[0].device_boot_time);
            $("#device-off").val(data[0].device_off_time);
            $("#device-start-work").val(data[0].device_photo_starttime)
        });
    });

    dialog.querySelector('#apply-this').addEventListener('click', function() {
        
        var device_id = currentNode.text;
        var device_photo_starttime = $('input#device-start-work').val();
        $.post("device_info_save", {device_id:device_id ,device_photo_starttime:device_photo_starttime},function(data,status){
            
        });
        alert("设置成功");
        var a = $('input#sample11');

        //$("#mdl-spinner").remove();
    });

    dialog.querySelector('#apply-all-line').addEventListener('click', function() {
        alert("设置成功");
        var a = $('input#sample11');

        //$("#mdl-spinner").remove();
    });

    dialog.querySelector('#apply-reboot').addEventListener('click', function() {
        alert("设置成功");
        var a = $('input#sample11');
    });

    dialog.querySelector('#apply-close').addEventListener('click', function() {
        // $.get("device_info", {device_id:device_id},function(data,status){

        // });

        dialog.close();
        //$("#mdl-spinner").remove();
    });
}
function takePhotoDialogAdd()
{
    var dialog = document.querySelector('#take-photo-dialog');
    var showDialogButton = document.querySelector('#device-take-photo-dialog');
    if (! dialog.showModal) {
        dialogPolyfill.registerDialog(dialog);
    }
    showDialogButton.addEventListener('click', function() {
        dialog.showModal();
        var device_id = currentNode.text;

        if ($("#online").text() == "在线") {
            var spinner = '<div class="mdl-spinner mdl-js-spinner is-active" id="mdl-spinner"></div>';
            dialog.querySelector('.mdl-dialog__content').innerHTML = spinner;

            var take = {"CommandCode":CommandCode.TAKE_PHOTO, 
            "DeviceID":device_id, 
            "CommandMessage":"take photo"};
            conn.send(JSON.stringify(take));
        } else {
            dialog.querySelector('.mdl-dialog__content').innerHTML = "设备离线不能拍照";

        }

        
    });

    dialog.querySelector('#apply-close').addEventListener('click', function() {
        dialog.close();
    });

    

}

$(document).ready(mainFunction);
