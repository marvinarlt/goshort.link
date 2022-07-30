<script setup lang="ts">
  import { ref } from 'vue';
  import { CopyIcon } from 'vue-tabler-icons';
  import IconButton from '@app/components/IconButton.vue';

  const backendOrigin = import.meta.env.VITE_BACKEND_ORIGIN;

  const linkInput = ref('');
  const link = ref(new URL('/RaNdOmId', backendOrigin));
  const isGenerated = ref(false);

  /**
   * Check if string is url.
   * 
   * @param {string} url
   * @returns {boolean}
   */
  const isUrl = (url: string): boolean => {
    try {
      new URL(url);
    } catch {
      return false;
    }

    return true;
  }

  /**
   * Send url to api to receive an id.
   * 
   * @async
   * @returns {Promise<void>}
   */
  const onSubmit = async (): Promise<void> => {
    if (! isUrl(linkInput.value)) {
      console.log('Invalid URL');
      return;
    }

    try {
      const backendUrl = new URL('/api/link', backendOrigin);
      const response = await fetch(backendUrl, {
        method: 'POST',
        body: JSON.stringify({
          url: linkInput.value
        })
      });
      const data = await response.json();

      isGenerated.value = true;
      link.value = new URL(`/${data.id}`, backendOrigin);
    } catch (error) {
      console.error(error);
    }
  }

  /**
   * Copy shorten link to clipboard.
   * 
   * @async
   * @returns {Promise<void>}
   */
  const copyLink = async (): Promise<void> => {
    if (! isGenerated.value) {
      console.log('Generate a shortlink first');
      return;
    }

    try {
      await navigator.clipboard.writeText(link.value.toString());
      console.log('Copied to clipboard')
    } catch {
      console.log('Copy to clipboard failed');
    }
  }
</script>

<template>
  <div class="index">
    <div class="index-result">
      <IconButton @click="copyLink">
        <CopyIcon />
      </IconButton>
      <a class="index-result-link" target="_blank" title="Generated shortlink" :class="{ 'disabled': ! isGenerated }" :href="link.toString()">{{ link.toString() }}</a>
    </div>
    <form class="index-form" @submit.prevent="onSubmit">
      <input
        class="index-form-input"
        type="url"
        name="url"
        id="url"
        placeholder="https://www.example.com/your/super/long/url"
        required
        v-model="linkInput"
      >
      <button class="index-form-submit button" type="submit">Shorten</button>
    </form>
  </div>
</template>