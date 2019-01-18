<div class="row">
    <div class="col-md-12">
        <ol class="breadcrumb">
            <li>概况</li>
        </ol>
    </div>
</div>
<div class="row">
{{if v .user .powers.概况.今日注册}}
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel panel-primary text-center no-boder">
            <div class="panel-left pull-left green">
                <i class="fa fa-user fa-5x"></i>
            </div>
            <div class="panel-right">
                <h3>{{.player_reg_count}}</h3>
                <strong> 今日注册</strong>
            </div>
        </div>
    </div>
{{end}}
{{if v .user .powers.概况.今日活跃}}
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel panel-primary text-center no-boder">
            <div class="panel-left pull-left brown">
                <i class="fa fa-user fa-5x"></i>
            </div>
            <div class="panel-right">
                <h3>{{.player_login_count}}</h3>
                <strong> 今日活跃</strong>
            </div>
        </div>
    </div>
{{end}}
{{if v .user .powers.概况.今日有效}}
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel panel-primary text-center no-boder">
            <div class="panel-left pull-left red">
                <i class="fa fa-user fa-5x"></i>
            </div>
            <div class="panel-right">
                <h3>{{.player_play_count}}</h3>
                <strong> 今日有效</strong>
            </div>
        </div>
    </div>
{{end}}
{{if v .user .powers.概况.今日充值}}
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel panel-primary text-center no-boder">
            <div class="panel-left pull-left green">
                <i class="fa fa-dollar fa-5x"></i>
            </div>
            <div class="panel-right">
                <h3>{{.player_pay_gold}}</h3><span>
                <strong> 今日充值</strong>({{.player_pay_count}}人</span>)
            </div>
        </div>
    </div>
{{end}}
{{if v .user .powers.概况.今日税收}}
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel panel-primary text-center no-boder">
            <div class="panel-left pull-left brown">
                <i class="fa fa-dollar fa-5x"></i>
            </div>
            <div class="panel-right">
                <h3>{{.player_tax_gold}}</h3><span>
                <strong> 今日税收</strong>({{.player_tax_count}}人</span>)
            </div>
        </div>
    </div>
{{end}}
{{if v .user .powers.概况.玩家提现}}
    <div class="col-md-4 col-sm-12 col-xs-12">
        <div class="panel panel-primary text-center no-boder">
            <div class="panel-left pull-left red">
                <i class="fa fa-dollar fa-5x"></i>
            </div>
            <div class="panel-right">
                <h3>{{.player_cash_gold}}</h3><span>
                <strong> 今日玩家提现</strong>({{.player_cash_count}}人</span>)
            </div>
        </div>
    </div>
{{end}}
</div>
