<template>
  <v-container>
    <namespace-selector
      @change="(newNamespaceName) => (namespaceName = newNamespaceName)"
    />

    <v-card v-if="namespaceName">
      <v-tabs v-model="tab" bg-color="primary">
        <v-tab value="configmap">ConfigMap</v-tab>
        <v-tab value="secret">Secret</v-tab>
      </v-tabs>

      <v-card-text>
        <v-window v-model="tab">
          <v-window-item value="configmap">
            <config-map-selector
              @change="(newConfigMapName) => (configmapName = newConfigMapName)"
              v-if="namespaceName"
              :namespaceName="namespaceName"
            />
            <config-map-editor
              v-if="namespaceName && configmapName"
              :configmapName="configmapName"
              :namespaceName="namespaceName"
            />
          </v-window-item>

          <v-window-item value="secret"> Two </v-window-item>
        </v-window>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from "vue";

const namespaceName = ref("");
const secretName = ref("");
const configmapName = ref("");
const tab = ref("configmap");

import NamespaceSelector from "./NamespaceSelector.vue";
import ConfigMapSelector from "./ConfigMapSelector.vue";
import ConfigMapEditor from "./ConfigMapEditor.vue";
</script>
