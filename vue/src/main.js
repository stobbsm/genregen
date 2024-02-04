import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import axios from 'axios'
import VueAxios from 'axios-vue'

loadFonts()

createApp(App)
  .use(vuetify)
  .use(VueAxios, axios)
  .mount('#app')
