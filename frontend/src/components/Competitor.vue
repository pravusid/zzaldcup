<template>
  <v-layout row wrap="" align-center class="my-5">
    <v-flex xs4 text-xs-center class="pa-4">
      <v-img
        :src="image"
        :lazy-src="image"
        aspect-ratio="1"
        class="grey lighten-2"
      ></v-img>
    </v-flex>
    <v-flex xs8>
      <v-text-field prepend-icon="create" label="사진 설명" type="text"></v-text-field>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  props: {
    image: String,
  },

  data: () => ({
    //
  }),
};
</script>
