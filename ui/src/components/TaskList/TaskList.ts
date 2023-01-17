import { ref } from "vue";
import { reqDeleteTask } from "@/api/requests";
import { useTasks } from "./task";

export function useTaskList() {
  const { tasks, updateTasks, updateTask } = useTasks();
  const canHandleClick = ref(true);
  const showConfirmModal = ref(false);

  const currTaskId = ref(0);
  const currAction = ref("");
  const currConfirmText = ref("");

  const menuOptions = ref([
    {
      label: "Delete task",
      key: "delete-task",
    },
  ]);

  const handleMenuButtonClick = (id: number) => {
    canHandleClick.value = false;
    currTaskId.value = id;
  };

  const handleMenuSelect = (key: string | number) => {
    switch (key) {
      case "delete-task": {
        currAction.value = key;
        currConfirmText.value = "Are you sure you want to delete this task?";
        showConfirmModal.value = true;
        break;
      }
    }
  };

  const handleConfirmPositiveClick = () => {
    switch (currAction.value) {
      case "delete-task": {
        reqDeleteTask(currTaskId.value).then(() => {
          updateTasks();
        });
      }
    }
  };

  const handleTaskClick = (id: number, handleFunc: (id: number) => void) => {
    if (canHandleClick.value) {
      handleFunc(id);
    }
    canHandleClick.value = true;
  };

  return {
    tasks,
    updateTasks,
    updateTask,

    showConfirmModal,
    currConfirmText,
    menuOptions,
    handleMenuButtonClick,
    handleMenuSelect,
    handleConfirmPositiveClick,
    handleTaskClick,
  };
}
