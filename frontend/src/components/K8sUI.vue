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
            <resource-selector
              v-if="namespaceName"
              v-model="configMapName"
              :namespaceName="namespaceName"
              resourceType="configmap"
            />
            <resource-editor
              v-if="namespaceName && configMapName"
              :resourceName="configMapName"
              :namespaceName="namespaceName"
              @deleted="configMapName = ''"
              resourceType="configmap"
            />
          </v-window-item>

          <v-window-item value="secret">
            <resource-selector
              v-if="namespaceName"
              v-model="secretName"
              :namespaceName="namespaceName"
              resourceType="secret"
            />
            <resource-editor
              v-if="namespaceName && secretName"
              :resourceName="secretName"
              :namespaceName="namespaceName"
              @deleted="secretName = ''"
              resourceType="secret"
            />
          </v-window-item>
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
import ResourceSelector from "./ResourceSelector.vue";
import ResourceEditor from "./ResourceEditor.vue";

const namespaceName = ref<string>("default");
const configMapName = ref<string>("");
const secretName = ref<string>("");
const tab = ref("secret");
</script>
