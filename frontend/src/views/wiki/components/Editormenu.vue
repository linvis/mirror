<template>
  <div class="menu">
    <el-button type="text" icon="el-icon-s-home" class="menu-button" @click="backHome" />
    <el-button type="text" icon="el-icon-upload" class="menu-button" @click="save" />
    <el-dropdown @command="handleReminder">
      <el-button type="text" icon="el-icon-bell" class="menu-button" />
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item command="reviewed">Reviewed</el-dropdown-item>
        <el-dropdown-item command="later">Later</el-dropdown-item>
        <el-dropdown-item command="start">Start</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
  </div>
</template>

<script>

import { bus } from '@/utils/bus'
import { saveDocument } from '@/api/editor'

const reviewDay = [1, 3, 5, 7, 14, 30]

export default {
  data() {
    return {
    }
  },
  created() {
  },
  beforeDestroy() {
  },
  methods: {
    updateActiveFile(file) {
      this.activeFile = file
      this.$log.debug(file)
    },
    backHome(event) {
      bus.$emit('show-reminder', true)
      bus.$emit('show-editor', false)
    },
    addDays(days) {
      var date = new Date(this.valueOf())
      date.setDate(date.getDate() + days)
      return date
    },
    handleReminder(command) {
      var activeFile = this.$store.state.editor.activeFile

      this.$log.debug(command, activeFile)

      if (command === 'start') {
        activeFile.reminder.enable = true
        activeFile.reminder.count = 1
        activeFile.reminder.last_time = 0
        activeFile.reminder.next_time = this.addDays(reviewDay[activeFile.reminder.count - 1])
      } else if (command === 'reviewed') {
        activeFile.reminder.count++
        if (activeFile.reminder.count >= reviewDay.length) {
          activeFile.reminder.count = reviewDay.length - 1
        }
        activeFile.reminder.last_time = activeFile.reminder.next_time
        activeFile.reminder.next_time = this.addDays(reviewDay[activeFile.reminder.count - 1])
      } else {
        activeFile.reminder.next_time = this.addDays(reviewDay[activeFile.reminder.count - 1])
      }
      this.$log.debug(this.$store.state.editor.activeFile)
      this.$log.debug(this.$store.state.editor.catalog)
      this.$store.dispatch('editor/submitCatalog')

      bus.$emit('update-reminder')
    },
    save(event) {
      var activeFile = this.$store.state.editor.activeFile
      var content = this.$store.state.editor.contents[activeFile.key]
      var doc = {
        'id': activeFile.id,
        'doc_key': activeFile.key,
        'content': content,
        'html': this.html
      }

      this.$log.debug('udpate tree', activeFile)
      saveDocument(doc).then(response => {
        activeFile.id = response.data.id
        var title = activeFile.metadata.title
        if (activeFile.metadata.title === 'untitle') {
          if (content.includes('\n')) {
            title = content.split('\n')[0]
            if (title.includes('#')) {
              title = title.split('#')[1]
            }
          }
        }
        activeFile.metadata.title = title

        this.$store.dispatch('editor/submitCatalog')

        this.$notify({
          title: '成功',
          message: '',
          type: 'success',
          duration: 800
        })
      })
    }

  }
}
</script>

<style>
.menu-button {
  padding: 0%;
  font-size: 20px;
  margin-left: 10px;
}
.el-dropdown + .el-dropdown {
  margin-left: 15px;
}
.el-icon-arrow-down {
  font-size: 12px;
}
</style>
