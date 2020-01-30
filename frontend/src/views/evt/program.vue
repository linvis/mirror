<template>
  <div class="mixin-components-container">
    <leetcode />
  </div>
</template>

<script>

import Leetcode from './components/Leetcode'

export default {
  components: {
    Leetcode
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
