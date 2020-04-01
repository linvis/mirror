<template>
  <v-toolbar dense width="100%" flat>
    <!-- <v-breadcrumbs :items="items" class="pa-0"></v-breadcrumbs> -->
    <v-chip
      v-if="activeNote.reminder.enable === false"
      label
      color="grey"
      outlined
    >
      Not Reviewed
    </v-chip>
    <v-chip v-else label color="grey" outlined>
      Reviewed
    </v-chip>

    <v-spacer></v-spacer>

    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <v-btn icon v-on="on" @click="handleClickHome">
          <v-icon>mdi-home</v-icon>
        </v-btn>
      </template>
      <span>home</span>
    </v-tooltip>

    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <v-btn icon v-on="on" @click="edit">
          <v-icon>mdi-pencil</v-icon>
        </v-btn>
      </template>
      <span>edit</span>
    </v-tooltip>

    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <v-btn icon v-on="on" @click="handleClickHome">
          <v-icon>mdi-calendar-clock</v-icon>
        </v-btn>
      </template>
      <span>start review</span>
    </v-tooltip>

    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <v-btn icon v-on="on" @click="handleClickHome">
          <v-icon>mdi-calendar-remove</v-icon>
        </v-btn>
      </template>
      <span>review later</span>
    </v-tooltip>

    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <v-btn icon v-on="on" @click="handleClickHome">
          <v-icon>mdi-calendar-check</v-icon>
        </v-btn>
      </template>
      <span>review done</span>
    </v-tooltip>
  </v-toolbar>
</template>

<script>
import { bus } from "@/utils/bus";

export default {
  data: () => ({
    showTagInput: false,
    newTag: null,
    showInput: false,
    dropItems: [
      { title: "Reviewd" },
      { title: "Later" },
      { title: "Start Review" }
    ]
  }),
  computed: {
    activeNote() {
      return this.$store.state.editor.activeNote;
    }
  },
  methods: {
    remove(item) {
      this.chips.splice(this.chips.indexOf(item), 1);
      this.chips = [...this.chips];
    },
    handleClickHome() {
      bus.$emit("editor-show", false);
      bus.$emit("reminder-show", true);
    },
    handleTagInput() {
      this.showTagInput = true;
      this.$refs.tagInput.$refs.input.focus();
    },
    edit() {
      bus.$emit("editor-edit");
    }
  }
};
</script>

<style>
/* .v-toolbar__content {
  padding: 0%;
} */
</style>
