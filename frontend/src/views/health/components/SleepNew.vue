<template>
  <el-card class="box-card" shadow="hover">
    <div slot="header" class="clearfix">
      <span>日常活动</span>
    </div>
    <el-form ref="form" label-position="left" :model="form" label-width="80px">
      <el-form-item label="日期">
        <el-date-picker
          v-model="record_date"
          type="date"
          placeholder="选择日期"
          :picker-options="pickerOptions"
        />
      </el-form-item>
      <el-form-item label="开始时间">
        <el-time-picker
          v-model="start_utc"
          format="HH:mm"
          placeholder="开始时间"
          @change="handleStartTimeChange"
        />
      </el-form-item>
      <el-form-item label="结束时间">
        <el-time-picker
          v-model="end_utc"
          format="HH:mm"
          placeholder="结束时间"
          @change="handleEndTimeChange"
        />
      </el-form-item>
      <el-form-item label="持续时间">
        <el-time-select
          v-model="duration_utc"
          :picker-options="{
            start: '00:00',
            step: '00:10',
            end: '04:00'
          }"
          placeholder="持续时间"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleSubmit">立即创建</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script>

import { submitSleepRec } from '@/api/sleep'
// import { parse } from 'path'

export default {
  data() {
    return {
      start_time: 0,
      end_time: 0,
      duration: 0,
      start_utc: 0,
      end_utc: 0,
      duration_utc: 0,
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now()
        }
      },
      record_date: new Date(),
      form: {

      }
    }
  },
  created() {
    // this.evt_date = Date.UTC(new Date().getFullYear(), new Date().getMonth(), new Date().getDate()) / 1000
    // this.evt_date = this.evt_date + (new Date().getTimezoneOffset() * 60)
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
      // var utcDate = Date.UTC(this.record_date.getFullYear(), this.record_date.getMonth(), this.record_date.getDate()) +
      // this.record_date.getTimezoneOffset() * 60 * 1000
      var utcDate = Date.UTC(this.record_date.getFullYear(), this.record_date.getMonth(), this.record_date.getDate())
      console.log(this.record_date.getTimezoneOffset())
      var record = {
        'record_date': utcDate,
        'start_time': this.start_time,
        'end_time': this.end_time,
        'duration': this.duration
      }
      submitSleepRec(record).then(response => {
        this.$notify({
          title: '成功',
          message: '',
          type: 'success',
          duration: 800
        })
        this.$bus.emit('update-sleeprecord')
      })
    }
  }
}
</script>

<style scoped>
.el-col {
  border-radius: 4px;
}
</style>
