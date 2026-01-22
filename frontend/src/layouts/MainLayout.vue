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

        <sidebar-server-link
          v-for="server in pb.servers"
          v-bind:key="server.id"
          :serverID="server.id"
          :serverName="server.name"
          :iconURL="server.icon"
          :description="server.description"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import {defineComponent, ref} from 'vue';
// import type {Ref} from 'vue';
import SidebarServerLink from "components/SidebarServerLink.vue";
import {usePocketBase} from "src/stores/pocketbase";
import {useRouter} from "vue-router";

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
    this.pb.getServers()
  },
  setup () {
    // const servers: Ref<Server[], Server[]> = ref([])
    const sidebarOpen = ref(false)
    const pb = usePocketBase();
    const router = useRouter()
    return { sidebarOpen, pb, router };
  }
});
</script>
