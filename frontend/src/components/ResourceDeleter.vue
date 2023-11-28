<template>
  <v-card>
    <v-card-title>
      <span class="headline"
        >Confirm Deletion of {{ props.resourceType }}
        {{ props.resourceName }}</span
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
              v-model="confirmResourceName"
              :label="`Type the name of the ${props.resourceType} to confirm deletion.`"
              :readonly="false"
              :counter="63"
              :counter-value="confirmResourceName?.length"
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
        @click="deleteResourceAndRefreshSelector"
        :disabled="confirmResourceName != resourceName"
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

import ResourceService from "./Resource.service";
const cms = new ResourceService();

const loading = ref<boolean>(false);

const props = defineProps(["namespaceName", "resourceName", "resourceType"]);

const emit = defineEmits(["close", "deleted"]);

const confirmResourceName = ref<string>("");
const deleteResourceAndRefreshSelector = () => {
  loading.value = true;
  cms
    .deleteResource(
      props.namespaceName,
      props.resourceType,
      confirmResourceName.value
    )
    .then(() => {
      confirmResourceName.value = "";
      emit("deleted");
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
