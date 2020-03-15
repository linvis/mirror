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
    handleReminder(command) {
      this.$log.debug(command)
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
