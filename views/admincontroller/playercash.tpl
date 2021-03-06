<script src="/static/js/fish.js"></script>
<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">首页</a></li>
            <li><a href="#"> 玩家管理</a></li>
            <li> 玩家提现</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <button class="navbar-toggle blue" data-toggle="collapse" data-target=".data-collapse"
                    style="border: 0; color: #fff; font-size: 8px">
                <span class="fa fa-long-arrow-down"></span>
                <span class="fa fa-long-arrow-up"></span>
            </button>
            <div class="panel-heading">
                玩家提现
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder="请输入玩家ID">
                        </div>
                        <div class="form-group input-group">
                            <select id="state" name="state" class="form-control">
                                <option value="-1">全部</option>
                                <option value="0">申请中</option>
                                <option value="1">等待付款</option>
                                <option value="2">拒绝兑换</option>
                                <option value="3">成功兑换</option>
                                <option value="4">退款</option>
                            </select>
                        </div>
                        <button class="btn btn-info form-control" type="submit"><span
                                class="fa fa-search"></span></button>
                    </form>
                </div>
                <table class="table table-striped table-bordered table-hover" id="dataTable">
                </table>
            </div>
        </div>
    </div>
</div>
<script>
    datacolumns = [
        {title: '玩家ID', field: 'user_id'},
        {title: '昵称', field: 'username'},
        {title: '提取金额', field: 'gold', formatter: fishApp.formatter.gold},
        {title: '获得金额', field: 'get_money', formatter: fishApp.formatter.gold},
        {title: '手续费', field: 'fee', formatter: fishApp.formatter.gold},
        {title: '提现类型', field: 'tx_type', formatter: fishApp.formatter.cashType},
        {title: '提现信息', formatter: fishApp.formatter.cashInfo},
        {title: '提现时间', field: 'withdrawals_log_time'},
        {title: '状态', field: 'state', formatter: fishApp.formatter.cashState},
        {
            title: '操作', formatter: function (value, row) {
                if (row.state < 2) {
                    return '<a class="btn btn-warning" onclick="fishApp.cashAction(' + row.id + ',4)">拒绝</a>' +
                            '<a class="btn btn-success" onclick="fishApp.cashAction(' + row.id + ',3)">同意</a>'
                }
            }
        }
    ];
    dataurl = location.href;
    showFooter = false;
    fishApp.dataPage();
</script>
