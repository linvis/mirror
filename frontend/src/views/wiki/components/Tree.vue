<template>
  <div class="app-container">
    <el-tree
      ref="tree"
      :data="catalog"
      :props="defaultProps"
      :filter-node-method="filterNode"
      class="tree"
      :highlight-current="true"
      :expand-on-click-node="true"
      :draggable="isDrag"
      @node-contextmenu="handleMenuClick"
      @node-click="handleNodeClick"
      @node-drop="handleDrop"
    >
      <span slot-scope="{ node, data }" class="custom-tree-node">
        <span v-show="!node.isEdit">
          <span v-if="data.metadata.filetype === 'nb'">
            <i class="el-icon-notebook-1" />
          </span>
          <span v-else-if="data.metadata.filetype === 'md'">
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
            v-model="data.metadata.title"
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
      defaultProps: {
        children: 'children',
        label: function(data, node) {
          return data.metadata.title
        }
      }
    }
  },
  computed: {
    catalog() {
      return this.$store.state.editor.catalog
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
    this.getEditorCatalog()
  },

  beforeDestroy() {
    bus.$off('update-catalog', this.updateCatalog)
  },

  methods: {
    getEditorCatalog() {
      queryEditorCatalog().then(response => {
        // this.catalog = response.data
        this.$store.dispatch('editor/changeCatalog', response.data)
      })
      // bus.$emit('get-catalog', this.catalog)
    },
    filterNode(value, data) {
      if (!value) return true
      return data.metadata.title.indexOf(value) !== -1
    },
    handleMenuClick(event, obj, node, components) {
      this.$log.debug(node)
      this.$refs.menu.open(event, { value: node })
    },
    handleNodeClick(data, node) {
      // node edit toggle click
      this.$log.debug(node)
      if (node.hasOwnProperty('isEdit') === true && node.isEdit === true) {
        return
      }

      if (data.metadata.filetype === 'md') {
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
      this.$log.debug(node)

      var data = node.data
      var key = this.newRandomKey()

      const newChild = {
        id: 0,
        key: key,
        parent: data.key,
        metadata: {
          title: 'untitle',
          level: data.level + 1,
          filetype: 'md',
          tag: []
        },
        reminder: {
          enable: false,
          count: 0,
          last_time: 0,
          next_time: 0
        },
        children: []
      }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)

      this.$store.dispatch('editor/submitCatalog', this.catalog)

      bus.$emit('show-reminder', false)
      bus.$emit('open-new-doc', newChild)
    },
    handleNewNoteBook(node) {
      var data = node.data

      if (data.metadata.level >= 2) {
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
        parent: data.key,
        metadata: {
          title: 'notebook',
          level: data.level + 1,
          filetype: 'nb',
          tag: []
        },
        reminder: {
          enable: false,
          count: 0,
          last_time: 0,
          next_time: 0
        },
        children: []
      }
      if (!data.children) {
        this.$set(data, 'children', [])
      }
      data.children.push(newChild)

      this.$store.dispatch('editor/submitCatalog', this.catalog)
    },
    NodeBlur(node, data) { // 输入框失焦
      this.$log.debug('lose focus', node, data)

      if (node.isEdit) {
        this.$set(node, 'isEdit', false)
      }
      this.isDrag = true

      this.$store.dispatch('editor/submitCatalog', this.catalog)
    },
    handleRename(node) {
      // alert(`You clicked on: "${node.data.metadata.title}"`)
      this.$log.debug(node)
      if (!node.isEdit) {
        this.$set(node, 'isEdit', true)
      }
      this.isDrag = false
      this.$nextTick(() => {
        this.$refs[node.data.key].$refs.input.focus()
      })
    },
    handleDelete(node) {
      this.$log.debug(node)

      const parent = node.parent
      var data = node.data
      const children = parent.data.children || parent.data
      const index = children.findIndex(d => d.id === data.id)
      children.splice(index, 1)

      this.$store.dispatch('editor/submitCatalog', this.catalog)
    },
    handleDrop(draggingNode, dropNode, dropType, ev) {
      this.$store.dispatch('editor/submitCatalog', this.catalog)
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

.tree {
  background: #f8f8f8;
}
</style>

