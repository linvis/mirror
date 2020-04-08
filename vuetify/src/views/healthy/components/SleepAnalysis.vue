<template>
  <v-container>
    <v-row id="sleep_analysis_chart" />
  </v-container>
</template>

<script>
import moment from "moment";
import Highcharts from "highcharts";
import stockInit from "highcharts/modules/stock";
import { querySleepRecAnalysis } from "@/api/sleep";
import { bus } from "@/utils/bus";
stockInit(Highcharts);

export default {
  data() {
    return {
      sleep_analysis_chart: null,
      avgLastWeek: [],
      avgLastTwoWeek: [],
      avgLastMonth: []
    };
  },
  created() {
    bus.$on("update-sleep-analysis", this.getSleepAnalysis);
  },
  mounted() {
    this.getSleepAnalysis();
  },
  destroyed() {
    bus.$on("update-sleep-analysis", this.getSleepAnalysis);
  },
  methods: {
    formatTime(duration) {
      // var time = new Date('2020, 0, 1, 0, 0, 0').getTime()
      var time = Date.UTC(2020, 0, 1, 0, 0, 0);
      if (duration < 0) {
        time = time + duration * 60 * 1000 + 24 * 60 * 60 * 1000;
      } else {
        time = time + duration * 60 * 1000;
      }
      return time;
    },
    getSleepAnalysis() {
      querySleepRecAnalysis().then(response => {
        this.avgLastWeek = [
          this.formatTime(response.data.duration[0]),
          this.formatTime(response.data.start_time[0]),
          this.formatTime(response.data.end_time[0])
        ];
        this.avgLastTwoWeek = [
          this.formatTime(response.data.duration[1]),
          this.formatTime(response.data.start_time[1]),
          this.formatTime(response.data.end_time[1])
        ];
        this.avgLastMonth = [
          this.formatTime(response.data.duration[2]),
          this.formatTime(response.data.start_time[2]),
          this.formatTime(response.data.end_time[2])
        ];

        this.$log.debug(this.avgLastWeek);
        this.$log.debug(this.avgLastTwoWeek);
        this.$log.debug(this.avgLastMonth);
        this.initChart();
      });
    },
    initChart() {
      this.sleep_analysis_chart = Highcharts.chart("sleep_analysis_chart", {
        chart: {
          type: "bar"
        },
        title: {
          text: "Sleep Analysis"
        },
        xAxis: {
          categories: ["平均睡眠时长", "平均入睡时间", "平均苏醒时间"],
          title: {
            text: null
          }
        },
        yAxis: {
          min: Date.UTC(2020, 0, 1, 0, 0, 0),
          max: Date.UTC(2020, 0, 2, 0, 0, 0),
          type: "datetime",
          tickPositioner: function() {
            var info = this.tickPositions.info;
            var positions = [];
            for (
              var i = Date.UTC(2020, 0, 1, 0, 0, 0);
              i <= Date.UTC(2020, 0, 2, 0, 0, 0);
              i += 3600 * 1000
            ) {
              positions.push(i);
            }
            positions.info = info;
            return positions;
          },
          lineWidth: 1,
          dateTimeLabelFormats: {
            day: "%H:%M"
          },

          title: {
            text: "Hours"
          }
        },

        tooltip: {
          formatter: function() {
            return moment.utc(this.y).format("HH:MM");
          }
        },
        series: [
          {
            name: "Last Week",
            data: this.avgLastWeek
          },
          {
            name: "Last Two Week",
            data: this.avgLastTwoWeek
          },
          {
            name: "Last Month",
            data: this.avgLastMonth
          }
        ]
      });
    }
  }
};
</script>
