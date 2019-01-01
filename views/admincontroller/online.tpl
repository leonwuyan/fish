<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> 首页</a></li>
            <li><a href="#"> 数据统计</a></li>
            <li>在线数据</li>
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
                在线数据
            </div>
            <div id="toolbar" class="data-collapse">
                <form id="search-form" class="form-inline" role="form">
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
                    <button id="showChart" class="btn btn-info form-control" type="button"><span
                            class="fa fa-table"></span></button>
                </form>
            </div>
            <div class="panel-body">
                <div class="panel">
                    <span class="badge">最大：<span id="max"></span></span>
                    <span class="badge">最小：<span id="min"></span></span>
                    <span class="badge">平均：<span id="avg"></span></span>
                </div>
                <table class="table table-striped table-bordered table-hover" id="dataTable" style="display: none;">
                </table>
                <div class="chart">
                    <div id="player_statistic"></div>
                </div>
            </div>
        </div>
    </div>
</div>
<script src="/static/js/morris/raphael-2.1.0.min.js"></script>
<script src="/static/js/morris/morris.js"></script>
<script>
    datacolumns = [
        {title: '时间', field: 'StatDate', switchable: false},
        {title: '人数', field: 'AllCount'},
    ];
    dataurl = location.href;
    showFooter = false;
</script>
<script src="/static/js/query_data.js?{{.rand}}"></script>
<script src="/static/js/statistic-chart.js"></script>
<script>
    pagination = false;
    showColumns = true;
    dataTable = $('#dataTable');
    elementName = 'player_statistic';
    dataTable.on('column-switch.bs.table', function () {
        drawLine(elementName, $(this));
    });
    dataTable.on('load-success.bs.table', function () {
        drawLine(elementName, $(this));
        data = dataTable.bootstrapTable('getData');
        onlines = statisticTable(data);
        $("#max").html(onlines.max);
        $("#min").html(onlines.min);
        $("#avg").html(onlines.avg);
    });
    $("#showChart").click(function () {
        if ($(this).find('span').attr('class') === "fa fa-table") {
            $(this).find('span').attr('class', "fa fa-line-chart");
            $("#" + elementName).hide();
            dataTable.show();
        } else {
            $(this).find('span').attr('class', "fa fa-table");
            $("#" + elementName).show();
            dataTable.hide()
        }
    });

    function statisticTable(data) {
        totalPlayer = 0;
        max = 0;
        min = 0;
        for (i in data) {
            if (i === "exists")
                break;
            totalPlayer += Number(data[i].AllCount);
            cur = Number(data[i].AllCount);
            if (i === '0') {
                max = min = cur
            }
            if (max < cur) {
                max = cur
            }
            if (min > cur) {
                min = cur
            }
        }
        avg = (totalPlayer / data.length).toFixed(2);
        return {min: min, max: max, avg: avg};
    }
</script>