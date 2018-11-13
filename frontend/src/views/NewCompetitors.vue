<template>
  <v-container>
    <h1 class="mb-4">
      Match: {{ match.matchName }}
      <small>( {{ match.competitors.length }} / {{ match.quota }})</small>
    </h1>
    <file-pond
      v-if="uploadable"
      name="zzal"
      ref="pond"
      label-idle="이 곳을 클릭 하거나, 그림파일(들)을 여기에 드래그해서 놓아주세요"
      label-tap-to-cancel="눌러서 취소"
      label-tap-to-retry="눌러서 재시도"
      label-tap-to-undo="눌러서 되돌리기"
      label-file-processing="업로드 중"
      label-file-processing-error="업로드 중 오류 발생"
      label-file-processing-complete="업로드 완료"
      label-button-remove-item="삭제"
      label-button-retry-item-processing="재시도"
      allow-multiple="true"
      allow-revert="false"
      accepted-file-types="image/jpeg, image/png"
      :server="host"
      @addfile="addFile"
      @processfile="processFile"
    />
    <competitor v-for="(item, index, key) in match.competitors"
      v-bind:idx="index" :key="key" :competitor="item" v-on:remove="removeCompetitor"/>
  </v-container>
</template>

<script>
import _ from 'lodash/fp';
import debounce from 'lodash/fp/debounce';

import vueFilePond from 'vue-filepond';
import 'filepond/dist/filepond.min.css';
import FilePondPluginFileValidateType from 'filepond-plugin-file-validate-type';
import FilePondPluginImagePreview from 'filepond-plugin-image-preview';
import 'filepond-plugin-image-preview/dist/filepond-plugin-image-preview.min.css';

import axios from '../libs/axios';
import Vue from '../main';
import Competitor from '../components/Competitor.vue';

export default {
  components: {
    FilePond: vueFilePond(FilePondPluginFileValidateType, FilePondPluginImagePreview),
    Competitor,
  },

  data: () => ({
    host: `${process.env.VUE_APP_SERVER}/api/competitor/image`,
    match: {
      competitors: [],
    },
    cursor: 0,
    files: [],
  }),

  methods: {
    addFile(error, file) {
      if (!this.uploadable) {
        this.$toasted.error('더 이상 추가할 수 없습니다');
        this.$refs.pond.removeFile(file);
        return;
      }
      file.setMetadata('matchId', this.match.id);
    },

    processFile(error, file) {
      if (error) return;
      this.files.push(file);
      setTimeout(() => {
        this.$refs.pond.removeFile(file);
      }, 2000);
    },

    removeCompetitor(val) {
      this.match.competitors.splice(this.match.competitors.indexOf(val), 1);
      this.$toasted.success('삭제되었습니다');
    },
  },

  computed: {
    uploadable() {
      return this.match.competitors.length < this.match.quota;
    },
  },

  watch: {
    // eslint-disable-next-line
    'files': debounce(2000, function() {
      const criteria = _.flow([
        _.map(c => JSON.parse(this.files[c].serverId).id),
        _.min,
      ])(_.range(this.cursor, this.files.length));

      axios.get('/api/competitor', {
        params: {
          matchId: this.match.id,
          id: criteria,
        },
      }).then((res) => {
        if (res.status === 200) {
          _.forEach(c => this.match.competitors.push(c), res.data);
          const added = res.data.length;
          this.cursor += added;
          this.$toasted.info(`사진 ${added}개가 추가되었습니다`);
        }
      });
    }),
  },

  // Router 객체에 위치하면 next()의 cb 작동하지 않음 (2.5.17 bug)
  beforeRouteEnter(to, from, next) {
    axios.get(`/api/match/detail/${to.params.matchName}`, {
      params: {
        related: true,
      },
    }).then((res) => {
      if (res.status === 200) {
        next((vm) => {
          const data = vm.$data;
          data.match = res.data;
        });
      }
    }).catch(() => {
      Vue.$toasted.error('자료 또는 권한이 없습니다');
      Vue.$router.push('/');
    });
  },
};
</script>
