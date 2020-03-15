<template>
  <el-container v-show="show" v-loading="loading" class="app">
    <el-container class="navbar">
      <el-col :span="18">
        <tag />
      </el-col>
      <el-col :span="3" :offset="3" class="editor-menu">
        <editormenu />
      </el-col>
    </el-container>
    <simplebar>
      <div ref="editor" class="editor" />
    </simplebar>
    <span style="height:=10vh;" />
  </el-container>
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

import Tag from './Tag'
import Editormenu from './Editormenu'

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
    simplebar,
    Tag,
    Editormenu
  },
  data() {
    return {
      show: false,
      loading: false,
      editor: null,
      activeFile: null,
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

      this.activeFile = node
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
      this.activeFile = node
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
        'id': this.activeFile.id,
        'content': this.content,
        'html': this.html
      }
      this.activeFile.title = this.title

      this.$log.debug('udpate tree', this.activeFile)

      saveDocument(doc).then(response => {
        this.activeFile.id = response.data.id
        if (this.title === '') {
          this.title = 'untitle'
          if (this.content.includes('\n')) {
            var title = this.content.split('\n')[0]
            if (title.includes('#')) {
              this.title = title.split('#')[1]
            }
          }
        }
        this.activeFile.metadata.title = this.title

        this.$log.debug('udpate tree', this.activeFile)

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
.navbar {
  width: 100%;
  height: 30px;
  padding: 0%;
  margin: 0%;
}
.editor-menu {
  display: flex;
  flex-direction: row;
  align-items: flex-end;
  height: 30px;
  /* transform: scale(2); */
}
.path {
  margin-top: 3vh;
}
.editor {
  /* float: left; */
  width: 100%;
  padding: 0px;
  /* margin: 10px; */
  margin-bottom: 40px;
  /* height: calc(100vh 30px); */
  height: 80vh;
  /* height: 90%; */
  /* height: 100px; */
  /* justify-content: left; */
  /* position: left; */
}
.app {
  display: flex;
  flex-direction: column;
  align-items: left;
  justify-content: flex-start;
}
</style>
