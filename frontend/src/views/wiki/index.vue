<template>
  <div class="components-container">
    <el-header style="margin-top:20px">
      <search />
    </el-header>
    <split-pane
      split="vertical"
      :default-percent="25"
      @resize="resize"
    >
      <template slot="paneL">
        <split-pane split="vertical" @resize="resize">
          <template slot="paneL">
            <tree />
          </template>
          <template slot="paneR">
            <tree />
          </template>
        </split-pane>
      </template>
      <template slot="paneR">
        <editor style="margin-left:40px" />
      </template>
    </split-pane>
    </el-header></div>
</template>

<script>
import splitPane from 'vue-splitpane'

import Search from './components/Search'
import Tree from './components/Tree'
import Editor from './components/Editor'

export default {
  components: {
    splitPane,
    Search,
    Tree,
    Editor
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
