<template>
  <v-card>
    <v-card-title>
      <span class="headline">Create ConfigMap</span>
    </v-card-title>

    <v-card-text>
      <v-container>
        <v-row>
          <v-col>
            <v-text-field
              :loading="loading"
              v-model="newConfigmapName"
              label="Name"
              :readonly="false"
              :counter="63"
              :counter-value="newConfigmapName?.length"
            />
          </v-col>
        </v-row>
      </v-container>
    </v-card-text>

    <v-card-actions>
      <v-spacer />
      <v-btn color="blue darken-1" @click="$emit('close')"> Cancel </v-btn>
      <v-btn color="blue darken-1" @click="createConfigMapAndRefreshSelector">
        Save
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref, defineProps, defineEmits } from "vue";

import ConfigMapService from "./ConfigMap.service";
const cms = new ConfigMapService();

const loading = ref<boolean>(false);

const props = defineProps(["namespaceName"]);

const emit = defineEmits(["close"]);

const newConfigmapName = ref<string>("");
const createConfigMapAndRefreshSelector = () => {
  loading.value = true;
  cms
    .createConfigMap(props.namespaceName, newConfigmapName.value)
    .then(() => {
      newConfigmapName.value = "";
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
