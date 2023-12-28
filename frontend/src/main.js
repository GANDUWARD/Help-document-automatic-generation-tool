import { createApp } from 'vue'
import App from './App.vue'
//需要在此处挂在路由配置
createApp(App).use(App.router).mount('#app')
