<template>
  <v-container>
    <namespace-selector v-model="namespaceName" />

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
                  ref="configMapSelector"
                  v-model="configMapName"
                  v-if="namespaceName"
                  :namespaceName="namespaceName"
                />
              </v-col>
            </v-row>
            <config-map-editor
              v-if="namespaceName && configMapName"
              :configMapName="configMapName"
              :namespaceName="namespaceName"
              @deleted="configMapName = ''"
            />
          </v-window-item>

          <v-window-item value="secret"> Two </v-window-item>
        </v-window>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">
export default {};
</script>

<script setup lang="ts">
import { ref } from "vue";

import NamespaceSelector from "./NamespaceSelector.vue";
import ConfigMapSelector from "./ConfigMapSelector.vue";
import ConfigMapEditor from "./ConfigMapEditor.vue";

const namespaceName = ref<string>("default");
const secretName = ref<string>("");
const configMapName = ref<string>("test");
const tab = ref("configmap");
</script>
