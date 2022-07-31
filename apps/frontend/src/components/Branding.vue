<script setup lang="ts">
  import { computed } from 'vue';

  const props = defineProps<{
    theme: string
  }>();

  const src = computed(() => `/assets/images/branding-${props.theme}.svg`);
</script>

<template>
  <RouterLink to="/" class="branding">
    <img class="branding-item" :src="src" alt="GoShort.link branding">
  </RouterLink>
</template>