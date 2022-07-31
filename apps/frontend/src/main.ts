import { createApp } from 'vue';
import { createI18n } from 'vue-i18n';
import { resolveLocale } from '@app/services/i18n';
import { createRouter, createWebHistory } from 'vue-router'
import en from '@app/locales/en.json';
import de from '@app/locales/de.json';
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
      path: '/legal-notice',
      component: () => import('@app/pages/LegalNotice.vue')
    },
    {
      path: '/privacy-policy',
      component: () => import('@app/pages/PrivacyPolicy.vue')
    }
  ]
});

const i18n = createI18n({
  locale: resolveLocale(),
  fallbackLocale: 'en',
  globalInjection: true,
  messages: { en, de }
});

createApp(App)
  .use(router)
  .use(i18n)
  .mount('#app');
