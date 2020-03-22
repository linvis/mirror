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
import { bus } from '@/utils/bus'

export default {
  data() {
    return {
      activeTab: 'first',
      sleep_chart: null,
      dur_data: [],
      start_time_data: [],
      end_time_data: [],
      chart_data: [],
      chart_label: 'sleep about'
    }
  },
  created() {
    bus.$on('update-sleep-record', this.getSleepRecord)
  },
  mounted() {
    this.getSleepRecord()
  },
  beforeDestroy() {
    bus.$off('update-sleep-record', this.getSleepRecord)
  },
  methods: {
    getSleepRecord() {
      querySleepRec().then(response => {
        console.log(response.data)

        this.dur_data = []
        this.start_time_data = []
        this.end_time_data = []

        for (var i = 0; i < response.data.date.length; i++) {
          this.dur_data.push([response.data.date[i], response.data.duration[i]])
          this.start_time_data.push([response.data.date[i], response.data.start_time[i] + 24 * 60])
          this.end_time_data.push([response.data.date[i], response.data.end_time[i] + 24 * 60])
        }

        // this.chart_data = this.dur_data.slice(0, this.dur_data.length)
        this.chart_data = this.dur_data
        console.log(this.chart_data)

        this.initChart()
      })
      // this.initChart()
    },
    handleTabClick(tab, event) {
      console.log(tab.name, this.activeTab)
      if (this.activeTab === tab.name) {
        // do nothing
        return
      } else if (tab.name === 'first') {
        this.chart_data = this.dur_data
        this.chart_label = 'sleep about'
      } else if (tab.name === 'second') {
        this.chart_data = this.start_time_data
        this.chart_label = 'goto sleep at'
      } else if (tab.name === 'third') {
        this.chart_data = this.end_time_data
        this.chart_label = 'wakeup at'
      }

      this.$log.debug(this.chart_data)

      this.activeTab = tab.name
      this.sleep_chart.series[0].update({
        data: this.chart_data
      }, false)
      this.sleep_chart.redraw()
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
            text: 'last 7 days'
          }, {
            type: 'week',
            count: 2,
            text: 'last 14 days'
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
              this.value = this.value % (24 * 60)
              var hour = parseInt(this.value / 60)
              var min = parseInt(this.value % 60)
              return parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
            }
          },
          title: {
            text: 'Hours'
          },
          opposite: false
        },

        tooltip: {
          formatter: function() {
            // yesterday
            var tmpY = this.y % (24 * 60)
            var hour = parseInt(tmpY / 60)
            var min = parseInt(tmpY % 60)
            var time = parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
            return 'sleep about ' + time
          }
        },

        series: [{
          name: 'sleep time',
          data: this.chart_data
        }
        ]
      })
    }

  }
}
</script>
