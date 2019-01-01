var fishApp = {
    init: function () {
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
    },
    indexPage: function () {
        $.post(location.href, function (result) {
            if (checkLogin(result.state)) {
                if (result.state === 0) {
                    //toastr.success(result.msg)
                    $("#rate").html(result.data.rate);
                    $("#gold").html(result.data.gold);
                    $("#today_reg").html(parseInt(result.data.today_reg_player) + parseInt(result.data.today_reg_player_child));
                    $("#today_reg1").html(result.data.today_reg_player);
                    $("#today_reg2").html(result.data.today_reg_player_child);
                    $("#yesterday_reg").html(parseInt(result.data.yesterday_reg_player) + parseInt(result.data.yesterday_reg_player_child));
                    $("#yesterday_reg1").html(result.data.yesterday_reg_player);
                    $("#yesterday_reg2").html(result.data.yesterday_reg_player_child);
                    $("#total_reg").html(parseInt(result.data.total_reg_player) + parseInt(result.data.total_reg_player_child));
                    $("#total_reg1").html(result.data.total_reg_player);
                    $("#total_reg2").html(result.data.total_reg_player_child);
                    $("#today_tax").html(result.data.today_tax);
                    $("#yesterday_tax").html(result.data.yesterday_tax);
                    $("#total_tax").html(result.data.total_tax);
                    $("#children").html(result.data.children);
                    $("#children_children").html(result.data.children_children);
                } else {
                    toastr.error(result.msg)
                }
            }
        })
    },
    dataPage: function () {
        pagination = true;
        showFooter = false;
        showColumns = false;
        $('#dataTable').bootstrapTable({
            method: 'POST',
            contentType: "application/x-www-form-urlencoded",
            url: dataurl,
            showColumns: showColumns,
            striped: true,
            dataField: "data",
            pageNumber: 1,
            pagination: pagination,//是否分页
            queryParamsType: 'limit',
            queryParams: queryParams,
            sidePagination: 'server',
            pageSize: 15,
            pageList: [15, 30, 50, 100],
            clickToSelect: true,
            toolbar: '#toolbar',
            locale: $.cookie('lang'),
            columns: datacolumns,
            showFooter: showFooter,
        });

        function queryParams(params) {
            searchParams = JSON.stringify($("#search-form").toJSON());
            return {
                pageSize: params.limit,
                pageIndex: params.pageNumber,
                searchParams: searchParams
            }
        }

        $('#search-form').submit(function () {
            $('#dataTable').bootstrapTable('refresh', {url: dataurl});
            return false;
        });
    },
    changepwd: function () {
        $.post(location.href, $("form").serializeArray(), function (result) {
            if (checkLogin(result.state)) {
                if (result.state === 0) {
                    toastr.success("修改成功")
                } else {
                    toastr.error("修改失败，" + result.msg)
                }
            }
        });
        return false
    },
    cashAction: function (id, state) {
        confirmMsg = "";
        switch (state) {
            case 2:
                confirmMsg = "确认拒绝本次申请？";
                break;
            case 3:
                confirmMsg = "确认同意本次申请？";
                break;
        }
        if (confirm(confirmMsg)) {
            this.putAction(location.href, {"action": "action", "id": id, "state": state});
        }
    },
    rechargeAction: function (params) {
        $.post(location.href, params, function (result) {
            if (checkLogin(result.state)) {
                if (result.state === 0) {
                    toastr.success("充值成功");
                } else {
                    toastr.error("充值失败，" + result.msg)
                }
            }
        });
    },
    putAction: function (url, params) {
        $.ajax({
            url: url,
            method: "put",
            data: params,
            success: function (result) {
                if (checkLogin(result.state)) {
                    if (result.state === 0) {
                        toastr.success("提交成功");
                        $('#data-modal').modal('hide');
                        $('#dataTable').bootstrapTable('refresh', {url: dataurl});
                    } else {
                        toastr.error("提交失败，" + result.msg)
                    }
                }
            }
        });
    },
    formatter: {
        cashType: function (t) {
            switch (t) {
                case 0:
                    return "支付宝";
                case 1:
                    return "银行卡";
                case 2:
                    return "代理商";
            }
        },
        cashState: function (value) {
            switch (value) {
                case 0:
                    return "<b class='text-info'>申请中</b>";
                case 1:
                    return "<b class='text-primary'>等待付款</b>";
                case 2:
                    return "<b class='text-warning'>拒绝兑换</b>";
                case 3:
                    return "<b class='text-success'>成功兑换</b>";
                case 4:
                    return "<b class='text-danger'>退款</b>";
            }
        },
        cashInfo: function (value, row) {
            switch (row.tx_type) {
                case 0:
                    return "支付宝账号：" + row.alipay + "<br>支付宝姓名：" + row.alipay_name;
                case 1:
                    return "银行：" + fishApp.formatter.bankType(row.bank_card_type_id) + "<br/>卡号：" + row.bank_card_no + "<br>真实姓名：" + row.real_name;
            }
        },
        bankType: function (b) {
            switch (b) {
                case 1:                    return "工商银行";
                case   2:                    return "农业银行";
                case   3    :                    return "中国银行";
                case   4    :                    return "建设银行";
                case   5    :                    return "招商银行";
                case  6    :                    return "交通银行";
                case   7    :                    return "光大银行";
                case 8    :                    return "华夏银行";
                case 9    :                    return "广发银行";
                case 10    :                    return "北京银行";
                case 11    :                    return "北京农商行";
                case 12    :                    return "上海银行";
                case 13    :                    return "上海农商银行";
                case  14    :                    return "渤海银行";
                case 15    :                    return "杭州银行";
                case 16    :                    return "广州市商业银行";
                case 17    :                    return "中信银行";
                case 18    :                    return "中国邮储银行";
                case 19    :                    return "兴业银行";
                case 20    :                    return "民生银行";
                case  21    :                    return "平安银行";
                case  22    :                    return "浦发银行";
                case    23    :                    return "杭州联合银行";
                case   24    :                    return "宁波银行";
                case   25    :                    return "南京银行";
                case   26    :                    return "温州市商业银行";
                case    27    :                    return "长沙银行";
                case    28    :                    return "集友银行";
                case   29    :                    return "浙商银行";
                case   30    :                    return "浙江稠州商业银行";
                case   31    :                    return "广州市农信社";
                case   32    :                    return "汉口银行";
            }
        },
        gold: function (value) {
            return (value / 100).toFixed(2);
        },
        rate: function (value) {
            return value + "%"
        }
    },
    getGameName: function (id) {
        switch (id) {
            case 1:
                return "捕鱼";
            case 2:
                return "扎金花";
            case 3:
                return "斗地主";
            case 4:
                return "百人牛牛";
            case 5:
                return "红黑大战";
            case 6:
                return "抢庄牛牛";
            case 7:
                return "龙虎斗";
        }
    },
    bankInfo: function () {
        $.post(location.href, function (result) {
            if (checkLogin(result.state)) {
                if (result.data.length === 0) {
                    //添加
                } else {
                    bank_info
                }
            }
        })
    }
};

function checkLogin(state) {
    if (state === 20000) {
        location.href = "./login";
        return false
    }
    return true
}

$(document).ready(function () {
    fishApp.init()
});

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