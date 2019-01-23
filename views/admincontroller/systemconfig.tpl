<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> 首页</a></li>
            <li>系统配置</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel-heading">
                站点配置
            </div>
            <div class="panel-body">
                <table class="table table-striped table-hover table-bordered">
                    <tbody>
                    {{range $key,$value := .site}}
                    <tr>
                        <td style="width: 200px">{{$key}}</td>
                        <td>
                            <input type="text" class="form-control" id="site::{{$key}}" value="{{$value}}"
                                   placeholder=""
                                   required>
                        </td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel-heading">
                域配置
            </div>
            <div class="panel-body">
                <table class="table table-striped table-hover table-bordered">
                    <tbody>
                    {{range $key,$value := .domain}}
                    <tr>
                        <td style="width: 200px">{{$key}}</td>
                        <td>
                            <input type="text" class="form-control" id="domain::{{$key}}" value="{{$value}}"
                                   placeholder=""
                                   required>
                        </td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel-heading">
                数据库
            </div>
            <div class="panel-body">
                <table class="table table-striped table-hover table-bordered">
                    <tbody>
                    {{range $key,$value := .db}}
                    <tr>
                        <td style="width: 200px">{{$key}}</td>
                        <td>
                            <input type="text" class="form-control" id="db::{{$key}}" value="{{$value}}"
                                   placeholder=""
                                   required>
                        </td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel-heading">
                短信配置
            </div>
            <div class="panel-body">
                <table class="table table-striped table-hover table-bordered">
                    <tbody>
                    {{range $key,$value := .sms}}
                    <tr>
                        <td style="width: 200px">{{$key}}</td>
                        <td>
                            <input type="text" class="form-control" id="sms::{{$key}}" value="{{$value}}"
                                   placeholder=""
                                   required>
                        </td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel-heading">
                支付渠道配置
            </div>
            <div class="panel-body">
                <p>渠道说明：1、汇易 2、沃汇宝 3、弘佳 4、纵横 5、易佳 7、SunApi</p>
                <table class="table table-striped table-hover table-bordered">
                    <tbody>
                    {{range $key,$value := .payment}}
                    <tr>
                        <td style="width: 200px">{{$key}}</td>
                        <td>
                            <input type="text" class="form-control" id="payment::{{$key}}" value="{{$value}}"
                                   placeholder=""
                                   required>
                        </td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
<script>
    $(function () {
        $("input").focus(function () {
            oldValue = this.value
        }).blur(function () {
            if (this.value !== oldValue) {
                var key = this.id;
                var value = this.value;
                $.post(document.URL, {k: this.id, v: this.value}, function (res) {
                    if (res.state === 0) {
                        toastr.success("修改成功："+key + "=" + value )
                    } else {
                        toastr.success("修改失败："+key + "=" + value)
                    }
                })
            }
        })
    });
</script>
