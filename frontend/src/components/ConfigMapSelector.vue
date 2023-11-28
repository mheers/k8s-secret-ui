<template>
  <label for="configMapDropdown" :id="componentID">Select ConfigMap:</label>
  <v-select
    :clearable="true"
    id="configMapDropdown"
    :items="configItems"
    @update:modelValue="handleConfigMapChange"
    :loading="loading"
    v-model="configMapName"
  >
    <template #append>
      <v-btn color="primary" @click="getConfigMaps()">
        <v-icon>mdi-refresh</v-icon>
      </v-btn>
      <v-btn color="secondary" @click="newNameEntering = true">
        <v-icon>mdi-plus</v-icon>

        <v-overlay
          v-model="newNameEntering"
          :attach="`#${componentID}`"
          contained
        >
          <v-card class="pa-2" min-width="300px">
            <v-text-field
              v-model="newName"
              label="New Name"
              required
              variant="outlined"
              clearable
              autofocus
            >
              <template #append>
                <v-btn @click="create()" color="primary" :disabled="!newName">
                  <v-icon>mdi-check</v-icon>
                </v-btn>
              </template>
            </v-text-field>
          </v-card>
        </v-overlay>
      </v-btn>
    </template>
  </v-select>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, computed } from "vue";

import ConfigMapService from "./ConfigMap.service";
const cms = new ConfigMapService();

const create = () => {
  loading.value = true;
  cms
    .createConfigMap(props.namespaceName, newName.value)
    .then(() => {
      newName.value = "";
      newNameEntering.value = false;
      handleConfigMapChange(newName.value);
    })
    .catch((error) => {
      console.error(error);
    })
    .finally(() => {
      loading.value = false;
    });
};

const props = defineProps({
  namespaceName: { type: String, default: "" },
  modelValue: { type: String, default: "" },
});
const emit = defineEmits(["update:modelValue"]);

const componentID = `component_${Math.floor(Math.random() * 1000)}`;

// Define reactive variables
const configs = ref([]);
const configMapName = ref<string>("");
const newName = ref<string>("");
const newNameEntering = ref<boolean>(false);
const loading = ref<boolean>(false);

watch(
  () => props.namespaceName,
  () => {
    getConfigMaps();
  }
);

watch(
  () => props.modelValue,
  () => {
    configMapName.value = props.modelValue;
  },
  { immediate: true }
);

// Fetch configs on component mount
onMounted(() => {
  getConfigMaps();
});

const getConfigMaps = () => {
  loading.value = true;
  cms
    .getConfigMaps(props.namespaceName)
    .then((response) => {
      configs.value = response;
    })
    .catch((error) => {
      console.error(error);
    })
    .finally(() => {
      loading.value = false;
    });
};

// Event handler for config change
const handleConfigMapChange = (name: string) => {
  configMapName.value = name;
  emit("update:modelValue", name);
};

const configItems = computed(() => {
  // just returns a list of names of the configs
  return configs.value.map((config: any) => {
    return config.metadata.name;
  });
});
</script>
