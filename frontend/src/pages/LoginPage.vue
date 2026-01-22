<template>
  <q-page class="row items-center justify-evenly">
    <q-card class="col-4">
      <q-card-section class="text-h5 text-center bg-primary text-white">
        Log In
      </q-card-section>
      <q-card-section class="text-center">
        <q-btn color="purple" icon="sports_esports" label="Log In With Discord" @click="handleDiscordLogin"/>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script lang="ts">

import {defineComponent} from "vue";
import {useRouter} from "vue-router";

export default defineComponent({
  name: 'LoginPage',
  components: {
  },
  methods: {
    async handleDiscordLogin() {
      await this.$pb.collection('users').authWithOAuth2({ provider: 'discord' });

      await this.router.push('/')
    }
  },
  async mounted() {
    await this.$pb.collection('users').authRefresh();
    if (this.$pb.authStore.isValid) {
      await this.router.push('/')
    }
  },
  setup () {
    const router = useRouter();

    return { router };
  }
});
</script>
