<template>
  <v-container>
    <v-btn bottom color="pink" dark fab fixed right @click="handleDialogClick">
      <v-icon>mdi-plus</v-icon>
    </v-btn>

    <dialog-sleep />
  </v-container>
</template>

<script>
import DialogSleep from "@/views/healthy/components/DialogSleep";

export default {
  components: {
    DialogSleep
  },

  methods: {
    handleDialogClick() {
      if (this.$store.state.menu.activeMenu === "Sleep") {
        this.$store.state.show.config.dialogSleep = true;
      }
    }
  }
};
</script>
