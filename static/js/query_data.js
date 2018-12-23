pagination = true;
showFooter = false;
showColumns = false;
$(function () {
    $(document).ready(function () {
        $('#dataTable').bootstrapTable({
            method: 'POST',
            contentType: "application/x-www-form-urlencoded",
            url: dataurl,
            showColumns:showColumns,
            striped: true,
            dataField: "data",
            pageNumber: 1,
            pagination: pagination,//是否分页
            queryParamsType: 'limit',
            queryParams: queryParams,
            sidePagination: 'server',
            pageSize: 15,
            pageList: [15, 30, 50, 100],
            clickToSelect: true,
            toolbar: '#toolbar',
            locale: $.cookie('lang'),
            columns: datacolumns,
            showFooter: showFooter,
        });
    });

    function queryParams(params) {
        searchParams = JSON.stringify($("#search-form").toJSON());
        return {
            pageSize: params.limit,
            pageIndex: params.pageNumber,
            searchParams: searchParams
        }
    }

    $('#search-form').submit(function () {
        $('#dataTable').bootstrapTable('refresh', {url: dataurl});
        return false;
    });
});