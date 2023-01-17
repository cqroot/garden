<script setup lang="ts">
import { ref } from "vue";
import {
  CalendarSharp as CalendarIcon,
  ListSharp as ListIcon,
  AddCircle as AddIcon,
} from "@vicons/ionicons5";
import type { Task } from "@/api/types";
import { reqPutTask } from "@/api/requests";

const task = ref({} as Task);
const showNewTaskPanel = ref(false);

const emit = defineEmits<{
  (e: "onTaskPut"): void;
}>();

const putTask = () => {
  reqPutTask(task.value).then(() => {
    emit("onTaskPut");
    showNewTaskPanel.value = false;
  });
};
</script>

<template>
  <div>
    <n-button
      type="primary"
      strong
      @click="showNewTaskPanel = !showNewTaskPanel"
      style="min-width: 100%"
      class="mb-2"
    >
      <template #icon>
        <n-icon><add-icon /></n-icon>
      </template>
      New Task
    </n-button>
    <n-collapse-transition :show="showNewTaskPanel" class="mb-2">
      <n-input
        class="mb-2"
        type="text"
        placeholder="Input task title"
        v-model:value="task.title"
      />

      <n-input
        class="mb-2"
        type="textarea"
        placeholder="Input task note"
        v-model:value="task.note"
      />

      <p style="padding: 0; margin: 0">
        <n-button class="mr-2">
          <template #icon>
            <n-icon>
              <calendar-icon />
            </n-icon>
          </template>
          Due date
        </n-button>

        <n-button>
          <template #icon>
            <n-icon>
              <list-icon />
            </n-icon>
          </template>
          Inbox
        </n-button>

        <n-button style="float: right" type="primary" @click="putTask">
          <template #icon>
            <n-icon>
              <add-icon />
            </n-icon>
          </template>
          Add
        </n-button>
      </p>
    </n-collapse-transition>
  </div>
</template>
