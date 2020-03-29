<template>
  <v-toolbar dense width="100%" flat>
    <v-breadcrumbs :items="items" class="pa-0"></v-breadcrumbs>

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
        <v-btn icon v-on="on" @click="handleClickHome">
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
    items: [
      {
        text: "Dashboard",
        disabled: false,
        href: "breadcrumbs_dashboard"
      },
      {
        text: "Link 1",
        disabled: false,
        href: "breadcrumbs_link_1"
      },
      {
        text: "Link 2",
        disabled: true,
        href: "breadcrumbs_link_2"
      }
    ],
    dropItems: [
      { title: "Reviewd" },
      { title: "Later" },
      { title: "Start Review" }
    ]
  }),
  methods: {
    handleClickHome() {
      bus.$emit("editor-show", false);
      bus.$emit("reminder-show", true);
    }
  }
};
</script>

<style>
/* .v-toolbar__content {
  padding: 0%;
} */
</style>
