import axios from "axios";
import type { Task } from "@/api/types";
import { errorNotification } from "@/api/notification";

function request() {
  const instance = axios.create({
    baseURL:
      process.env.NODE_ENV === "production"
        ? "/v1/"
        : "http://localhost:8000/v1",
  });

  instance.interceptors.response.use(
    function (response) {
      return response;
    },
    function (error) {
      errorNotification(error.message);
      return Promise.reject(error);
    }
  );

  return instance;
}

export function reqGetTasks() {
  return request().get<Task[]>("/task");
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

export function reqMarkTaskDone(id: number) {
  return request().put(`/task/status/${id}`);
}
