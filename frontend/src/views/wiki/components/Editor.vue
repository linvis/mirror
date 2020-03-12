<template>
  <!-- <div id="app"> -->
  <el-container v-show="show" v-loading="loading" class="app">
    <div class="navbar">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>首页</el-breadcrumb-item>
        <el-breadcrumb-item>活动管理</el-breadcrumb-item>
        <el-breadcrumb-item>活动列表</el-breadcrumb-item>
        <el-breadcrumb-item>活动详情</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <!-- <el-header class="head">
      <el-col :span="12">
      </el-col>
      <el-col :span="12">
        <div class="right-menu">
          <el-button type="text" icon="el-icon-check" @click="save" />
          <el-button type="primary" plain size="small" @click="edit">edit</el-button>
          <el-button type="primary" plain size="small" @click="reminder">reminder</el-button>
        </div>
      </el-col>
    </el-header>-->
    <simplebar class="editor">
      <div ref="editor" />
    </simplebar>
  </el-container>
  <!-- </div> -->
</template>

<script>
import Muya from '@/muya/lib'
import TablePicker from '@/muya/lib/ui/tablePicker'
import QuickInsert from '@/muya/lib/ui/quickInsert'
import CodePicker from '@/muya/lib/ui/codePicker'
import EmojiPicker from '@/muya/lib/ui/emojiPicker'
import ImagePathPicker from '@/muya/lib/ui/imagePicker'
import ImageSelector from '@/muya/lib/ui/imageSelector'
import FormatPicker from '@/muya/lib/ui/formatPicker'
import FrontMenu from '@/muya/lib/ui/frontMenu'
import '@/muya/themes/default.css'
import './assets/index.css'

import simplebar from 'simplebar-vue'
import 'simplebar/dist/simplebar.min.css'

Muya.use(TablePicker)
Muya.use(QuickInsert)
Muya.use(CodePicker)
Muya.use(EmojiPicker)
Muya.use(ImagePathPicker)
Muya.use(ImageSelector)
Muya.use(FormatPicker)
Muya.use(FrontMenu)

import { saveDocument, queryDocumentByID } from '@/api/editor'
import { bus } from '@/utils/bus'

export default {
  components: {
    simplebar
  },
  data() {
    return {
      show: false,
      loading: false,
      editor: null,
      activeNode: null,
      content: '',
      html: '',
      title: ''
    }
  },
  // name: 'App',
  created() {
    bus.$on('show-editor', this.showEditor)
    bus.$on('open-document', this.openDoc)
    bus.$on('open-new-doc', this.openNewDoc)
    // this.$nextTick(() => {
    // })
  },
  mounted() {
    if (this.editor === null) {
      const ele = this.$refs.editor
      this.editor = new Muya(ele)
    }
    this.editor.on('change', changes => {
      console.log(changes.markdown)
      this.content = changes.markdown
    })
  },
  beforeDestroy() {
    bus.$off('show-editor', this.showEditor)
    bus.$off('open-document', this.openDoc)
    bus.$off('open-new-doc', this.openNewDoc)
    this.editor.destroy()
  },
  methods: {
    showEditor(state) {
      this.show = true
    },
    setContent(content) {
      this.content = content
      this.editor.setMarkdown(content)
    },
    openDoc(node) {
      this.$log.debug(node)

      if (node.id <= 0) {
        return
      }

      this.activeNode = node
      this.loading = true

      queryDocumentByID(node.id).then(response => {
        var docs = response.data
        this.$log.debug('update markdown', docs.content)
        this.setContent(docs.content)
        this.loading = false
      })

      this.loading = false
    },
    openNewDoc(node) {
      this.activeNode = node
      this.showEditor(true)
    },
    edit(event) {
      this.setContent('Welcome to use muya...\n\nhello fuck\n')
    },
    reminder(event) {
      this.setContent('Welcome to use muya...\n\nhello fuck\n')
    },
    save(event) {
      var doc = {
        'id': this.activeNode.id,
        'content': this.content,
        'html': this.html
      }
      this.activeNode.label = this.title

      this.$log.debug('udpate tree', this.activeNode)

      saveDocument(doc).then(response => {
        this.activeNode.id = response.data.id
        if (this.title === '') {
          this.title = 'untitle'
          if (this.content.includes('\n')) {
            var title = this.content.split('\n')[0]
            if (title.includes('#')) {
              this.title = title.split('#')[1]
            }
          }
        }
        this.activeNode.label = this.title

        this.$log.debug('udpate tree', this.activeNode)

        bus.$emit('update-catalog')

        this.$notify({
          title: '成功',
          message: '',
          type: 'success',
          duration: 800
        })
      })
    }
  }
}
</script>

<style>
.components-container {
  /* display: flex;
  flex-direction: column;
  justify-content: flex-start; */
  /* position: relative; */
}
.right-menu {
  float: right;
  height: 100%;
  line-height: 10px;
  justify-content: left;
}
.navbar {
  width: 100%;
  height: 20px;
}
.editor {
  float: left;
  width: 100%;
  /* height: 300px; */
  padding: 0%;
  /* height: 100%; */
  height: calc(100vh - 20px);
  /* height: 100px; */
  justify-content: left;
}
.app {
  display: flex;
  flex-direction: column;
  align-items: left;
}
</style>
