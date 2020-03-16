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
  computed: {
    catalog() {
      return this.$store.state.editor.catalog
    }
  },
  watch: {
    catalog: function(newVal, oldVal) {
      this.$log.debug('watch catalog change', newVal)
      this.reminderList = this.filterReminder(newVal)
    }
  },
  created() {
    bus.$on('show-reminder', this.showReminder)
    bus.$on('update-reminder', this.updateReminder)
  },
  beforeDestroy() {
    bus.$off('show-reminder', this.showReminder)
  },
  methods: {
    showReminder(state) {
      this.show = state
    },
    updateReminder() {
      this.reminderList = this.filterReminder(this.catalog)
    },
    filterReminder(catalog) {
      var list = []

      if (catalog === null || typeof catalog === 'undefined') {
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
        }
        if ('children' in catalog[i]) {
          var subList = this.filterReminder(catalog[i].children)
          list = list.concat(subList)
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
