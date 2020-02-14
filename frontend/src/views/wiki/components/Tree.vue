<template>
  <div class="app-container">
    <el-tree
      ref="tree2"
      :data="catalog"
      :props="defaultProps"
      :filter-node-method="filterNode"
      class="filter-tree"
      default-expand-all
      @node-contextmenu="handleClick"
    >
      <span slot-scope="{ node, data }" class="custom-tree-node">
        <span>{{ node.label }}</span>
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

import { queryEditorCatalog } from '@/api/editor'

export default {
  components: {
    VueContext
  },

  data() {
    return {
      filterText: '',
      // catalog: [],
      catalog: [{
        id: '1-0',
        label: 'Level one 1',
        children: [{
          id: '1-1',
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
      console.log(node.data.id)
      console.log(node.data.children)

      var data = node.data

      const newChild = { id: data.id++, label: 'testtest', children: [] }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)
    },
    handleRename(node) {
      alert(`You clicked on: "${node.label}"`)
      console.log(node)
    },
    handleDelete(node) {
      alert(`You clicked on: "${node.label}"`)
      console.log(node)
    }
  }
}
</script>

