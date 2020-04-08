<template>
  <v-navigation-drawer
    ref="drawer"
    v-model="show"
    :clipped="$vuetify.breakpoint.lgAndUp"
    app
    :width="navigation.width"
  >
    <item />
    <tree />
  </v-navigation-drawer>
</template>

<script>
import Item from "./Item";
import Tree from "./Tree";

export default {
  name: "SideBar",
  components: {
    Item,
    Tree
  },
  data: () => ({
    navigation: {
      width: 320,
      borderSize: 3
    }
  }),
  computed: {
    show: {
      get: function() {
        return this.$store.state.show.config.nav;
      },
      set: function(newVal) {
        this.$store.state.show.config.nav = newVal;
      }
    }
  },
  mounted() {
    // this.setBorderWidth();
    // this.setEvents();
  },
  methods: {
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
