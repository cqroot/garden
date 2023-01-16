import { useNotificationStore } from "@/stores/notification";

export const errorNotification = (content: string) => {
  useNotificationStore().errorContent = content;
};
