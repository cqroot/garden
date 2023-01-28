import { ref } from "vue";
import type { Task } from "@/api/types";
import { reqGetTask, reqPutTask, reqUpdateTask } from "@/api/requests";

export function useTaskDialog(emit: any) {
  const timestamp = ref(0);
  const task = ref({} as Task);
  const isModalVisible = ref(false);
  const isNewTask = ref(true);

  const createTask = () => {
    task.value = {} as Task;
    isNewTask.value = true;
    isModalVisible.value = true;
  };

  const editTask = (id: number) => {
    reqGetTask(id).then((response) => {
      task.value = response.data;
      timestamp.value = task.value.due;
      isNewTask.value = false;
      isModalVisible.value = true;
    });
  };

  function submitTask() {
    task.value.due = timestamp.value;
    if (isNewTask.value) {
      reqPutTask(task.value).then(() => {
        emit("onTaskCreate");
      });
    } else {
      reqUpdateTask(task.value.id, task.value).then(() => {
        emit("onTaskUpdate", task.value.id);
      });
    }
    isModalVisible.value = false;
  }

  return {
    isModalVisible,
    timestamp,
    task,
    submitTask,
    createTask,
    editTask,
  };
}
