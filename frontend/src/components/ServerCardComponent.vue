<template>
  <q-card-section class="bg-primary text-white text-h5">
    <q-icon :name="'img:'+ iconURL" size="32px" v-if="iconURL"/>
    <q-icon name="sports_esports" size="32px" v-else/>
    {{ serverName }}
  </q-card-section>

  <q-card-section>
    {{ description }}
  </q-card-section>

  <q-card-actions align="evenly">
    <q-btn icon="settings" label="Manage" color="accent" rounded @click="handleClickSettings"/>
    <q-btn icon="person" label="Live View" color="accent" rounded />
  </q-card-actions>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import {useRouter} from "vue-router";

export default defineComponent({
  name: 'ServerCardComponent',
  components: {},
  methods: {
    async handleClickSettings() {
      await this.router.push("/servers/" + this.serverID)
    },
  },
  props: {
    serverID: {
      type: String,
      required: true,
    },
    serverName: {
      type: String,
      required: true,
    },
    iconURL: {
      type: String,
      required: false,
    },
    description: {
      type: String,
      required: false,
    }
  },
  setup () {
    const router = useRouter()
    return { router };
  }
});
</script>
