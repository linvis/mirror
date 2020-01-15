<template>
  <div class="mixin-components-container">
    <el-row>
      <sleep-new />
    </el-row>
    <el-row style="margin-top:10px;">
      <sleep-chart />
    </el-row>
  </div>
</template>

<script>

import SleepChart from './components/SleepChart'
import SleepNew from './components/SleepNew'

export default {
  components: {
    SleepNew,
    SleepChart
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
