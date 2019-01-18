<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> 概况</a></li>
            <li><a href="#"> 数据统计</a></li>
            <li>游戏数据</li>
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
                游戏数据
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <span class="input-group-addon">开始</span>
                            <input id="begin" name="begin" class="form-control" type="date"
                                   value="{{thismonth}}" placeholder="" required>
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
                <hr>
                <div id="s_poundage" class="panel panel-danger">
                    <div class="panel-heading">
                        <span>玩家输赢（分）
                            <span class="badge pull-right" id="s_poundage_i"></span>
                        </span>
                    </div>
                    <div class="panel-body easypiechart-panel">
                    </div>
                </div>
                <div id="s_players" class="panel panel-info">
                    <div class="panel-heading">
                        <span>玩家人数
                            <span class="badge pull-right" id="s_players_i"></span>
                        </span>
                    </div>
                    <div class="panel-body easypiechart-panel">
                    </div>
                </div>
                <div id="s_playtimes" class="panel panel-warning">
                    <div class="panel-heading">
                        <span>游戏次数
                            <span class="badge pull-right" id="s_playtimes_i"></span>
                        </span>
                    </div>
                    <div class="panel-body easypiechart-panel">
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<script src="/static/js/easypiechart.js"></script>
<script>
    $(function () {
        $("#search-form").submit(function () {
            getData();
            return false
        });
        getData();
    });
    chart_styles = ["easypiechart-orange", "easypiechart-teal", "easypiechart-red", "easypiechart-blue", "easypiechart-green", "easypiechart-yellow"];

    function chartPie(id, name, total, num) {
        percent = (num / total * 100).toFixed(2);
        chart_title = document.createElement("h3");
        chart_title.innerText = name;
        chart_percent = document.createElement("span");
        chart_percent.innerText = percent + "%";
        chart_chart_value = document.createElement("span");
        chart_chart_value.className = "value";
        chart_chart_value.innerText = num;
        chart_chart = document.createElement("div");
        chart_chart.className = "easypiechart " + chart_styles[id % 6];
        chart_chart.setAttribute("data-percent", percent);
        chart_chart.appendChild(chart_chart_value);
        chart = document.createElement("div");
        chart.className = "col-lg-2 col-md-3 col-sm-6";
        chart.appendChild(chart_title);
        chart.appendChild(chart_chart);
        chart.appendChild(chart_percent);
        return chart
    }

    function chartRender() {
        $('.easypiechart-teal').easyPieChart({
            scaleColor: false,
            barColor: '#1ebfae'
        });
        $('.easypiechart-orange').easyPieChart({
            scaleColor: false,
            barColor: '#ffb53e'
        });
        $('.easypiechart-red').easyPieChart({
            scaleColor: false,
            barColor: '#f9243f'
        });
        $('.easypiechart-blue').easyPieChart({
            scaleColor: false,
            barColor: '#30a5ff'
        });
        $('.easypiechart-green').easyPieChart({
            scaleColor: false,
            barColor: '#9acd32'
        });
        $('.easypiechart-yellow').easyPieChart({
            scaleColor: false,
            barColor: '#e9da00'
        });
    }

    function getData() {
        s_players = $("#s_players .panel-body");
        s_playtimes = $("#s_playtimes .panel-body");
        s_poundage = $("#s_poundage .panel-body");
        s_players.html('加载中...');
        s_playtimes.html('加载中...');
        s_poundage.html('加载中...');
        $.post(location.href, $("#search-form").serialize(), function (res) {
            s_players.html('');
            s_playtimes.html('');
            s_poundage.html('');
            var total_win = total_times = 0;
            res.data.forEach(function (item) {
                total_win += item.win_or_lose;
                total_times += item.play_times;
            });
            $("#s_poundage_i").html('<span class="fa fa-diamond"></span> ' + total_win);
            $("#s_players_i").html('<span class="fa fa-user"></span> ' + res.total);
            $("#s_playtimes_i").html('<span class="fa fa-bell"></span> ' + total_times);

            res.data.forEach(function (item) {
                gameName = fishApp.getGameName(item.game_id);
                s_poundage.append(chartPie(item.game_id, gameName, total_win, item.win_or_lose));
                s_players.append(chartPie(item.game_id, gameName, res.total, item.play_players));
                s_playtimes.append(chartPie(item.game_id, gameName, total_times, item.play_times));
                chartRender()
            });
        })
    }
</script>