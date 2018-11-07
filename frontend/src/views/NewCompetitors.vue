<template>
  <v-layout row>
    <v-flex>
      <competitor></competitor>
    </v-flex>
  </v-layout>
</template>

<script>
import Competitor from '../components/Competitor.vue';

export default {
  components: {
    Competitor,
  },
};
</script>
