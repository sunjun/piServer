var conn = null;
$(document).ready(main);

function main()
{
    startWebsocket();
    if (conn) {
        waitForSocketConnection(conn, function() {
            getDevices();
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

function getDevices()
{
    var getAll = {"CommandCode":2, "DeviceID":"123456", "CommandMessage":"all"};
    conn.send(JSON.stringify(getAll));
}
function startWebsocket()
{
    conn = new WebSocket("ws://localhost:23456/webControlServer");
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
    $("#online").text(data[0].id);
 
}  
function onError(e) {
}  

