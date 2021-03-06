import Vue from 'vue';
import Router from 'vue-router';

import Home from '@/views/Home.vue';
import NewMatch from '@/views/NewMatch.vue';
import NewCompetitors from '@/views/NewCompetitors.vue';
import UserMatches from '@/views/UserMatches.vue';

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
      path: '/match/edit/:matchName',
      name: 'newCompetitors',
      component: NewCompetitors,
    },
    {
      path: '/match/user',
      name: 'myMatches',
      component: UserMatches,
    },
  ],
  scrollBehavior() {
    return { x: 0, y: 0 };
  },
});
