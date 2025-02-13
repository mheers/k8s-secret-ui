import './assets/main.css'


import { createApp } from 'vue'
import App from './App.vue'

// Vuetify
import "@mdi/font/css/materialdesignicons.css";
import 'vuetify/styles'
import { createVuetify, type ThemeDefinition } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'


const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: 'light'
    },
    icons: {
        defaultSet: "mdi",
    },
})


createApp(App).use(vuetify).mount('#app')
