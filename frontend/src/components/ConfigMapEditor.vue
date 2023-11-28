<template>
  <v-card>
    <v-card-title> Labels </v-card-title>

    <value-editor v-model="labels" />
  </v-card>

  <v-card>
    <v-card-title> Data </v-card-title>
    <value-editor v-model="data" text-area />
  </v-card>

  <v-dialog v-model="dialog" max-width="500px">
    <config-map-deleter
      v-if="namespaceName"
      :namespaceName="namespaceName"
      :configMapName="configMapName"
      @close="dialog = false"
    />
  </v-dialog>

  <v-sheet class="d-flex mb-6">
    <v-sheet class="ma-2 pa-2 me-auto">
      <v-btn color="primary" @click="saveConfigMap">
        <v-icon>mdi-content-save</v-icon> Save
      </v-btn>
    </v-sheet>
    <v-sheet class="ma-2 pa-2">
      <v-btn color="blue darken-1" @click="getConfigMap">
        <v-icon>mdi-lock-reset</v-icon> Reset
      </v-btn>
    </v-sheet>
    <v-sheet class="ma-2 pa-2">
      <v-btn color="error" @click="dialog = true">
        <v-icon>mdi-delete</v-icon> Delete
      </v-btn>
    </v-sheet>
  </v-sheet>
</template>

<script lang="ts">
export default {};
</script>

<script setup lang="ts">
import { ref, onMounted, watch, toRaw } from "vue";

import ValueEditor from "./ValueEditor.vue";
import ConfigMapDeleter from "./ConfigMapDeleter.vue";

import ConfigMapService from "./ConfigMap.service";
const cms = new ConfigMapService();

const props = defineProps(["namespaceName", "configMapName"]);

// Define reactive variables
const data = ref<Map<string, string>>();
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
    .then((cm) => {
      labels.value = new Map();
      data.value = new Map();

      data.value = new Map(Object.entries(cm.data || {}));

      if (cm.metadata) {
        configmapNameRO.value = cm.metadata.name;

        if (cm.metadata.labels) {
          labels.value = new Map(Object.entries(cm.metadata.labels));
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
      toRaw(labels.value ?? new Map<string, string>()),
      toRaw(data.value ?? new Map<string, string>())
    )
    .then(() => {})
    .catch((error) => {
      console.error(error);
    });
};
</script>
