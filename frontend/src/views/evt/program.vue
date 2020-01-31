<template>
  <div class="mixin-components-container">
    <el-row>
      <leetcode />
    </el-row>
    <el-row style="margin-top:10px;">
      <github />
    </el-row>
  </div>
</template>

<script>

import Leetcode from './components/Leetcode'
import Github from './components/Github'

export default {
  components: {
    Leetcode,
    Github
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
