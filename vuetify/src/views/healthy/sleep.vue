<template>
  <v-container>
    <sleep-chart />
    <sleep-analysis />
  </v-container>
</template>

<script>
import SleepChart from "./components/SleepChart";
import SleepAnalysis from "./components/SleepAnalysis";

export default {
  components: {
    SleepChart,
    SleepAnalysis
  }
};
</script>
