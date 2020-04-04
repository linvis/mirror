<template>
  <v-container>
    <div slot="header" class="clearfix">
      <span>Leetcode统计</span>
    </div>
    <div
      id="cal-heatmap"
      style="margin-top:10px;margin-left:30px; max-width: 50%;"
    />
  </v-container>
</template>

<script>
import { queryLeetcodeRec } from "@/api/program";
import CalHeatMap from "cal-heatmap";
import "cal-heatmap/cal-heatmap.css";
import moment from "moment";

export default {
  data() {
    return {
      calData: { "2020-01-27": 7, "2020-01-28": 1 },
      transition: "scale-transition"
      // calData: {}
    };
  },
  created() {},
  mounted() {
    this.featchData();
  },
  methods: {
    featchData() {
      queryLeetcodeRec().then(
        response => {
          var obj = JSON.parse(response);
          this.calData = obj;
          this.initHeatMap();
        },
        error => {
          this.initHeatMap();
        }
      );
    },
    initHeatMap() {
      var dateOffset = 24 * 60 * 60 * 1000 * 365; // 365 days
      var myDate = new Date();
      myDate.setTime(myDate.getTime() - dateOffset);

      var currentMonth = myDate.getMonth();

      var cal = new CalHeatMap();
      cal.init({
        data: this.calData,
        itemSelector: "#cal-heatmap",
        itemName: ["submissions"],
        start: myDate, // January, 1st 2000
        dataType: "json",
        range: 53,
        domain: "week",
        tooltip: true,
        legend: [2, 4, 6, 8],
        domainLabelFormat: function(date) {
          if (date.getMonth() === currentMonth) {
            return "";
          }

          currentMonth = date.getMonth();

          return currentMonth + 1;
        },
        subDomainDateFormat: function(date) {
          return moment(date).format("LL"); // Use the moment library to format the Date
        }
      });
    }
  }
};
</script>
