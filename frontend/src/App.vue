<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>ZZALDCUP&nbsp;&nbsp;&nbsp;</span>
      </v-toolbar-title>
      <v-toolbar-items>
        <v-btn flat :to="'/'">
          <span class="mr-2">전체 보기</span>
          <v-icon>list</v-icon>
        </v-btn>
        <v-btn flat :to="'/match/new'">
          <span class="mr-2">새로 만들기</span>
          <v-icon>note_add</v-icon>
        </v-btn>
      </v-toolbar-items>
      <v-spacer></v-spacer>
      <v-btn flat>
        <span class="mr-2">Contact us</span>
        <v-icon>open_in_new</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    //
  }),
};
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Noto+Sans+KR:400,700');

#app {
  font-size: 1em;
  font-family: 'Noto Sans KR';
}
.bold {
  font-weight: bold;
}
.inline {
  display: inline;
}
</style>
