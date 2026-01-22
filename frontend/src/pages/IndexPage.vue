<template>
  <q-page class="row items-center justify-evenly">
    <q-card class="col-8 col-sm-12 col-md-8 col-lg-4" v-for="server in pb.servers" v-bind:key="server.id" style="min-height: 400px">
      <server-card-component :serverID="server.id" :serverName="server.name" :iconURL="server.icon" :description="server.description"/>
    </q-card>
  </q-page>
</template>

<script lang="ts">
import ServerCardComponent from "components/ServerCardComponent.vue";
import {defineComponent, ref} from "vue";
// import SidebarServerLink from "components/SidebarServerLink.vue";
import {usePocketBase} from "stores/pocketbase";

export default defineComponent({
  name: 'IndexPage',
  props: {
  },
  components: {
    ServerCardComponent
  },
  methods: {

  },
  mounted() {
    this.pb.getServers()
  },
  setup () {
    // const servers: Ref<Server[], Server[]> = ref([])
    const sidebarOpen = ref(false)
    const pb = usePocketBase();
    return { sidebarOpen, pb };
  }
});
</script>
