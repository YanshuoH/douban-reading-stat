$(function () {
    var publisherPie = [];
    var publisherStat = dataset["result"]["2015"]["publisher"];
    for (var publisher in publisherStat) {
        publisherPie.push({
            name: publisher,
            y: publisherStat[publisher]
        });
    }

    $('#publisherChart').highcharts({
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
            name: 'Publisher',
            colorByPoint: true,
            data: publisherPie
        }]
    });
});