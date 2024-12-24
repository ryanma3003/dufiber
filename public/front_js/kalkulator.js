function handleUtang()
{
    var maalEmas = parseInt($('#maalEmas').val() || 0);
    var maalUang = parseInt($('#maalUang').val() || 0);
    var maalAset = parseInt($('#maalAset').val() || 0);
    var logamEmas = parseInt($('#logamEmas').val() || 0);
    var logamPerak = parseInt($('#logamPerak').val() || 0);
    var logamLain = parseInt($('#logamLain').val() || 0);
    var niagaAset = parseInt($('#niagaAset').val() || 0);
    var niagaLaba = parseInt($('#niagaLaba').val() || 0);
    var utang = parseInt($('#utang').val() || 0);

    var zakat = (((maalEmas+maalUang+maalAset+logamEmas+logamPerak+logamLain+niagaAset+niagaLaba) - utang) * 2.5) / 100;
    $('#nilaiZakat').val(zakat);
}

function handleZakatMaal()
{
    var maalEmas = parseInt($('#maalEmas').val());
    var maalUang = parseInt($('#maalUang').val());
    var maalAset = parseInt($('#maalAset').val());

    var zakat = ((maalEmas+maalUang+maalAset) * 2.5) / 100;
    $('#nilaiZakat').val(zakat);
}

function handleZakatLogam()
{
    var logamEmas = parseInt($('#logamEmas').val());
    var logamPerak = parseInt($('#logamPerak').val());
    var logamLain = parseInt($('#logamLain').val());

    var zakat = ((logamEmas+logamPerak+logamLain) * 2.5) / 100;
    $('#nilaiZakat').val(zakat);
}

function handleZakatNiaga()
{
    var niagaAset = parseInt($('#niagaAset').val());
    var niagaLaba = parseInt($('#niagaLaba').val());

    var zakat = ((niagaAset+niagaLaba) * 2.5) / 100;
    $('#nilaiZakat').val(zakat);
}

$(document).ready(function () {
    $('#jenisZakat').on('change', function() {
        var jns = $(this).val();

        if(jns == 1)
        {
            $('#1').show().find("input").attr("required", true);
            $('#2').hide().find("input").attr("required", false);
            $('#3').hide().find("input").attr("required", false);
            $('#4').hide().find("input").attr("required", false);
            $('#5').hide().find("input").attr("required", false);
            $('#6').hide().find("input").attr("required", false);
            $('#7').hide().find("input").attr("required", false);
            $('#8').hide().find("input").attr("required", false);
            $('#9').hide().find("input").attr("required", false);
            $('#10').hide().find("input").attr("required", false);
            $('#11').hide().find("input").attr("required", false);
            $('#12').hide().find("input").attr("required", false);

            $('#penghasilan').val(0);
            $('#nilaiZakat').val(0);
        }
        else if(jns == 2)
        {
            $('#1').hide().find("input").attr("required", false);
            $('#2').show().find("input").attr("required", true);
            $('#3').show().find("input").attr("required", true);
            $('#4').show().find("input").attr("required", true);
            $('#5').hide().find("input").attr("required", false);
            $('#6').hide().find("input").attr("required", false);
            $('#7').hide().find("input").attr("required", false);
            $('#8').hide().find("input").attr("required", false);
            $('#9').hide().find("input").attr("required", false);
            $('#10').show().find("input").attr("required", true);
            $('#11').hide().find("input").attr("required", false);
            $('#12').hide().find("input").attr("required", false);

            $('#maalEmas').val(0);
            $('#maalUang').val(0);
            $('#maalAset').val(0);
            $('#utang').val(0);
            $('#nilaiZakat').val(0);
        }
        else if(jns == 3)
        {
            $('#1').hide().find("input").attr("required", false);
            $('#2').hide().find("input").attr("required", false);
            $('#3').hide().find("input").attr("required", false);
            $('#4').hide().find("input").attr("required", false);
            $('#5').hide().find("input").attr("required", false);
            $('#6').hide().find("input").attr("required", false);
            $('#7').hide().find("input").attr("required", false);
            $('#8').hide().find("input").attr("required", false);
            $('#9').hide().find("input").attr("required", false);
            $('#10').hide().find("input").attr("required", false);
            $('#11').show().find("input").attr("required", true);
            $('#12').show().find("input").attr("required", true);

            $('#orang').val(0);
            $('#nilaiZakat').val(0);
        }
        else if(jns == 4)
        {
            $('#1').hide().find("input").attr("required", false);
            $('#2').hide().find("input").attr("required", false);
            $('#3').hide().find("input").attr("required", false);
            $('#4').hide().find("input").attr("required", false);
            $('#5').show().find("input").attr("required", true);
            $('#6').show().find("input").attr("required", true);
            $('#7').show().find("input").attr("required", true);
            $('#8').hide().find("input").attr("required", false);
            $('#9').hide().find("input").attr("required", false);
            $('#10').show().find("input").attr("required", true);
            $('#11').hide().find("input").attr("required", false);
            $('#12').hide().find("input").attr("required", false);

            $('#logamEmas').val(0);
            $('#logamPerak').val(0);
            $('#logamLain').val(0);
            $('#utang').val(0);
            $('#nilaiZakat').val(0);
        }
        else if(jns == 5)
        {
            $('#1').hide().find("input").attr("required", false);
            $('#2').hide().find("input").attr("required", false);
            $('#3').hide().find("input").attr("required", false);
            $('#4').hide().find("input").attr("required", false);
            $('#5').hide().find("input").attr("required", false);
            $('#6').hide().find("input").attr("required", false);
            $('#7').hide().find("input").attr("required", false);
            $('#8').show().find("input").attr("required", true);
            $('#9').show().find("input").attr("required", true);
            $('#10').show().find("input").attr("required", true);
            $('#11').hide().find("input").attr("required", false);
            $('#12').hide().find("input").attr("required", false);

            $('#niagaAset').val(0);
            $('#niagaLaba').val(0);
            $('#utang').val(0);
            $('#nilaiZakat').val(0);
        }
    });

    $('#jenisWakaf').on('change', function() {
        var wkf = $(this).val();
        if(wkf == 13) {
            $('#hargawakafpohon').show().find("select").attr("required", true);
            $('#jumlahpohon').show().find("input").attr("required", true);

            $('#pohon').val(0);
            $('#nominalPohon').val(0);
        }
    });

    $('#hariFidyah').on('keyup change input', function() {
        var nominalFidyah = parseInt($(this).val() * fidyah);
        $('#nominalFidyah').val(nominalFidyah);
    });

    $('#penghasilan').on('keyup change input', function() {
        var penghasilan = parseInt(($(this).val() * 2.5) / 100);
        $('#nilaiZakat').val(penghasilan);
    });

    $('#orang').on('keyup change input', function() {
        var fitrah = parseInt($(this).val() * $('#nominalZakatFitrah').val());
        $('#nilaiZakat').val(fitrah);
    });

    $('#nominalZakatFitrah').on('change', function() {
        var fitrah = parseInt($(this).val() * $('#orang').val());
        $('#nilaiZakat').val(fitrah);
    });

    $('#pohon').on('keyup change input', function() {
        var nominalPohon = parseInt($(this).val() * $('#wakafPohon').val());
        $('#nominalPohon').val(nominalPohon);
    });

    $('#wakafPohon').on('change', function() {
        var nominalPohon = parseInt($(this).val() * $('#pohon').val());
        $('#nominalPohon').val(nominalPohon);
    });
});
