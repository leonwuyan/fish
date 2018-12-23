<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <title>{{.site.title}}</title>
    <link href="/static/css/bootstrap.css" rel="stylesheet"/>
    <link href="/static/css/font-awesome.css" rel="stylesheet"/>
    <link href="/static/css/custom-styles.css" rel="stylesheet"/>
    <script src="/static/js/jquery-1.10.2.js"></script>
    <script src="/static/js/jquery.cookie.js"></script>
</head>
<style>
    #login-wrapper {
        text-align: center;
        margin: auto;
    }

    #login_page {
        padding-top: 100px;
        max-width: 680px;
        margin: auto;
        text-align: center;
    }

    .inner-tittle {
        margin: 15px;
    }

    .t-inner {
        color: #444444;
    }

    .login {
        background-color: #fff;
        border-radius: 10px;
        width: 90%;
        margin: auto;
        padding: 10px;
        text-align: center;
    }

    .form-control {
        position: relative;
        height: auto;
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;
        box-sizing: border-box;
        padding: 10px;
        font-size: 16px;
    }
</style>
<body>
<div id="wrapper">
    <nav class="navbar navbar-default top-navbar" role="navigation">
        <div class="navbar-header">
            <a class="navbar-brand" href="#"><i class="fa fa-gear"></i> <strong>{{.site.title}}</strong></a>
        </div>
    </nav>
    <div id="login-wrapper">
        <div id="login_page">
            <div class="login">
                <h3 class="inner-tittle t-inner">用户登录</h3>
                <form role="form" method="post">
                    <div class="alert-danger"></div>
                    <div class="form-group input-group">
                        <span class="input-group-addon"><a class="fa fa-user" style="width: 20px"></a></span>
                        <input class="form-control" id="name" name="name" type="text" placeholder="请输入账号"
                               required/>
                    </div>
                    <div class="form-group input-group">
                        <span class="input-group-addon"><a class="fa fa-lock" style="width: 20px"></a></span>
                        <input class="form-control" id="pwd" name="pwd" type="password" placeholder="请输入密码"
                               required/>
                    </div>
                    <div class="form-group input-group">
                        <span class="input-group-addon"><a class="fa fa-eye" style="width: 20px"></a></span>
                        <input class="form-control" id="captcha" name="captcha" type="text" placeholder="请输入验证码"
                               required/>
                        <span id="cap_img" class="input-group-addon" style="padding: 0">{{create_captcha}}</span>
                    </div>
                    <div class="form-group">
                        <button class="btn btn-block btn-info">立即登录</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
    <footer>
        <p class="text-center">Copyright &copy;RenQi Technology Ltd.Co All rights reserved {{year}} </p>
    </footer>
    <script>
        $("form").submit(function (o) {
            params = $("form").serialize();
            $.post(location.href, params, function (result) {
                if (result.state !== 0) {
                    $(".alert-danger").html(result.msg);
                    $(".captcha-img").click();
                }else{
                    location.href = "./"
                }
            });
            return false
        });
    </script>
</div>
</body>
</html>