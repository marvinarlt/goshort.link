<script setup lang="ts">
  import { ref } from 'vue';
  import { CopyIcon } from 'vue-tabler-icons';
  import { useNotification } from '@app/composables/notification';
  import IconButton from '@app/components/IconButton.vue';

  const backendOrigin = import.meta.env.VITE_BACKEND_ORIGIN;
  const notification = useNotification();

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
      notification.add({
        type: 'error',
        title: 'notifications.invalid-url.title',
        body: 'notifications.invalid-url.body'
      });
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
      
      copyLink();
    } catch (error) {
      console.error(error);
      notification.add({
        type: 'error',
        title: 'notifications.generating-failed.title',
        body: 'notifications.generating-failed.body'
      });
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
      notification.add({
        type: 'warning',
        title: 'notifications.not-copied.title',
        body: 'notifications.not-copied.body'
      });
      return;
    }

    try {
      await navigator.clipboard.writeText(link.value.toString());
      notification.add({
        type: 'success',
        title: 'notifications.copied.title',
        body: 'notifications.copied.body'
      });
    } catch {
      notification.add({
        type: 'error',
        title: 'notifications.copying-failed.title',
        body: 'notifications.copying-failed.body'
      });
    }
  }
</script>

<template>
  <div class="index">
    <div class="index-result">
      <IconButton @click="copyLink">
        <CopyIcon />
      </IconButton>
      <a
        class="index-result-link"
        target="_blank"
        :title="$t('pages.index.shortend-title')"
        :class="{ 'disabled': ! isGenerated }"
        :href="link.toString()"
      >{{ link.toString() }}</a>
    </div>
    <form class="index-form" @submit.prevent="onSubmit">
      <input
        class="index-form-input"
        type="url"
        name="url"
        id="url"
        :placeholder="$t('pages.index.placeholder')"
        required
        v-model="linkInput"
      >
      <button class="index-form-submit button" type="submit">{{ $t('pages.index.submit') }}</button>
    </form>
  </div>
</template>