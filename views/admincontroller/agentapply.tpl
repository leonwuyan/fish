<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> "menu.home"}}</a></li>
            <li><a href="#"> "menu.info"}}</a></li>
            <li> "menu.players"}}</li>
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
                "menu.players"}}
            </div>
            <div class="panel-body">
                <div id="toolbar" class="data-collapse">
                    <form id="search-form" class="form-inline" role="form">
                        <div class="form-group input-group">
                            <input id="id" name="id" class="form-control" type="text"
                                   placeholder=" tips.accountid">
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
        {title: 'ID', field: 'UserId'},
        {title: '11', field: 'nickname'},
        {title: '22', field: 'diamonds'},
        {title: '22', field: 'room_cards'},
        {title: '22', field: 'room_cards'},
        {title: '33', field: 'bind_time'},
        {title: '44', field: 'last_login_time'},
        {
            title: ' "data.action"}}', formatter: function (value, row) {
                if (row.groupId !== 1) {
                    return '<a class="btn btn-info" href="{{.host}}add?id=' + row.id + '"> "content.set_to_agent"}}</a> '
                }
            }
        }
    ];
    dataurl =location.href;
    showFooter = false;
</script>

<script src="/static/js/fish.js?{{.rand}}"></script>
<script>

    fishApp.dataPage();
</script>
