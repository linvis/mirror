<template>
  <el-card class="box-card" shadow="hover">
    <div slot="header" class="clearfix">
      <span>睡眠分析</span>
    </div>
    <el-row id="sleep_analysis_chart" />
  </el-card>
</template>

<script>

import Highcharts from 'highcharts'
import stockInit from 'highcharts/modules/stock'
stockInit(Highcharts)

export default {
  data() {
    return {
      sleep_analysis_chart: null,
      sleep_duration_x: [],
      sleep_duration_y: [],
      sleep_x: [],
      sleep_y: [],
      xydata: [
        [
          1516631400000,
          177
        ],
        [
          1516717800000,
          177
        ],
        [
          1516804200000,
          174
        ],
        [
          1516890600000,
          171
        ],
        [
          1516977000000,
          171
        ],
        [
          1517236200000,
          167
        ],
        [
          1517322600000,
          166
        ],
        [
          1517409000000,
          167
        ],
        [
          1517495400000,
          167
        ],
        [
          1517581800000,
          160
        ],
        [
          1517841000000,
          156
        ],
        [
          1517927400000,
          163
        ]]
    }
  },
  created() {
  },
  mounted() {
    this.featchData()
  },
  methods: {
    featchData() {
      // querySleepRec().then(response => {
      //   this.xydata = response.data

      //   this.initChart()
      // })
      this.initChart()
    },
    timeToString(time) {
      var hour = parseInt(time / 60)
      var min = parseInt(time % 60)
      return parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
    },
    initChart() {
      this.sleep_analysis_chart = Highcharts.chart('sleep_analysis_chart', {
        chart: {
          type: 'bar'
        },
        title: {
          text: 'Sleep Analysis'
        },
        xAxis: {
          categories: ['平均睡眠时长', '平均入睡时间', '平均苏醒时间'],
          title: {
            text: null
          }
        },
        yAxis: {
          min: 0,
          title: {
            text: 'Population (millions)',
            align: 'high'
          },
          labels: {
            overflow: 'justify'
          }
        },
        tooltip: {
          valueSuffix: ' millions'
        },
        plotOptions: {
          bar: {
            dataLabels: {
              enabled: true
            }
          }
        },
        legend: {
          layout: 'vertical',
          align: 'right',
          verticalAlign: 'top',
          x: -40,
          y: 80,
          floating: true,
          borderWidth: 1,
          backgroundColor:
            Highcharts.defaultOptions.legend.backgroundColor || '#FFFFFF',
          shadow: true
        },
        credits: {
          enabled: false
        },
        series: [{
          name: 'This Week',
          data: [107, 31, 635]
        }, {
          name: 'Last Week',
          data: [133, 156, 947]
        }]
      })
    }
  }
}
</script>
