<template>
  <v-card>
    <v-card-title>
      <span class="headline"
        >Confirm Deletion of ConfigMap {{ props.configMapName }}</span
      >
    </v-card-title>

    <v-card-subtitle>
      <span class="subtitle">This action cannot be undone.</span>
    </v-card-subtitle>

    <v-card-text>
      <v-container>
        <v-row>
          <v-col>
            <v-text-field
              :loading="loading"
              v-model="confirmConfigMapName"
              label="Type the name of the ConfigMap to confirm deletion."
              :readonly="false"
              :counter="63"
              :counter-value="confirmConfigMapName?.length"
            />
          </v-col>
        </v-row>
      </v-container>
    </v-card-text>

    <v-card-actions>
      <v-spacer />
      <v-btn color="blue darken-1" @click="$emit('close')"> Cancel </v-btn>
      <v-btn
        color="blue darken-1"
        @click="deleteConfigMapAndRefreshSelector"
        :disabled="confirmConfigMapName != configMapName"
      >
        Delete
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
export default {};
</script>

<script setup lang="ts">
import { ref } from "vue";

import ConfigMapService from "./ConfigMap.service";
const cms = new ConfigMapService();

const loading = ref<boolean>(false);

const props = defineProps(["namespaceName", "configMapName"]);

const emit = defineEmits(["close"]);

const confirmConfigMapName = ref<string>("");
const deleteConfigMapAndRefreshSelector = () => {
  loading.value = true;
  cms
    .deleteConfigMap(props.namespaceName, confirmConfigMapName.value)
    .then(() => {
      confirmConfigMapName.value = "";
    })
    .catch((error) => {
      console.error(error);
    })
    .finally(() => {
      loading.value = false;
      emit("close");
    });
};
</script>

<style></style>
