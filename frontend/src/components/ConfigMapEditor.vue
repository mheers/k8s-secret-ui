<template>
  <div>
    <v-ace-editor
      v-model:value="content"
      lang="json"
      theme="chrome"
      style="height: 300px"
    />
  </div>
</template>

<script setup lang="ts">
import { VAceEditor } from "vue3-ace-editor";

import { ref, onMounted, watch } from "vue";

const props = defineProps(["namespace", "configmap"]);

// Define reactive variables
const content = ref("");

watch(
  () => props.configmap,
  () => {
    updateValue();
  }
);

// Fetch configs on component mount
onMounted(async () => {
  updateValue();
});

const updateValue = async () => {
  try {
    const response = await fetch(
      `http://localhost:8000/api/configs/${props.namespace}/${props.configmap}`
    );
    const data = await response.json();
    content.value = JSON.stringify(data.data); // Assuming data is an array of configs
  } catch (error) {
    console.error("Error fetching configs:", error);
  }
};
</script>
