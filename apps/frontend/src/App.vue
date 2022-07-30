<script setup lang="ts">
  import { MoonIcon } from 'vue-tabler-icons';
  import Branding from '@app/components/Branding.vue';
  import IconButton from '@app/components/IconButton.vue';
</script>

<template>
  <div class="app-wrapper">
    <nav class="app-navigation">
      <Branding />
      <ul class="app-controls">
        <li class="app-controls-item">
          <IconButton>
            <MoonIcon />
          </IconButton>
        </li>
      </ul>
    </nav>
    <main class="app-main">
      <RouterView />
    </main>
    <footer class="app-footer">
      <p class="app-footer-copyright">Copyright &copy; {{ new Date().getFullYear() }} GoShort.link. All Rights Reserved</p>
      <ul class="app-footer-navigation">
        <li class="app-footer-navigation-item">
          <RouterLink to="/imprint" class="app-footer-navigation-link">Imprint</RouterLink>
          </li>
        <li class="app-footer-navigation-item">
          <RouterLink to="/privacy-policy" class="app-footer-navigation-link">Privacy Policy</RouterLink>
        </li>
      </ul>
    </footer>
  </div>
</template>
