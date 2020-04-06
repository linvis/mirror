<template>
  <v-container v-show="show">
    <v-toolbar flat>
      <v-btn icon class="hidden-xs-only" @click="handleBackClick">
        <v-icon>mdi-arrow-left</v-icon>
      </v-btn>

      <v-toolbar-title class="pa-0">Library</v-toolbar-title>
    </v-toolbar>

    <!-- <el-tree
      ref="tree"
      :data="catalog"
      :props="defaultProps"
      class="tree"
      :highlight-current="true"
      :expand-on-click-node="true"
      :draggable="isDrag"
      @node-contextmenu="handleShowMenu"
      @node-click="handleNodeClick"
      @node-drop="handleDrop"
    >
      <span slot-scope="{ node, data }" class="custom-tree-node">
        <span v-show="!node.isEdit">
          <v-icon v-if="data.metadata.filetype === 'nb'">mdi-folder</v-icon>
          <v-icon v-else>mdi-note-text</v-icon>
          <span style="margin-left:10px;">{{ data.metadata.title }}</span>
        </span>
        <v-text-field
          :ref="data.key"
          v-model="data.metadata.title"
          autofocus
          @blur.stop="NodeBlur(node)"
          @keyup.enter.native="NodeBlur(node)"
          v-show="node.isEdit"
          single-line
        ></v-text-field>
      </span>
    </el-tree> -->
    <v-treeview
      v-model="tree"
      :open="open"
      :items="catalog"
      activatable
      item-key="key"
      open-on-click
    >
      <template v-slot:label="{ item }">
        <div
          @click="openNote(item)"
          @contextmenu="handleShowMenu($event, item)"
          v-show="!item.isEdit"
        >
          {{ item.metadata.title }}
        </div>

        <v-text-field
          :ref="item.key"
          v-model="item.metadata.title"
          autofocus
          @blur.stop="NodeBlur(item)"
          @keyup.enter.native="NodeBlur(item)"
          v-show="item.isEdit"
          single-line
        ></v-text-field>
      </template>
      <template v-slot:prepend="{ item, open }">
        <v-icon v-if="item.metadata.filetype === 'nb'">
          {{ open ? "mdi-folder-open" : "mdi-folder" }}
        </v-icon>
        <v-icon v-else>
          {{ "mdi-note-text" }}
        </v-icon>
      </template>
    </v-treeview>

    <v-menu
      v-model="showMenu"
      :position-x="x"
      :position-y="y"
      absolute
      offset-y
    >
      <v-list>
        <v-list-item
          v-for="menuItem in menuItems"
          :key="menuItem"
          @click="clickAction($event, menuItem)"
        >
          <v-list-item-title>{{ menuItem }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-container>
</template>

<script>
import { v4 as uuidv4 } from "uuid";
import { bus } from "@/utils/bus";
import { queryEditorCatalog } from "@/api/editor";
export default {
  computed: {
    show: {
      get: function() {
        return this.$store.state.show.config.tree;
      },
      set: function(newVal) {
        this.$store.state.show.config.tree = newVal;
      }
    },
    catalog: {
      get: function() {
        return this.$store.state.editor.catalog;
      },
      set: function(newVal) {
        this.$store.state.editor.catalog = newVal;
      }
    }
  },
  data: () => ({
    open: ["public"],
    active: [],
    tree: [],
    showMenu: false,
    x: 0,
    y: 0,
    menuItems: ["New Notebook", "New File", "Rename", "Delete"],
    rightActiveItem: null
  }),
  mounted() {
    this.getEditorCatalog();
    bus.$on("get-catalog", this.getEditorCatalog);
  },
  beforeDestroy() {
    bus.$off("get-catalog", this.getEditorCatalog);
  },
  methods: {
    getEditorCatalog() {
      queryEditorCatalog().then(response => {
        // this.catalog = response.data
        this.$store.dispatch("editor/changeCatalog", response.data);
      });
      // bus.$emit('get-catalog', this.catalog)
    },
    newRandomKey() {
      var uuid = uuidv4();
      return uuid.split("-").join("");
    },
    findItem(items, key) {
      for (var i = 0; i < items.length; i++) {
        if (items[i].key === key) {
          return items[i];
        } else {
          var item = this.findItem(items[i].children, key);
          if (item !== null) {
            return item;
          }
        }
      }

      return null;
    },
    handleBackClick() {
      this.$store.state.show.config.menu = true;
      this.$store.state.show.config.tree = false;
    },
    handleShowMenu(e, item) {
      this.$log.debug(e, item);
      this.rightActiveItem = item;
      e.preventDefault();
      this.showMenu = false;
      this.x = e.clientX;
      this.y = e.clientY;
      this.$nextTick(() => {
        this.showMenu = true;
      });
    },
    clickAction(e, menuItem) {
      if (menuItem === "Rename") {
        this.handleRename(this.rightActiveItem);
      } else if (menuItem === "Delete") {
        this.handleDelete(this.rightActiveItem);
      } else if (menuItem === "New Notebook") {
        this.handleNewNoteBook(this.rightActiveItem);
      } else if (menuItem === "New File") {
        this.handleNewFile(this.rightActiveItem);
      } else if (menuItem === "Delete") {
        this.handleDelete(this.rightActiveItem);
      }
    },
    NodeBlur(item) {
      // 输入框失焦
      this.$log.debug("lose focus", item);

      if (item.isEdit) {
        this.$set(item, "isEdit", false);
      }

      this.$store.dispatch("editor/updateCatalog", this.catalog);
    },
    handleNewFile(item) {
      this.$log.debug(item);

      var data = item;
      var key = this.newRandomKey();

      const newChild = {
        id: 0,
        key: key,
        parent: data.key,
        metadata: {
          title: "untitle",
          level: data.level + 1,
          filetype: "md",
          tag: []
        },
        reminder: {
          enable: false,
          count: 0,
          last_time: 0,
          next_time: 0
        },
        children: []
      };
      if (!data.children) {
        this.$set(data, "children", []);
      }
      data.children.push(newChild);

      this.$store.dispatch("editor/updateCatalog", this.catalog);

      //   bus.$emit("show-reminder", false);
      bus.$emit("open-new-file", newChild);
    },
    handleNewNoteBook(item) {
      var data = item;

      if (data.metadata.level >= 2) {
        this.$message({
          message: "仅支持2级nodebook",
          type: "warning"
        });
        return;
      }

      var key = this.newRandomKey();

      const newChild = {
        id: 0,
        key: key,
        parent: data.key,
        metadata: {
          title: "nodebook",
          level: data.level + 1,
          filetype: "nb",
          tag: []
        },
        reminder: {
          enable: false,
          count: 0,
          last_time: 0,
          next_time: 0
        },
        children: []
      };
      if (!data.children) {
        this.$set(data, "children", []);
      }
      data.children.push(newChild);

      //   this.$store.dispatch("editor/updateCatalog", this.catalog);
    },
    handleRename(item) {
      this.$log.debug(item);

      if (!item.isEdit) {
        this.$set(item, "isEdit", true);
      }
      this.$nextTick(() => {
        this.$refs[item.key].$refs.input.focus();
      });
    },
    handleDelete(item) {
      this.$log.debug(item);

      var parent = this.findItem(this.catalog, item.parent);
      const children = parent.children || parent;
      const index = children.findIndex(d => d.id === item.id);
      children.splice(index, 1);

      //   this.$store.dispatch("editor/updateCatalog", this.catalog);
    },
    openNote(item) {
      this.$log.debug(item);
      if (item.metadata.filetype === "md") {
        bus.$emit("open-file", item);
      }
    },
    handleNodeClick(data, node) {
      // node edit toggle click
      this.$log.debug(node);
      if (
        Object.prototype.hasOwnProperty.call(node, "isEdit") === true &&
        node.isEdit === true
      ) {
        return;
      }

      if (data.metadata.filetype === "md") {
        // bus.$emit("show-reminder", false);
        // bus.$emit("show-editor", true);
        // bus.$emit("open-document", data);
      }
    },
    handleDrop(draggingNode, dropNode, dropType, ev) {
      this.$store.dispatch("editor/submitCatalog", this.catalog);
    }
  }
};
</script>

<style>
.el-tree-node__label {
  font-size: 16px;
}
.el-tree-node__content {
  margin: 4px;
  padding: 8px;
}
</style>
