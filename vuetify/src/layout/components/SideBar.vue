<template>
  <v-navigation-drawer
    ref="drawer"
    v-model="navigation.show"
    :clipped="$vuetify.breakpoint.lgAndUp"
    app
    :width="navigation.width"
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
              <v-list-item-title>
                {{ item.children[0].meta.text }}
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </div>
      </template>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
import { bus } from "@/utils/bus";

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
    navigation: {
      show: true,
      width: 256,
      borderSize: 3
    }
  }),
  created() {
    bus.$on("nav-show", this.updateNavShow);
  },
  mounted() {
    this.setBorderWidth();
    this.setEvents();
  },
  beforeDestroy() {
    bus.$off("nav-show", this.updateNavShow);
  },
  methods: {
    updateNavShow(show) {
      this.navigation.show = show;
    },
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
      return urljoin(item.path, child.path);
    },
    setBorderWidth() {
      let i = this.$refs.drawer.$el.querySelector(
        ".v-navigation-drawer__border"
      );
      i.style.width = this.navigation.borderSize + "px";
      i.style.cursor = "ew-resize";
      // i.style.backgroundColor = "#e7e7e7";
    },
    setEvents() {
      const minSize = this.navigation.borderSize;
      const el = this.$refs.drawer.$el;
      const drawerBorder = el.querySelector(".v-navigation-drawer__border");
      const direction = el.classList.contains("v-navigation-drawer--right")
        ? "right"
        : "left";

      function resize(e) {
        document.body.style.cursor = "ew-resize";
        let f =
          direction === "right"
            ? document.body.scrollWidth - e.clientX
            : e.clientX;
        el.style.width = f + "px";
      }

      drawerBorder.addEventListener(
        "mousedown",
        e => {
          if (e.offsetX < minSize) {
            el.style.transition = "initial";
            document.addEventListener("mousemove", resize, false);
          }
          drawerBorder.style.backgroundColor = "blue";
        },
        false
      );

      document.addEventListener(
        "mouseup",
        () => {
          el.style.transition = "";
          this.navigation.width = el.style.width;
          document.body.style.cursor = "";
          document.removeEventListener("mousemove", resize, false);
          drawerBorder.style.backgroundColor = "#e7e7e7";

          // bus.$emit("nav-width-change", this.navigation.width);
        },
        false
      );
    }
  }
};
</script>
