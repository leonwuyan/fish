<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="renderer" content="webkit">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
    <title>自由娱乐-下载</title>
    <script src="/static/js/jquery-1.10.2.js"></script>
    <script src="/static/js/clipboard.js"></script>
    <script src="/static/js/swiper.4.3.3.min.js"></script>
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
    <link rel="stylesheet" href="/static/css/swiper.4.3.3.min.css">
    <link rel="stylesheet" href="/static/css/qp_style.css">
</head>
<body onload="initialization()">
<div class="warp">
    <a href="#" data-clipboard-text="###{{.agentId}}">
        <div class="header">
            <div class="game-logo">
                <img src="/static/img/down/logo.png" alt="自由娱乐">
            </div>
            <div class="game-name">
                <p>自由娱乐</p>
                <p>火爆全球超刺激的棋牌游戏</p>
            </div>
            <div class="game-down">
                <img src="/static/img/down/game-down.png" alt="" class="gameDown">
            </div>
        </div>
    </a>
    <div class="game-cont">
        <div class="game-down-2">
            <a href="#" data-clipboard-text="###{{.agentId}}">
                <img src="/static/img/down/game-down-2.png" alt="" class="gameDown">
            </a>
        </div>
        <div class="game-num swiper-container">
            <ul class="swiper-wrapper">
                <li class="swiper-slide"><img src="/static/img/down/1.png" alt="捕鱼"></li>
                <li class="swiper-slide"><img src="/static/img/down/2.png" alt="斗地主"></li>
                <li class="swiper-slide"><img src="/static/img/down/3.png" alt="炸金花"></li>
                <li class="swiper-slide"><img src="/static/img/down/4.png" alt="金典牛牛"></li>
                <li class="swiper-slide"><img src="/static/img/down/5.png" alt="红黑大战"></li>
                <li class="swiper-slide"><img src="/static/img/down/6.png" alt="龙虎斗"></li>
                <li class="swiper-slide"><img src="/static/img/down/7.png" alt="百人牛牛"></li>
            </ul>
        </div>
        <img src="/static/img/down/img-03.jpg" alt="">
        <img src="/static/img/down/img-04.jpg" alt="">
        <img src="/static/img/down/img-05.jpg" alt="">
        <img src="/static/img/down/img-06.jpg" alt="">
    </div>
</div>
<script src="/static/js/tool.js"></script>
<script>
    var swiper = new Swiper('.swiper-container', {
        slidesPerView: 'auto',
        freeMode: true
    });
    var clip = new Clipboard('a');
</script>
</body>
</html>
