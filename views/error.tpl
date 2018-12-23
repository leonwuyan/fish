<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <title>错误信息</title>
    <link href="/static/css/bootstrap.css" rel="stylesheet"/>
    <link href="/static/css/custom-styles.css" rel="stylesheet"/>
</head>
<style>
    #error-wrapper {
        text-align: center;
        margin: auto;
    }

    #error_page {
        padding-top: 100px;
        max-width: 680px;
        margin: auto;
        text-align: center;
    }

    .login {
        background-color: #fff;
        border-radius: 10px;
        width: 90%;
        margin: auto;
        padding: 10px;
        text-align: center;
    }
</style>
<body>
<div id="wrapper">
    <div id="error-wrapper">
        <div id="error_page">
            <div class="login">
                <h3 class="inner-tittle t-inner">错误码：{{.code}}</h3>
                <br/>
                <p class="sorry">{{.content}}</p>
                <div class="error-btn">
                    <a class="read fourth" href="javascript:history.back()">返回</a>
                </div>
            </div>
        </div>
    </div>
    <footer>
        <p class="text-center">Copyright &copy; Technology Ltd.Co All rights reserved {{year}} </p>
    </footer>
</div>
</body>
</html>