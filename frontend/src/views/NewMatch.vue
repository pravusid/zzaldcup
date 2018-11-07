<template>
  <v-container>
    <v-layout row wrap align-center justify-center>
      <v-flex xs12 sm12 md8>
        <h1>New Match</h1>
        <v-form>
          <v-text-field prepend-icon="label" label="대결명" v-model="match.matchName"
            type="text"></v-text-field>
          <v-select prepend-icon="more_vert" v-model="match.quota" :items="items"
            label="최대 개수"></v-select>
          <v-layout row>
            <v-flex xs3>
              <v-switch label="공개여부" v-model="match.isPublic"></v-switch>
            </v-flex>
            <v-flex xs9>
              <v-text-field label="접근주소" type="text" v-if="match.isPublic"></v-text-field>
            </v-flex>
          </v-layout>
        </v-form>
      </v-flex>
    </v-layout>
    <v-layout row wrap="" align-center justify-center>
      <v-flex xs12 sm12 md8>
        <v-btn color="primary" @click="save">생성</v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data: () => ({
    match: {},
    items: [16, 32, 64, 128],
  }),

  methods: {
    save() {
      axios.post('/api/match', this.match)
        .then((res) => {
          console.log(res);
          this.match = {};
        });
    },
  },
};
</script>

<style scoped>
</style>
