<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li>我的玩家</li>
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
                我的玩家
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder="请输入玩家ID">
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
        {title: '昵称', field: 'nick_name'},
        {title: '所属代理', field: 'agent_id'},
        {title: '身上余额', field: 'global_num', formatter: fishApp.formatter.gold},
        {title: '保险箱余额', field: 'bank_num', formatter: fishApp.formatter.gold},
        {title: '累计税收', field: 'tax', formatter: fishApp.formatter.gold},
        {title: '充值金额', field: 'total_recharge_sum', formatter: fishApp.formatter.gold},
        {title: '提现金额', field: 'all_withdraw_amount', formatter: fishApp.formatter.gold},
    ];
    dataurl =location.href;
    fishApp.dataPage();
</script>
