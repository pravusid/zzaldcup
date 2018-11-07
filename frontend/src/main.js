import '@babel/polyfill';
import Vue from 'vue';

import Toasted from 'vue-toasted';
import ImgInputer from 'vue-img-inputer';
import 'vue-img-inputer/dist/index.css';
import './plugins/vuetify';
import './libs/axios';

import App from './App.vue';
import router from './router';
import store from './store';

Vue.use(Toasted, {
  position: 'bottom-center',
  duration: 2000,
});
Vue.component('ImgInputer', ImgInputer);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
