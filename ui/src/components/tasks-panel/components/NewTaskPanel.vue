<script setup lang="ts">
import { ref } from "vue";
import {
  ListSharp as ListIcon,
  AddCircle as AddIcon,
  ArrowUpCircle as CollapseIcon,
} from "@vicons/ionicons5";
import type { Task } from "@/api/types";
import { reqPutTask } from "@/api/requests";
import DatePickButton from "./DatePickButton.vue";

const task = ref({} as Task);
const showNewTaskPanel = ref(false);
const datePickButtonRef = ref();

const emit = defineEmits<{
  (e: "onTaskPut"): void;
}>();

const handlePutTask = () => {
  const timestamp = datePickButtonRef.value.timestamp;
  if (timestamp != 0 && timestamp != null) {
    task.value.due = timestamp;
  }
  reqPutTask(task.value).then(() => {
    emit("onTaskPut");
    showNewTaskPanel.value = false;
  });
};
</script>

<template>
  <div>
    <n-button
      v-if="showNewTaskPanel"
      strong
      :focusable="false"
      @click="showNewTaskPanel = !showNewTaskPanel"
      style="min-width: 100%"
      class="mb-2"
    >
      <template #icon>
        <n-icon><collapse-icon /></n-icon>
      </template>
      Collapse
    </n-button>
    <n-button
      v-else
      type="primary"
      strong
      :focusable="false"
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

      <date-pick-button ref="datePickButtonRef" />

      <n-button>
        <template #icon>
          <n-icon>
            <list-icon />
          </n-icon>
        </template>
        Inbox
      </n-button>

      <n-button style="float: right" type="primary" @click="handlePutTask">
        <template #icon>
          <n-icon>
            <add-icon />
          </n-icon>
        </template>
        Add
      </n-button>
    </n-collapse-transition>
  </div>
</template>
