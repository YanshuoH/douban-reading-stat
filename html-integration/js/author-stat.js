$(function () {
    var authorPie = [];
    var authorStat = dataset["result"]["2015"]["author"];
    for (var author in authorStat) {
        authorPie.push({
            name: author,
            y: authorStat[author]
        });
    }

    $('#authorChart').highcharts({
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
            formatter: function() {
                return this.key + ': <b>' + this.y + ' book</b>';
            },
            //pointFormat: '<b>{point.y}</b>',
            style: {
                fontSize: '15px'
            }
        },
        plotOptions: {
            pie: {
                allowPointSelect: true,
                cursor: 'pointer',
                dataLabels: {
                    enabled: true,
                    format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                    style: {
                        color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black',
                        fontSize: '15px'
                    }
                }
            }
        },
        series: [{
            name: 'Author',
            colorByPoint: true,
            data: authorPie
        }]
    });
});