<template>
  <v-container>
    <v-tabs @change="handleTabClick">
      <v-tab name="first">睡眠时长</v-tab>
      <v-tab name="second">入睡时间</v-tab>
      <v-tab name="third">苏醒时间</v-tab>
    </v-tabs>
    <v-row id="sleep_chart" />
  </v-container>
</template>

<script>
import Highcharts from "highcharts";
import stockInit from "highcharts/modules/stock";
stockInit(Highcharts);

import { querySleepRec } from "@/api/sleep";
import { bus } from "@/utils/bus";

export default {
  data() {
    return {
      activeTab: 0,
      sleep_chart: null,
      dur_data: [],
      start_time_data: [],
      end_time_data: [],
      chart_data: [],
      chart_label: "sleep about"
    };
  },
  created() {
    bus.$on("update-sleep-record", this.getSleepRecord);
  },
  mounted() {
    this.getSleepRecord();
  },
  beforeDestroy() {
    bus.$off("update-sleep-record", this.getSleepRecord);
  },
  methods: {
    copyData(ori) {
      var target = [];
      for (var i = 0; i < ori.length; i++) {
        target.push(ori[i]);
      }

      return target;
    },
    getSleepRecord() {
      querySleepRec().then(response => {
        this.dur_data = [];
        this.start_time_data = [];
        this.end_time_data = [];

        for (var i = 0; i < response.data.date.length; i++) {
          this.dur_data.push([
            response.data.date[i],
            response.data.duration[i]
          ]);
          this.start_time_data.push([
            response.data.date[i],
            response.data.start_time[i] + 24 * 60
          ]);
          this.end_time_data.push([
            response.data.date[i],
            response.data.end_time[i] + 24 * 60
          ]);
        }

        this.chart_data = Object.assign([], this.dur_data);
        this.$log.debug(
          "print time",
          this.dur_data,
          this.start_time_data,
          this.end_time_data
        );
        this.$log.debug(this.chart_data);

        this.initChart();
      });
    },
    handleTabClick(tab) {
      console.log(tab, this.activeTab);
      if (this.activeTab === tab) {
        // do nothing
        return;
      } else if (tab === 0) {
        this.chart_data = Object.assign([], this.dur_data);
        this.chart_label = "sleep about";
      } else if (tab === 1) {
        this.chart_data = Object.assign([], this.start_time_data);
        this.chart_label = "goto sleep at";
      } else if (tab === 2) {
        this.chart_data = Object.assign([], this.end_time_data);
        this.chart_label = "wakeup at";
      }

      this.$log.debug(this.chart_data);

      this.activeTab = tab;
      this.sleep_chart.series[0].update(
        {
          data: this.chart_data
        },
        false
      );
      this.sleep_chart.redraw();
    },
    timeToString(time) {
      var hour = parseInt(time / 60);
      var min = parseInt(time % 60);
      return (
        parseInt(hour / 10).toString() +
        (hour % 10).toString() +
        ":" +
        parseInt(min / 10).toString() +
        (min % 10).toString()
      );
    },
    initChart() {
      this.sleep_chart = Highcharts.stockChart("sleep_chart", {
        rangeSelector: {
          enabled: true,
          allButtonsEnabled: true,
          buttons: [
            {
              type: "week",
              count: 1,
              text: "last 7 days"
            },
            {
              type: "week",
              count: 2,
              text: "last 14 days"
            },
            {
              type: "month",
              text: "Month"
            }
          ],
          buttonTheme: {
            width: 60
          },
          selected: 0
        },

        title: {
          text: "Sleep Analysis"
        },

        yAxis: {
          labels: {
            formatter: function() {
              this.value = this.value % (24 * 60);
              var hour = parseInt(this.value / 60);
              var min = parseInt(this.value % 60);
              return (
                parseInt(hour / 10).toString() +
                (hour % 10).toString() +
                ":" +
                parseInt(min / 10).toString() +
                (min % 10).toString()
              );
            }
          },
          title: {
            text: "Hours"
          },
          opposite: false
        },

        tooltip: {
          formatter: function() {
            // yesterday
            var tmpY = this.y % (24 * 60);
            var hour = parseInt(tmpY / 60);
            var min = parseInt(tmpY % 60);
            var time =
              parseInt(hour / 10).toString() +
              (hour % 10).toString() +
              ":" +
              parseInt(min / 10).toString() +
              (min % 10).toString();
            return "sleep about " + time;
          }
        },

        series: [
          {
            name: "sleep time",
            data: this.chart_data
          }
        ]
      });
    }
  }
};
</script>
