<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> 概况</a></li>
            <li><a href="#"> 代理管理</a></li>
            <li> 代理申请</li>
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
                代理申请
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder="请输入玩家ID">
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
        {title: '申请时间', field: 'apply_time'},
        {title: '姓名', field: 'name'},
        {title: '电话', field: 'phone'},
        {title: '邮箱', field: 'email'},
        {title: 'QQ', field: 'qq'},
        {title: '微信', field: 'wei_xin'},
        {title: '消息内容', field: 'message'},
        {title: '是否处理', field: 'is_deal'},
        {title: '处理时间', field: 'deal_time'},
        {title: '回复内容', field: 'reply'},
    ];
    dataurl =location.href;
    showFooter = false;
    fishApp.dataPage();
</script>
