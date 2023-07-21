// echart configuration

// github.com/apache/echarts
// github.com/ecomfe/echarts-stat
// github.com/ecomfe/awesome-echarts
// echarts.apache.org/examples/en/editor.html?c=bar-histogram

// import does not work; possibly because of stackoverflow.com/questions/71022803
// import * as echarts from './echarts.min.js';
// import ecStat from './ecStat.js';



var chartDom = document.getElementById('chart_container');
// console.log(chartDom);
var myChart = echarts.init(chartDom);
var opt1;
var opt2;



// var colorPalette = ['#d87c7c', '#919e8b', '#d7ab82', '#6e7074', '#61a0a8', '#efa18d', '#787464', '#cc7e63', '#724e58', '#4b565b'];
var colorPalette = [
    '#229',
    '#22b',
    '#22c',
    '#22d',
    'var(--clr-pri-hov);',
    ];
function getColor() {
    let idx = colorPalette.length % counterDraws;
    return colorPalette[seriesIdx];
}


// histogram config
// ======================

const w  = 10;   // width
const wh = w/2;  // width half

var maxXHisto = 0;
function getMax() {
    return maxXHisto + 2;
}



// risky asset random draws
var ds1Example = {
    source: [
        [10.3, 143],
        [10.6, 214],
        [10.8, 251],
        [10.0, 176],
        [10.1, 221],
        [10.2, 188],
        [10.4, 191],
        [10.0, 196],
        [10.9, 177],
        [10.9, 153],
        [10.3, 201],
    ]
};

// risky asset histogram
var ds2Example = {
    source: [
        [ 25, 0],
        [ 75, 0],
        [125, 0],
        [175, 0],
        [225, 0],
        [275, 0],
    ]
};


// riskless asset histogram
var ds3 = {
    source: [
        [25,  0],
        [75, 10],
        [125, 0],
        [175, 0],
        [225, 0],
        [275, 0],
    ]
};


var ds1 = {
    source: []
};
var counterDrawsInit = 4 ;
var counterDraws = counterDrawsInit;  // counter for getData



// Carolin-01-start
var mn = 0.0; // mean
var sd = 1.0; // standard deviation

var mn = 153.0; // mean
var sd = 15.89; // standard deviation

// github.com/chen0040/js-stats
var normDist = new jsstats.NormalDistribution(mn, sd);


var ds2 = {
    source: []
};
for (let i = 0; i <= 300/w; i++) {
    ds2.source.push([wh + i*w, 0]);
}

// console.log(ds2.source);

// getData compiles data for eChart options object
// usage:
//       myChart.setOption({
//          dataset: getData(),
//       });
function getData() {

    counterDraws++;

    if (false) {
        try {
            // both cases lead to infinity
            console.log( "icnp(0.0):", normDist.invCumulativeProbability(0)    );
            console.log( "icnp(1.0):", normDist.invCumulativeProbability(1.0)  );
        } catch (error) {
            console.error(error);
        }
    }

    //
    // random draws - mapped to normal dist.
    for (let i = ds1.source.length; i < (counterDraws+1); i++) {
        let linDraw = Math.random(); // a number from 0 to <1

        while (linDraw == 0.0) {
            // just avoid 0.0, because it creates infinity below
            linDraw = Math.random();
        }

        let draw  = normDist.invCumulativeProbability(linDraw)
        // console.log(`   lin draw ${linDraw} => draw  ${draw}`);

        let subAr = ["draw", draw];
        ds1.source.push(subAr);
    }

    // console.log(`counterDraws ${counterDraws} - ds1a: `, ds1a.source );


    //
    // histogram data
    let i0 = ds1.source.length - 1
    if (counterDraws == counterDrawsInit+1) {
        i0 = 0;
    }
    for (let i = i0; i < ds1.source.length; i++) {
        let val    = Math.floor(ds1.source[i][1]);
        let binId  = Math.round(val/w)*w + wh;
        let binIdx = (binId - wh) / w;
        // console.log(`   val ${val} => binId ${binId} - => binIdx ${binIdx}`);
        ds2.source[binIdx][1]++;
        if (ds2.source[binIdx][1] > maxXHisto) {
            maxXHisto = ds2.source[binIdx][1];
        }

    }

    ds3.source[2] = [75, maxXHisto];

    // console.log(`counterDraws ${counterDraws} - ds2a: `, ds2a.source);

    return [
        ds1,
        ds2,
        ds3,
        {
            transform: {
                type: 'ecStat:histogram',
                // print: true,
                config: { dimensions: [1] }
            }
        },
    ];
}

