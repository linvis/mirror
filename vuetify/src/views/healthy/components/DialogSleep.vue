<template>
  <v-dialog v-model="show" max-width="600px">
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
                    label="date"
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
            <v-col cols="12">
              <v-text-field
                v-model="duration"
                label="duration"
                prepend-icon="mdi-alarm"
                readonly
              ></v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="show = false">Close</v-btn>
        <v-btn color="blue darken-1" text @click="save">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import moment from "moment";
import { submitSleepRec } from "@/api/sleep";
import { bus } from "@/utils/bus";

export default {
  data: () => ({
    dateMenu: false,
    date: new Date().toISOString().substr(0, 10),
    startMenu: false,
    startTime: null,
    endMenu: false,
    endTime: null,
    dialog1: false,
    dialog2: false,
    dialog3: false,
    startTimeMin: 0,
    endTimeMin: 0,
    durationMin: 0,
    duration: "00:00"
  }),
  computed: {
    show: {
      get: function() {
        return this.$store.state.show.config.dialogSleep;
      },
      set: function(newVal) {
        this.$store.state.show.config.dialogSleep = newVal;
      }
    }
  },
  methods: {
    handleDateChange() {
      this.dateMenu = false;
      this.$refs.dialog1.save(this.date);

      this.$log.debug(typeof this.date);
    },

    handleStartChange() {
      this.startMenu = false;
      this.$refs.dialog2.save(this.startTime);

      var now = moment().format("YYYY-MM-DD");
      var time = moment(now + " " + this.startTime);
      this.startTimeMin = time.minute() + time.hour() * 60;

      this.$log.debug(this.startTimeMin);

      if (this.endTime === null) {
        return;
      }

      if (this.startTimeMin > this.endTimeMin) {
        this.startTimeMin = this.startTimeMin - 24 * 60;
      }

      var timeToString = function(time) {
        var hour = parseInt(time / 60);
        var min = parseInt(time % 60);
        return (
          parseInt(hour / 10).toString() +
          (hour % 10).toString() +
          ":" +
          parseInt(min / 10).toString() +
          (min % 10).toString()
        );
      };

      this.duration = this.endTimeMin - this.startTimeMin;
      this.duration = timeToString(this.duration);
    },
    handleEndChange() {
      this.endMenu = false;
      this.$refs.dialog3.save(this.endTime);

      var now = moment().format("YYYY-MM-DD");
      var time = moment(now + " " + this.endTime);
      this.endTimeMin = time.minute() + time.hour() * 60;

      this.$log.debug(this.endTimeMin);

      if (this.startTime === null) {
        return;
      }

      if (this.startTimeMin > this.endTimeMin) {
        this.startTimeMin = this.startTimeMin - 24 * 60;
      }

      var timeToString = function(time) {
        var hour = parseInt(time / 60);
        var min = parseInt(time % 60);
        return (
          parseInt(hour / 10).toString() +
          (hour % 10).toString() +
          ":" +
          parseInt(min / 10).toString() +
          (min % 10).toString()
        );
      };

      this.durationMin = this.endTimeMin - this.startTimeMin;
      this.duration = timeToString(this.durationMin);
    },
    save() {
      this.show = false;

      var day = moment(this.date);

      var utcDate = Date.UTC(day.year(), day.month(), day.date());

      var record = {
        record_date: utcDate,
        start_time: this.startTimeMin,
        end_time: this.endTimeMin,
        duration: this.durationMin
      };
      this.$log.debug(record);
      submitSleepRec(record).then(response => {
        bus.$emit("update-sleep-record");
        bus.$emit("update-sleep-analysis");
      });
    }
  }
};
</script>
