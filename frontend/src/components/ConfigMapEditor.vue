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

import ConfigMapService from "./ConfigMap.service";
const cms = new ConfigMapService();

const props = defineProps(["namespaceName", "configMapName"]);

// Define reactive variables
const content = ref("");
const labels = ref<Map<string, string>>();
const configmapNameRO = ref<string>("");
const dialog = ref<boolean>(false);

watch(
  () => props.configMapName,
  () => {
    getConfigMap();
  }
);

// Fetch configs on component mount
onMounted(async () => {
  getConfigMap();
});

const getConfigMap = () => {
  cms
    .getConfigMap(props.namespaceName, props.configMapName)
    .then((data) => {
      labels.value = new Map();

      content.value = JSON.stringify(data.data); // Assuming data is an array of configs

      if (data.metadata) {
        configmapNameRO.value = data.metadata.name;

        if (data.metadata.labels) {
          labels.value = new Map(Object.entries(data.metadata.labels));
        }
      }
    })
    .catch((error) => {
      console.error(error);
    });
};

const saveConfigMap = () => {
  cms
    .saveConfigMap(
      props.namespaceName,
      props.configMapName,
      labels.value,
      JSON.parse(content.value)
    )
    .then(() => {})
    .catch((error) => {
      console.error(error);
    });
};

import LabelEditor from "./LabelEditor.vue";
import ConfigMapDeleter from "./ConfigMapDeleter.vue";
</script>
