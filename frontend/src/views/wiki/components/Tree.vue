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
      @node-contextmenu="handleClick"
    >
      <span slot-scope="{ node, data }" class="custom-tree-node">
        <div v-if="data.filetype === 0">
          <i class="el-icon-folder" />
        </div>
        <div v-else-if="data.filetype === 1">
          <i class="el-icon-document" />
        </div>
        <div v-else />
        <span style="margin-left:5px;">{{ node.label }}</span>
        <span>
          <el-button type="text" size="mini" @click="() => handleNew(node)">+</el-button>
        </span>
      </span>
    </el-tree>
    <vue-context ref="menu">
      <template v-if="child.data" slot-scope="child">
        <li>
          <a @click.prevent="handleNew(child.data.value)">New</a>
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

import { queryEditorCatalog, submitNewCatalog } from '@/api/editor'

export default {
  components: {
    VueContext
  },

  data() {
    return {
      filterText: '',
      // catalog: [],
      catalog: [{
        id: 1,
        label: 'Level one 1',
        filetype: 0,
        level: 1,
        children: [{
          id: 2,
          filetype: 1,
          label: 'Level two 1-1'
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
    handleClick(event, obj, node, components) {
      console.log(node)
      this.$refs.menu.open(event, { value: node })
    },

    handleNew(node) {
      // alert(`You clicked on: "${node.label}"`)
      console.log(node)

      var data = node.data

      const newChild = { id: data.id++, label: 'notebook', children: [] }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)

      submitNewCatalog(newChild).then(response => {
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

