<template>
  <v-app id="inspire" flex>
    <header-bar />

    <side-bar ref="side" />

    <v-content ref="app">
      <v-container class="fill-height" fluid align-start>
        <app-main height="100%" width="100%" />
      </v-container>
    </v-content>
    <v-btn bottom color="pink" dark fab fixed right>
      <v-icon>mdi-plus</v-icon>
    </v-btn>
  </v-app>
</template>

<script>
import AppMain from "./components/AppMain";
import HeaderBar from "./components/HeaderBar";
import SideBar from "./components/SideBar";

// import { bus } from "@/utils/bus";

export default {
  name: "Layout",
  components: {
    AppMain,
    HeaderBar,
    SideBar
  },
  props: {
    source: String
  },
  data() {
    return {
      sideWidth: null
      // appWidth: "calc(100vh-256px)"
    };
  },
  computed: {
    // sideWidth() {
    //   return this.sideWidth;
    // }
  },
  watch: {
    sideWidth: function(newVal, oldVal) {
      this.$nextTick(() => {
        console.log(newVal, oldVal);
      });
    }
  },
  created() {
    // bus.$on("nav-width-change", this.updateAppWidth);
  },
  mounted() {
    this.sideWidth = this.$refs.side.$el.style.width;
    console.log(this.sideWidth);
  },
  beforeDestroy() {
    // bus.$off("nav-width-change", this.updateAppWidth);
  },
  methods: {
    updateAppWidth(width) {
      this.appWidth = "calc(100vh-" + width + ")";
      console.log(width);
      console.log(this.appWidth);
    }
  }
};
</script>
