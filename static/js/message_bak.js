var wsUrl = "ws://" + location.host + "/admin/ws";
var lastMsgTime = new Date();
var timeout = 10;
var webSocket;

function webConnection() {
    webSocket = new WebSocket(wsUrl);
    webSocket.onmessage = function (e) {
        lastMsgTime = new Date();
        jsonData = JSON.parse(e.data);
        if (checkLogin(jsonData.state)) {
            switch (jsonData.data.msg_type) {
                case 0:
                    toastr.info(jsonData.data.msg_content);
                    break;
                case 1:
                    toastr.warning(jsonData.data.msg_content);
                    break;
                case 2://没新消息，不处理
                    //toastr.warning("没新消息，不处理");
                    break;
                default:
                    toastr.error("收到外星来的消息");
                    break
            }
        }
    };
}

function checkConn() {
    if (new Date() - lastMsgTime > timeout * 1000) {
        console.log("close and reconnection");
        if (webSocket.readyState === 1) {
            webSocket.close();
        }
        webConnection();
    }
}

webConnection();
setInterval(checkConn, 1000);
