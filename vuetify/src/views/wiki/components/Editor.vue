<template>
  <v-container v-show="show" height="100%" width="100%" class="pt-2 pl-4">
    <v-row>
      <v-col cols="22">
        <div id="editor" />
      </v-col>
      <v-col cols="2">
        <!-- <v-navigation-drawer ref="drawer" :right="true" clipped :width="256">
          <div id="toc"></div>
          <ol>
            <a href="#hh" class="toc-level1">
              hh
            </a>
          </ol>
          <ol>
            <a href="#hh" class="toc-level2">
              hh
            </a>
          </ol>
        </v-navigation-drawer> -->
        <toc />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Prism from "prismjs";

import "@/styles/base.scss";

Prism.highlightAll();

import Stackedit from "stackedit-js";
import { bus } from "@/utils/bus";
import { queryDocumentByID } from "@/api/editor";
import { saveDocument } from "@/api/editor";

import Toc from "./Toc";

export default {
  components: {
    Toc
  },
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

    // <h1 id="aaa"><a class="markdownIt-Anchor" href="#aaa">#</a> AAA</h1>
    // <h1 id="bbb"><a class="markdownIt-Anchor" href="#bbb">#</a> BBB</h1>
    // <h1 id="ccc"><a class="markdownIt-Anchor" href="#ccc">#</a> CCC</h1>
    // var a = markdownIt({
    //   html: true,
    //   linkify: true,
    //   typographer: true
    // })
    //   .use(markdownItTocAndAnchor, {
    //     // ...options
    //     toc
    //   })
    //   .render("# AAA\n# BBB\n# CCC\nfoo\nbar\nbaz");

    // var toc = document.getElementById("toc");
    // toc.innerHTML = a;

    // console.log(a);
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
    generateToc(content) {
      var regex = /#+.*\n/g;
      // content = "# AAA\n## BBB\n### CCC\nfoo\nbar\nbaz ## DD\n"
      var tocList = content.match(regex);
      if (tocList.length <= 0) {
        return;
      }
    },
    openFile(item) {
      queryDocumentByID(item.id).then(response => {
        var docs = response.data;
        this.$log.debug("item", item);
        this.$log.debug("update markdown from remote", docs.content);
        // this.open(item.metadata.title, docs.content);
        this.setHTML(docs.content);
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

<style>
.toc-level1 {
  padding-left: 4px;
}
.toc-level2 {
  padding-left: 8px;
}
.toc-level3 {
  padding-left: 12px;
}
</style>
