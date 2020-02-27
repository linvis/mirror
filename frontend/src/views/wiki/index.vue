<template>
  <div class="components-container">
    <split-pane split="vertical" :default-percent="25" @resize="resize">
      <template slot="paneL">
        <el-header style="margin-top:10px">
          <search />
        </el-header>
        <tree />
      </template>
      <template slot="paneR">
        <reminder style="margin-left:20px;margin-top:20px" />
        <editor style="margin-left:20px;margin-top:20px" />
      </template>
    </split-pane>
  </div>
</template>

<script>
import splitPane from 'vue-splitpane'

import Search from './components/Search'
import Tree from './components/Tree'
import Editor from './components/Editor'
import Reminder from './components/Reminder'

export default {
  components: {
    splitPane,
    Search,
    Tree,
    Editor,
    Reminder
  },
  data() {
    return {
      showReminder: true,
      showEditor: false
    }
  },
  methods: {
    created() {
      // this.$refs.created()
      console.log('create')
    },
    refresh() {
      console.log('refresh')
      this.showReminder = false
      this.showEditor = true
    },
    resize() {
    }

  }

}
</script>

<style>
.el-header,
.el-footer {
  /* background-color: #b3c0d1; */
  color: #333;
  border-radius: 2px;
  /* text-align: center;
  line-height: 60px; */
}
.components-container {
  position: relative;
  height: 100vh;
}
</style>
