<template>
  <v-container>
    <leetcode />
    <github />
  </v-container>
</template>

<script>
import Leetcode from "./components/Leetcode";
import Github from "./components/Github";

export default {
  components: {
    Leetcode,
    Github
  }
};
</script>
