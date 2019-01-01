<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> 概况</a></li>
            <li>提现申请</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-lg-6">
        <div class="panel panel-default">
            <div class="panel-heading">
                提现申请
            </div>
            <div class="panel-body">
                <form role="form" id="form_apply" method="post">
                    <div class="form-group input-group">
                        <span class="input-group-addon">提取金额</span>
                        <input class="form-control" id="amount" name="amount" type="number"
                               placeholder="请输提取金额" required/>
                    </div>
                    <div class="form-group">
                        <span class="alert-dismissable">可提取金额</span><span id="can_cash" class="text-info">88.52</span>
                    </div>
                    <div class="form-group input-group">
                        <span class="input-group-addon">提取方式</span>
                        <select id="bank_info_id" name="bank_info_id" title="" class="form-control">
                        </select>
                        <a id="show_add" class="input-group-addon">添加提取方式</a>
                    </div>
                    <input type="hidden" id="action" name="action" value="apply">
                </form>
                <div id="add_form" style="margin: 20px;padding:20px;background-color: #dca7a7">
                    <form id="form_add_bank" class="form-group">
                        <div class="form-group input-group">
                            <span class="input-group-addon">提取方式</span>
                            <select id="cash_type" name="cash_type" title="" class="form-control">
                                <option value="1">银行卡</option>
                                <option value="0">支付宝</option>
                            </select>
                        </div>
                        <div class="form-group input-group">
                            <span class="input-group-addon">选择银行</span>
                            <select id="bank_code" name="bank_code" title="" class="form-control">
                                <option value="1">工商银行</option>
                                <option value="2">农业银行</option>
                                <option value="3">中国银行</option>
                                <option value="4">建设银行</option>
                                <option value="5">招商银行</option>
                                <option value="6">交通银行</option>
                                <option value="7">光大银行</option>
                                <option value="8">华夏银行</option>
                                <option value="9">广发银行</option>
                                <option value="10">北京银行</option>
                                <option value="11">北京农商行</option>
                                <option value="12">上海银行</option>
                                <option value="13">上海农商银行</option>
                                <option value="14">渤海银行</option>
                                <option value="15">杭州银行</option>
                                <option value="16">广州市商业银行</option>
                                <option value="17">中信银行</option>
                                <option value="18">中国邮储银行</option>
                                <option value="19">兴业银行</option>
                                <option value="20">民生银行</option>
                                <option value="21">平安银行</option>
                                <option value="22">浦发银行</option>
                                <option value="23">杭州联合银行</option>
                                <option value="24">宁波银行</option>
                                <option value="25">南京银行</option>
                                <option value="26">温州市商业银行</option>
                                <option value="27">长沙银行</option>
                                <option value="28">集友银行</option>
                                <option value="29">浙商银行</option>
                                <option value="30">浙江稠州商业银行</option>
                                <option value="31">广州市农信社</option>
                                <option value="32">汉口银行</option>
                            </select>
                        </div>
                        <div class="form-group input-group">
                            <span class="input-group-addon">卡号</span>
                            <input class="form-control" id="bank_card" name="bank_card" type="text"
                                   placeholder="请输入银行卡号" required/>
                        </div>
                        <div class="form-group input-group">
                            <span class="input-group-addon">开户行</span>
                            <input class="form-control" id="bank_name" name="bank_name" type="text"
                                   placeholder="请输入开户行" required/>
                        </div>
                        <div class="form-group input-group">
                            <span class="input-group-addon">姓名</span>
                            <input class="form-control" id="real_name" name="real_name" type="text"
                                   placeholder="请输入姓名" required/>
                        </div>
                        <div class="form-group input-group" style="width: 100%">
                            <a id="btn_close" class="btn btn-default" style="float: left">关闭</a>
                            <a id="btn_save" class="btn btn-info" style="float: right">保存</a>
                        </div>
                        <input type="hidden" id="action" name="action" value="save">
                    </form>
                </div>
                <div class="form-group" style="text-align: center">
                    <button id="btn_apply" class="form-control btn btn-info" type="button">确认提取</button>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    var bankInfos = [];
    var canCash = fishApp.formatter.gold({{.user.Gold}});
    $("#cash_type").change(function () {
        form = this.parentNode.parentNode;
        if (this.value === "0") {
            form.children[2].children[0].innerText = "支付宝";
            $(form.children[2].children[1]).attr("placeholder", "请输入支付宝账号");
            $(form.children[1]).hide();
            $(form.children[3]).hide();
        } else {
            form.children[2].children[0].innerText = "卡号";
            $(form.children[2].children[1]).attr("placeholder", "请输入银行卡号");
            $(form.children[1]).show();
            $(form.children[3]).show();
        }
    });
    $("#btn_close").click(function () {
        $("#add_form").hide();
    });
    $("#show_add").click(function () {
        $("#add_form").show();
    });
    $("#btn_save").click(function () {
        $.ajax({
            type: "put",
            url: location.href,
            data: $("#form_add_bank").serializeArray() ,
            success: function (result) {
                if (result.state === 0) {
                    toastr.success("保存成功");
                    $("#form_add_bank")[0].reset();
                    loadData();
                } else {
                    toastr.error("保存失败");
                }
            }
        });
        $("#add_form").hide();
    });
    $("#btn_apply").click(function () {
        $.ajax({
            type: "put",
            url: location.href,
            data: $("#form_apply").serializeArray() ,
            success: function (result) {
                if (result.state === 0) {
                    toastr.success("申请成功");
                    location.href = "../"
                } else {
                    toastr.error(result.msg);
                }
            }
        });
    });
    function loadBankInfos() {
        $.ajax({
            type: "post",
            url: location.href,
            async: false,
            success: function (result) {
                if (result.state === 0) {
                    bankInfos = result.data;
                }
            }
        });
    }
    function loadCanCash() {
        $.ajax({
            type: "post",
            url: location.href,
            async: false,
            success: function (result) {
                if (result.state === 0) {
                    bankInfos = result.data;
                }
            }
        });
    }
    function loadData() {
        loadBankInfos();
        $("#add_form").hide();
        bank_info_id = $("#bank_info_id");
        $("#can_cash").html(canCash);
        bank_info_id.html('');
        if (bankInfos.length === 0) {
            bank_info_id.append("<option value=\"0\">请先绑定银行卡或支付宝</option>")
        } else {
            for (i in bankInfos) {
                if (bankInfos[i].cash_type === 0) {
                    bank_info_id.append("<option value=\"" + bankInfos[i].id + "\">支付宝：" + bankInfos[i].bank_card_no + "</option>")
                } else {
                    bank_info_id.append("<option value=\"" + bankInfos[i].id + "\">银行卡：" + bankInfos[i].bank_card_no + "</option>")
                }
            }
        }
    }
    loadData()
</script>