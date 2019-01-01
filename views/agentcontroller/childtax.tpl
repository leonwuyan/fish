<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li><a href="{{.host}}agents">我的代理</a></li>
            <li>收入明细</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel-heading">
                收入明细
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
        {title: '时间', field: 'log_time'},
        {title: '玩家ID', field: 'user_id'},
        {title: '玩家昵称', field: 'user_name'},
        {
            title: '游戏', formatter: function (value, row) {
                return fishApp.getGameName(row.game_id) + ":房间:" + row.room_id;
            }
        },
        {title: '总税收', field: 'tax', formatter: fishApp.formatter.gold},
        {title: '提成比例', field: 'rate', formatter: fishApp.formatter.rate},
        {title: '下级分成比例', field: 'child_rate', formatter: fishApp.formatter.rate},
        {title: '我的收入', field: 'fee', formatter: fishApp.formatter.gold},
    ];
    dataurl = location.href;
    fishApp.dataPage();
</script>
