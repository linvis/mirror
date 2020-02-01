<template>
  <el-card class="box-card" shadow="hover">
    <div slot="header" class="clearfix">
      <span>Leetcode统计</span>
    </div>
    <!-- <el-row id="sleep_chart" /> -->
    <el-row id="cal-heatmap" v-loading="leetcodeLoading" style="margin-top:10px;margin-left:30px" />
  </el-card>
</template>

<script>

import { queryLeetcodeRec } from '@/api/program'
import CalHeatMap from 'cal-heatmap'
import 'cal-heatmap/cal-heatmap.css'
import moment from 'moment'

export default {
  data() {
    return {
      calHeatmap: null,
      calData: { '2020-01-27': 7, '2020-01-28': 1 },
      leetcodeLoading: true
      // calData: {}
    }
  },
  created() {
  },
  mounted() {
    this.featchData()
  },
  methods: {
    featchData() {
      queryLeetcodeRec().then(
        response => {
          console.log(response)
          var obj = JSON.parse(response)
          console.log(typeof (obj))
          this.calData = obj
          this.leetcodeLoading = false
          this.initHeatMap()
        },
        error => {
          this.leetcodeLoading = false
          this.initHeatMap()
        }
      )
    //   this.initHeatMap()
    },
    initHeatMap() {
      var dateOffset = (24 * 60 * 60 * 1000) * 365 // 365 days
      var myDate = new Date()
      myDate.setTime(myDate.getTime() - dateOffset)

      var currentMonth = myDate.getMonth()

      // this.calHeatmap = new CalHeatmap()
      // this.calHeatmap.init({})
      var cal = new CalHeatMap()
      cal.init({
        data: this.calData,
        // data: 'http://localhost:9528/dev-api/record/query/leetcode',
        itemSelector: '#cal-heatmap',
        itemName: ['submissions'],
        start: myDate, // January, 1st 2000
        dataType: 'json',
        range: 53,
        domain: 'week',
        tooltip: true,
        legend: [2, 4, 6, 8],
        domainLabelFormat: function(date) {
          if (date.getMonth() === currentMonth) {
            return ''
          }

          currentMonth = date.getMonth()

          return currentMonth + 1
        },
        subDomainDateFormat: function(date) {
          return moment(date).format('LL') // Use the moment library to format the Date
        }
      })
    }
  }
}
</script>

<style>
body {
  margin: 0;
}
</style>
