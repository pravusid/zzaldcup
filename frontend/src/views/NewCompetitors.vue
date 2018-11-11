<template>
  <v-container>
    <h1 class="mb-4">Match #{{ match.matchName }}</h1>
    <file-pond
      name="zzal"
      ref="pond"
      label-idle="그림파일(들)을 여기에 드래그해서 놓아주세요"
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
      server="http://localhost:8080/api/competitor/image"
      @init="init"
      @processfile="processFile"
    />
    <competitor/>
    <v-layout>
      <v-flex xs12 text-xs-center>
        <v-btn color="primary" large>전송</v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import axios from 'axios';

import vueFilePond from 'vue-filepond';
import 'filepond/dist/filepond.min.css';
import FilePondPluginFileValidateType from 'filepond-plugin-file-validate-type';
import FilePondPluginImagePreview from 'filepond-plugin-image-preview';
import 'filepond-plugin-image-preview/dist/filepond-plugin-image-preview.min.css';

import Vue from '../main';
import Competitor from '../components/Competitor.vue';

export default {
  components: {
    FilePond: vueFilePond(FilePondPluginFileValidateType, FilePondPluginImagePreview),
    Competitor,
  },

  data: () => ({
    match: {},
    files: [],
  }),

  methods: {
    init() {
      console.log(this.$refs.pond.getFiles());
    },

    processFile(error, file) {
      if (error) return;
      this.files.push(file);
      setTimeout(() => {
        this.$refs.pond.removeFile(file);
      }, 1000);
    },
  },

  // Router 객체에 위치하면 next() cb 작동하지 않음 (2.5.17 bug)
  beforeRouteEnter(to, from, next) {
    axios.get(`/api/match/${to.params.matchName}`)
      .then((res) => {
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
