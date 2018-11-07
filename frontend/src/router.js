import Vue from 'vue';
import Router from 'vue-router';

import Home from '@/views/Home.vue';
import NewMatch from '@/views/NewMatch.vue';
import NewCompetitors from '@/views/NewCompetitors.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/match/new',
      name: 'newMatch',
      component: NewMatch,
    },
    {
      path: '/competitor/new',
      name: 'NewCompetitors',
      component: NewCompetitors,
    },
  ],
});
