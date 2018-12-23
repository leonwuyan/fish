/*------------------------------------------------------
    Author : www.webthemez.com
    License: Commons Attribution 3.0
    http://creativecommons.org/licenses/by/3.0/
---------------------------------------------------------  */

(function ($) {
    var mainApp = {
        initFunction: function () {
            $(window).bind("load resize", function () {
                if ($(this).width() < 768) {
                    $('div.sidebar-collapse').addClass('collapse')
                } else {
                    $('div.sidebar-collapse').removeClass('collapse')
                }
                if ($(this).width() < 768) {
                    $('div.topbar-collapse').addClass('collapse')
                } else {
                    $('div.topbar-collapse').removeClass('collapse')
                }
            });
        }
    }
    $(document).ready(function () {
        mainApp.initFunction();
    });
}(jQuery));
