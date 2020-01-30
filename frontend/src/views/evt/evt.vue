<template>
  <div class="mixin-components-container">
    <!-- <el-row> -->
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
        <el-form-item label="活动类型">
          <el-cascader
            v-model="evt_type"
            placeholder="活动类型"
            :options="options"
            :props="{ expandTrigger: 'hover' }"
            filterable
            clearable
            @change="handleChange"
          />
        </el-form-item>
        <div v-if="evt_type === 'evt1'">
          <leetcode />
        </div>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">立即创建</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>

// import { submit } from '@/api/record'
import Leetcode from './components/Leetcode'

export default {
  components: {
    Leetcode
  },
  data() {
    return {
      options: [{
        value: 'evt1',
        label: 'leetcode'
      }],
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now()
        }
      },
      form: {},
      record_date: new Date(),
      evt_type: null

    }
  },
  methods: {
    handleSubmit() {
      console.log('submit!')
    },
    handleChange(value) {
      console.log(value)
      this.evt_type = value[0]
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
