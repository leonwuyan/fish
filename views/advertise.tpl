<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="renderer" content="webkit">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
    <title>自由娱乐-下载</title>
    <script src="https://code.jquery.com/jquery-1.12.4.min.js"></script>
    <script src="https://cdn.bootcss.com/clipboard.js/1.5.12/clipboard.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Swiper/4.3.3/js/swiper.min.js"></script>
    <script>
        adaptation(750);

        function adaptation(size) {
            if (document.documentElement.clientWidth > size) {
                document.documentElement.style.fontSize = size / 7.5 + "px"
            } else {
                document.documentElement.style.fontSize = document.documentElement.clientWidth / 7.5 + "px"
            }
        }

        window.onresize = function () {
            adaptation(750)
        };
    </script>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/Swiper/4.3.3/css/swiper.min.css">
    <style>
        html,body {line-height:1.5;}
        body {background:#121633;}
        body::-webkit-scrollbar{display: none}
        body,h1,h2,h3,h4,h5,h6,ul,ol,dl,dd,p {margin:0;}
        body,h1,input,button,select,textarea {font-family:"Microsoft Yahei","PingFang SC","Helvetica Neue",serif;font-size:.28rem;color:#1a1a1a;}
        input,button,select,textarea {border:none;outline:none;}
        input,button,textarea,th,td {padding:0;}
        i,em,cite {font-style:normal;}
        ul,ol {padding:0;list-style:none;}
        img {vertical-align:top;border:none;}
        a {text-decoration:none;color:inherit;}

        .warp,.header,.step,.footer{margin:0 auto;width:7.5rem;}
        .warp:after{content:'';display:table;clear:both;}

        .header{position:fixed;display:table;background:rgba(0,0,0,.6);z-index:3;}
        .game-logo,.game-name,.game-down{display:table-cell;height:1.5rem;vertical-align:middle;}
        .game-logo{width:1.27rem;text-align:right;}
        .game-logo img{border-radius:.2rem;width:1rem;height:1rem;}
        .game-name{padding:0 .17rem;font-size:.24rem;font-weight:300;color:#fff;}
        .game-name p:first-child{font-size:.36rem;font-weight:bold;}
        .game-down{width:2.17rem;}
        .game-down img{width:1.8rem;height:.7rem;}
        .game-cont img{width:100%;}
        .game-down-2{position:relative;height:7.4rem;background:url("https://github.com/freegame1010403936/freegame/raw/master/img-01.jpg") no-repeat top / contain;}
        .game-down-2 img{position:absolute;bottom:0.3rem;left:50%;transform:translate(-50%);width:3.45rem}
        .game-num{height:3.3rem;background:url("https://github.com/freegame1010403936/freegame/raw/master/img-02.jpg") no-repeat top / contain;}
        .game-num li{float:left;margin:1.8rem .4rem 0 0;width:1.28rem;height:1.28rem;}
        .game-num li:first-child{margin:1.8rem .4rem 0;}

        .step{background:url("https://github.com/freegame1010403936/freegame/raw/master/step-header.jpg") no-repeat top / contain;text-align:center;}
        .step-header{padding:.5rem 0;font-size:.36rem;color:#fff;}
        .step img{margin-bottom:.3rem;width:6.9rem;}

        .weixin{position:fixed;top:0;right:0;bottom:0;left:0;text-align:center;background:rgba(0,0,0,.8);z-index:9;}
        .weixin img{margin:.3rem 0 0 .4rem;width:5.73rem;height:2.65rem;}
    </style>
</head>
<body>
<div class="warp">
    <a href="#">
        <div class="header">
            <div class="game-logo">
                <img src="https://github.com/freegame1010403936/freegame/raw/master/logo.png" alt="自由娱乐">
            </div>
            <div class="game-name">
                <p>自由娱乐</p>
                <p>火爆全球超刺激的棋牌游戏</p>
            </div>
            <div class="game-down">
                <img src="https://github.com/freegame1010403936/freegame/raw/master/game-down.png" alt="" class="gameDown">
            </div>
        </div>
    </a>
    <div class="game-cont">
        <div class="game-down-2">
            <a href="#">
                <img src="https://github.com/freegame1010403936/freegame/raw/master/game-down-2.png" alt="" class="gameDown">
            </a>
        </div>
        <div class="game-num swiper-container">
            <ul class="swiper-wrapper">1.png
                <li class="swiper-slide"><img src="https://github.com/freegame1010403936/freegame/raw/master/1.png" alt="捕鱼"></li>
                <li class="swiper-slide"><img src="https://github.com/freegame1010403936/freegame/raw/master/2.png" alt="斗地主"></li>
                <li class="swiper-slide"><img src="https://github.com/freegame1010403936/freegame/raw/master/3.png" alt="炸金花"></li>
                <li class="swiper-slide"><img src="https://github.com/freegame1010403936/freegame/raw/master/4.png" alt="金典牛牛"></li>
                <li class="swiper-slide"><img src="https://github.com/freegame1010403936/freegame/raw/master/5.png" alt="红黑大战"></li>
                <li class="swiper-slide"><img src="https://github.com/freegame1010403936/freegame/raw/master/6.png" alt="龙虎斗"></li>
                <li class="swiper-slide"><img src="https://github.com/freegame1010403936/freegame/raw/master/7.png" alt="百人牛牛"></li>
            </ul>
        </div>
        <img src="https://github.com/freegame1010403936/freegame/raw/master/img-03.jpg" alt="">
        <img src="https://github.com/freegame1010403936/freegame/raw/master/img-04.jpg" alt="">
        <img src="https://github.com/freegame1010403936/freegame/raw/master/img-05.jpg" alt="">
        <img src="https://github.com/freegame1010403936/freegame/raw/master/img-06.jpg" alt="">
    </div>
</div>+
<script>
    function request(name) {
    reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
}
    var swiper = new Swiper('.swiper-container', {
        slidesPerView: 'auto',
        freeMode: true
    });
    var clip = new Clipboard('a', {
        text: function () {
            return '###' + request("id")
        }
    });
    const tool = {
        linkFlag: '#',
        userAgent: navigator.userAgent,

        isiOSAutoDownload: false, // if true, auto download ios package
        iosDownloadUrl: 'itms-services://?action=download-manifest&url=https://raw.githubusercontent.com/freegame1010403936/freegame/master/manifest.plist',

        isAndroidAutoDownload: false, // if true, auto download android package
        androidDownloadUrl: 'http://res-f556678-com.oss-cn-hongkong.aliyuncs.com/freegame.apk',

        isPCAutoDownload: false,// if true, auto download android package on PC

        initialization: function (){
            this.createExplainStyle();
            this.createExplainImageDOM();
            this.createWeiXinDOM();

            // 置換 ios download url 如果有設置downurl
            this.iosDownloadUrl = this.getQueryString('downurl') || this.iosDownloadUrl;

            let is_weixin = this.is_weixin();
            let is_ios = this.is_ios();
            if (is_weixin){
                if (is_ios){
                    document.getElementsByName('wbllq')[1].style.display="block";
                } else{
                    document.getElementsByName('wbllq')[0].style.display="block";
                }
                return;
            }
            this.spreadDownloadLink();
            this.changeurl();
            this.autoDownload();
        },
        // 將所有標籤為a的連結換掉
        spreadDownloadLink: function() {
            let as = document.getElementsByTagName('a');
            let url = this.downloadgame();
            for(let i=0,j=as.length;i < j;i++){
                if(as[i].href.indexOf(this.linkFlag)>=0){
                    as[i].href=url;
                }
            }
            this.linkFlag = url;
        },
        loadHtml: function (){
            let step = this.getQueryString('enterstep');
            if(step=='1'){
                document.getElementsByName('dsdsd')[0].style.display="block";
            }
        },
        closeHtml: function (){
            document.getElementsByName('dsdsd')[0].style.display="none";
        },
        changeurl: function (){
            let url = window.location.href;
            let step = this.getQueryString('enterstep');
            if(step=='1'){
                let newurl = this.changeUrlArg(url,'enterstep',2);
                let stateObj = {};
                history.pushState(stateObj, '', newurl);
            }
        },
        autoDownload: function () {

            let isAndroid = this.is_android();
            let isiOS = this.is_ios();
            let isAPPStorePackage = this.is_appstore_package();

            // 企业包不能自动下载
            if(this.isiOSAutoDownload && isiOS && isAPPStorePackage)  {
                location.href = this.iosDownloadUrl;
            }
            else if(this.isAndroidAutoDownload && isAndroid) {
                location.href = this.androidDownloadUrl;
            }
            else if(this.isPCAutoDownload){
                location.href = this.androidDownloadUrl;
            }
        },
        downloadgame: function () {
            let new_title = this.getQueryString('title');
            if(''!=new_title&&undefined!=new_title&&'null'!=new_title){
                document.title = new_title;
            }
            let isAndroid = this.is_android();
            let isiOS = this.is_ios();
            let ispc = this.is_pc();

            if(isAndroid){
                return this.androidDownloadUrl;
            }else if(isiOS){
                this.loadHtml();
                return this.iosDownloadUrl;
            } else if (ispc) {
                return this.androidDownloadUrl;
            }else{
                return this.androidDownloadUrl;
            }
        },
        getQueryString: function (name) {
            let reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
            let r = window.location.search.substr(1).match(reg);

            if (r != null) {
                return decodeURI(r[2]);
            }
            return null;
        },
        changeUrlArg: function (url, arg, val){
            let pattern = arg+'=([^&]*)';
            let replaceText = arg+'='+val;
            return url.match(pattern) ? url.replace(eval('/('+ arg+'=)([^&]*)/gi'), replaceText) : (url.match('[\?]') ? url+'&'+replaceText : url+'?'+replaceText);
        },
        // 載入圖
        createExplainImageDOM: function() {
            let createView = `
                <div name="dsdsd" class="max_width" style="display: none;z-index:1000;">
                    <div class="wh_float">
                        <div class="wh_opacity" style="width:100%;height: 100%;"></div>
                    </div>
                    <div class="wh_float">
                        <div style="width:100%;height: 7%;z-index:1000;"></div>
                        <div style="width:100%;height: 93%;z-index:1000;"><img style="width:100%;height: 100%;" src="https://github.com/freegame1010403936/freegame/raw/master/wh_main.png"></div>
                    </div>
                    <div class="wh_float2">
                        <div style="width:100%;height: 6%;z-index:1000;"><img style="height: 100%;margin-left:15%;margin-top: 1.3%;width:70%;float: left" src="https://github.com/freegame1010403936/freegame/raw/master/wh_detail.png">
                            <img style="height: 75%;width:7%;margin-left:1.2%;margin-top: 2.2%;float: left;z-index:1000;" src="https://github.com/freegame1010403936/freegame/raw/master/wh_close.png" alt="关闭" onclick="javascript:tool.closeHtml();">
                        </div>
                        <div style="width: 100%;height: 73%;z-index:1000;"></div>
                        <div style="width: 100%;height: 20%;text-align: center;vertical-align: middle;z-index:1000;"><img style="margin-top:6%;height:60%;width:5.4%;z-index:1000;" src="https://github.com/freegame1010403936/freegame/raw/master/wh_output.gif"/></div>
                    </div>
                </div>`;
            document.getElementsByTagName('body')[0].innerHTML = createView + document.getElementsByTagName('body')[0].innerHTML;
        },
        //載入CSS
        createExplainStyle: function () {
            let style = `
                    <style type="text/css">
                        .wh_float {
                            max-width: 720px;
                            width: 100%;
                            height: 100%;
                            position: fixed;
                            display: inline-block;
                            z-index: 1001 !important;
                        }
                        .wh_opacity{
                            opacity: 0.75;
                            background-color: rgba(0, 0, 0, 1) !important;
                            background-color: #000;
                            filter: alpha(opacity=100);
                            z-index: 1001;
                        }

                        .wh_float2 {
                            max-width: 720px;
                            height: 100%;
                            width: 100%;
                            position: fixed;
                            display: inline-block;
                            opacity: 1;
                            z-index: 1001;
                        }
                    </style>`;
            document.getElementsByTagName('head')[0].innerHTML += style;
        },
        //微信圖(指示以其他瀏覽器開啟)
        createWeiXinDOM: function() {
            let createView = `
                <div class="weixin" name="wbllq" style="display:none;z-index:1000;position:fixed;top:0;right:0;bottom:0;left:0;text-align:center;background:rgba(0,0,0,.8);">
                    <img src="https://github.com/freegame1010403936/freegame/raw/master/weixinTip.png" alt="" style="margin:.3rem 0 0 .4rem;width:90%;height:auto;margin: .3rem 0 0 .4rem;">
                </div>
                <div class="weixin" name="wbllq" style="display:none;z-index:1000;position:fixed;top:0;right:0;bottom:0;left:0;text-align:center;background:rgba(0,0,0,.8);">
                    <img src="https://github.com/freegame1010403936/freegame/raw/master/weixinTip2.png" alt="" style="margin:.3rem 0 0 .4rem;width:90%;height:auto;margin: .3rem 0 0 .4rem;">
                </div>`;
            document.getElementsByTagName('body')[0].innerHTML += createView;
        },
        is_weixin: function () {
            let ua = this.userAgent.toLowerCase();
            if (ua.match(/MicroMessenger/i) == 'micromessenger') {
                return true;
            } else {
                return false;
            }
        },
        is_ios: function() {
            let u = this.userAgent;
            return !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/)  || u.indexOf('iphone') > -1  ||u.indexOf('ipad') > -1;
        },
        is_android: function(){
            let u = this.userAgent;
            return u.indexOf('Android') > -1 || u.indexOf('Adr') > -1  ||  u.indexOf('android') > -1 || u.indexOf('linux') > -1;
        },
        is_pc: function() {
            let u = this.userAgent;
            return u.indexOf('Windows') > -1;
        },
        is_appstore_package: function () {
            let url = this.getQueryString('downurl') || this.iosDownloadUrl;
            if (!!url && url.match(/itunes.apple.com/i) == 'itunes.apple.com'){
                return true;
            }else{
                return false;
            }
        }
    };

    document.body.onload = tool.initialization();
</script>
</body>
</html>
