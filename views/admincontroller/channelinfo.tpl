<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li><a href="#"> 系统配置</a></li>
            <li><a href="{{.host}}channel/list"> 渠道列表</a></li>
            <li>渠道信息</li>
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
                渠道信息
            </div>
            <div class="panel-body">
                <table class="table table-striped table-hover table-bordered">
                    <tbody>
                    <tr>
                        <td>渠道ID</td>
                        <td>
                            <img src='{{.player.Header}}' class="header-img"/>
                        </td>
                    </tr>
                    <tr>
                        <td>) "data.nickname"}}</td>
                        <td>{{.player.Nickname}}</td>
                    </tr>
                    <tr>
                        <td>) "data.accountid"}}</td>
                        <td>{{.player.Id}}</td>
                    </tr>
                    <tr>
                        <td>) "data.diamonds"}}</td>
                        <td>{{.player.Diamonds}}</td>
                    </tr>
                    <tr>
                        <td>) "data.roomcards"}}</td>
                        <td>{{.player.RoomCards}}</td>
                    </tr>
                    <tr>
                        <td>) "data.createtime"}}</td>
                        <td>{{dateformat .player.CreateTime "2006-01-02 15:04:05"}}</td>
                    </tr>
                    <tr>
                        <td>) "data.lastlogin"}}</td>
                        <td>{{dateformat .player.LastLoginTime "2006-01-02 15:04:05"}}</td>
                    </tr>
                    <tr>
                        <td>) "data.status"}}</td>
                        <td id="status"></td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>