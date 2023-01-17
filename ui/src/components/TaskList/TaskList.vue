<script setup lang="ts">
import { MenuSharp as MenuIcon } from "@vicons/ionicons5";
import { useTaskList } from "./TaskList";

const {
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
} = useTaskList();

const emit = defineEmits<{
  (e: "onTaskClick", id: number): void;
}>();

const handleTaskClick_ = (id: number) => {
  emit("onTaskClick", id);
};

defineExpose({
  updateTasks,
  updateTask,
});
</script>

<template>
  <n-list hoverable clickable bordered>
    <n-list-item
      v-for="task in tasks"
      :key="task.id"
      @click="handleTaskClick(task.id, handleTaskClick_)"
    >
      <template #prefix>
        <n-button quaternary circle>
          <template #icon>
            <n-icon><menu-icon /></n-icon>
          </template>
        </n-button>
      </template>

      <template #suffix>
        <n-dropdown
          trigger="click"
          placement="bottom-end"
          :options="menuOptions"
          @select="handleMenuSelect"
        >
          <n-button quaternary circle @click="handleMenuButtonClick(task.id)">
            <template #icon>
              <n-icon><menu-icon /></n-icon>
            </template>
          </n-button>
        </n-dropdown>
      </template>

      <n-thing :title="task.title">
        {{ task.note }}
      </n-thing>
    </n-list-item>
  </n-list>

  <n-modal
    v-model:show="showConfirmModal"
    :mask-closable="false"
    preset="dialog"
    title="Dialog"
    :content="currConfirmText"
    positive-text="Confirm"
    negative-text="Cancel"
    @positive-click="handleConfirmPositiveClick"
  />
</template>
