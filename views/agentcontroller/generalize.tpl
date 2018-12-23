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
                <div id="qrcode" style="display: none">
                </div>
                <div id="canbox0" class="canbox">
                    <canvas id="myCanvas0" width="" height=""></canvas>
                    <button class="btn btn-info" onclick="downLoad(0)">立即保存</button>
                </div>
                <div id="canbox1" class="canbox">
                    <canvas id="myCanvas1" width="" height=""></canvas>
                    <button class="btn btn-info" onclick="downLoad(1)">立即保存</button>
                </div>
                <div id="canbox2" class="canbox">
                    <canvas id="myCanvas2" width="" height=""></canvas>
                    <button class="btn btn-info" onclick="downLoad(2)">立即保存</button>
                </div>
                <div id="canbox3" class="canbox">
                    <canvas id="myCanvas3" width="" height=""></canvas>
                    <button class="btn btn-info" onclick="downLoad(3)">立即保存</button>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    var clipboard = new Clipboard('.copy');
    $("#qrcode").qrcode({
        width: 140, //宽度
        height: 140, //高度
        text: "{{.ad_url}}", //任意内容
    });
    function darwQrCode(id) {
        //画海报
        var width = document.getElementById("canbox" + id).offsetWidth; //宽度
        var height = document.getElementById("canbox" + id).offsetHeight; // 高度
        var c = document.getElementById("myCanvas" + id);
        c.width = width;
        c.height = height;
        var ctx = c.getContext("2d");
        //首先画上背景图
        var img = new Image();
        img.src = "/static/img/bg" + id + ".png";
        img.setAttribute("crossOrigin", 'Anonymous');

        function convertCanvasToImage(canvas) {
            var image = new Image();
            image.src = canvas.toDataURL("image/jpeg");
            return image;
        }

        var mycans = $('canvas')[0];//二维码所在的canvas
        var codeimg = convertCanvasToImage(mycans);
        var xw = width - 72 - 29;
        var xh = height - 6 - 72;

        img.onload = function () { //必须等待图片加载完成
            ctx.drawImage(img, 0, 0, width, height); //绘制图像进行拉伸
            ctx.drawImage(codeimg, 67, 275, 105, 105);
        }
    }
    darwQrCode(0);
    darwQrCode(1);
    darwQrCode(2);
    darwQrCode(3);
    function downLoad(id){
        //cavas 保存图片到本地  js 实现
        //------------------------------------------------------------------------
        //1.确定图片的类型  获取到的图片格式 data:image/Png;base64,......
        var type ='image/jpeg';//你想要什么图片格式 就选什么吧
        var d=document.getElementById("myCanvas" + id);
        var imgdata=d.toDataURL(type);
        //2.0 将mime-type改为image/octet-stream,强制让浏览器下载
        var fixtype=function(type){
            type=type.toLocaleLowerCase().replace(/jpg/i,'jpeg');
            var r=type.match(/png|jpeg|bmp|gif/)[0];
            return 'image/'+r;
        };
        imgdata=imgdata.replace(fixtype(type),'image/octet-stream');
        //3.0 将图片保存到本地
        var saveFile=function(data,filename)
        {
            var save_link=document.createElementNS('http://www.w3.org/1999/xhtml', 'a');
            save_link.href=data;
            save_link.download=filename;
            var event=document.createEvent('MouseEvents');
            event.initMouseEvent('click',true,false,window,0,0,0,0,0,false,false,false,false,0,null);
            save_link.dispatchEvent(event);
        };
        var filename='game.jpg';
        //注意咯 由于图片下载的比较少 就直接用当前几号做的图片名字
        saveFile(imgdata,filename);
    }
</script>
