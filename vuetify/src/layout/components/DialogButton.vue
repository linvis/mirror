<template>
  <v-container>
    <v-btn bottom color="pink" dark fab fixed right @click="handleDialogClick">
      <v-icon>mdi-plus</v-icon>
    </v-btn>

    <v-dialog v-model="dialog" max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Sleep Record</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-dialog
                  ref="dialog1"
                  v-model="dateMenu"
                  :return-value.sync="date"
                  width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="date"
                      label="date1"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="date"
                    @change="handleDateChange"
                    scrollable
                  >
                  </v-date-picker>
                </v-dialog>
              </v-col>

              <v-col cols="12">
                <v-dialog
                  ref="dialog2"
                  v-model="startMenu"
                  :return-value.sync="startTime"
                  width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="startTime"
                      label="start time"
                      prepend-icon="mdi-alarm"
                      readonly
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                    v-if="startMenu"
                    v-model="startTime"
                    format="24hr"
                    full-width
                    @change="handleStartChange"
                  >
                  </v-time-picker>
                </v-dialog>
              </v-col>

              <v-col cols="12">
                <v-dialog
                  ref="dialog3"
                  v-model="endMenu"
                  :return-value.sync="endTime"
                  width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="endTime"
                      label="end time"
                      prepend-icon="mdi-alarm"
                      readonly
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                    v-if="endMenu"
                    v-model="endTime"
                    format="24hr"
                    full-width
                    @change="handleEndChange"
                  >
                  </v-time-picker>
                </v-dialog>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="dialog = false"
            >Close</v-btn
          >
          <v-btn color="blue darken-1" text @click="dialog = false">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    dateMenu: false,
    date: new Date().toISOString().substr(0, 10),
    startMenu: false,
    startTime: null,
    endMenu: false,
    endTime: null,
    dialog: false,
    dialog2: false,
    dialog3: false
  }),
  methods: {
    handleDialogClick() {
      this.dialog = true;
    },
    handleDateChange() {
      this.dateMenu = false;
      this.$refs.dialog.save(this.date);
    },
    handleStartChange() {
      this.startMenu = false;
      this.$refs.dialog2.save(this.startTime);
    },
    handleEndChange() {
      this.endMenu = false;
      this.$refs.dialog3.save(this.endTime);
      console.log(this.endTime);
    }
  }
};
</script>
