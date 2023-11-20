<template>
  <div>
    <label for="namespaceDropdown">Select Namespace:</label>
    <v-select
      id="namespaceDropdown"
      :items="namespaceItems"
      @update:modelValue="handleNamespaceChange"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, defineEmits } from "vue";

const emit = defineEmits(["change"]);

// Define reactive variables
const namespaces = ref([]);

// Fetch namespaces on component mount
onMounted(async () => {
  try {
    const response = await fetch("http://localhost:8000/api/namespaces");
    const data = await response.json();
    namespaces.value = data; // Assuming data is an array of namespaces
  } catch (error) {
    console.error("Error fetching namespaces:", error);
  }
});

// Event handler for namespace change
const handleNamespaceChange = (namespace: string) => {
  console.log("handleNamespaceChange", namespace);
  emit("change", namespace);
};

const namespaceItems = computed(() => {
  // just returns a list of names of the namespaces
  return namespaces.value.map((namespace: any) => {
    return namespace.metadata.name;
  });
});
</script>
