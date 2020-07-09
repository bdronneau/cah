import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css';
import VueLogger from 'vuejs-logger';
import { getConfig } from './api/config'
import '@mdi/font/css/materialdesignicons.css';

Vue.config.productionTip = false;

const isProduction = process.env.NODE_ENV === 'production';
const options = {
  isEnabled: true,
  logLevel : isProduction ? 'error' : 'debug',
  stringifyArguments : false,
  showLogLevel : true,
  showMethodName : true,
  separator: '|',
  showConsoleColors: true
};

Vue.use(VueLogger, options);

getConfig().then(data =>
  store.commit('updateConfig', data)
).finally(() => {
  new Vue({
    router,
    store,
    vuetify,
    render: (h) => h(App),
  }).$mount('#app');
})
