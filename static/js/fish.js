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
            $.post(location.href, {"action": "action", "id": id, "state": state}, function (result) {
                if (checkLogin(result.state)) {
                    if (result.state === 0) {
                        toastr.success("提交成功");
                        $('#dataTable').bootstrapTable('refresh', {url: dataurl});
                    } else {
                        toastr.error("提交失败，" + result.msg)
                    }
                }
            })
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
    putAction: function (utl, params) {
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
                case 1:
                    return "支付宝";
                case 2:
                    return "银行卡";
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
        bankType: function (b) {
            switch (b) {
                case 1:
                    return "工商银行";
            }
        },
        gold: function (value) {
            return (value / 100).toFixed(2);
        }
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