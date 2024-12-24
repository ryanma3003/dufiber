
$(document).on("change", '#jenisDonasi', function () {
    var id = $("#jenisDonasi").val();
    $.ajax({
        url: $(this).attr('data-url'),
        method: "POST",
        data: { id: id, _token: $('input[name=_token]').val() },
        dataType: "json",
        success: function (data) {
            if (data != '') {
                // console.log(data.html);
                $('#khususDonasi').html(data.html);
            }
        }
    });
});