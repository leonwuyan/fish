<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="#"> 游戏配置</a></li>
            <li>代理列表</li>
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
                代理列表
            </div>
            <div class="panel-body">
                <table class="table table-striped table-bordered table-hover" id="dataTable">
                </table>
            </div>
        </div>
    </div>
</div>
<script>
    datacolumns = [
        {title: 'ID', field: 'agent_id'},
        {title: '名称', field: 'agent_name'},
        {title: '描述', field: 'remarks'},
        {title: 'QQ', field: 'QQ'},
        {title: '微信', field: 'wen_xin'},
        {
            title: '操作', formatter: function (value, row) {
                return '<a href="{{.host}}sys/agent?id="' + row.agent_id + '>详细信息</a>'
            }
        }
    ];
    dataurl = location.href;
    fishApp.dataPage();
</script>
