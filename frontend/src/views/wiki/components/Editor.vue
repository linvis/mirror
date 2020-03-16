<template>
  <el-container v-show="show" v-loading="loading" class="app">
    <el-container class="navbar">
      <el-col :span="16">
        <tag />
      </el-col>
      <el-col :span="5" :offset="3" class="editor-menu">
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

import { queryDocumentByID } from '@/api/editor'
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
      var key = this.$store.state.editor.activeFile.key
      var content = changes.markdown
      if (content === '\n') {
        return
      }
      this.$log.debug('markdown callback', key, content)
      this.$store.dispatch('editor/changeFileContent', { key, content })
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
    setContent(key, content) {
      this.$log.debug('markdown set content')
      this.$store.dispatch('editor/changeFileContent', { key, content })
      this.editor.setMarkdown(content)
    },
    openDoc(file) {
      this.$log.debug(file)

      if (file.metadata.filetype !== 'md') {
        return
      }

      this.$store.dispatch('editor/changeActiveFile', file)
      this.loading = true

      var localContents = this.$store.state.editor.contents
      this.$log.debug('local', localContents, file.key in localContents)
      if (file.key in localContents) {
        this.editor.setMarkdown(localContents[file.key])
      } else {
        queryDocumentByID(file.id).then(response => {
          var docs = response.data
          this.$log.debug('update markdown from remote', docs.content)
          this.setContent(file.key, docs.content)
        })
      }

      this.loading = false
    },
    openNewDoc(file) {
      this.$store.dispatch('editor/changeActiveFile', file)
      this.setContent(file.key, '')
      this.showEditor(true)
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
