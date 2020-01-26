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
import { querySleepRecAnalysis } from '@/api/sleep'
stockInit(Highcharts)

export default {
  data() {
    return {
      sleep_analysis_chart: null,
      avgLastWeek: null,
      avgLastTwoWeek: null,
      avgLastMonth: null
    }
  },
  created() {
  },
  mounted() {
    this.featchDataj()
  },
  methods: {
    featchDataj() {
      querySleepRecAnalysis().then(response => {
        this.avgLastWeek = [response.data.duration[0], response.data.start_time[0], response.data.end_time[0]]
        this.avgLastTwoWeek = [response.data.duration[1], response.data.start_time[1], response.data.end_time[1]]
        this.avgLastMonth = [response.data.duration[2], response.data.start_time[2], response.data.end_time[2]]
        console.log(this.avgLastWeek)
        console.log(this.avgLastTwoWeek)
        console.log(this.avgLastMonth)
        this.initChart()
      })
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
          labels: {
            formatter: function() {
              this.value = this.value % (24 * 60)
              var hour = parseInt(this.value / 60)
              var min = parseInt(this.value % 60)
              return parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
            }
          },
          title: {
            text: 'Hours'
          }
        },

        tooltip: {
          formatter: function() {
            // yesterday
            var tmpY = this.y % (24 * 60)
            var hour = parseInt(tmpY / 60)
            var min = parseInt(tmpY % 60)
            var time = parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
            return 'Sleep about ' + time
          }
        },
        // yAxis: {
        //   min: 0,
        //   title: {
        //     text: 'Population (millions)',
        //     align: 'high'
        //   },
        //   labels: {
        //     overflow: 'justify'
        //   }
        // },
        // tooltip: {
        //   valueSuffix: ' millions'
        // },
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
          data: this.avgLastWeek
        }, {
          name: 'Last Two Week',
          data: this.avgLastTwoWeek
        }, {
          name: 'Last Month',
          data: this.avgLastMonth
        }]
      })
    }
  }
}
</script>
