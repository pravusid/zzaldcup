<template>
  <v-layout row wrap="" align-center>
    <v-flex xs1 text-xs-center>{{ idx + 1 }}</v-flex>
    <v-flex xs2 text-xs-center class="pa-3">
      <v-img :src="image" :lazy-src="image" aspect-ratio="1" class="grey lighten-2"></v-img>
    </v-flex>
    <v-flex xs9>
      <v-layout row>
        <v-flex xs9>
          <v-text-field prepend-icon="create" label="사진 설명" type="text"
            v-model="caption" @input="updateCompetitor"/>
        </v-flex>
        <v-flex xs1 text-xs-center>
          <v-icon color="green darken-2" size="35" v-if="!inprogress">done</v-icon>
          <v-progress-circular indeterminate color="red" size="25" v-else/>
        </v-flex>
        <v-flex xs2>
          <v-btn color="error" small @click="removeCompetitor">삭제</v-btn>
        </v-flex>
      </v-layout>
    </v-flex>
  </v-layout>
</template>

<script>
import debounce from 'lodash/fp/debounce';
import axios from '../libs/axios';

export default {
  props: [
    'idx',
    'competitor',
  ],

  mounted() {
    this.caption = this.competitor.caption;
  },

  data: () => ({
    caption: '',
    inprogress: false,
  }),

  methods: {
    updateCompetitor(val) {
      this.inprogress = true;
      this.debounceInput(val);
    },

    removeCompetitor() {
      axios.delete(`api/competitor/${this.competitor.id}`)
        .then(({ status }) => {
          if (status === 200) {
            this.$emit('remove', this.competitor);
          }
        }).catch(() => {
          this.$toasted.error('삭제 실패: 다시 시도해 주세요');
        });
    },

    // eslint-disable-next-line
    debounceInput: debounce(1000, function(val) {
      axios.put(`/api/competitor/${this.competitor.id}`, {
        id: this.competitor.id,
        caption: this.caption,
      }).then(({ status }) => {
        if (status === 200) this.$toasted.info('수정되었습니다');
        this.inprogress = false;
      }).catch((err) => {
        if (err.response) this.$toasted.error('수정중 오류 발생');
      });
    }),
  },

  computed: {
    image() {
      return `${process.env.VUE_APP_SERVER}/api/competitor/${this.competitor.imageUrl}`;
    },
  },
};
</script>
