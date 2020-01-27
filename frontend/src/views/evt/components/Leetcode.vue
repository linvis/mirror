<template>
  <el-form-item label="nums">
    <el-input-number v-model="num" :min="1" :max="50" label="数量" />
  </el-form-item>
</template>

<script>
export default {
  data: function() {
    return {
      num: 1
    }
  }
}
</script>
