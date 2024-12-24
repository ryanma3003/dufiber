
$(document).ready(function () {
    var agree = $("#agreement");

    agree.click(function () {
        if ($(this).is(":checked")) {
            $("#submit").removeAttr("disabled");
        } else {
            $("#submit").attr("disabled", "disabled");
        }
    });
}); 