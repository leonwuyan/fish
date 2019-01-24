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
                        <div style="width:100%;height: 93%;z-index:1000;"><img style="width:100%;height: 100%;" src="https://gitee.com/freegame1010403936/fish/raw/master/img/tool/wh_main.png"></div>
                    </div>
                    <div class="wh_float2">
                        <div style="width:100%;height: 6%;z-index:1000;"><img style="height: 100%;margin-left:15%;margin-top: 1.3%;width:70%;float: left" src="https://gitee.com/freegame1010403936/fish/raw/master/img/tool/wh_detail.png">
                            <img style="height: 75%;width:7%;margin-left:1.2%;margin-top: 2.2%;float: left;z-index:1000;" src="https://gitee.com/freegame1010403936/fish/raw/master/img/tool/wh_close.png" alt="关闭" onclick="javascript:tool.closeHtml();">
                        </div>
                        <div style="width: 100%;height: 73%;z-index:1000;"></div>
                        <div style="width: 100%;height: 20%;text-align: center;vertical-align: middle;z-index:1000;"><img style="margin-top:6%;height:60%;width:5.4%;z-index:1000;" src="https://gitee.com/freegame1010403936/fish/raw/master/img/tool/wh_output.gif"/></div>
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
                    <img src="https://gitee.com/freegame1010403936/fish/raw/master/img/tool/weixinTip.png" alt="" style="margin:.3rem 0 0 .4rem;width:90%;height:auto;margin: .3rem 0 0 .4rem;">
                </div>
                <div class="weixin" name="wbllq" style="display:none;z-index:1000;position:fixed;top:0;right:0;bottom:0;left:0;text-align:center;background:rgba(0,0,0,.8);">
                    <img src="https://gitee.com/freegame1010403936/fish/raw/master/img/tool/weixinTip2.png" alt="" style="margin:.3rem 0 0 .4rem;width:90%;height:auto;margin: .3rem 0 0 .4rem;">
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