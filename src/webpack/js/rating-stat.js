$(function () {
    var ratingPie = [];
    for (var i = 1; i <= 5; i ++) {
        var currentRatingCount = dataset["result"]["2015"]["rating"][i];

        ratingPie.push({
            name: i + " star" + (i == 1 ? "" : "s"),
            y: currentRatingCount
        });
    }

    $('#ratingChart').highcharts({
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
            name: 'Rating',
            colorByPoint: true,
            data: ratingPie
        }]
    });
});