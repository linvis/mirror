<template>
  <v-app id="inspire">
    <header-bar />

    <side-bar />

    <v-content>
      <v-container class="fill-height" fluid>
        <app-main />
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import AppMain from "./components/AppMain";
import HeaderBar from "./components/HeaderBar";
import SideBar from "./components/SideBar";

export default {
  name: "Layout",
  components: {
    AppMain,
    HeaderBar,
    SideBar
  },
  props: {
    source: String
  }
};
</script>
