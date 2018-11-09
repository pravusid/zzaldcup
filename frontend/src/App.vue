<template>
  <v-app>
    <v-toolbar dark color="blue-grey">
      <v-toolbar-title class="headline text-uppercase">
        <span>ZZALDCUP&nbsp;&nbsp;&nbsp;</span>
      </v-toolbar-title>
      <v-toolbar-items>
        <v-btn flat :to="{ path: '/' }" @click="key = Date.now()">
          <v-icon>list</v-icon>
          <span class="ml-2">전체 보기</span>
        </v-btn>
        <v-btn flat :to="{ path: '/match/new' }">
          <v-icon>add</v-icon>
          <span class="ml-2">새로 만들기</span>
        </v-btn>
        <v-btn flat :to="{ path: '/match/user' }" @click="key = Date.now()">
          <v-icon>assessment</v-icon>
          <span class="ml-2">전체 Match + 미완성 + private</span>
          <!-- <span class="ml-2">나의 Match</span> -->
        </v-btn>
      </v-toolbar-items>
      <v-spacer></v-spacer>
      <v-btn flat>
        <span class="mr-2">회원등록</span>
        <v-icon>open_in_new</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content>
      <router-view :key="key"></router-view>
    </v-content>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    key: null,
  }),
};
</script>

<style>
@import url("https://fonts.googleapis.com/css?family=Noto+Sans+KR:400,700");

#app {
  font-size: 1em;
  font-family: "Noto Sans KR";
}
.bold {
  font-weight: bold;
}
.inline {
  display: inline;
}
</style>
