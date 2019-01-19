<link href="/static/css/switch.css" rel="stylesheet">
<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li><a href="#"> 游戏配置</a></li>
            <li><a href="{{.host}}channel/list"> 代理列表</a></li>
            <li>代理信息</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel-heading">
                代理信息
            </div>
        {{if .err}}
            <div class="panel-body">
                <div class="alert alert-danger">{{.err}}</div>
            </div>
        {{else}}
            <div class="panel-body">
                <form class="form-group" method="post" action="{{.host}}sys/agent">
                    <table class="table table-striped table-hover table-bordered">
                        <tbody>
                        <tr>
                            <td>代理ID</td>
                            <td>
                                <input class="form-control" id="agent_id" name="agent_id" value="{{.data.AgentId}}" readonly placeholder=""/>
                            </td>
                        </tr>
                        <tr>
                            <td>描述</td>
                            <td><input class="form-control" id="remark" name="remark" value="{{.data.Remarks}}"
                                       placeholder=""/></td>
                        </tr>
                        <tr>
                            <td>QQ</td>
                            <td><input class="form-control" id="qq" name="qq" value="{{.data.QQ}}"
                                       placeholder=""/></td>
                        </tr>
                        <tr>
                            <td>WenXin</td>
                            <td><input class="form-control" id="wx" name="wx" value="{{.data.WenXin}}"
                                       placeholder=""/></td>
                        </tr>
                        <tr>
                            <td>是否显示</td>
                            <td>
                                <div class="switch-container">
                                    <input id="show" name="show" type="checkbox"
                                           {{if .data.ShowInGame }}checked{{end}} class="switch"/>
                                    <label for="show"></label>
                                </div>
                            </td>
                        </tr>
                        </tbody>
                    </table>
                    <button type="reset">重置</button>
                    <button type="button" onclick="save()">保存
                    </button>
                </form>
            </div>
        {{end}}
        </div>
    </div>
</div>
<script>
    dataurl = location.href;

    function save() {
        fishApp.putAction(location.href, $('form').serializeArray());
    }
</script>