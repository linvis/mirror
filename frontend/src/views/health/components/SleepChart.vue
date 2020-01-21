<template>
  <el-card class="box-card" shadow="hover">
    <div slot="header" class="clearfix">
      <span>睡眠统计</span>
    </div>
    <!-- <el-row id="sleep_chart" /> -->
    <el-tabs @tab-click="handleTabClick">
      <!-- <el-tabs> -->
      <el-tab-pane label="睡眠时长" name="first" />
      <el-tab-pane label="入睡时间" name="second" />
      <el-tab-pane label="苏醒时间" name="third" />
    </el-tabs>
    <el-row id="sleep_chart" />
  </el-card>
</template>

<script>

import Highcharts from 'highcharts'
import stockInit from 'highcharts/modules/stock'
stockInit(Highcharts)

import { querySleepRec } from '@/api/sleep'

export default {
  data() {
    return {
      sleep_chart: null,
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
    this.featchData('month')
  },
  methods: {
    featchData(time) {
      // querySleepRec(time).then(response => {
      //   this.options = response.data
      //   this.sleep_duration_x = response.data.date
      //   this.sleep_duration_y = response.data.duration

      //   this.initChart()
      // })
      this.initChart()
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
      this.sleep_chart = Highcharts.stockChart('sleep_chart', {
        rangeSelector: {
          enabled: true,
          allButtonsEnabled: true,
          buttons: [{
            type: 'week',
            count: 1,
            text: 'Week'
          }, {
            type: 'week',
            count: 2,
            text: '2Week'
          }, {
            type: 'month',
            text: 'Month'
          }],
          buttonTheme: {
            width: 60
          },
          selected: 0
        },

        title: {
          text: 'Sleep Analysis'
        },

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

        tooltip: {
          formatter: function() {
            var hour = parseInt(this.y / 60)
            var min = parseInt(this.y % 60)
            var time = parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
            return 'Sleep about ' + time
          }
        },

        series: [{
          name: 'sleep time',
          data: this.xydata
        }
        ]
      })
    }

  }
}
</script>
