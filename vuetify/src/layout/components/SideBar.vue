<template>
  <v-navigation-drawer
    v-model="drawer"
    :clipped="$vuetify.breakpoint.lgAndUp"
    app
  >
    <v-list dense>
      <template v-for="item in routes">
        <div v-if="isMenuItemValid(item)" :key="item.name">
          <v-list-group
            v-if="isMoreChildren(item)"
            :key="item.name"
            :value="false"
            :prepend-icon="item.meta.icon"
          >
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title>{{ item.meta.text }}</v-list-item-title>
              </v-list-item-content>
            </template>
            <v-list-item v-for="(child, i) in item.children" :key="i" link>
              <v-list-item-action v-if="child.meta.icon">
                <v-icon>{{ child.meta.icon }}</v-icon>
              </v-list-item-action>
              <v-list-item-content>
                <v-list-item-title>{{ child.meta.text }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-group>
          <v-list-item
            v-else
            :key="item.name"
            :to="getFullPath(item, item.children[0])"
            link
          >
            <v-list-item-action>
              <v-icon>{{ item.children[0].meta.icon }}</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>{{
                item.children[0].meta.text
              }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </div>
      </template>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  name: "SideBar",
  computed: {
    routes() {
      console.log(this.$router.options.routes);
      return this.$router.options.routes;
    }
  },
  props: {
    source: String
  },
  data: () => ({
    dialog: false,
    drawer: null
  }),
  methods: {
    isMenuItemValid(item) {
      if (item.hidden) {
        return false;
      } else {
        return true;
      }
    },
    isMoreChildren(item) {
      if (item.children.length > 1) {
        return true;
      } else {
        return false;
      }
    },
    getFullPath(item, child) {
      var urljoin = require("url-join");
      console.log(urljoin(item.path, child.path));
      return urljoin(item.path, child.path);
    }
  }
};
</script>
