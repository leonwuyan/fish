<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li><a href="#">日志</a></li>
            <li>游戏日志</li>
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
                游戏日志
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder="请输入玩家ID">
                        </div>
                        <div class="form-group input-group">
                            <select id="game_type" name="game_type" class="form-control">
                                <option value="0">全部</option>
                                <option value="1">捕鱼</option>
                                <option value="2">扎金花</option>
                                <option value="3">斗地主</option>
                                <option value="4">百人牛牛</option>
                                <option value="5">红黑大战</option>
                                <option value="6">抢庄牛牛</option>
                                <option value="7">龙虎斗</option>
                            </select>
                        </div>
                        <div class="form-group input-group">
                            <span class="input-group-addon">开始</span>
                            <input id="begin" name="begin" class="form-control" type="datetime-local"
                                   value="{{.beginT}}" placeholder="" required>
                        </div>
                        <div class="form-group input-group">
                            <span class="input-group-addon">结束</span>
                            <input id="end" name="end" class="form-control" type="datetime-local"
                                   value="{{.endT}}" placeholder="" required>
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
        {title: '游戏', field: 'game_id',formatter:fishApp.getGameName},
        {title: '房间', field: 'room_id'},
        {title: '时间', field: 'create_time'},
        {title: '开始金额', field: 'start_gold', formatter: fishApp.formatter.gold},
        {title: '变化金额', field: 'gold_change', formatter: fishApp.formatter.gold},
    ];
    dataurl =location.href;
    fishApp.dataPage();
</script>
