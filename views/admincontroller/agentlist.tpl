<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li>我的代理</li>
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
                我的代理
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder="请输入代理ID">
                        </div>
                        <div class="form-group input-group">
                            <input id="name" name="name" class="form-control" type="text"
                                   placeholder="请输入代理名称">
                        </div>
                        <div class="form-group input-group">
                            <input id="mobile" name="mobile" class="form-control" type="text"
                                   placeholder="请输入手机号码">
                        </div>
                        <button class="btn btn-info form-control" type="submit"><span
                                class="fa fa-search"></span></button>
                        <button id="btn-add" type="button" class="btn-info form-control"
                                data-toggle="modal"
                                data-target="#data-modal"
                                data-remote="{{.domain}}forms/?a=add_agent"><span
                                class="fa fa-plus"></span>添加代理</button>
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
        {title: 'ID', field: 'id'},
        {title: '名称', field: 'name'},
        {title: '代理等级', field: 'level'},
        {title: '上级代理', field: 'parent_id'},
        {title: '电话', field: 'mobile'},
        {title: '旗下代理数', field: 'total_children_immediate'},
        {title: '旗下玩家数', field: 'total_players_immediate'},
        {title: '玩家总收入', field: 'total_tax', formatter: goldFormatter},
        {title: '提成比例', field: 'rate'},
        {title: '余额', field: 'gold', formatter: goldFormatter},
        {title: '总收入', field: 'total_tax', formatter: goldFormatter},
        {title: '注册时间', field: 'register_time'},
        {
            title: '操作', formatter: function (value, row) {
                return '<a data-toggle="modal" data-target="#data-modal" data-remote="{{.domain}}forms/?a=change_rate&id='+value+'&rate=70">修改提成</a>'
            }
        }
    ];

    function goldFormatter(value) {
        return (value / 100).toFixed(2)
    }
    dataurl = location.href;
</script>
<script src="/static/js/fish.js?{{.rand}}"></script>
<script>
    fishApp.dataPage();
</script>
