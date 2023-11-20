<template>
  <namespace-selector @change="(newNamespace) => (namespace = newNamespace)" />

  <v-card v-if="namespace">
    <v-tabs v-model="tab" bg-color="primary">
      <v-tab value="configmap">ConfigMap</v-tab>
      <v-tab value="secret">Secret</v-tab>
    </v-tabs>

    <v-card-text>
      <v-window v-model="tab">
        <v-window-item value="configmap">
          <config-map-selector
            @change="(newConfigMap) => (configmap = newConfigMap)"
            v-if="namespace"
            :namespace="namespace"
          />
          <config-map-editor
            v-if="namespace && configmap"
            :configmap="configmap"
            :namespace="namespace"
          />
        </v-window-item>

        <v-window-item value="secret"> Two </v-window-item>
      </v-window>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from "vue";

const namespace = ref("");
const secret = ref("");
const configmap = ref("");
const tab = ref("configmap");

import NamespaceSelector from "./NamespaceSelector.vue";
import ConfigMapSelector from "./ConfigMapSelector.vue";
import ConfigMapEditor from "./ConfigMapEditor.vue";
</script>
