$(function () {
    var translatorPie = [];
    var translatorStat = dataset["result"]["2015"]["translator"];
    for (var translator in translatorStat) {
        translatorPie.push({
            name: translator,
            y: translatorStat[translator]
        });
    }

    $('#translatorChart').highcharts({
        chart: {
            plotBackgroundColor: null,
            plotBorderWidth: null,
            plotShadow: false,
            type: 'pie'
        },
        title: {
            text: ''
        },
        tooltip: {
            pointFormat: '{series.name}: <b>{point.y}</b>'
        },
        plotOptions: {
            pie: {
                allowPointSelect: true,
                cursor: 'pointer',
                dataLabels: {
                    enabled: true,
                    format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                    style: {
                        color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                    }
                }
            }
        },
        series: [{
            name: 'Translator',
            colorByPoint: true,
            data: translatorPie
        }]
    });
});