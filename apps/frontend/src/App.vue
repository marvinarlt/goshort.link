<script setup lang="ts">
  import { MoonIcon } from 'vue-tabler-icons';
  import Branding from '@app/components/Branding.vue';
  import IconButton from '@app/components/IconButton.vue';
  import TheLocaleSelect from '@app/components/TheLocaleSelect.vue';
  import Notifications from '@app/components/Notifications.vue';
  import { useTheme } from '@app/composables/theme';

  const { theme, toggle } = useTheme();
</script>

<template>
  <div class="app-wrapper">
    <nav class="app-navigation">
      <Branding :theme="theme" />
      <ul class="app-controls">
        <li class="app-controls-item">
          <IconButton @click="toggle">
            <MoonIcon />
          </IconButton>
        </li>
        <li class="app-controls-item">
          <TheLocaleSelect />
        </li>
      </ul>
    </nav>
    <main class="app-main">
      <router-view />
    </main>
    <footer class="app-footer">
      <p class="app-footer-copyright">&copy; {{ new Date().getFullYear() }} GoShort.link</p>
      <ul class="app-footer-navigation">
        <li class="app-footer-navigation-item">
          <RouterLink to="/legal-notice" class="app-footer-navigation-link" :title="$t('navigation.legal-notice.title')">{{ $t('navigation.legal-notice.label') }}</RouterLink>
          </li>
        <li class="app-footer-navigation-item">
          <RouterLink to="/privacy-policy" class="app-footer-navigation-link" :title="$t('navigation.privacy-policy.label')">{{ $t('navigation.privacy-policy.label') }}</RouterLink>
        </li>
      </ul>
    </footer>
  </div>
  <Notifications />
</template>
