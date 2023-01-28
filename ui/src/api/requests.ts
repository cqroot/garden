import axios from "axios";
import type { Task } from "@/api/types";

function request() {
  const instance = axios.create({
    baseURL:
      process.env.NODE_ENV === "production"
        ? "/v1"
        : "http://localhost:8000/v1",
  });

  instance.interceptors.response.use(
    function (response) {
      return response;
    },
    function (error) {
      window.$notification.error({
        content: error.message,
        duration: 2500,
        keepAliveOnHover: true,
      });
      return Promise.reject(error);
    }
  );

  return instance;
}

export function reqGetTasks(dueStart: number, dueEnd: number) {
  return request().get<Task[]>("/task", {
    params: {
      start: dueStart,
      end: dueEnd,
    },
  });
}

export function reqGetTask(id: number) {
  return request().get<Task>(`/task/${id}`);
}

export function reqPutTask(task: Task) {
  return request().put(`/task`, task);
}

export function reqUpdateTask(id: number, task: Task) {
  return request().put(`/task/${id}`, task);
}

export function reqDeleteTask(id: number) {
  return request().delete(`/task/${id}`);
}
