<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card>
          <v-card-title> Manage Values </v-card-title>
          <v-card-text>
            <!-- Form for creating and editing values -->
            <v-form @submit.prevent="saveValue">
              <v-row>
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model="editing.key"
                    label="Key"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model="editing.value"
                    label="Value"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="4">
                  <v-btn type="submit" color="primary">Save Value</v-btn>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Display existing values -->
    <v-table>
      <template v-slot:top>
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
            <th>Actions</th>
          </tr>
        </thead>
      </template>
      <tbody>
        <tr v-for="v in values" :key="v[0]">
          <td>{{ v[0] }}</td>
          <td>{{ v[1] }}</td>
          <td>
            <v-btn @click="editValue(v[0])" color="primary">Edit</v-btn>
            <v-btn @click="deleteValue(v[0])" color="error">Delete</v-btn>
          </td>
        </tr>
      </tbody>
    </v-table>
  </v-container>
</template>

<script setup lang="ts">
import { ref, defineProps, watch, onMounted } from "vue";

import type { Value } from "./types";

const props = defineProps({ modelValue: { type: Map<string, string> } });

watch(
  () => props.modelValue,
  () => {
    updateValues();
  }
);

// Fetch configs on component mount
onMounted(async () => {
  updateValues();
});

const updateValues = () => {
  values.value = props.modelValue;
};

const editing = ref<Value>({
  key: "t",
  value: "v",
});
const values = ref<Map<string, string>>();
let editingKey = ref<string>();

const saveValue = () => {
  values.value?.set(editing.value.key, editing.value.value);

  // Clear form
  editing.value = { key: "", value: "" };
};

const editValue = (key: string) => {
  // Set form values for editing
  editing.value.key = key || "";
  editing.value.value = values.value?.get(key) || "";
  editingKey.value = key;
};

const deleteValue = (key: string) => {
  // Remove value at the specified key
  values.value?.delete(key);
  // Clear form if the deleted value was being edited
  if (editingKey.value === key) {
    editing.value = { key: "", value: "" };
    editingKey.value = undefined;
  }
};
</script>
