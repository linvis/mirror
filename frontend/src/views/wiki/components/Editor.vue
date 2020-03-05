<template>
  <el-container v-show="show" v-loading="loading" style="height:100%;">
    <el-header class="head">
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
          <el-button type="primary" plain size="small" @click="saveDoc">save</el-button>
          <el-button type="primary" plain size="small" @click="edit">edit</el-button>
        </div>

        <el-dialog title="markdown" :visible.sync="dialogFormVisible">
          <el-input
            v-model="markdownText"
            type="textarea"
            :autosize="{ minRows: 8 }"
            placeholder="pleas input text"
          />
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="mdToHTML">确 定</el-button>
          </div>
        </el-dialog>
      </el-col>
      <!-- </el-row> -->
    </el-header>

    <el-main class="main">
      <div class="editor">
        <editor-menu-bar v-slot="{ commands, isActive }" :editor="editor">
          <div class="menubar">
            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.bold() }"
              @click="commands.bold"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_bold" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.italic() }"
              @click="commands.italic"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_italic" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.strike() }"
              @click="commands.strike"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_strike" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.underline() }"
              @click="commands.underline"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_underline" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.code() }"
              @click="commands.code"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_code" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.paragraph() }"
              @click="commands.paragraph"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_paragraph" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.heading({ level: 1 }) }"
              @click="commands.heading({ level: 1 })"
            >H1</button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.heading({ level: 2 }) }"
              @click="commands.heading({ level: 2 })"
            >H2</button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.heading({ level: 3 }) }"
              @click="commands.heading({ level: 3 })"
            >H3</button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.bullet_list() }"
              @click="commands.bullet_list"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_ul" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.ordered_list() }"
              @click="commands.ordered_list"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_ol" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.blockquote() }"
              @click="commands.blockquote"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_quote" />
              </span>
            </button>

            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.code_block() }"
              @click="commands.code_block"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_code" />
              </span>
            </button>

            <button class="menubar__button" @click="commands.horizontal_rule">
              <span class="svg-container">
                <svg-icon icon-class="editor_hr" />
              </span>
            </button>

            <button class="menubar__button" @click="commands.undo">
              <span class="svg-container">
                <svg-icon icon-class="editor_undo" />
              </span>
            </button>

            <button class="menubar__button" @click="commands.redo">
              <span class="svg-container">
                <svg-icon icon-class="editor_redo" />
              </span>
            </button>
            <button
              class="menubar__button"
              :class="{ 'is-active': isActive.todo_list() }"
              @click="commands.todo_list"
            >
              <span class="svg-container">
                <svg-icon icon-class="editor_checklist" />
              </span>
            </button>
          </div>
        </editor-menu-bar>
        <editor-content class="editor__content" :editor="editor" />
      </div>
    </el-main>
  </el-container>
</template>

<script>
import Stackedit from 'stackedit-js'
import { Editor, EditorContent, EditorMenuBar } from 'tiptap'
import {
  Placeholder,
  Bold,
  Blockquote,
  CodeBlock,
  HardBreak,
  Heading,
  HorizontalRule,
  OrderedList,
  BulletList,
  ListItem,
  TodoItem,
  TodoList,
  Code,
  Italic,
  Link,
  Strike,
  Underline,
  History,
  Image
} from 'tiptap-extensions'
import Doc from './Doc'
import Title from './Title'

import { saveDocument, queryAllDocument, queryDocumentByID } from '@/api/editor'
import { bus } from '@/utils/bus'

export default {
  components: {
    EditorContent,
    EditorMenuBar
  },
  data() {
    return {
      show: false,
      loading: false,
      editor: null,
      activeNode: null,
      title: '',
      contents: '',
      dialogFormVisible: false,
      markdownText: ''
    }
  },
  created() {
    bus.$on('show-editor', this.showEditor)
    bus.$on('open-document', this.openDocument)
    bus.$on('open-new-doc', this.openNewDoc)
  },
  mounted() {
    this.initEditor()
    // this.getAllDoc()
  },
  beforeDestroy() {
    bus.$off('show-editor', this.showEditor)
    bus.$off('open-document', this.openDocument)
    bus.$off('open-new-doc', this.openNewDoc)
    this.editor.destroy()
  },
  methods: {
    initEditor() {
      this.editor = new Editor({
        autoFocus: true,
        extensions: [
          new Doc(),
          new Title(),
          new Placeholder({
            showOnlyCurrent: false,
            emptyNodeText: node => {
              if (node.type.name === 'title') {
                return 'Title'
              }

              return 'Write something'
            }
          }),
          new Bold(),
          new Blockquote(),
          new BulletList(),
          new CodeBlock(),
          new HardBreak(),
          new Heading({ levels: [1, 2, 3] }),
          new HorizontalRule(),
          new ListItem(),
          new OrderedList(),
          new TodoItem({
            nested: true
          }),
          new TodoList(),
          new Link(),
          new Bold(),
          new Code(),
          new Italic(),
          new Strike(),
          new Underline(),
          new History(),
          new Image()
        ]
      })
      this.editor.on('update', ({ getJSON, getHTML }) => {
        // get new content on update
        const json = getJSON()
        if ('text' in json.content[0].content[0]) {
          this.title = json.content[0].content[0].text
        }

        const newContent = getHTML()
        this.contents = newContent
      })
    },
    setContent(content) {
      this.$log.debug(content)
      this.contents = content
      this.editor.setContent(content)
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
      this.setContent(`
          <h2>
            Todo List
          </h2>
          <p>
            There is always something to do. Thankfully, there are checklists for that. Don't forget to call mom.
          </p>
          <ul data-type="todo_list">
            <li data-type="todo_item" data-done="true">
              Buy beer
            </li>
            <li data-type="todo_item" data-done="true">
              Buy meat
            </li>
            <li data-type="todo_item" data-done="true">
              Buy milk
            </li>
            <li data-type="todo_item" data-done="false">
              Call mom
            </li>
          </ul>
        `)
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
      this.dialogFormVisible = true
    },
    saveDoc(event) {
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
    },
    mdToHTML(event) {
      this.dialogFormVisible = false

      const stackedit = new Stackedit()
      stackedit.openFile({
        name: 'Filename',
        content: { text: this.markdownText }
      }, true /* silent mode */)

      stackedit.on('fileChange', (file) => {
        this.setContent(file.content.html)
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
@import "~@/styles/editor.scss";
@import "~@/styles/main.scss";
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
  justify-content: left;
}
// .header {
//   display: flex;
//   padding: 20px;
// }

// .main {
//   display: flex;
//   justify-content: center;
//   // padding: 50px;
// }
ul[data-type="todo_list"] {
  padding-left: 0;
}

li[data-type="todo_item"] {
  display: flex;
  flex-direction: row;
}

.todo-checkbox {
  border: 2px solid $color-black;
  height: 0.9em;
  width: 0.9em;
  box-sizing: border-box;
  margin-right: 10px;
  margin-top: 0.3rem;
  user-select: none;
  -webkit-user-select: none;
  cursor: pointer;
  border-radius: 0.2em;
  background-color: transparent;
  transition: 0.4s background;
}

.todo-content {
  flex: 1;
  > p:last-of-type {
    margin-bottom: 0;
  }
  > ul[data-type="todo_list"] {
    margin: 0.5rem 0;
  }
}

li[data-done="true"] {
  > .todo-content {
    > p {
      text-decoration: line-through;
    }
  }
  > .todo-checkbox {
    background-color: $color-black;
  }
}

li[data-done="false"] {
  text-decoration: none;
}
</style>
