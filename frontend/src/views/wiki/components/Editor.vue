<template>
  <el-container v-show="show" v-loading="loading">
    <el-header height="20px">
      <!-- <el-row> -->
      <el-col :span="12">
        <el-breadcrumb separator="/">
          <el-breadcrumb-item>首页</el-breadcrumb-item>
          <el-breadcrumb-item>活动管理</el-breadcrumb-item>
          <el-breadcrumb-item>活动列表</el-breadcrumb-item>
          <el-breadcrumb-item>活动详情</el-breadcrumb-item>
        </el-breadcrumb>
      </el-col>
      <el-col :span="12">
        <div class="right-menu">
          <el-button type="primary" plain size="small" @click="edit">edit</el-button>
          <el-button type="primary" plain size="small" @click="save">save</el-button>
        </div>
      </el-col>
      <!-- </el-row> -->
    </el-header>
    <el-main>
      <!-- <editor-content ref="editor" class="editor__content" :editor="editor" /> -->
      <div class="content" />
    </el-main>
  </el-container>
</template>

<script>
import { Editor } from 'tiptap'
import Stackedit from 'stackedit-js'

import { saveDocument, queryAllDocument, queryDocumentByID } from '@/api/editor'
import { bus } from '@/utils/bus'

export default {
  components: {
  },
  data() {
    return {
      show: false,
      loading: false,
      eleContent: null,
      stackedit: null,
      activeNode: null,
      title: '',
      contentHTML: '',
      contentText: ''
    }
  },
  created() {
    bus.$on('show-editor', this.showEditor)
    bus.$on('open-document', this.openDocument)
    bus.$on('open-new-doc', this.openNewDoc)
  },
  mounted() {
    this.initEditor()
  },
  beforeDestroy() {
    bus.$off('show-editor', this.showEditor)
    bus.$off('open-document', this.openDocument)
    bus.$off('open-new-doc', this.openNewDoc)
    this.editor.destroy()
  },
  methods: {
    initEditor() {
      this.stackedit = new Stackedit()
      this.eleContent = document.getElementsByClassName('content')
    },
    setContent(html) {
      this.contentHTML = html
      this.eleContent.innerHTML = html
    },
    showEditor(state) {
      this.show = true
    },
    openDocument(node) {
      this.$log.debug(node)

      if (node.id <= 0) {
        return
      }

      this.activeNode = node
      this.loading = true

      queryDocumentByID(node.id).then(response => {
        var docs = response.data
        this.setContent(docs.html)
        this.loading = false
      })
      this.loading = false
    },
    openNewDoc(node) {
      this.activeNode = node
      this.showEditor(true)
    },
    updateContent(content) {
      this.contents = content
    },
    edit(event) {
      this.stackedit.openFile({
        name: 'Filename', // with a filename
        content: {
          // text: this.contentText // and the Markdown content.
          text: 'Hello **Markdown** writer!'
        }
      })

      this.stackedit.on('fileChange', (file) => {
        // el.value = file.content.text
        this.contentHTML = file.content.html
        this.contentText = file.content.text
        this.$log.debug(file.content.text)
        this.$log.debug(file.content.html)
      })

      this.stackedit.on('close', () => {
        this.eleContent.innerHTML = this.contentHTML
        this.$log.debug(this.contentText)
        this.$log.debug(this.contentHTML)
      })
    },
    save(event) {
      var file = {
        'id': this.activeNode.id,
        'doc_key': this.activeNode.key,
        // 'title': this.title,
        'raw': '',
        'html': this.contents
      }
      saveDocument(file).then(response => {
        this.activeNode.id = response.data.id
        if (this.title !== '') {
          this.activeNode.label = this.title
        }

        bus.$emit('update-catalog')

        this.$notify({
          title: '成功',
          message: '',
          type: 'success',
          duration: 800
        })
      })
    },
    getAllDoc() {
      queryAllDocument().then(response => {
        var docs = response.data
        this.contents = docs[0].html
        this.editor.setContent(this.contents)
      })
    }
  }
}
</script>

<style lang="scss">
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
.el-button {
  padding: 2 20px;
}

@import "~@/styles/variables.scss";
// @import "~@/styles/editor.scss";
// @import "~@/styles/main.scss";
@import "~@/styles/menubar.scss";
@import "~@/styles/menububble.scss";

.editor *.is-empty:nth-child(1)::before,
.editor *.is-empty:nth-child(2)::before {
  content: attr(data-empty-text);
  float: left;
  color: #aaa;
  pointer-events: none;
  height: 0;
  font-style: italic;
}

.right-menu {
  float: right;
  height: 100%;
  line-height: 10px;
}
</style>
