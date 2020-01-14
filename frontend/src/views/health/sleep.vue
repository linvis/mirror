<template>
  <div class="mixin-components-container">
    <!-- <el-row> -->
    <el-row :gutter="10">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>日常活动</span>
          <!-- <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button> -->
          <el-button style="float: right;" type="primary" @click="handleSubmit">提交</el-button>
        </div>
        <el-row style="margin-top:10px;">
          <el-col :span="4">
            <el-time-picker
              v-model="start_utc"
              format="HH:mm"
              placeholder="开始时间"
              @change="handleStartTimeChange"
            />
          </el-col>

          <el-col :span="4" :offset="4">
            <el-time-picker
              v-model="end_utc"
              format="HH:mm"
              placeholder="结束时间"
              @change="handleEndTimeChange"
            />
          </el-col>

          <el-col :span="4" :offset="4">
            <el-time-select
              v-model="duration_utc"
              :picker-options="{
                start: '00:00',
                step: '00:10',
                end: '04:00'
              }"
              placeholder="持续时间"
            />
          </el-col>
        </el-row>
        <!-- </div> -->
      </el-card>
    </el-row>
    <el-row id="sleep_chart" />
  </div>
</template>

<script>

import Highcharts from 'highcharts'
import { submit, getEvtType } from '@/api/daily_evt'
import { getSleepData } from '@/api/daily_evt'
// import { parse } from 'path'

export default {
  data() {
    return {
      evt_type: [],
      options: [],
      evt_date: 0,
      start_time: 0,
      end_time: 0,
      duration: 0,
      num: 1,
      comment: '',

      start_utc: 0,
      end_utc: 0,
      duration_utc: 0,

      sleep_chart: null,
      sleep_duration_x: [],
      sleep_duration_y: []
    }
  },
  created() {
    this.featchData()
    this.evt_date = Date.UTC(new Date().getFullYear(), new Date().getMonth(), new Date().getDate()) / 1000
    this.evt_date = this.evt_date + (new Date().getTimezoneOffset() * 60)
    this.options = []
  },
  methods: {
    featchData() {
      getEvtType().then(response => {
      // this.options = response.data.items
        console.log(response.data)
        this.options = response.data
      })

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
    },

    handleChange(type) {
      // console.log(type)
    },
    handleStartTimeChange(time) {
      this.start_time = this.start_utc.getHours() * 60 + this.start_utc.getMinutes()

      if (this.end_time === 0) {
        return
      }

      var timeToString = function(time) {
        var hour = parseInt(time / 60)
        var min = parseInt(time % 60)
        return parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
      }

      if (this.start_time > this.end_time) {
        this.start_time = this.start_time - 24 * 60
      }

      this.duration = this.end_time - this.start_time
      this.duration_utc = timeToString(this.duration)
    },
    handleEndTimeChange(time) {
      this.end_time = this.end_utc.getHours() * 60 + this.end_utc.getMinutes()

      var timeToString = function(time) {
        var hour = parseInt(time / 60)
        var min = parseInt(time % 60)
        return parseInt(hour / 10).toString() + (hour % 10).toString() + ':' + parseInt(min / 10).toString() + (min % 10).toString()
      }

      if (this.start_time === 0) {
        return
      }

      if (this.start_time > this.end_time) {
        this.start_time = this.start_time - 24 * 60
      }

      this.duration = this.end_time - this.start_time
      this.duration_utc = timeToString(this.duration)
    },
    handleSubmit: function(event) {
      console.log(this.value)
      var evt = {
        'evt_type': this.evt_type[0],
        'evt_item': this.evt_type[1],
        'evt_date': this.evt_date,
        'start_time': this.start_time,
        'end_time': this.end_time,
        'duration': this.duration,
        'num': this.num,
        'comment': this.comment
      }
      submit(evt).then(response => {
        // alert('submit OK')
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
