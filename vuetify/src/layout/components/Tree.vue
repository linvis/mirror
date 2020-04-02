<template>
  <v-container v-show="show">
    <v-toolbar flat>
      <v-btn icon class="hidden-xs-only" @click="handleBackClick">
        <v-icon>mdi-arrow-left</v-icon>
      </v-btn>

      <v-toolbar-title class="pa-0">Library</v-toolbar-title>
    </v-toolbar>

    <el-tree
      :data="data"
      :props="defaultProps"
      @node-click="handleNodeClick"
    ></el-tree>
    <!-- <v-treeview
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
    </v-treeview> -->

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
    menuItems: ["New Notebook", "New File", "Rename"],
    rightActiveItem: null,
    data: [
      {
        label: "一级 1",
        children: [
          {
            label: "二级 1-1",
            children: [
              {
                label: "三级 1-1-1"
              }
            ]
          }
        ]
      },
      {
        label: "一级 2",
        children: [
          {
            label: "二级 2-1",
            children: [
              {
                label: "三级 2-1-1"
              }
            ]
          },
          {
            label: "二级 2-2",
            children: [
              {
                label: "三级 2-2-1"
              }
            ]
          }
        ]
      },
      {
        label: "一级 3",
        children: [
          {
            label: "二级 3-1",
            children: [
              {
                label: "三级 3-1-1"
              }
            ]
          },
          {
            label: "二级 3-2",
            children: [
              {
                label: "三级 3-2-1"
              }
            ]
          }
        ]
      }
    ],
    defaultProps: {
      children: "children",
      label: "label"
    }
  }),
  methods: {
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

      //   const parent = item.parent;
      //   var data = node.data;
      //   const children = parent.data.children || parent.data;
      //   const index = children.findIndex(d => d.id === data.id);
      //   children.splice(index, 1);

      //   this.$store.dispatch("editor/submitCatalog", this.catalog);
    },
    openNote(item) {
      this.$log.debug(item);
      this.$store.state.editor.activeNote = item;
    }
  }
};
</script>

<style>
.el-tree-node__label {
  font-size: 16px;
}
.treeitem {
  height: 24px;
  width: 24px;
  margin: 6px;
  padding: 8px;
}
</style>
