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
            <v-row>
              <v-col>
                <config-map-selector
                  @change="
                    (newConfigMapName) => (configMapName = newConfigMapName)
                  "
                  v-if="namespaceName"
                  :namespaceName="namespaceName"
                />
              </v-col>
            </v-row>
            <config-map-editor
              v-if="namespaceName && configMapName"
              :configMapName="configMapName"
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

const namespaceName = ref("default");
const secretName = ref("");
const configMapName = ref("");
const tab = ref("configmap");

import NamespaceSelector from "./NamespaceSelector.vue";
import ConfigMapSelector from "./ConfigMapSelector.vue";
import ConfigMapEditor from "./ConfigMapEditor.vue";
</script>
