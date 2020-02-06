<template>
  <el-input v-model="filterText" placeholder="Filter keyword" style="margin-bottom:30px;" />
</template>

<script>
export default {
  data() {
    return {
      filterText: ''
    }
  }
}
</script>
