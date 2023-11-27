<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer">
      <v-list>
        <v-list-item>
          <v-list-item-title class="title"> Editor </v-list-item-title>
          <v-list-item-subtitle> secrets and configmaps </v-list-item-subtitle>
        </v-list-item>

        <v-divider></v-divider>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-app-bar-title>Edit k8s secrets and configmaps</v-app-bar-title>
    </v-app-bar>

    <v-main>
      <K8sUI />
    </v-main>
  </v-app>
</template>

<style scoped></style>

<script setup lang="ts">
import K8sUI from "./components/K8sUI.vue";
import { ref } from "vue";

const drawer = ref(true);
</script>
