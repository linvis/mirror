<template>
  <v-container height="100%" width="100%" flex>
    <tool-bar />
    <reminder />
    <editor />
  </v-container>
</template>

<script>
import ToolBar from "./components/ToolBar";
import Editor from "./components/Editor";
import Reminder from "./components/Reminder";

export default {
  components: {
    ToolBar,
    Editor,
    Reminder
  },
  mounted() {
    this.$store.state.show.config.menu = false;
    this.$store.state.show.config.tree = true;
  }
};
</script>
