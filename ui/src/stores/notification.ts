import { ref } from "vue";
import { defineStore } from "pinia";

export const useNotificationStore = defineStore("message", () => {
  const warningContent = ref("");
  const errorContent = ref("");

  return { warningContent, errorContent };
});
