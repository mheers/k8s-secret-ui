<template>
  <v-tabs v-model="tab" bg-color="primary">
    <v-tab value="data">data</v-tab>
    <v-tab value="labels">labels</v-tab>
  </v-tabs>

  <v-window v-model="tab">
    <v-window-item value="data">
      <value-editor v-model="data" text-area />
    </v-window-item>
    <v-window-item value="labels">
      <value-editor v-model="labels" />
    </v-window-item>
  </v-window>

  <v-dialog v-model="dialog" max-width="500px">
    <resource-deleter
      v-if="namespaceName"
      :namespaceName="namespaceName"
      :resourceName="resourceName"
      :resourceType="resourceType"
      @close="dialog = false"
      @deleted="$emit('deleted')"
    />
  </v-dialog>

  <v-sheet class="d-flex mb-6">
    <v-sheet class="ma-2 pa-2 me-auto">
      <v-btn color="primary" @click="saveResource">
        <v-icon>mdi-content-save</v-icon> Save
      </v-btn>
    </v-sheet>
    <v-sheet class="ma-2 pa-2">
      <v-btn color="blue darken-1" @click="getResource">
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
import ResourceDeleter from "./ResourceDeleter.vue";

import ResourceService from "./Resource.service";
const cms = new ResourceService();

const props = defineProps(["namespaceName", "resourceName", "resourceType"]);
const emit = defineEmits(["deleted"]);

// Define reactive variables
const data = ref<Map<string, string>>();
const labels = ref<Map<string, string>>();
const configmapNameRO = ref<string>("");
const dialog = ref<boolean>(false);
const tab = ref("data");

watch(
  () => props.resourceName,
  () => {
    getResource();
  }
);

// Fetch configs on component mount
onMounted(async () => {
  getResource();
});

const getResource = () => {
  cms
    .getResource(props.namespaceName, props.resourceType, props.resourceName)
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

const saveResource = () => {
  cms
    .saveResource(
      props.namespaceName,
      props.resourceType,
      props.resourceName,
      toRaw(labels.value ?? new Map<string, string>()),
      toRaw(data.value ?? new Map<string, string>())
    )
    .then(() => {})
    .catch((error) => {
      console.error(error);
    });
};
</script>
