<template>
  <div class="menu">
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

const reviewDay = [1, 3, 5, 7, 14, 30]

export default {
  data() {
    return {
      activeFile: null
    }
  },
  created() {
    bus.$on('open-document', this.updateActiveFile)
    bus.$on('open-new-doc', this.updateActiveFile)
    // this.$nextTick(() => {
    // })
  },
  beforeDestroy() {
    bus.$off('open-document', this.updateActiveFile)
    bus.$off('open-new-doc', this.updateActiveFile)
  },
  methods: {
    updateActiveFile(file) {
      this.activeFile = file
      this.$log.debug(file)
    },
    save(event) {

    },
    addDays(days) {
      var date = new Date(this.valueOf())
      date.setDate(date.getDate() + days)
      return date
    },
    handleReminder(command) {
      this.$log.debug(command)
      if (command === 'start') {
        this.activeFile.reminder.enable = true
        this.activeFile.reminder.count = 1
        this.activeFile.reminder.last_time = 0
        this.activeFile.reminder.next_time = this.addDays(reviewDay[this.activeFile.reminder.count - 1])
      } else if (command === 'reviewed') {
        this.activeFile.reminder.count++
        if (this.activeFile.reminder.count >= reviewDay.length()) {
          this.activeFile.reminder.count = reviewDay.length() - 1
        }
        this.activeFile.reminder.last_time = this.activeFile.reminder.next_time
        this.activeFile.reminder.next_time = this.addDays(reviewDay[this.activeFile.reminder.count - 1])
      } else {
        this.activeFile.reminder.next_time = this.addDays(reviewDay[this.activeFile.reminder.count - 1])
      }
      bus.$emit('update-catalog')
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
