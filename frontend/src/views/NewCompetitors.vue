<template>
  <v-container>
    <h1 class="mb-4">Match #{{ match.matchName }}</h1>
    <file-pond
      name="zzal"
      ref="pond"
      label-idle="그림파일(들)을 여기에 드래그해서 놓아주세요"
      allow-multiple="true"
      accepted-file-types="image/jpeg, image/png"
      server="http://localhost:8080/api/competitor/image"
      :files="files"
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
      this.files.push(file);
      console.log(this.files);
    },
  },

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
