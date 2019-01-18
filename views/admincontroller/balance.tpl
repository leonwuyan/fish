<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li>概况</li>
        </ol>
    </div>
</div>
<div class="panel panel-default" style="padding: 10px">
    <form id="search-form" class="form-inline" role="form" method="post">
        <div class="form-group input-group">
            <span class="input-group-addon">开始</span>
            <input id="begin" name="begin" class="form-control" type="date"
                   value="{{thismonth}}" placeholder="" required>
        </div>
        <div class="form-group input-group">
            <span class="input-group-addon">结束</span>
            <input id="end" name="end" class="form-control" type="date"
                   value="{{today}}" placeholder="" required>
        </div>
        <button class="btn btn-info form-control" type="submit"><span
                class="fa fa-search"></span></button>
    </form>
</div>
{{if v .user .powers.结算.代理结算}}
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel panel-heading">
                代理结算
            </div>
            <div class="panel-body">
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left green">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="sysWinA"></h3>
                            <strong> 系统输赢</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left brown">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="sysFeeA"></h3>
                            <strong> 系统税收</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left red">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="agentFee"></h3>
                            <strong> 代理税收</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left blue">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="agentTax"></h3>
                            <strong> 代理收益</strong>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
{{end}}
{{if v .user .powers.结算.系统结算}}
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel panel-heading">
                系统结算
            </div>
            <div class="panel-body">
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left green">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="sysWin"></h3>
                            <strong> 系统输赢</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left brown">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="sysFee"></h3>
                            <strong> 系统税收</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left red">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="sysWinO"></h3>
                            <strong> 非代理系统输赢</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left blue">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="sysFeeO"></h3>
                            <strong> 非代理系统税收</strong>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
{{end}}
{{if v .user .powers.结算.其他结算}}
<div class="row">
    <div class="col-md-12">
        <div class="panel panel-default">
            <div class="panel panel-heading">
                其他结算
            </div>
            <div class="panel-body">
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left green">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="playerRecharge"></h3>
                            <strong> 玩家充值</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left brown">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="sysAward"></h3>
                            <strong> 系统赠送</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left red">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="playerCash"></h3>
                            <strong> 玩家提现</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left blue">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="sysPunish"></h3>
                            <strong> 系统惩罚</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left green">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="playerRechargeAward"></h3>
                            <strong> 活动赠送</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left brown">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="playerRechargeForTest"></h3>
                            <strong> 测试赠送</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left red">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="playerRemain"></h3>
                            <strong> 玩家当前余额</strong>
                        </div>
                    </div>
                </div>
                <div class="col-md-3 col-sm-6 col-xs-12">
                    <div class="panel panel-primary text-center no-boder">
                        <div class="panel-left pull-left blue">
                            <i class="fa fa-dollar fa-5x"></i>
                        </div>
                        <div class="panel-right">
                            <h3 id="dev"></h3>
                            <strong> 结算</strong>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
{{end}}
<script>
    fishApp.balance();
    $("form").submit(function () {
        fishApp.balance();
        return false
    })
</script>
