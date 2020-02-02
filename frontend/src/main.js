import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import axios from 'axios';
import md5 from 'js-md5';
import Vuex from 'vuex'
import VueResource from 'vue-resource';

Vue.use(Vuex)
Vue.config.productionTip = false
Vue.use(VueResource)
Vue.prototype.$md5 = md5;

Vue.prototype.$ip = "127.0.0.1";
Vue.prototype.$port = "8000/api";


//开启debug模式
Vue.config.debug = true;
axios.defaults.withCredentials = true;
Vue.prototype.$axios = axios;


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
