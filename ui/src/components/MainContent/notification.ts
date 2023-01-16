import { watch } from "vue";
import { storeToRefs } from "pinia";
import { useNotification as useNotification_ } from "naive-ui";
import { useNotificationStore } from "@/stores/notification";

export function useNotification() {
  const notification = useNotification_();
  const { warningContent, errorContent } = storeToRefs(useNotificationStore());

  watch(warningContent, (content) => {
    notification.warning({
      content: content,
      duration: 5000,
    });
  });
  watch(errorContent, (content) => {
    notification.error({
      content: content,
      duration: 5000,
    });
  });
}
