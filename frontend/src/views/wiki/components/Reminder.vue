<template>
  <!-- <el-container v-show="show" style="height:100%;"> -->
  <el-container v-show="show" style="height: 100%; border: 1px solid #eee">
    <el-main>
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="date" label="日期" width="180" />
        <el-table-column prop="name" label="姓名" width="180" />
        <el-table-column prop="address" label="地址" />
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
import { bus } from '@/utils/bus'
export default {
  data() {
    const item = {
      date: '2016-05-02',
      name: '王小虎',
      address: '上海市普陀区金沙江路 1518 弄'
    }
    return {
      show: true,
      tableData: Array(30).fill(item)
    }
  },
  created() {
    bus.$on('show-reminder', this.showReminder)
  },
  beforeDestroy() {
    bus.$off('show-reminder', this.showReminder)
  },
  methods: {
    showReminder(state) {
      this.show = state
    }
  }
}
</script>
