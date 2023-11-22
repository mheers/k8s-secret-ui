<template>
  <div>
    <v-text-field
      v-model:value="configmapNameRO"
      readonly
      active
      label="Name"
      :counter="63"
      :counter-value="configmapNameRO.length"
    />

    <label-editor :labels="labels" />

    <v-ace-editor
      v-model:value="content"
      lang="json"
      theme="chrome"
      style="height: 300px"
      active
    />
  </div>
</template>

<script setup lang="ts">
import { VAceEditor } from "vue3-ace-editor";
import LabelEditor from "./LabelEditor.vue";
import { ref, onMounted, watch } from "vue";

const props = defineProps(["namespaceName", "configmapName"]);

// Define reactive variables
const content = ref("");
const labels = ref<Map<string, string>>();
const configmapNameRO = ref<string>("");

watch(
  () => props.configmapName,
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
      `http://localhost:8000/api/configs/${props.namespaceName}/${props.configmapName}`
    );
    const data = await response.json();

    labels.value = new Map();

    configmapNameRO.value = data.metadata.name;
    labels.value = new Map(Object.entries(data.metadata.labels));
    content.value = JSON.stringify(data.data); // Assuming data is an array of configs
  } catch (error) {
    console.error("Error fetching configs:", error);
  }
};
</script>
