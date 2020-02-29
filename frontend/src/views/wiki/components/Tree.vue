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
      :draggable="isDrag"
      @node-contextmenu="handleMenuClick"
      @node-click="handleNodeClick"
      @node-drop="handleDrop"
    >
      <span slot-scope="{ node, data }" class="custom-tree-node">
        <span v-show="!node.isEdit">
          <span v-if="data.filetype === 'nb'">
            <i class="el-icon-notebook-1" />
          </span>
          <span v-else-if="data.filetype === 'md'">
            <i class="el-icon-document" />
          </span>
          <span v-else />
          <span style="margin-left:5px;">{{ node.label }}</span>
          <span>
            <el-button type="text" size="mini" @click="() => handleNewFile(node)">+</el-button>
          </span>
        </span>
        <span v-show="node.isEdit">
          <el-input
            :ref="data.key"
            v-model="data.label"
            class="slot-t-input"
            size="mini"
            autofocus
            @blur.stop="NodeBlur(node, data)"
            @keyup.enter.native="NodeBlur(node, data)"
          />
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
      isDrag: true,
      // catalog: [],
      catalog: [{
        id: 0,
        key: '0',
        label: 'root',
        filetype: 'nb',
        level: 0,
        parent: '0',
        children: [{
          id: 1,
          key: '1',
          label: 'level 1',
          filetype: 'md',
          level: 1,
          parent: '0'
        }, {
          id: 2,
          key: '2',
          label: 'level 2',
          filetype: 'md',
          level: 1,
          parent: '0'
        }
        ]
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
  created() {
    bus.$on('update-catalog', this.updateCatalog)
  },
  mounted() {
    this.featchData()
  },

  beforeDestroy() {
    bus.$off('update-catalog', this.updateCatalog)
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
    handleMenuClick(event, obj, node, components) {
      console.log(node)
      this.$refs.menu.open(event, { value: node })
    },
    handleNodeClick(data, node) {
      // node edit toggle click
      console.log('node click')
      console.log(node)
      if (node.hasOwnProperty('isEdit') === true && node.isEdit === true) {
        return
      }

      if (data.filetype === 'md') {
        bus.$emit('show-reminder', false)
        bus.$emit('show-editor', true)
        bus.$emit('open-document', data)
      }
    },

    newRandomKey() {
      var uuid = uuidv4()
      return uuid.split('-').join('')
    },
    updateCatalog() {
      updateEditorCatalog(this.catalog).then(response => {
      })
    },

    handleNewFile(node) {
      // alert(`You clicked on: "${node.label}"`)
      console.log(node)

      var data = node.data

      var key = this.newRandomKey()

      const newChild = {
        id: 0,
        key: key,
        label: 'file',
        level: data.level + 1,
        filetype: 'md',
        parent: data.key,
        children: []
      }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)

      bus.$emit('show-reminder', false)
      bus.$emit('open-new-doc', newChild)
    },
    handleNewNoteBook(node) {
      var data = node.data

      if (data.level >= 2) {
        this.$message({
          message: '仅支持2级notebook',
          type: 'warning'
        })
        return
      }

      var key = this.newRandomKey()

      const newChild = {
        id: 0,
        key: key,
        label: 'notebook',
        level: data.level + 1,
        filetype: 'nb',
        parent: data.key,
        children: []
      }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)

      updateEditorCatalog(this.catalog).then(response => {
      })
    },
    NodeBlur(node, data) { // 输入框失焦
      console.log('lose focus')
      if (node.isEdit) {
        this.$set(node, 'isEdit', false)
      }
      this.isDrag = true
      updateEditorCatalog(this.catalog).then(response => {
      })
    },
    handleRename(node) {
      // alert(`You clicked on: "${node.label}"`)
      console.log(node)
      if (!node.isEdit) {
        this.$set(node, 'isEdit', true)
      }
      this.isDrag = false
      this.$nextTick(() => {
        this.$refs[node.data.key].$refs.input.focus()
      })
    },
    handleDelete(node) {
      console.log(node)
      const parent = node.parent
      var data = node.data
      const children = parent.data.children || parent.data
      const index = children.findIndex(d => d.id === data.id)
      children.splice(index, 1)

      updateEditorCatalog(this.catalog).then(response => {
      })
    },
    handleDrop(draggingNode, dropNode, dropType, ev) {
      updateEditorCatalog(this.catalog).then(response => {
      })
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

