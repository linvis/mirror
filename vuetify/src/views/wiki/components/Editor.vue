<template>
  <v-container v-show="editorShow" height="100%" width="100%" class="pt-2 pl-4">
    <div id="editor" />
  </v-container>
</template>

<script>
import Stackedit from "stackedit-js";
import { bus } from "@/utils/bus";

export default {
  data() {
    return {
      editorShow: true,
      eleContent: null,
      stackedit: null,
      activeNode: null,
      title: "",
      contentHTML: "",
      contentText: ""
    };
  },
  created() {
    bus.$on("editor-show", this.updateShow);
  },
  mounted() {
    this.initEditor();
  },
  beforeDestroy() {
    bus.$off("editor-show", this.updateShow);
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
    setContent(html) {
      this.contentHTML = html;
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
        // el.value = file.content.text
        this.contentHTML = file.content.html;
        this.contentText = file.content.text;
        console.log(file.content.text);
        console.log(file.content.html);
      });
      this.stackedit.on("close", () => {
        console.log("close");
        this.eleContent.innerHTML = this.contentHTML;
        console.log(this.contentText);
        console.log(this.contentHTML);
      });
    }
  }
};
</script>

<style>
.editor {
  /* padding-left: 16px; */
}
</style>
