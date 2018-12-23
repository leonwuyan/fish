<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li>概况</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-warning text-center no-boder red" style="border-radius: 10px">
            <div class="panel-body">
                <div class="info-header">
                    <div class="header-con ">
                        <ul>
                            <li><strong>代理ID：</strong>{{.user.Id}}</li>
                            <li><strong>代理名称：</strong>{{.user.Name}} (
                                <i class="fa fa-key fa-fw"></i>
                                <a href="{{.host}}changepwd" style="color: #fff"> 修改密码</a>)
                            </li>
                            <li><strong>提成比例：</strong><span id="rate"></span>%</li>
                            <li>
                                <strong>可提现金额：</strong>
                                <span id="gold" class="fa fa-diamond" style="margin:0 10px 0 10px;font-size: 18px"> <strong id="gold"></strong></span>
                                <a class="btn btn-danger" style="margin-right: 15px" href="{{.host}}cash">立即提取</a></li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<div class="row">
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel green" style="border-radius: 10px;padding: 10px;height: 120px;">
            <h3><strong id="today_reg"></strong></h3>
            <h3>今日注册</h3>
            <h5>直属：<span id="today_reg1"></span> 非直属：<span id="today_reg2"></span></h5>
        </div>
    </div>
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel blue" style="border-radius: 10px;padding: 10px;height: 120px;">
            <h3><strong id="yesterday_reg">10</strong></h3>
            <h3>昨日注册</h3>
            <h5>直属：<span id="yesterday_reg1">1</span> 非直属：<span id="yesterday_reg2">2</span></h5>
        </div>
    </div>
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel red" style="border-radius: 10px;padding: 10px;height: 120px;">
            <h3><strong id="total_reg">10</strong></h3>
            <h3>总注册</h3>
            <h5>直属：<span id="total_reg1">1</span> 非直属：<span id="total_reg2">2</span></h5>
        </div>
    </div>
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel green" style="border-radius: 10px;padding: 10px;height: 120px;">
            <h3><strong id="today_tax">10</strong></h3>
            <h3>今日收益</h3>
        </div>
    </div>
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel blue" style="border-radius: 10px;padding: 10px;height: 120px;">
            <h3><strong id="yesterday_tax">10</strong></h3>
            <h3>昨日收益</h3>
        </div>
    </div>
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel red" style="border-radius: 10px;padding: 10px;height: 120px;">
            <h3><strong id="total_tax">10</strong></h3>
            <h3>总收益</h3>
        </div>
    </div>
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel green" style="border-radius: 10px;padding: 10px;height: 120px;">
            <h3><strong id="children">10</strong></h3>
            <h3>直属代理</h3>
        </div>
    </div>
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel blue" style="border-radius: 10px;padding: 10px;height: 120px;">
            <h3><strong id="children_children">10</strong></h3>
            <h3>非直属代理</h3>
        </div>
    </div>
</div>
<script>
    $(document).ready(function () {
        fishApp.indexPage();
    });
</script>