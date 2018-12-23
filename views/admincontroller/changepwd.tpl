<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}"> 概况</a></li>
            <li>修改密码</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-lg-6">
        <div class="panel panel-default">
            <div class="panel-heading">
                修改密码
            </div>
            <div class="panel-body">
                <form role="form" method="post">
                    <div class="form-group input-group">
                        <span class="input-group-addon">原密码</span>
                        <input class="form-control" id="old" name="old" type="password"
                               placeholder="请输入原密码" required/>
                    </div>
                    <div class="form-group input-group">
                        <span class="input-group-addon">新密码</span>
                        <input class="form-control" id="new" name="new" type="password"
                               placeholder="请输入新密码" required/>
                    </div>
                    <div class="form-group input-group">
                        <span class="input-group-addon">新密码</span>
                        <input class="form-control" id="new1" name="new1" type="password"
                               placeholder="请再次输入新密码" required/>
                    </div>
                    <div class="form-group" style="text-align: center">
                        <button class="btn btn-info" type="submit">修改密码</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>
<script>
    $("form").submit(function () {
        return fishApp.changepwd();
    })
</script>
