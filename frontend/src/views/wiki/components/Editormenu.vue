<template>
  <div class="menu">
    <el-button
      type="text"
      icon="el-icon-s-home"
      class="menu-button"
      @click="backHome"
    />
    <el-button
      type="text"
      icon="el-icon-upload"
      class="menu-button"
      @click="save"
    />
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
import { bus } from '@/utils/bus';
import { saveDocument } from '@/api/editor';

const reviewDay = [1, 3, 5, 7, 14, 30]

export default {
  data() {
    return {}
  },
  computed: {
    activeFile() {
      return this.$store.state.editor.activeFile
    }
  },
  created() {},
  beforeDestroy() {},
  methods: {
    backHome(event) {
      bus.$emit('show-reminder', true)
      bus.$emit('show-editor', false)
    },
    addDays(days) {
      var date = new Date()
      date.setDate(date.getDate() + days)
      return date.getTime()
    },
    handleReminder(command) {
      this.$log.debug(command, this.activeFile)

      if (command === 'start') {
        this.activeFile.reminder.enable = true
        this.activeFile.reminder.count = 1
        this.activeFile.reminder.last_time = 0
        this.activeFile.reminder.next_time = this.addDays(
          reviewDay[this.activeFile.reminder.count - 1]
        )
      } else if (command === 'reviewed') {
        this.activeFile.reminder.count++
        if (this.activeFile.reminder.count >= reviewDay.length) {
          this.activeFile.reminder.count = reviewDay.length - 1
        }
        this.activeFile.reminder.last_time = this.activeFile.reminder.next_time
        this.activeFile.reminder.next_time = this.addDays(
          reviewDay[this.activeFile.reminder.count - 1]
        )
      } else {
        this.activeFile.reminder.next_time = this.addDays(
          reviewDay[this.activeFile.reminder.count - 1]
        )
      }
      this.$log.debug(this.$store.state.editor.activeFile)
      this.$log.debug(this.$store.state.editor.catalog)
      this.$store.dispatch('editor/submitCatalog')

      bus.$emit('update-reminder')
    },
    save(event) {
      var content = this.$store.state.editor.contents[this.activeFile.key]
      var doc = {
        id: this.activeFile.id,
        doc_key: this.activeFile.key,
        content: content,
        html: this.html
      }

      this.$log.debug('udpate tree', this.activeFile)
      saveDocument(doc).then(response => {
        this.activeFile.id = response.data.id
        var title = this.activeFile.metadata.title
        if (this.activeFile.metadata.title === 'untitle') {
          if (content.includes('\n')) {
            title = content.split('\n')[0]
            if (title.includes('#')) {
              title = title.split('#')[1]
            }
          }
        }
        this.activeFile.metadata.title = title

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
