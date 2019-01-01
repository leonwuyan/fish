<script src="/static/js/clipboard.js"></script>
<script src="http://cdn.staticfile.org/jquery.qrcode/1.0/jquery.qrcode.min.js"></script>
<style type="text/css">
    .canbox {
        float: left;
        margin: 5px;
        margin-bottom: 35px;
        width: 240px;
        height: 400px;
    }
</style>
<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li><a href="{{.host}}">概况</a></li>
            <li>我的推广</li>
        </ol>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <div class="panel no-boder">
            <div class="panel-body">
                <h4>您的推广链接：</h4>
                <span class="url">{{.ad_url}}</span>
                <a class="btn btn-default copy" href="#" data-clipboard-action="copy"
                   data-clipboard-target=".url">复制链接</a>
            </div>
            <hr>
            <div class="panel-body">
                <h4>您的推广二维码：</h4>
                <div id="canbox0" class="canbox">
                    <img src="{{.domain}}advertise/{{.user.Id}}_0.png" width="240" height="400">
                    <button class="btn btn-info" onclick="downLoad(0)">立即保存</button>
                </div>
                <div id="canbox1" class="canbox">
                    <img src="{{.domain}}advertise/{{.user.Id}}_1.png" width="240" height="400">
                    <button class="btn btn-info" onclick="downLoad(1)">立即保存</button>
                </div>
                <div id="canbox2" class="canbox">
                    <img src="{{.domain}}advertise/{{.user.Id}}_2.png" width="240" height="400">
                    <button class="btn btn-info" onclick="downLoad(2)">立即保存</button>
                </div>
                <div id="canbox3" class="canbox">
                    <img src="{{.domain}}advertise/{{.user.Id}}_3.png" width="240" height="400">
                    <button class="btn btn-info" onclick="downLoad(3)">立即保存</button>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    var clipboard = new Clipboard('.copy');

    function downLoad(id) {
        window.open("{{.domain}}advertise/{{.user.Id}}_" + id + ".png")
    }
</script>
