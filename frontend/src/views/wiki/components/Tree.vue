<template>
  <div class="app-container">
    <el-tree
      ref="tree2"
      :data="data2"
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
          <a @click.prevent="alertName(child.data.value)">Alert name</a>
        </li>
      </template>
    </vue-context>
  </div>
</template>

<script>

import VueContext from 'vue-context'
import 'vue-context/src/sass/vue-context.scss'

export default {
  components: {
    VueContext
  },

  data() {
    return {
      options: [
        {
          name: 'Duplicate',
          slug: 'duplicate'
        },
        {
          name: 'Edit',
          slug: 'edit'
        },
        {
          name: 'Delete',
          slug: 'delete'
        }
      ],
      filterText: '',
      data2: [{
        id: '1-0',
        label: 'Level one 1',
        children: [{
          id: '1-1',
          label: 'Level two 1-1'
        }]
      }, {
        id: '2-0',
        label: 'Level one 2',
        children: [{
          id: '2-1',
          label: 'Level two 2-1'
        }, {
          id: '2-2',
          label: 'Level two 2-2'
        }]
      }, {
        id: '3-0',
        label: 'Level one 3',
        children: [{
          id: '3-1',
          label: 'Level two 3-1'
        }, {
          id: '3-2',
          label: 'Level two 3-2'
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

  methods: {
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    handleClick(event, obj, node, components) {
      console.log(node)
      this.$refs.menu.open(event, { value: node })
    },
    alertName(node) {
      alert(`You clicked on: "${node.label}"`)
      console.log(node)
    }
  }
}
</script>

