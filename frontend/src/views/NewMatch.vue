<template>
  <v-container>
    <v-layout row wrap="" align-center justify-center>
      <v-flex xs12 sm12 md8>
        <h1>New Match</h1>
        <v-form>
          <v-text-field prepend-icon="label" label="대결명" v-model="match.matchName"
            type="text"></v-text-field>
          <span class="red--text">{{ msg }}</span>
          <v-select prepend-icon="more_vert" v-model="match.quota" :items="items"
            label="최대 개수"></v-select>
          <v-switch label="비공개" color="primary" v-model="match.private"></v-switch>
        </v-form>
      </v-flex>
    </v-layout>
    <v-layout row wrap="" align-center justify-center>
      <v-flex xs12 sm12 md8 text-xs-center>
        <v-progress-circular indeterminate color="green" v-if="submitting"></v-progress-circular>
        <v-btn color="primary" @click="save" v-else>생성</v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import axios from 'axios';
import _ from 'lodash';

export default {
  data: () => ({
    match: {},
    items: [16, 32, 64, 128],
    msg: '',
    submitting: false,
  }),

  methods: {
    save() {
      this.submitting = true;
      if (this.msg) {
        this.$toasted.error('다른 이름을 선택해 주세요');
        return;
      }
      axios.post('/api/match', this.match)
        .then((res) => {
          if (res.status === 201) {
            const { matchName } = this.match;
            this.match = {};
            this.$toasted.success('생성 되었습니다');
            this.$router.push(`/match/edit/${matchName}`);
          }
        }).catch(() => {
          this.$toasted.error('잘못된 값을 입력했습니다');
        }).then(() => {
          this.submitting = false;
        });
    },
  },

  watch: {
    // eslint-disable-next-line
    'match.matchName': _.debounce(function (val) {
      if (val === undefined) return;
      const duplicate = '이미 존재하는 이름입니다';
      axios.get(`/api/match/${this.match.matchName}`)
        .then(() => {
          this.msg = duplicate;
        }).catch(({ response }) => {
          this.msg = (response.status === 404) ? '' : duplicate;
        });
    }, 500),
  },
};
</script>

<style scoped>
</style>
