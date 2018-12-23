<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">首页</a></li>
            <li>后台充值</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-lg-6">
        <div class="panel panel-default">
            <div class="panel-heading">
                后台充值
            </div>
            <div class="panel panel-body">
                <form role="form" method="post">
                    <div class="form-group input-group">
                        <span class="input-group-addon">玩家ID</span>
                        <input class="form-control" id="id" name="id" type="number"
                               placeholder="请输入玩家ID" required/>
                    </div>
                    <div class="form-group input-group">
                        <span class="input-group-addon">充值金额</span>
                        <input class="form-control" id="amount" name="amount" type="number"
                               placeholder="请输入充值金额" required/>
                    </div>
                    <div class="form-group" style="text-align: center">
                        <button class="btn btn-default" type="reset">重置</button>
                        <button class="btn btn-info" type="submit">确认提交</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>
<script>
    $(function () {
        if (request("id")) {
            $("#id").val(request("id"))
        }
    });
    $("form").submit(function () {
        var params = $(this).toJSON();
        confirmMsg = '请确认，玩家账号ID：{0}，购买金额：{1}';
        if (confirm(confirmMsg.format(params.id, params.amount))) {
            fishApp.rechargeAction(params);
        }
        return false
    })
</script>
