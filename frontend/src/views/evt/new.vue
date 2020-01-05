<template>
  <div class="mixin-components-container">
    <!-- <el-row> -->
    <el-row :gutter="20">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>日常活动</span>
        </div>
        <el-row>
          <el-col :span="4" class="text-center">
            <el-cascader
              v-model="evt_type"
              placeholder="活动类型"
              :options="options"
              :props="{ expandTrigger: 'hover' }"
              filterable
              clearable
              @change="handleChange"
            />
          </el-col>

          <el-col :span="4" class="text-center">
            <el-time-picker
              v-model="start_utc"
              format="HH:mm"
              placeholder="开始时间"
              @change="handleStartTimeChange"
            />
          </el-col>

          <el-col :span="4" class="text-center">
            <el-time-picker
              v-model="end_utc"
              format="HH:mm"
              placeholder="结束时间"
              @change="handleEndTimeChange"
            />
          </el-col>

          <el-col :span="4" class="text-center">
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

          <el-col :span="4" class="text-center">
            <el-input-number v-model="num" :min="1" :max="10" label="数量" @change="handleChange" />
          </el-col>
          <el-row>
            <el-row :gutter="20" style="margin-top:50px;">
              <el-col :span="4" class="text-center">
                <el-input v-model="comment" placeholder="请输入内容" />
              </el-col>

              <el-col :span="4" class="text-center">
                <el-button type="primary" @click="handleSubmit">提交</el-button>
              </el-col>
            </el-row>
          </el-row>
        </el-row>
      </el-card>
    </el-row>
  </div>
</template>

<script>

import { submit, getEvtType } from '@/api/daily_evt'
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
      duration_utc: 0

    }
  },
  created() {
    this.featchData()
    this.evt_date = Date.UTC(new Date().getFullYear(), new Date().getMonth(), new Date().getDate()) / 1000
    this.options = []
  },
  methods: {
    featchData() {
    // this.loading = true
      getEvtType().then(response => {
      // this.options = response.data.items
        console.log(response.data)
        this.options = response.data
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
        this.start_time = 24 * 60 - this.start_time
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
        this.start_time = 24 * 60 - this.start_time
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
