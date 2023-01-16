import { reqGetTask, reqGetTasks } from "@/api/requests";
import type { Task } from "@/api/types";
import { copyTask } from "@/api/types";
import { ref } from "vue";

export function useTasks() {
  const tasks = ref([] as Task[]);

  const updateTasks = () => {
    reqGetTasks().then((response) => {
      tasks.value = response.data;
    });
  };

  const updateTask = (id: number) => {
    reqGetTask(id).then((response) => {
      const _task = response.data;
      for (const task of tasks.value) {
        if (task.id === _task.id) {
          copyTask(task, _task);
        }
      }
    });
  };

  return {
    tasks,
    updateTasks,
    updateTask,
  };
}
