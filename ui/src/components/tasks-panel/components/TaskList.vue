<script setup lang="ts">
import {
  CheckmarkSharp as CheckIcon,
  MenuSharp as MenuIcon,
  CalendarSharp as CalendarIcon,
} from "@vicons/ionicons5";
import { formattedDate } from "@/api/date";
import { useTaskList } from "./task-list";

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
        <n-button tertiary circle>
          <template #icon>
            <n-icon><check-icon /></n-icon>
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
        <template #description>
          <p class="text-zinc-500 m-0 mb-1 p-0">{{ task.note }}</p>
        </template>
      </n-thing>
      <n-tag size="small" :bordered="false" type="success" v-if="task.due != 0">
        <template #icon>
          <n-icon>
            <calendar-icon />
          </n-icon>
        </template>
        {{ formattedDate(task.due) }}
      </n-tag>
    </n-list-item>
  </n-list>

  <n-modal
    v-model:show="showConfirmModal"
    :mask-closable="false"
    preset="dialog"
    title="Confirm"
    :content="currConfirmText"
    positive-text="Confirm"
    negative-text="Cancel"
    @positive-click="handleConfirmPositiveClick"
  />
</template>
