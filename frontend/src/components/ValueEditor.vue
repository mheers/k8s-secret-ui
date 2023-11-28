<template>
  <v-container @click="editingKey = ''" :id="`${componentID}`">
    <!-- Display existing values -->
    <v-table>
      <thead>
        <tr>
          <th>Key</th>
          <th>Value</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="v in values" :key="v[0]">
          <th @dblclick="editValue(v[0])">
            {{ v[0] }}
          </th>
          <td>
            <div>
              <span v-if="editingKey === v[0]">
                <v-textarea
                  v-if="textArea"
                  v-model="v[1]"
                  required
                  variant="outlined"
                  clearable
                  rows="1"
                  @update:model-value="values?.set(v[0], v[1])"
                  select-all="true"
                  @click.stop
                  @keydown="
                    $event.shiftKey && $event.key === 'Enter' && saveValue()
                  "
                ></v-textarea>
                <v-text-field
                  v-else
                  v-model="v[1]"
                  required
                  variant="outlined"
                  clearable
                  rows="1"
                  @update:model-value="values?.set(v[0], v[1])"
                  select-all="true"
                  @click.stop
                  @keydown="
                    $event.shiftKey && $event.key === 'Enter' && saveValue()
                  "
                ></v-text-field>
              </span>
              <span
                v-else
                @dblclick="editValue(v[0])"
                style="white-space: pre"
                >{{ v[1] }}</span
              >
            </div>
          </td>
          <td>
            <v-btn @click="deleteValue(v[0])" color="error">
              <v-icon>mdi-delete</v-icon>
            </v-btn>
          </td>
        </tr>
      </tbody>
    </v-table>
    <v-btn color="primary" @click="newKeyEntering = true">
      <v-icon>mdi-plus</v-icon>
      <v-overlay v-model="newKeyEntering" :attach="`#${componentID}`" contained>
        <v-card class="pa-2" min-width="300px">
          <v-text-field
            v-model="newKey"
            label="New Key"
            required
            variant="outlined"
            clearable
            autofocus
          >
            <template #append>
              <v-btn
                @click="
                  addValue(newKey);
                  newKeyEntering = false;
                "
                color="primary"
                :disabled="!newKey"
              >
                <v-icon>mdi-check</v-icon>
              </v-btn>
            </template>
          </v-text-field>
        </v-card>
      </v-overlay>
    </v-btn>
  </v-container>
</template>

<script lang="ts">
export default {};
</script>

<script setup lang="ts">
import { ref, watch, onMounted } from "vue";

import type { Value } from "./types";

const props = defineProps({
  modelValue: { type: Map<string, string> },
  textArea: { type: Boolean, default: false },
});

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

const componentID = `component_${Math.floor(Math.random() * 1000)}`;

const editing = ref<Value>({
  key: "t",
  value: "v",
});
const values = ref<Map<string, string>>();
let editingKey = ref<string>();
let newKey = ref<string>("");
let newKeyEntering = ref<boolean>(false);

const addValue = (key: string) => {
  if (!values.value?.has(key)) {
    values.value?.set(key, "");
  }
};

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
