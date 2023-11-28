<template>
  <div>
    <label for="namespaceDropdown">Select Namespace:</label>
    <v-select
      id="namespaceDropdown"
      :items="namespaceItems"
      v-model="namespaceName"
      @update:modelValue="handleNamespaceChange"
      :loading="loading"
      clearable
    >
      <template #append>
        <v-btn color="primary" @click="getNamespaces()">
          <v-icon>mdi-refresh</v-icon>
        </v-btn>
      </template>
    </v-select>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, computed } from "vue";

import NamespaceService from "./Namespace.service";
const nss = new NamespaceService();

const emit = defineEmits(["update:modelValue"]);
const props = defineProps({
  modelValue: { type: String, default: "" },
});

// Define reactive variables
const namespaces = ref([]);
const namespaceName = ref<string>("");
const loading = ref<boolean>(false);

// Fetch namespaces on component mount
onMounted(() => {
  getNamespaces();
});

watch(
  () => props.modelValue,
  () => {
    namespaceName.value = props.modelValue;
  },
  { immediate: true }
);

const getNamespaces = () => {
  loading.value = true;
  nss
    .getNamespaces()
    .then((response) => {
      namespaces.value = response;
    })
    .catch((error) => {
      console.error(error);
    })
    .finally(() => {
      loading.value = false;
    });
};

// Event handler for namespace update
const handleNamespaceChange = () => {
  console.log("handleNamespaceChange", namespaceName.value);
  emit("update:modelValue", namespaceName.value);
};

const namespaceItems = computed(() => {
  // just returns a list of names of the namespaces
  return namespaces.value.map((namespace: any) => {
    return namespace.metadata.name;
  });
});
</script>
