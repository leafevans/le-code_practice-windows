import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import router from './router'

import ECharts from 'vue-echarts'
import 'echarts'

const app = createApp(App)

app.use(router)
app.use(ElementPlus)

app.component('v-chart', ECharts)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.mount('#app')
