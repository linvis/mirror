<template>
  <!-- <div class="components-container"> -->
  <el-container>
    <splitpanes>
      <pane min-size="10" size="20" max-size="30" class="sidebar">
        <search />
        <tree />
      </pane>
      <pane>
        <reminder style="margin-left:20px;margin-top:5px" />
        <editor style="margin-left:20px;margin-top:5px" />
      </pane>
    </splitpanes>
    <!-- <split-pane split="vertical" :default-percent="25" @resize="resize">
      <template slot="paneL">
        <el-container>
          <el-header style="margin-top:10px">
            <search />
          </el-header>
          <el-main>
            <tree />
          </el-main>
        </el-container>
      </template>
      <template slot="paneR">
        <reminder style="margin-left:20px;margin-top:20px" />
        <editor style="margin-left:20px;margin-top:20px" />
      </template>
    </split-pane>-->
    <!-- </div> -->
  </el-container>
</template>

<script>
// import splitPane from 'vue-splitpane'
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'

import Search from './components/Search'
import Tree from './components/Tree'
import Editor from './components/Editor'
import Reminder from './components/Reminder'

export default {
  components: {
    // splitPane,
    Splitpanes,
    Pane,
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
  color: #333;
  border-radius: 2px;
}
.el-container {
  height: calc(100vh - 50px);
}
.splitpanes__pane span {
  font-family: Helvetica, Arial, sans-serif;
  height: 100%;
}
.splitpanes--vertical > .splitpanes__splitter {
  min-width: 5px;
  height: 100%;
}
.splitpanes--vertical {
  height: 100%;
}
.sidebar {
  background: #f8f8f8;
  height: 100%;
}
.search {
  margin-top: 5%;
  margin-bottom: 0%;
  margin-left: 10%;
  margin-right: 10%;
}
.tree {
  margin: 0%;
}
</style>
