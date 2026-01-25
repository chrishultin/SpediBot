<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="sidebarOpen = !sidebarOpen"
        />

        <q-toolbar-title>
          SpediBot
        </q-toolbar-title>

      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="sidebarOpen"
      show-if-above
      bordered
    >
      <q-list>
        <q-item clickable @click="handleClickHome">
          Home
        </q-item>
        <q-item-label header>
          Servers
        </q-item-label>
        <span v-if="!loading">
          <sidebar-server-link
            v-for="server in pb.servers"
            v-bind:key="server.id"
            :serverID="server.id"
            :serverName="server.name"
            :iconURL="server.icon"
            :description="server.description"
          />
        </span>
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view v-if="!loading" :key="route.fullPath"/>
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import {defineComponent, ref} from 'vue';
// import type {Ref} from 'vue';
import SidebarServerLink from "components/SidebarServerLink.vue";
import {usePocketBase} from "src/stores/pocketbase";
import {useRoute, useRouter} from "vue-router";

// interface Server {
//   name: string;
//   id: string;
//   description: string;
//   icon: string;
// }

export default defineComponent({
  name: 'MainLayout',
  props: {
  },
  components: {
    SidebarServerLink
  },
  methods: {
    async handleClickHome() {
      await this.router.push("/")
    }
  },
  mounted() {
    this.pb.getServers().then(() => {this.loading = false}).catch((err) => {console.log(err)})
  },
  setup () {
    const sidebarOpen = ref(false)
    const pb = usePocketBase();
    const router = useRouter()
    const route = useRoute();
    const loading = ref(true)
    return { sidebarOpen, pb, router, route, loading };
  }
});
</script>
