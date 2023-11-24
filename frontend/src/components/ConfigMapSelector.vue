<template>
  <label for="configMapDropdown">Select ConfigMap:</label>
  <v-select
    :clearable="true"
    id="configMapDropdown"
    :items="configItems"
    @update:modelValue="handleConfigMapChange"
  >
    <template #append>
      <v-btn color="primary" @click="updateValues()"> refresh </v-btn>
      <v-btn color="secondary" @click="dialog = true"> + </v-btn>
    </template>
  </v-select>

  <v-dialog v-model="dialog" max-width="500px">
    <config-map-creator
      v-if="namespaceName"
      :namespaceName="namespaceName"
      @close="dialog = false"
    />
  </v-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, computed, defineProps, defineEmits } from "vue";

const props = defineProps(["namespaceName"]);
const emit = defineEmits(["change"]);

// Define reactive variables
const configs = ref([]);
const selectedConfig = ref("");
const dialog = ref<boolean>(false);

watch(
  () => props.namespaceName,
  () => {
    updateValues();
  }
);

// Fetch configs on component mount
onMounted(async () => {
  updateValues();
});

const updateValues = async () => {
  try {
    const response = await fetch(
      `http://localhost:8000/api/configs/${props.namespaceName}`
    );
    const data = await response.json();
    configs.value = data; // Assuming data is an array of configs
  } catch (error) {
    console.error("Error fetching configs:", error);
  }
};

// Event handler for config change
const handleConfigMapChange = (configMapName: string) => {
  emit("change", configMapName);
};

const configItems = computed(() => {
  // just returns a list of names of the configs
  return configs.value.map((config: any) => {
    return config.metadata.name;
  });
});

import ConfigMapCreator from "./ConfigMapCreator.vue";
</script>
