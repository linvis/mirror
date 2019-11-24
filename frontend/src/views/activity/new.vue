<template>
  <div class="mixin-components-container">
    <!-- <el-row> -->
    <el-row :gutter="20">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>日常活动</span>
        </div>
        <div style="margin-bottom:80px;">
          <el-col :span="4" class="text-center">
            <el-select v-model="type" placeholder="活动类型">
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-col>

          <el-col :span="4" class="text-center">
            <el-time-picker
              v-model="start_time"
              format="HH:mm"
              value-format="timestamp"
              placeholder="开始时间"
            />
          </el-col>

          <el-col :span="4" class="text-center">
            <el-time-picker
              v-model="end_time"
              format="HH:mm"
              value-format="timestamp"
              placeholder="结束时间"
            />
          </el-col>

          <el-col :span="4" class="text-center">
            <el-time-select
              v-model="duration"
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
        </div>

        <el-row :gutter="20" style="margin-top:50px;">
          <el-col :span="4" class="text-center">
            <el-input v-model="comment" placeholder="请输入内容" />
          </el-col>

          <el-col :span="4" class="text-center">
            <el-button type="primary" @click="handleSubmit">提交</el-button>
          </el-col>
        </el-row>
      </el-card>
    </el-row>

    <!-- <el-row> -->
    <el-row :gutter="20" style="margin-top:50px;">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>Buttons</span>
        </div>
        <div style="margin-bottom:50px;">
          <el-col :span="4" class="text-center">
            <router-link class="pan-btn blue-btn" to="/documentation/index">Documentation</router-link>
          </el-col>
        </div>
      </el-card>
    </el-row>
  </div>
</template>

<script>

import { submit } from '@/api/activity'

export default {
  data() {
    return {
      options: [
        {
          value: 'sleep',
          label: '睡眠'
        },
        {
          value: 'read',
          label: '阅读'
        },
        {
          value: 'fitness',
          label: '健身'
        },
        {
          value: 'entertainment',
          label: '娱乐'
        }
      ],
      type: '',
      start_time: new Date(),
      end_time: new Date(),
      duration: '',
      num: 1,
      comment: ''
    }
  },
  methods: {
    handleChange(type) {
      console.log(type)
    },
    handleSubmit: function(event) {
      var activity = {
        'type': this.type,
        'start_time': this.start_time,
        'end_time': this.end_time,
        'duration': this.duration,
        'num': this.num,
        'comment': this.comment
      }
      submit(activity).then(response => {
        alert('submit OK')
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
