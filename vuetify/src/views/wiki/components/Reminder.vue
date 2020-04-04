<template>
  <v-data-table
    :headers="headers"
    :items="reminderList"
    :sort-by="['overdue']"
    :sort-desc="[true]"
    class="elevation-1"
    v-show="show"
  >
    <template v-slot:top>
      <v-toolbar flat color="white">
        <v-toolbar-title>Review</v-toolbar-title>
        <v-divider class="mx-4" inset vertical></v-divider>
        <v-spacer></v-spacer>
      </v-toolbar>
    </template>
    <template v-slot:item.actions="{ item }">
      <!-- <v-icon small class="mr-2" @click="openFile(item)" color="red">
        "mdi-close-circle-outline"
      </v-icon>
      <v-icon @click="openFile(item)" color="green">
        "mdi-check-circle-outline"
      </v-icon> -->
      <v-icon
        small
        v-show="!item.checked"
        class="mr-2"
        @click="openFile(item)"
        color="red"
      >
        mdi-close-circle-outline
      </v-icon>
      <v-icon small v-show="item.checked" @click="openFile(item)" color="green">
        mdi-check-circle-outline
      </v-icon>
    </template>
  </v-data-table>
</template>

<script>
import { bus } from "@/utils/bus";
import moment from "moment";

export default {
  data: () => ({
    headers: [
      {
        text: "File",
        align: "start",
        sortable: false,
        value: "name"
      },
      { text: "Overdue", value: "overdue" },
      { text: "Reviewd-time", value: "reviewed" },
      { text: "Actions", value: "actions", sortable: false }
    ],
    reminderList: []
  }),
  computed: {
    show() {
      return this.$store.state.show.config.reminder;
    },
    catalog() {
      return this.$store.state.editor.catalog;
    }
  },

  watch: {
    catalog: function(newValue, oldValue) {
      this.reminderList = this.filterReminder(newValue);
    }
  },

  created() {
    bus.$on("update-reminder", this.updateReminder);
    this.updateReminder();
  },

  beforeDestroy() {
    bus.$off("update-reminder", this.updateReminder);
  },

  methods: {
    updateReminder() {
      this.reminderList = this.filterReminder(this.catalog);
      this.$log.debug(this.reminderList);
    },
    diffDay(timestamp) {
      if (timestamp === "0") {
        return 1;
      }
      var reviewTime = moment(parseInt(timestamp));
      var today = moment();

      return reviewTime.diff(today, "days");
    },
    filterReminder(catalog) {
      var list = [];

      if (catalog === null || typeof catalog === "undefined") {
        return list;
      }

      this.$log.debug(catalog);
      for (var i = 0; i < catalog.length; i++) {
        if (catalog[i].reminder.enable === true) {
          var next = this.diffDay(catalog[i].reminder.next_time);
          var last = this.diffDay(catalog[i].reminder.last_time);

          if (next <= 0) {
            list.push({
              name: catalog[i].metadata.title,
              overdue: -next,
              reviewed: catalog[i].reminder.count,
              file: catalog[i],
              checked: false
            });
          }

          if (last === 0) {
            list.push({
              name: catalog[i].metadata.title,
              overdue: 0,
              reviewed: catalog[i].reminder.count,
              file: catalog[i],
              checked: true
            });
          }
        }
        if ("children" in catalog[i]) {
          var subList = this.filterReminder(catalog[i].children);
          list = list.concat(subList);
        }
      }

      return list;
    },

    openFile(item) {
      this.$log.debug(item);
      this.$store.state.show.config.editor = !this.$store.state.show.config
        .editor;
      this.$store.state.show.config.reminder = !this.$store.state.show.config
        .reminder;
      bus.$emit("open-file", item.file);
    }
  }
};
</script>
