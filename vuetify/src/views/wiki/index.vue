<template>
  <!-- <div id="editor" height="100%" width="100%" /> -->
  <v-container height="100%" width="100%">
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
  }
};
</script>
