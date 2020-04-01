<template>
  <v-container v-show="editorShow" height="100%" width="100%" class="pt-2 pl-4">
    <div id="editor" />
  </v-container>
</template>

<script>
import Prism from "prismjs";
import "katex/dist/katex.css";

Prism.highlightAll();

import Stackedit from "stackedit-js";
import { bus } from "@/utils/bus";

export default {
  data() {
    return {
      editorShow: true,
      eleContent: null,
      stackedit: null,
      title: "",
      html: "",
      contentText: "",
      observer: null
    };
  },
  created() {
    bus.$on("editor-show", this.updateShow);
    bus.$on("editor-edit", this.edit);
  },
  mounted() {
    this.initEditor();
    // this.renderCodeBlock();
  },
  beforeDestroy() {
    bus.$off("editor-show", this.updateShow);
    bus.$off("editor-edit", this.edit);
  },
  methods: {
    updateShow(show) {
      this.editorShow = show;
    },
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
    edit() {
      this.stackedit.openFile({
        name: "Filename", // with a filename
        content: {
          // text: this.contentText // and the Markdown content.
          text: "Hello **Markdown** writer!"
        }
      });
      this.stackedit.on("fileChange", file => {
        this.setHTML(file.content.html);
        this.contentText = file.content.text;
        console.log(this.contentText);
        console.log(file.content.html);
      });
      this.stackedit.on("close", () => {
        console.log("close");
        Prism.highlightAll();
        // this.eleContent.innerHTML = this.contentHTML;
      });
    }
  }
};
</script>

<style></style>
