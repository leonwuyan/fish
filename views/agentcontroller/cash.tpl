<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li>账户提现</li>
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
                账户提现
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder="请输入代理ID">
                        </div>
                        <div class="form-group input-group">
                            <select id="state" name="state" class="form-control">
                                <option value="-1">全部</option>
                                <option value="1">等待付款</option>
                                <option value="2">拒绝兑换</option>
                                <option value="3">成功兑换</option>
                            </select>
                        </div>
                        <button class="btn btn-info form-control" type="submit"><span
                                class="fa fa-search"></span></button>
                        <a id="btn-add" type="button" href="{{.host}}cash/apply" class="btn-danger form-control"><span
                                class="fa fa-plus"></span>发起新提现
                        </a>
                    </form>
                </div>
                <table class="table table-striped table-bordered table-hover" id="dataTable">
                </table>
            </div>
        </div>
    </div>
</div>
<div class="modal fade" id="data-modal" tabindex="-1" role="dialog"
     aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
        </div><!-- /.modal-content -->
    </div><!-- /.modal -->
</div>
<script>
    datacolumns = [
        {title: '代理ID', field: 'agent_id'},
        {title: '订单号', field: 'order_id'},
        {title: '姓名', field: 'real_name'},
        {title: '金额', field: 'gold',formatter: fishApp.formatter.gold},
        {title: '支付方式', field: 'tx_type', formatter: fishApp.formatter.cashType},
        {title: '支付信息', formatter: fishApp.formatter.cashInfo},
        {title: '时间', field: 'withdrawals_log_time'},
        {title: '状态', field: 'state', formatter: fishApp.formatter.cashState},
    ];
    dataurl = location.href;
    fishApp.dataPage();
</script>
