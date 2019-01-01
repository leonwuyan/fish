<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> 首页</a></li>
            <li>管理员列表</li>
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
                管理员列表
            </div>
            <div id="toolbar" class="data-collapse">
                <div class="btn-group">
                    <button id="btn-add" type="button" class="btn btn-info"
                            data-toggle="modal"
                            data-target="#data-modal"
                            data-remote="{{.domain}}forms/?a=add_admin"><span
                            class="fa fa-plus"></span>添加
                    </button>
                </div>
            </div>
            <div class="panel-body">
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
        {title: '账号', field: 'name'},
        {title: '权限', field: 'permissions'},
        {title: '创建时间', field: 'create_time'},
        {title: '状态', field: 'frozen_time',formatter: stateFormatter},
        {title: '操作', field: 'id', formatter: actionFormatter}
    ];

    function actionFormatter(value) {
        if (value !== 1) {
            edit = '<button class="btn btn-warning" data-toggle="modal" data-target="#data-modal"data-remote="{{.domain}}forms/?a=edit_admin&id=' + value + '"><span class="fa fa-edit"></span></button> ';
            del = '<button class="btn btn-danger" data-account-id="' + value + '" onclick="delete_account(this)"><span class="fa fa-close"></span></button> ';
            return edit + del;
        }
        return ""
    }
    function stateFormatter(value) {
        if (new Date(value) < new Date()){
            return '正常';
        }
        return '封停';
    }

    function delete_account(el) {
        $.ajax({
            url: location.href,
            method: "put",
            data: {action: "del", id: $(el).data("account-id")},
            success: function (res) {
                if (res.status !== 0) {
                    alert(res.error);
                } else {
                    $("#dataTable").bootstrapTable('refresh');
                }
            }
        });
    }
    dataurl = location.href;
    fishApp.dataPage();
</script>