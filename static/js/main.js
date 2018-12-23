function language(ln) {
    $.cookie('lang', ln);
    location.reload()
}

function request(name) {
    reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
}

String.prototype.format = function (args) {
    if (arguments.length > 0) {
        result = this;
        if (arguments.length === 1 && typeof (args) === "object") {
            for (key in args) {
                reg = new RegExp("({" + key + "})", "g");
                result = result.replace(reg, args[key]);
            }
        }
        else {
            for (i = 0; i < arguments.length; i++) {
                if (arguments[i] === undefined) {
                    return "";
                }
                else {
                    reg = new RegExp("({[" + i + "]})", "g");
                    result = result.replace(reg, arguments[i]);
                }
            }
        }
        return result;
    }
    else {
        return this;
    }
};
Array.prototype.exists = function (val) {
    for (i = 0; i < this.length; i++) {
        if (this[i] === val) {
            return true;
        }
    }
    return false;
};

$.fn.toJSON = function () {
    o = {};
    a = this.serializeArray();
    $.each(a, function () {
        if (o[this.name]) {
            if (!o[this.name].push) {
                o[this.name] = [o[this.name]];
            }
            o[this.name].push(this.value || '');
        } else {
            o[this.name] = this.value || '';
        }
    });
    return o;
};

mainApp = {
    initFunction: function () {
        $('#main-menu').metisMenu();
        $(window).bind("load resize", function () {
            if ($(this).width() < 768) {
                $('div.sidebar-collapse').addClass('collapse')
            } else {
                $('div.sidebar-collapse').removeClass('collapse')
            }
            if ($(this).width() < 768) {
                $('div.topbar-collapse').addClass('collapse')
            } else {
                $('div.topbar-collapse').removeClass('collapse')
            }
            if ($(this).width() < 768) {
                $('div.data-collapse').addClass('collapse')
            } else {
                $('div.data-collapse').removeClass('collapse')
            }
        });
        this.getMessage();
    },
    getMessage: function () {
        $("#msg-count").html("");
        $("#msg-read-all").hide();
        $("#msg-none").show();
        $.get(msgUrl, function (res) {
            if (res.status === 0) {
                if (res.total !== 0) {
                    for (i = 0; i < res.total; i++) {
                        title = res.data[i].title;
                        time = new Date(res.data[i].send_time);
                        content = res.data[i].content;
                        if (content.length > 30) {
                            content = content.substr(0, 29) + "...";
                        }
                        child = '<li><a href="#"> <strong>' + title + '</strong> <span class="pull-right text-muted"><em>' + time.toLocaleDateString() + '</em></span><div>' + content + '</div></a></li>';
                        child += '<li class="divider"></li>';
                        $("#msg-list").prepend(child)
                    }
                    $("#msg-count").html(res.total);
                    $("#msg-read-all").show();
                    $("#msg-none").hide();
                }
            }
        }, 'json');
    },
};
$(document).ready(function () {
    mainApp.initFunction();
    msg_interval = setInterval(mainApp.getMessage, 60000);
});

function ajaxSms(obj) {
    el = obj.parentNode;
    mobile = $(el).find("input[type='tel']").val();
    smsType = $(el).find("input[type='hidden']").val();
    if (!isPoneAvailable(mobile)) {
        alert("输入正确的手机号");
        return;
    }
    if (obj.innerHTML === "获取验证码") {
        $.get(window.location.origin + "/sms", {"mobile": mobile, "t": smsType}, function (ret) {
            if (ret ===0) {
                timerDisabled(obj);
            }else{
                alert(ret.error)
            }
        })
    }
}

function timerDisabled(obj) {
    //obj.setAttribute("disabled", true);
    var times = 60;
    obj.innerHTML = "(" + times + "s)再次获取";
    var timer = setInterval(function () {
        times--;
        if (times < 1) {
            obj.innerHTML = "获取验证码";
            clearInterval(timer);
        } else {
            obj.innerHTML = "(" + times + "s)再次获取";
        }
    }, 1000);
}

function isPoneAvailable(str) {
    if (str.length === 0) {
        return false
    }
    var myreg = /^[1][3,4,5,7,8][0-9]{9}$/;
    if (!myreg.test(str)) {
        return false;
    }
    return true
}