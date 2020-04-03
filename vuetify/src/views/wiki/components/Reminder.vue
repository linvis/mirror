<template>
  <v-data-table
    :headers="headers"
    :items="desserts"
    sort-by="overdue"
    class="elevation-1"
    v-show="show"
  >
    <template v-slot:top>
      <v-toolbar flat color="white">
        <v-toolbar-title>View</v-toolbar-title>
        <v-divider class="mx-4" inset vertical></v-divider>
        <v-spacer></v-spacer>
      </v-toolbar>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
      <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
    </template>
  </v-data-table>
</template>

<script>
import { bus } from "@/utils/bus";

export default {
  data: () => ({
    reminderShow: false,
    dialog: false,
    headers: [
      {
        text: "File",
        align: "start",
        sortable: false,
        value: "name"
      },
      { text: "Overdue", value: "overdue" },
      { text: "Next-time", value: "nextTime" },
      { text: "Actions", value: "actions", sortable: false }
    ],
    desserts: [],
    editedIndex: -1
  }),
  computed: {
    show() {
      return this.$store.state.show.config.reminder;
    }
  },

  created() {
    this.initialize();
  },

  beforeDestroy() {},

  methods: {
    initialize() {
      this.desserts = [
        {
          name: "Frozen Yogurt",
          overdue: 2,
          nextTime: 6.0
        }
      ];
    },

    editItem(item) {
      this.editedIndex = this.desserts.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true;
    },

    deleteItem(item) {
      const index = this.desserts.indexOf(item);
      confirm("Are you sure you want to delete this item?") &&
        this.desserts.splice(index, 1);
    },

    close() {
      this.dialog = false;
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1;
      }, 300);
    },

    save() {
      if (this.editedIndex > -1) {
        Object.assign(this.desserts[this.editedIndex], this.editedItem);
      } else {
        this.desserts.push(this.editedItem);
      }
      this.close();
    }
  }
};
</script>
