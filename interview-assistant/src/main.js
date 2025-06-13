import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)

// 全局错误处理
app.config.errorHandler = (err, vm, info) => {
  console.error('全局错误:', err);
  console.error('错误信息:', info);
}

app.use(router)
app.mount('#app')