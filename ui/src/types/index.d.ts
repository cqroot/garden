import type { MessageApi, NotificationApi } from "naive-ui";

export {};

declare global {
  interface Window {
    $message: MessageApi;
    $notification: NotificationApi;
  }
}
