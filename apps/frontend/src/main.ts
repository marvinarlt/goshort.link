import { createApp } from 'vue';
import { createRouter, createWebHistory } from 'vue-router'
import App from '@app/App.vue';
import "@app/scss/main.scss";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: () => import('@app/pages/Index.vue')
    },
    {
      path: '/imprint',
      component: () => import('@app/pages/Imprint.vue')
    },
    {
      path: '/privacy-policy',
      component: () => import('@app/pages/PrivacyPolicy.vue')
    }
  ]
});

createApp(App)
  .use(router)
  .mount('#app');
