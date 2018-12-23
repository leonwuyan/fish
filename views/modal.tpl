<form id="data-form" action="{{.url}}">
    <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
        <h4 class="modal-title">{{.title}}</h4>
    </div>
    <div class="modal-body">
        <div id="data-err" class="alert-danger"></div>
    {{.form}}
    </div>
    <div class="modal-footer">
        <div class="form-group">
            <button type="button" class="btn btn-default"
                    data-dismiss="modal">关闭
            </button>
            <button type="submit"
                    class="btn btn-primary">提交
            </button>
        </div>
    </div>
</form>
<script>
    $("#data-form").submit(function () {
        url = $(this).attr("action");
        params = $(this).serialize();
        fishApp.putAction(url, params);
        return false
    });
    $('#data-modal').on('hidden.bs.modal', function () {
        $(this).removeData("bs.modal");
    });
    {{if .hasSlider}}
    $('input[type="slider"]').slider().on('change', function (e) {
        if (e.value.newValue < $(this).data('slider-can-min')) {
            $(this).slider('setValue', $(this).data('slider-can-min'))
        }
        $("#" + this["id"] + "_val").html($(this).slider('getValue'))
    });
    {{end}}
</script>