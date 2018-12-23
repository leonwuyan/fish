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
                                class="fa fa-key fa-fw"></i> 修改密码
                        </a>
                        </li>
                        <li class="divider"></li>
                        <li><a href="{{.host}}logout"><i
                                class="fa fa-sign-out fa-fw"></i> 安全退出
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
                    <a href="{{.host}}"><i class="fa fa-home"></i> 首页</a>
                </li>
            {{if v .user .powers.Admin.AdminAccount_List}}
                <li>
                    <a href="{{.host}}account"><i class="fa fa-user-o"></i> 账号管理</a>
                </li>
            {{end}}
            {{if v .user .powers.Player.Player_List}}
                <li>
                    <a href="#"><i class="fa fa-user"></i> 玩家管理<span
                            class="fa arrow"></span></a>
                    <ul class="nav nav-second-level">
                        <li>
                            <a href="{{.host}}player/list">玩家列表</a>
                        </li>
                        <li>
                            <a href="{{.host}}player/cash">玩家提现</a>
                        </li>
                    </ul>
                </li>
            {{end}}
            {{if v .user .powers.Agent.Agent_List}}
                <li>
                    <a href="#"><i class="fa fa-user-circle"></i> 代理管理<span
                            class="fa arrow"></span></a>
                    <ul class="nav nav-second-level">
                        <li>
                            <a href="{{.host}}agent/list">代理列表</a>
                        </li>
                    {{if v .user .powers.Agent.Agent_Apply_List}}
                        <li>
                            <a href="{{.host}}agent/apply">代理申请</a>
                        </li>
                    {{end}}
                    {{if v .user .powers.Agent.Agent_Cash_List}}
                        <li>
                            <a href="{{.host}}agent/cash">佣金提取</a>
                        </li>
                    {{end}}
                    </ul>
                </li>
            {{end}}
            {{if v .user .powers.Recharge.Recharge}}
                <li>
                    <a href="{{.host}}recharge"><i class="fa fa-dollar"></i> 后台充值</a>
                </li>
            {{end}}
            {{if v .user .powers.System.SystemCfg}}
                <li>
                    <a href="{{.host}}sysconfig"><i class="fa fa-cog"></i> 系统配置</a>
                </li>
            {{end}}
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