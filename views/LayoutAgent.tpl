<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="zh-cn">
<head>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <title>{{.site.title}}</title>
    <link href="/static/css/bootstrap.css" rel="stylesheet"/>
    <link href="/static/js/bootstrap-table/bootstrap-table.css" rel="stylesheet"/>
    <link href="/static/css/bootstrap-slider.css" rel="stylesheet"/>
    <link href="/static/css/font-awesome.css" rel="stylesheet"/>
    <link href="/static/js/morris/morris-0.4.3.min.css" rel="stylesheet"/>
    <link href="/static/css/custom-styles.css" rel="stylesheet"/>
    <script src="/static/js/clipboard.js"></script>
    <script src="/static/js/jquery-1.10.2.js"></script>
    <script src="/static/js/bootstrap.js"></script>
    <script src="/static/js/bootstrap-slider/bootstrap-slider.js"></script>
    <script src="/static/js/jquery.cookie.js"></script>
    <script src="/static/js/jquery.metisMenu.js"></script>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/toastr.js/latest/toastr.css" rel="stylesheet" />
    <script src="https://cdnjs.cloudflare.com/ajax/libs/toastr.js/latest/toastr.min.js"></script>
    <script src="/static/js/bootstrap-table/bootstrap-table.js"></script>
    <script src="/static/js/bootstrap-table/i18n/bootstrap-table-zh-cn.js"></script>
    <script src="/static/js/fish.js?"></script>
    <script>
        toastr.options.positionClass = 'toast-center-center';
    </script>
</head>
<body>
<div id="wrapper">
    <nav class="navbar navbar-default top-navbar" role="navigation">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".sidebar-collapse">
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <button class="navbar-toggle" data-toggle="collapse" data-target=".topbar-collapse"
                    style="border: 0; color: #fff;">
                <span class="fa fa-long-arrow-down"></span>
                <span class="fa fa-long-arrow-up"></span>
            </button>
            <a class="navbar-brand" href="{{.host}}"><i class="fa fa-gear"></i>
                <strong>{{.site.title}}</strong></a>
        </div>
        <div class="topbar-collapse">
            <ul class="nav navbar-top-links navbar-right" id="top-menu">
                <li class="dropdown">
                    <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                    {{.user.Name}} <i class="fa fa-user fa-fw"></i> <i class="fa fa-caret-down"></i>
                    </a>
                    <ul class="dropdown-menu dropdown-user">
                        <li><a href="{{.host}}changepwd"><i
                                class="fa fa-key fa-fw"></i>修改密码
                        </a>
                        </li>
                        <li class="divider"></li>
                        <li><a href="{{.host}}logout"><i
                                class="fa fa-sign-out fa-fw"></i>安全退出
                        </a>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
    </nav>
    <nav class="navbar-default navbar-side" role="navigation">
        <div class="sidebar-collapse">
            <ul class="nav" id="main-menu">
                <li>
                    <a href="{{.host}}"><i class="fa fa-home"></i> 概况</a>
                </li>
                <li>
                    <a href="{{.host}}agents"><i class="fa fa-user"></i> 我的代理</a>
                </li>
                <li>
                    <a href="{{.host}}generalize"><i class="fa fa-share-alt"></i> 我的推广</a>
                </li>
                <li>
                    <a href="{{.host}}players"><i class="fa fa-user-circle"></i> 我的玩家</a>
                </li>
                <li>
                    <a href="{{.host}}tax"><i class="fa fa-diamond"></i> 收入明细</a>
                </li>
                <li>
                    <a href="{{.host}}cash"><i class="fa fa-dollar"></i> 账号提现</a>
                </li>
                <li>
                    <a href="{{.host}}changepwd"><i class="fa fa-key"></i> 修改密码</a>
                </li>
            </ul>
        </div>
    </nav>

    <div id="page-wrapper">
        <div id="page-inner">
        {{.LayoutContent}}
            <footer>
                <p class="text-center">Copyright &copy;RenQi Technology Ltd.Co All rights reserved {{year}} </p>
            </footer>
        </div>
    </div>
</div>
</body>
</html>