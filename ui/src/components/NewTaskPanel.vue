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

const emit = defineEmits<{
  (e: "onTaskPut"): void;
}>();

const putTask = () => {
  reqPutTask(task.value).then(() => {
    emit("onTaskPut");
  });
};
</script>

<template>
  <n-collapse accordion>
    <template #arrow>
      <n-icon>
        <add-icon />
      </n-icon>
    </template>
    <n-collapse-item title="New Task" name="1">
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

      <n-button style="float: right" type="info" @click="putTask">
        <template #icon>
          <n-icon>
            <add-icon />
          </n-icon>
        </template>
        Add
      </n-button>
    </n-collapse-item>
  </n-collapse>
</template>
