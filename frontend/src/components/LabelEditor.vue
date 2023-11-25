<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card>
          <v-card-title> Manage Labels </v-card-title>
          <v-card-text>
            <!-- Form for creating and editing labels -->
            <v-form @submit.prevent="saveLabel">
              <v-row>
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model="label.key"
                    label="Key"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model="label.value"
                    label="Value"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="4">
                  <v-btn type="submit" color="primary">Save Label</v-btn>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Display existing labels -->
    <v-row>
      <v-col v-for="label in labels" :key="label[0]" cols="12" md="4">
        <v-card>
          <v-card-title>{{ label[0] }}</v-card-title>
          <v-card-subtitle>{{ label[1] }}</v-card-subtitle>
          <v-card-actions>
            <v-btn @click="editLabel(label[0])" color="primary">Edit</v-btn>
            <v-btn @click="deleteLabel(label[0])" color="error">Delete</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref, defineProps, watch, onMounted } from "vue";

import type { Label } from "./types";

const props = defineProps({ modelValue: { type: Map<string, string> } });

watch(
  () => props.modelValue,
  () => {
    updateLabels();
  }
);

// Fetch configs on component mount
onMounted(async () => {
  updateLabels();
});

const updateLabels = () => {
  labels.value = props.modelValue;
};

const label = ref<Label>({
  key: "",
  value: "",
});
const labels = ref<Map<string, string>>();
let editingKey = ref<string>();

const saveLabel = () => {
  labels.value?.set(label.value.key, label.value.value);

  // Clear form
  label.value = { key: "", value: "" };
};

const editLabel = (key: string) => {
  // Set form values for editing
  label.value.key = key || "";
  label.value.value = labels.value?.get(key) || "";
  editingKey.value = key;
};

const deleteLabel = (key: string) => {
  // Remove label at the specified key
  labels.value?.delete(key);
  // Clear form if the deleted label was being edited
  if (editingKey.value === key) {
    label.value = { key: "", value: "" };
    editingKey.value = undefined;
  }
};
</script>
