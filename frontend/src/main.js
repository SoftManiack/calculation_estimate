import Vue from 'vue'
import App from './App.vue'
import { BootstrapVue } from 'bootstrap-vue';
import store from '@/store/index';

Vue.config.productionTip = false

Vue.use(BootstrapVue);

new Vue({
  render: h => h(App),
  store,
}).$mount('#app')
