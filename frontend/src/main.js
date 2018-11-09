import '@babel/polyfill';
import Vue from 'vue';

import Toasted from 'vue-toasted';
import './plugins/vuetify';
import './libs/axios';

import App from './App.vue';
import router from './router';
import store from './store';

Vue.use(Toasted, {
  position: 'top-center',
  duration: 2000,
});

Vue.config.productionTip = false;

export default new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
