var msgUrl = "http://" + location.host + "/admin/ws";
var intervalTime = 5000;

function getMessage() {
    $.post(msgUrl, function (result) {
        if (checkLogin(result.state)) {
            for (ind in result.data) {
                msg = result.data[ind];
                switch (msg.msg_type) {
                    case 0:
                        toastr.info(msg.msg_content);
                        break;
                    case 1:
                        toastr.warning(msg.msg_content);
                        break;
                    case 2:
                        break;
                    default:
                        toastr.error("收到外星来的消息");
                        break;
                }
            }
        }
    }, 'json')
}

setInterval(getMessage, intervalTime);
