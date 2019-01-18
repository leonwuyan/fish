<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li><a href="#">日志</a></li>
            <li>充值日志</li>
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
                充值日志
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder="请输入玩家ID">
                        </div>
                        <div class="form-group input-group">
                            <input id="agent_id" name="agent_id" class="form-control" type="text"
                                   placeholder="请输入代理ID">
                        </div>
                        <div class="form-group input-group">
                            <select id="recharge_type" name="recharge_type" class="form-control">
                                <option value="0">全部</option>
                                <option value="100">玩家充值</option>
                                <option value="101">活动赠送</option>
                                <option value="102">测试赠送</option>
                            </select>
                        </div>
                        <div class="form-group input-group">
                            <span class="input-group-addon">开始</span>
                            <input id="begin" name="begin" class="form-control" type="date"
                                   value="{{today}}" placeholder="" required>
                        </div>
                        <div class="form-group input-group">
                            <span class="input-group-addon">结束</span>
                            <input id="end" name="end" class="form-control" type="date"
                                   value="{{today}}" placeholder="" required>
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
        {title: '代理ID', field: 'agent_id'},
        {title: '单号', field: 'transaction_id'},
        {title: '方式', field: 'recharge_type', formatter: fishApp.formatter.rechargeType},
        {title: '金额', field: 'gold_change', formatter: fishApp.formatter.gold},
        {title: '充值时间', field: 'recharge_time'},
        {title: '是否到账', field: 'is_send'},
        {title: '到账时间', field: 'send_time'},
    ];
    dataurl = location.href;
    fishApp.dataPage();
</script>
