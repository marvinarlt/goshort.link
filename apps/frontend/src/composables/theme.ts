import { ref } from 'vue';

export function useTheme() {
  enum Theme {
    LIGHT = 'light',
    DARK = 'dark'
  };
  const theme = ref(Theme.LIGHT);
  
  /**
   * Set theme.
   * 
   * @param {string} name
   * @returns {void}
   */
  const set = (name: string) => {
    theme.value = name as Theme;
    document.documentElement.setAttribute('data-theme', name);
    localStorage.setItem('theme', name);
  }

  /**
   * Get theme from local storage or based on media query.
   * 
   * @returns {string}
   */
  const get = (): string => {
    let name = localStorage.getItem('theme');

    if (null === name) {
      name = window.matchMedia('(prefers-color-scheme: dark)').matches
        ? Theme.DARK
        : Theme.LIGHT;
      theme.value = name as Theme;
    }

    return name;
  }

  /**
   * Toggles between themes.
   * 
   * @returns {string}
   */
  const toggle = (): string => {
    const currentTheme = get();
    const newTheme = Theme.LIGHT === currentTheme ? Theme.DARK : Theme.LIGHT;

    set(newTheme);

    return newTheme;
  }

  set(get());

  return {
    Theme,
    theme,
    set,
    get,
    toggle
  }
}