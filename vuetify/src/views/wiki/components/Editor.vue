<template>
  <v-container v-show="show" height="100%" width="100%" class="pt-2 pl-4">
    <div id="editor" />
  </v-container>
</template>

<script>
import Prism from "prismjs";
import "katex/dist/katex.css";

Prism.highlightAll();

import Stackedit from "stackedit-js";
import { bus } from "@/utils/bus";
import { queryDocumentByID } from "@/api/editor";
import { saveDocument } from "@/api/editor";

export default {
  data() {
    return {
      editorShow: true,
      eleContent: null,
      stackedit: null,
      title: "",
      html: "",
      content: "",
      observer: null
    };
  },
  computed: {
    activeNote: {
      get() {
        return this.$store.state.editor.activeNote;
      },
      set(newVal) {
        this.$store.state.editor.activeNote = newVal;
      }
    },
    show() {
      return this.$store.state.show.config.editor;
    }
  },
  created() {
    bus.$on("open-file", this.openFile);
    bus.$on("open-new-file", this.openNewFile);
    bus.$on("edit-file", this.edit);
  },
  mounted() {
    this.initEditor();
  },
  beforeDestroy() {
    bus.$off("open-file", this.openFile);
    bus.$off("open-new-file", this.openNewFile);
    bus.$off("edit-file", this.edit);
  },
  methods: {
    initEditor() {
      this.stackedit = new Stackedit();
      this.eleContent = document.getElementById("editor");
      this.eleContent.innerHTML = "Hello Jarvis！";
      // this.edit();
    },
    setHTML(html) {
      this.html = html;
      this.eleContent.innerHTML = html;
    },
    updateContent(content) {
      this.contents = content;
    },
    openFile(item) {
      queryDocumentByID(item.id).then(response => {
        var docs = response.data;
        this.$log.debug("item", item);
        this.$log.debug("update markdown from remote", docs.content);
        this.open(item.metadata.title, docs.content);
        this.$store.state.editor.contents[item.key] = docs.content;
      });
      this.activeNote = item;
      //   this.open(item.metadata.title, "Hello **Markdown** writer!");
      //   this.$store.state.editor.contents[item.key] =
      //     "Hello **Markdown** writer!";
    },
    openNewFile(item) {
      this.activeNote = item;
      this.open(item.metadata.title, "");
    },
    edit() {
      var item = this.activeNote;
      var content = this.$store.state.editor.contents[item.key];
      this.open(item.metadata.title, content);
    },
    open(file, content) {
      this.stackedit.openFile({
        name: file, // with a filename
        content: {
          text: content // and the Markdown content.
          // text: "Hello **Markdown** writer!"
        }
      });
      this.stackedit.on("fileChange", file => {
        this.setHTML(file.content.html);
        this.content = file.content.text;
        console.log(this.content);
        console.log(file.content.html);
      });
      this.stackedit.on("close", () => {
        console.log("close");
        Prism.highlightAll();
        this.$store.state.editor.contents[this.activeNote.key] = this.content;
        // this.eleContent.innerHTML = this.contentHTML;
        this.save();
      });
    },
    save() {
      var content = this.$store.state.editor.contents[this.activeNote.key];
      var doc = {
        id: this.activeNote.id,
        doc_key: this.activeNote.key,
        content: content,
        html: this.html
      };

      saveDocument(doc).then(response => {
        this.activeNote.id = response.data.id;
        this.$store.dispatch("editor/updateCatalog");
        // var title = this.activeFile.metadata.title;
        // if (this.activeFile.metadata.title === "untitle") {
        //   if (content.includes("\n")) {
        //     title = content.split("\n")[0];
        //     if (title.includes("#")) {
        //       title = title.split("#")[1];
        //     }
        //   }
        // }
        // this.activeFile.metadata.title = title;

        // this.$notify({
        //   title: "成功",
        //   message: "",
        //   type: "success",
        //   duration: 800
        // });
      });
    }
  }
};
</script>

<style></style>
