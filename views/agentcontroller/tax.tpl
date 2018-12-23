<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
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
                return getGameName(row.game_id) + ":房间:" + row.room_id;
            }
        },
        {title: '总税收', field: 'tax', formatter: goldFormatter},
        {
            title: '提成比例', field: 'rate', formatter: function (value) {
                return value + '%'
            }
        },
        {
            title: '下级分成比例', field: 'child_rate', formatter: function (value) {
                return value + '%'
            }
        },
        {title: '我的收入', field: 'fee', formatter: goldFormatter},
    ];
    dataurl = location.href;
    showFooter = false;

    function getGameName(id) {
        switch (id) {
            case 1:
                return "捕鱼";
            case 2:
                return "扎金花";
            case 3:
                return "斗地主";
            case 4:
                return "百人牛牛";
            case 5:
                return "红黑大战";
            case 6:
                return "抢庄牛牛";
            case 7:
                return "龙虎斗";
        }
    }

    function goldFormatter(value) {
        return (value / 100).toFixed(2)
    }
</script>
<script src="/static/js/fish.js?{{.rand}}"></script>
<script>
    fishApp.dataPage();
</script>
