<template>
  <!-- <el-container v-show="show" style="height:100%;"> -->
  <el-container v-show="show" style="height: 100%; border: 1px solid #eee">
    <el-main>
      <el-table :data="reminderList" style="width: 100%">
        <el-table-column prop="fileName" label="filename" width="180" />
        <el-table-column prop="overdue" label="overdue" width="100" />
        <el-table-column prop="operate" label="operate">
          <template slot-scope="scope">
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">open</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
import { bus } from '@/utils/bus'
export default {
  data() {
    return {
      show: true,
      reminderList: []
    }
  },
  created() {
    bus.$on('show-reminder', this.showReminder)
    bus.$on('get-catalog', this.updateReminder)
  },
  beforeDestroy() {
    bus.$off('show-reminder', this.showReminder)
    bus.$off('get-catalog', this.updateReminder)
  },
  methods: {
    showReminder(state) {
      this.show = state
    },
    updateReminder(catalog) {
      this.reminderList = this.filterReminder(catalog)
      this.$log.debug(this.reminderList)
    },
    filterReminder(catalog) {
      var list = []

      if (catalog === null) {
        return list
      }

      this.$log.debug(catalog)
      for (var i = 0; i < catalog.length; i++) {
        if (catalog[i].reminder.enable === true) {
          list.push({
            'fileName': catalog[i].metadata.title,
            'overdue': 0,
            'file': catalog[i]
          })

          if ('children' in catalog[i]) {
            var subList = this.filterReminder(catalog[i].children)
            list = list.concat(subList)
          }
        }
      }

      return list
    },
    handleEdit(index, row) {
      this.$log.debug(index, row)
      bus.$emit('show-reminder', false)
      bus.$emit('show-editor', true)
      bus.$emit('open-document', row.file)
    }
  }
}
</script>
