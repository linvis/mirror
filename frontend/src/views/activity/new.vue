<template>
  <div class="mixin-components-container">
    <!-- <el-row> -->
    <el-row :gutter="20">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>日常活动</span>
        </div>
        <div style="margin-bottom:50px;">
          <el-col :span="4" class="text-center">
            <el-select v-model="value" placeholder="活动类型">
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              ></el-option>
            </el-select>
          </el-col>

          <el-col :span="4" class="text-center">
            <el-time-picker v-model="start_time" format="HH:mm" placeholder="开始时间"></el-time-picker>
          </el-col>

          <el-col :span="4" class="text-center">
            <el-time-picker v-model="end_time" format="HH:mm" placeholder="结束时间"></el-time-picker>
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
            ></el-time-select>
          </el-col>

          <el-col :span="4" class="text-center">
            <el-input-number v-model="num" @change="handleChange" :min="1" :max="10" label="数量"></el-input-number>
          </el-col>
        </div>
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
export default {
  data() {
    return {
      options: [
        {
          value: "选项1",
          label: "睡眠"
        },
        {
          value: "选项2",
          label: "阅读"
        },
        {
          value: "选项3",
          label: "健身"
        },
        {
          value: "选项4",
          label: "娱乐"
        }
      ],
      value: "",
      start_time: "",
      end_time: "",
      duration: "",
      num: 1
    };
  },
  methods: {
    handleChange(value) {
      console.log(value);
    }
  }
};
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
