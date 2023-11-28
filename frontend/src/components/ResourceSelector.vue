<template>
  <label for="resourceDropdown" :id="componentID"
    >Select {{ $props.resourceType }}:</label
  >
  <v-select
    :clearable="true"
    id="resourceDropdown"
    :items="configItems"
    @update:modelValue="handleResourceChange"
    :loading="loading"
    v-model="resourceName"
  >
    <template #append>
      <v-btn color="primary" @click="getResources()">
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

<script lang="ts">
export default {};
</script>

<script setup lang="ts">
import { ref, onMounted, watch, computed } from "vue";

import ResourceService from "./Resource.service";
const rs = new ResourceService();

const create = () => {
  loading.value = true;
  rs.createResource(props.namespaceName, props.resourceType, newName.value)
    .then(() => {
      getResources();
      handleResourceChange(newName.value);
      newName.value = "";
      newNameEntering.value = false;
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
  resourceType: { type: String, default: "" },
});
const emit = defineEmits(["update:modelValue"]);

const componentID = `component_${Math.floor(Math.random() * 1000)}`;

// Define reactive variables
const resources = ref([]);
const resourceName = ref<string>("");
const newName = ref<string>("");
const newNameEntering = ref<boolean>(false);
const loading = ref<boolean>(false);

const getResources = () => {
  loading.value = true;
  rs.getResources(props.namespaceName, props.resourceType)
    .then((response) => {
      resources.value = response;
    })
    .catch((error) => {
      console.error(error);
    })
    .finally(() => {
      loading.value = false;
    });
};

// Event handler for config change
const handleResourceChange = (name: string) => {
  resourceName.value = name;
  emit("update:modelValue", name);
};

const configItems = computed(() => {
  // just returns a list of names of the resources
  return resources.value.map((config: any) => {
    return config.metadata.name;
  });
});

watch(
  () => props.namespaceName,
  () => {
    getResources();
  }
);

watch(
  () => props.modelValue,
  () => {
    resourceName.value = props.modelValue;
    if (props.modelValue === "") {
      getResources();
    }
  },
  { immediate: true }
);

// Fetch resources on component mount
onMounted(() => {
  getResources();
});
</script>
