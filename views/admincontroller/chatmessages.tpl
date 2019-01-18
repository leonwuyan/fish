<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> 概况</a></li>
            <li><a href="#"> 客服</a></li>
            <li> 消息管理</li>
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
                消息管理
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder="请输入玩家ID">
                        </div>
                        <div class="form-group input-group">
                            <select id="msg_type" name="msg_type" class="form-control">
                                <option value="1">客服消息</option>
                                <option value="2">商务消息</option>
                                <option value="3">全民代理消息</option>
                                <option value="4">系统消息</option>
                            </select>
                        </div>
                        <button class="btn btn-info form-control" type="submit"><span
                                class=" fa fa-search"></span></button>
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
        {title: '消息ID', field: 'id'},
        {title: '时间', field: 'creation_time'},
        {title: '玩家ID', field: 'user_id'},
        {title: '消息类型', field: 'message_type', formatter: fishApp.formatter.msgType},
        {title: '内容', field: 'message'},
        {title: '是否处理', field: 'is_processed', formatter: processedFormatter},
        {title: '客服编号', field: 'kefu_id'},
        {title: '操作', formatter: actionFormatter}
    ];

    function actionFormatter(value, row) {
        if (!row.is_processed || row.is_user_message) {
            return '<a  class="btn btn-warning" onclick="setState(' + row.id + ')">设置为已处理</a>' +
                    '<a class="btn btn-info" data-toggle="modal" data-target="#data-modal" data-remote="{{.domain}}forms/?a=add_service_msg&id=' + row.id + '&uid=' + row.user_id + '">回复</a>';

        }
    }

    function processedFormatter(value) {
        return value ? '<span class="text-success">已处理</span>' : '<span class="text-danger">未处理</span>';
    }

    function setState(id) {
        fishApp.putAction(location.href, {"action": "set_state", "id": id})
    }

    dataurl = location.href;
    showFooter = false;
    fishApp.dataPage();
</script>
