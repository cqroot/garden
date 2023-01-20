import { ref } from "vue";
import { lightTheme, darkTheme } from "naive-ui";

export function useTheme() {
  const currentTheme = ref(lightTheme);

  const toggleTheme = () => {
    if (currentTheme.value.name === "light") {
      currentTheme.value = darkTheme;
    } else {
      currentTheme.value = lightTheme;
    }
  };

  return {
    currentTheme,
    toggleTheme,
  };
}
