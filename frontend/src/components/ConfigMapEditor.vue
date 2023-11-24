<template>
  <label-editor :labels="labels" />

  <v-ace-editor
    v-model:value="content"
    lang="json"
    theme="chrome"
    style="height: 300px"
    active
  />

  <v-dialog v-model="dialog" max-width="500px">
    <config-map-deleter
      v-if="namespaceName"
      :namespaceName="namespaceName"
      :configMapName="configMapName"
      @close="dialog = false"
    />
  </v-dialog>

  <v-btn color="primary" @click="saveConfigMap"> Save </v-btn>
  <v-btn color="error" @click="dialog = true"> Delete </v-btn>
</template>

<script setup lang="ts">
import { VAceEditor } from "vue3-ace-editor";
import { ref, onMounted, watch } from "vue";

const props = defineProps(["namespaceName", "configMapName"]);

// Define reactive variables
const content = ref("");
const labels = ref<Map<string, string>>();
const configmapNameRO = ref<string>("");
const dialog = ref<boolean>(false);

watch(
  () => props.configMapName,
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
      `http://localhost:8000/api/configs/${props.namespaceName}/${props.configMapName}`
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

import LabelEditor from "./LabelEditor.vue";
import ConfigMapDeleter from "./ConfigMapDeleter.vue";
</script>