opt1 = {
    dataset: getData(),
    tooltip: {},
    grid: [
        {
            top:    '04%',
            right:  '75%',
        },
        {
            top:    '04%',
            left:   '35%',
            width:  '50%',
        }
    ],
    xAxis: [
        {
            type:  'value',
            type:  'category',
            // do not include zero position
            scale:  true,
            gridIndex: 0
        },
        {
            scale: true,
            gridIndex: 1,
            inverse: true,
            min: 0,
            max: function(){
                return getMax();
            },
        },
    ],
    yAxis: [
        {
            gridIndex: 0,
            min: 0,
            max: 300,
        },
        {
            gridIndex: 1,
            // min: 0,
            // max: 350,
            // must be category: https://github.com/apache/echarts/issues/15960
            //      or https://echarts.apache.org/en/option.html#series-custom
            type: 'category',
            // axisTick:  { show: false },
            // axisLabel: { show: false },
            // axisLine:  { show: false },

            axisLine: {
                // necessary for position: right to take effect
                onZero: false,
            },
            position: 'right',
        },
        {
            gridIndex: 1,
            type: 'category',
            axisLine: {
                // necessary for position: right to take effect
                onZero: false,
            },
            position: 'right',
        }


    ],
    series: [
        {
            name: 'random draws',
            type: 'scatter',
            color: '#d87c7c',
            xAxisIndex: 0,
            yAxisIndex: 0,
            encode: { tooltip: [0, 1] },
            symbol: 'emptyCircle',
            symbol: 'circle',
            // symbolOffset only works for the entire series
            //   symbolOffset: [  -33, 10],
            //   symbolOffset: [ Math.floor((Math.random() *  44)) -22],
            symbolSize: function (value, params) {
                // console.log(`symbolSize`, params.data);
                // console.log(`symbolSize`, params.dataIndex, counterDraws);
                let a1 = params.dataIndex + 1;
                let a2 = counterDraws;
                if (a1 == a2) {
                    return 10;
                }
                params.color = '#919e8b';  // does not affect
                return 3;
                // return value;
            },
            itemStyle: {
                // borderWidth: 3,
                // borderColor: '#EE6666',

                // color function not possible
                // color: function (value, params) {
                //     return '#919e8b';
                // },
                // color: 'yellow',


                opacity: 0.4,
            },

            // color does not work as symbolSize
            // color: function (value, params) {
            //     return getColor();
            // },
            // color: getColor(),

            datasetIndex: 0,
        },
        {
            name: 'histogram',
            type: 'bar',
            color: '#d87c7c',
            xAxisIndex: 1,
            yAxisIndex: 1,
            barWidth: '99.3%',
            barWidth: '4px',
            label: {
                show: true,
                position: 'center',
                position: 'left',
                position: 'right',
                // distance to host graphic element
                distance: 55,
                offset: [10,0],
            },
            // label position - free func
            //   echarts.apache.org/en/option.html#series-bar
            labelLayout(params) {
                let fs = 12;
                if (params.rect.width < 1.0) {
                    fs = 0;
                }
                return {

                    // x:        params.rect.x + 1,
                    dx: 2,
                    y:        params.rect.y + 1,
                    fontSize: fs,
                    // not working
                    //   opacity: 0.2,
                    //   color: '#AA0101',
                };
            },
            encode: { x: 1, y: 0, itemName: 4 },
            datasetIndex: 1
        },
        {
            name: 'histogram2',
            type: 'bar',
            xAxisIndex: 1,
            yAxisIndex: 2,
            barWidth: '32px',
            encode: { x: 1, y: 0, itemName: 4 },
            datasetIndex: 2
        },

    ],
};

var dataXAxix = [];
let iStart = new Date().getFullYear()
for (let i = iStart; i <= iStart+15; i++) {
    dataXAxix.push(i);
}
console.log(dataXAxix)


var dataReturns = [];
for (let i = 0; i <= 15; i++) {
    dataReturns.push(250+i*2000);
}
console.log(dataReturns)


let seriesIdx = -1;
let animDuration = 800;

