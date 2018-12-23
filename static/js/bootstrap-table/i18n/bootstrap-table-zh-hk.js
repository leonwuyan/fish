!function (a) {
    "use strict";
    a.fn.bootstrapTable.locales["zh-hk"] = {
        formatLoadingMessage: function () {
            return "正在努力的加載數據中，請稍後……"
        }, formatRecordsPerPage: function (a) {
            return "每頁顯示 " + a + " 條記錄"
        }, formatShowingRows: function (a, b, c) {
            return "顯示第 " + a + " 到第 " + b + " 條記錄，總共 " + c + " 條記錄"
        }, formatSearch: function () {
            return "搜索"
        }, formatNoMatches: function () {
            return "沒有找到匹配的記錄"
        }, formatPaginationSwitch: function () {
            return "隱藏/顯示分頁"
        }, formatRefresh: function () {
            return "刷新"
        }, formatToggle: function () {
            return "切換"
        }, formatColumns: function () {
            return "列"
        }, formatExport: function () {
            return "導出數據"
        }, formatClearFilters: function () {
            return "清空過濾"
        }
    }, a.extend(a.fn.bootstrapTable.defaults, a.fn.bootstrapTable.locales["zh-hk"])
}(jQuery);