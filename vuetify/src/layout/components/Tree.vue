<template>
  <v-container v-show="show">
    <v-toolbar flat>
      <v-btn icon class="hidden-xs-only" @click="handleBackClick">
        <v-icon>mdi-arrow-left</v-icon>
      </v-btn>

      <v-toolbar-title class="pa-0">Library</v-toolbar-title>
    </v-toolbar>
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
        >
          {{ item.metadata.title }}
        </div>
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
          @click="clickAction"
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
    menuItems: ["create file", "create directory"]
  }),
  methods: {
    handleBackClick() {
      this.$store.state.show.config.menu = true;
      this.$store.state.show.config.tree = false;
    },
    handleShowMenu(e, item) {
      this.$log.debug(e, item);
      e.preventDefault();
      this.showMenu = false;
      this.x = e.clientX;
      this.y = e.clientY;
      this.$nextTick(() => {
        this.showMenu = true;
      });
    },
    clickAction(e) {
      alert("clicked");
    },
    openNote(item) {
      this.$log.debug(item);
      this.$store.state.editor.activeNote = item;
    }
  }
};
</script>
