<template>
  <div class="mixin-components-container">
    <el-row id="sleep_chart" />
  </div>
</template>

<script>
import Highcharts from 'highcharts'

import { getSleepData } from '@/api/daily_evt'

export default {
  data() {
    return {
      sleep_chart: null,
      sleep_duration_x: [],
      sleep_duration_y: []
    }
  },
  created() {
    this.featchData()
  },
  mounted() {
    // this.initChart()
  },
  methods: {
    featchData() {
    // this.loading = true
      getSleepData().then(response => {
        this.options = response.data
        this.sleep_duration_x = response.data.date
        this.sleep_duration_y = response.data.duration

        this.initChart()
      })
    },
    timeToString(time) {
      var hour = parseInt(time / 60)
      var min = parseInt(time % 60)
      return parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
    },
    initChart() {
      Highcharts.chart('sleep_chart', {

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

<style scoped>
.mixin-components-container {
  background-color: #f0f2f5;
  padding: 30px;
  min-height: calc(100vh - 84px);
}
.component-item {
  min-height: 100px;
}
</style>
