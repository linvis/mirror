<template>
  <el-card class="box-card" shadow="hover">
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
  </el-card>
</template>

<script>

import { submitSleepRec } from '@/api/sleep'
// import { parse } from 'path'

export default {
  data() {
    return {
      evt_date: 0,
      start_time: 0,
      end_time: 0,
      duration: 0,
      start_utc: 0,
      end_utc: 0,
      duration_utc: 0
    }
  },
  created() {
    this.evt_date = Date.UTC(new Date().getFullYear(), new Date().getMonth(), new Date().getDate()) / 1000
    this.evt_date = this.evt_date + (new Date().getTimezoneOffset() * 60)
  },
  methods: {
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
      var record = {
        'record_date': this.evt_date,
        'start_time': this.start_time,
        'end_time': this.end_time,
        'duration': this.duration
      }
      submitSleepRec(record).then(response => {
        alert('submit OK')
      })
    }
  }
}
</script>
