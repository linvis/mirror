<template>
  <el-card class="box-card" shadow="hover">
    <div slot="header" class="clearfix">
      <span>睡眠时长</span>
    </div>
    <!-- <el-row id="sleep_chart" /> -->
    <el-tabs @tab-click="handleTabClick">
      <!-- <el-tabs> -->
      <el-tab-pane label="最近1周" name="first" />
      <el-tab-pane label="最近2周" name="second" />
      <el-tab-pane label="最近1月" name="third" />
    </el-tabs>
    <el-row id="sleep_chart" />
  </el-card>
</template>

<script>

import Highcharts from 'highcharts'
import { querySleepRec } from '@/api/sleep'

export default {
  data() {
    return {
      sleep_chart: null,
      sleep_duration_x: [],
      sleep_duration_y: [],
      sleep_x: [],
      sleep_y: [],
      sleep_record_week: {
        'time': [],
        'start': [],
        'end': []
      },
      sleep_record_twoweek: {
        'time': [],
        'start': [],
        'end': []
      }
    }
  },
  created() {
    this.featchData('week')
  },
  methods: {
    featchData(time) {
      querySleepRec(time).then(response => {
        this.options = response.data
        this.sleep_duration_x = response.data.date
        this.sleep_duration_y = response.data.duration

        this.initChart()
      })
    },
    handleTabClick(tab, event) {
      console.log(tab.name)
      var time = ''
      if (tab.name === 'first') {
        time = 'week'
      } else if (tab.name === 'second') {
        time = 'twoweek'
      } else if (tab.name === 'third') {
        time = 'month'
      }
      querySleepRec(time).then(response => {
        this.options = response.data
        this.sleep_duration_x = response.data.date
        this.sleep_duration_y = response.data.duration

        this.sleep_chart.series[0].update({
          x: this.sleep_duration_x,
          data: this.sleep_duration_y
        }, false)
        this.sleep_chart.redraw()
      })
    },
    timeToString(time) {
      var hour = parseInt(time / 60)
      var min = parseInt(time % 60)
      return parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
    },
    initChart() {
      this.sleep_chart = Highcharts.chart('sleep_chart', {

        title: {
          text: 'Sleep Analysis'
        },

        // subtitle: {
        //   text: 'Source: thesolarfoundation.com'
        // },

        yAxis: {
          labels: {
            formatter: function() {
            //   return this.timeToString(this.value)
              var hour = parseInt(this.value / 60)
              var min = parseInt(this.value % 60)
              return parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
            }
          },
          title: {
            text: 'Hours'
          }
        },
        legend: {
          layout: 'vertical',
          align: 'right',
          verticalAlign: 'middle'
        },

        tooltip: {
          formatter: function() {
            var hour = parseInt(this.y / 60)
            var min = parseInt(this.y % 60)
            var time = parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
            return 'Sleep about ' + time
          }
        },

        xAxis: {
          categories: this.sleep_duration_x
        //   categories: ['2019-12-31', '2020-1-1', '2020-1-2', '2020-1-3', '2020-1-4', '2020-1-5']
        },

        series: [{
          name: 'sleep time',
          data: this.sleep_duration_y
        //   data: [540, 478, 464, 449, 531, 502]
        }],

        responsive: {
          rules: [{
            condition: {
              maxWidth: 500
            },
            chartOptions: {
              legend: {
                layout: 'horizontal',
                align: 'center',
                verticalAlign: 'bottom'
              }
            }
          }]
        }

      })
    }

  }
}
</script>
