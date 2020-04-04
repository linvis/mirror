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
        <v-btn icon v-on="on" @click="startReview">
          <v-icon>mdi-calendar-clock</v-icon>
        </v-btn>
      </template>
      <span>start review</span>
    </v-tooltip>

    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <v-btn icon v-on="on" @click="reviewLater">
          <v-icon>mdi-calendar-remove</v-icon>
        </v-btn>
      </template>
      <span>review later</span>
    </v-tooltip>

    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <v-btn icon v-on="on" @click="reviewDone">
          <v-icon>mdi-calendar-check</v-icon>
        </v-btn>
      </template>
      <span>review done</span>
    </v-tooltip>
  </v-toolbar>
</template>

<script>
import { bus } from "@/utils/bus";

const reviewDay = [1, 3, 5, 7, 14, 30];
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
    addDays(days) {
      var date = new Date();
      date.setDate(date.getDate() + days);
      return date.getTime();
    },
    remove(item) {
      this.chips.splice(this.chips.indexOf(item), 1);
      this.chips = [...this.chips];
    },
    handleClickHome() {
      this.$store.state.show.config.editor = !this.$store.state.show.config
        .editor;
      this.$store.state.show.config.reminder = !this.$store.state.show.config
        .reminder;
    },
    handleTagInput() {
      this.showTagInput = true;
      this.$refs.tagInput.$refs.input.focus();
    },
    edit() {
      bus.$emit("edit-file");
    },
    startReview() {
      this.activeNote.reminder.enable = true;
      this.activeNote.reminder.count = 1;
      this.activeNote.reminder.last_time = 0;
      this.activeNote.reminder.next_time = this.addDays(
        reviewDay[this.activeFile.reminder.count - 1]
      );
    },
    reviewLater() {
      this.activeNote.reminder.next_time = this.addDays(
        reviewDay[this.activeNote.reminder.count - 1]
      );
    },
    reviewDone() {
      this.activeNote.reminder.count++;
      if (this.activeNote.reminder.count >= reviewDay.length) {
        this.activeNote.reminder.count = reviewDay.length - 1;
      }
      this.activeNote.reminder.last_time = this.addDays(0);
      this.activeNote.reminder.next_time = this.addDays(
        reviewDay[this.activeNote.reminder.count - 1]
      );
      this.$store.dispatch("editor/updateCatalog");
      bus.$emit("update-reminder");
    }
  }
};
</script>

<style>
/* .v-toolbar__content {
  padding: 0%;
} */
</style>
