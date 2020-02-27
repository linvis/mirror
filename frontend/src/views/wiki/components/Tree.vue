<template>
  <div class="app-container">
    <el-tree
      ref="tree"
      :data="catalog"
      :props="defaultProps"
      :filter-node-method="filterNode"
      class="filter-tree"
      :highlight-current="true"
      :expand-on-click-node="false"
      @node-contextmenu="handleMenuClick"
      @node-click="handleNodeClick"
    >
      <span slot-scope="{ node, data }" class="custom-tree-node">
        <div v-if="data.filetype === 'nb'">
          <i class="el-icon-notebook-1" />
        </div>
        <div v-else-if="data.filetype === 'md'">
          <i class="el-icon-document" />
        </div>
        <div v-else />
        <span style="margin-left:5px;">{{ node.label }}</span>
        <span>
          <el-button type="text" size="mini" @click="() => handleNewFile(node)">+</el-button>
        </span>
      </span>
    </el-tree>
    <vue-context ref="menu">
      <template v-if="child.data" slot-scope="child">
        <li>
          <a @click.prevent="handleNewFile(child.data.value)">New File</a>
          <a @click.prevent="handleNewNoteBook(child.data.value)">New NoteBook</a>
          <a @click.prevent="handleRename(child.data.value)">Rename</a>
          <a @click.prevent="handleDelete(child.data.value)">Delete</a>
        </li>
      </template>
    </vue-context>
  </div>
</template>

<script>

import VueContext from 'vue-context'
import 'vue-context/src/sass/vue-context.scss'
import { v4 as uuidv4 } from 'uuid'

import { queryEditorCatalog, updateEditorCatalog } from '@/api/editor'
import { bus } from '@/utils/bus'

export default {
  components: {
    VueContext
  },

  data() {
    return {
      filterText: '',
      // catalog: [],
      catalog: [{
        id: '0',
        label: 'root',
        filetype: 'nb',
        level: 0,
        parentID: '0',
        children: [{
          id: '1',
          label: 'level 1',
          filetype: 'md',
          level: 1,
          parentID: '0'
        }]
      }],
      defaultProps: {
        children: 'children',
        label: 'label'
      }
    }
  },
  watch: {
    filterText(val) {
      this.$refs.tree2.filter(val)
    }
  },
  mounted() {
    this.featchData()
  },

  methods: {
    featchData() {
      queryEditorCatalog().then(response => {
        this.catalog = response.data
      })
    },
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    getMaxChildID(children) {
      var maxID = -1
      for (var i = 0; i < children.length; i++) {
        if (children[i].id > maxID) {
          maxID = children[i].id
        }
      }

      return maxID
    },
    handleMenuClick(event, obj, node, components) {
      console.log(node)
      this.$refs.menu.open(event, { value: node })
    },
    handleNodeClick(data) {
      if (data.filetype === 'md') {
        bus.$emit('show-reminder', false)
        bus.$emit('show-editor', true)
        bus.$emit('open-document', data.id)
      }
    },

    newID() {
      var uuid = uuidv4()
      return uuid.split('-').join('')
    },

    handleNewFile(node) {
      // alert(`You clicked on: "${node.label}"`)
      console.log(node)

      var data = node.data

      var id = this.newID()

      const newChild = {
        id: id,
        label: 'file',
        level: data.level + 1,
        filetype: 'md',
        parentID: data.id,
        children: []
      }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)

      updateEditorCatalog(this.catalog).then(response => {
        bus.$emit('show-reminder', false)
        bus.$emit('open-new-doc', id)
      })
    },
    handleNewNoteBook(node) {
      // alert(`You clicked on: "${node.label}"`)
      var data = node.data

      console.log('current node', data.level, data.id, data.label)

      if (data.level >= 2) {
        this.$message({
          message: '仅支持2级notebook',
          type: 'warning'
        })
        return
      }

      var id = this.newID()

      const newChild = {
        id: id,
        label: 'notebook',
        level: data.level + 1,
        filetype: 'nb',
        parentID: data.id,
        children: []
      }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)

      updateEditorCatalog(this.catalog).then(response => {
      })
    },
    handleRename(node) {
      alert(`You clicked on: "${node.label}"`)
      console.log(node)
    },
    handleDelete(node) {
      console.log(node)
      const parent = node.parent
      var data = node.data
      const children = parent.data.children || parent.data
      const index = children.findIndex(d => d.id === data.id)
      children.splice(index, 1)
    }
  }
}
</script>

<style>
.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}
</style>