opt2 = {
    title: {
        // text: 'ECharts Getting Started Example'
        text: 'Auszahlungen',
        left: '1%'
    },
    tooltip: {},
    toolbox: {
        show: true,
        right: 10,
        feature: {
            saveAsImage: { show: true },
            // magicType:   { show: true, type: ['stack', 'tiled'] },
            dataZoom: { yAxisIndex: 'none' },
            restore: {},
        }
    },
    grid: {
        left: '12%',
        left: '13%',
        right: '3%',
        bottom: '7%',
      },    
    legend: {
        // data: ['sales']
    },
    xAxis: {
        // type: 'category',
        type: 'time',
        type: 'value',

        // only in numerical axis, i.e., type: 'value'.
        //    show zero position only, if justified by data
        //    if min and max are set, this setting has no meaning
        scale: true,

        // animation only makes sense for series
        animation: false,

        axisLabel: {
            // compare  axisLabel. formatter
            formatter: function (vl, index) {
                return vl + ' ';
            },
            textStyle: {
                color: function (vl, index) {
                    return vl >= 2030 ? 'green' : 'red';
                }
            },
        },
        // data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        // data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri']
        // data: dataXAxix,
        // min: dataXAxix[0]-2,
        // min: 2000,
        // min: 'dataMin',
        min: function (vl) {
            // this returns dataReturns.min and max
            console.log(`min ${vl.min} max ${vl.max} `)
            return vl.min;
        },

        min: iStart+0,
        max: iStart+15,

        // show label of the min/max tick
        //    seems ineffective, labels are shown anyway
        showMinLabel: false,
        showMaxLabel: false,

        // number of ticks - recommendation
        splitNumber: 8,
        // number of ticks - mandatory
        interval: 3,
        interval: 2,


        axisTick: {
            show: false,
            show: true,
            length: 4,
        },
        minorTick: {
            show: true,
        },
        // inverse: true,
        axisLine: {
            // show: false,
            // whether axis lies on the other's origin position - where value is 0 on axis.
            onZero: false,
            onZero: true,
        },

    },
    yAxis: {
        type: 'value',
        name: 'y-axis-name',
        min: 0,
        max: 40*1000,

        axisLabel: {
            // compare  axisLabel.formatter
            formatter: function (vl, index) {
                // adjust grid.left
                let vl1 = vl.toFixed(0)
                vl1 = vl1 + ' €';
                vl1 = vl1.replace("000 €", ".000 €", )
                return vl1;
            },
        },

    },
    series: [
        {
            // name - only if we want it to be shown
            // name: 'series1',
            type: 'line',
            dummy: seriesIdx++,
            color: colorPalette[seriesIdx],
            symbol: 'emptyCircle',
            symbolSize: 6,
            showSymbol: true,
            animation: false,
            animation: true,
            animationDelay:    seriesIdx * animDuration,
            animationDuration: animDuration,

            // explanation for encode: 
            //      see 10 lines below - 'data'
            //      see     https://echarts.apache.org/en/option.html#
            //      search  'series-line. encode' 
            encode: { 
                x: 0, 
                y: 1, 
                itemName: 2, 
                tooltip: [0, 1, 2],
             },
            data: [
                // [col1, col2, col3 ... ]
                // [dimX, dimY, other dimensions ...
                // In cartesian (grid), "dimX" and "dimY" correspond to xAxis and yAxis respectively.
                //    see      https://echarts.apache.org/en/option.html#series-line
                //    search   'Relationship between "value" and axis.type'
                //
                [2023,     950, 'item-1'  ],
                [2024,    2900, 'item-2'  ],
                [2025,    4400, 'item-3'  ],
                [2026,    5000, 'item-4'  ],
                [2027,    6500, 'item-5'  ],
                [2029,   13500, 'item-6'  ],
                [2029.5, 13800, 'item-7'  ],
                [2030,        , 'item-8'  ],
                [2031,   22000, 'item-9'  ],
                [2034,   24000, 'item-10'  ],
                [2036,   26000, 'item-11'  ],
                [2037,   36000, 'item-12'  ],
            ],
        },

        {
            // name - only if we want it to be shown
            // name: 'series2',
            type: 'line',
            dummy: seriesIdx++,
            color: colorPalette[seriesIdx],

            symbol: 'emptyCircle',
            symbolSize: 4,
            showSymbol: true,
            animation: false,
            animation: true,
            animationDelay:    seriesIdx * animDuration,
            animationDuration: animDuration,

            // see first series for explanation of "encode" and "data" config
            data: [
                [2023,    175],
                [2024,   2200],
                [2025,   4000],
                [2026,   4000],
                [2027,   4500],
                [2030,   4500],
                [2036,  24000],
                [2037,  33000],
            ],
        },

    ]
};

// opt1 && myChart.setOption(opt1);
opt2 && myChart.setOption(opt2);
console.log(`echart config and creation complete`)

