<template>
  <q-page class="row items-center justify-evenly">
    <q-card class="col-8">
      <q-card-section class="bg-primary text-h4 text-white">
        {{ server?.name }}
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script lang="ts">
import {computed, defineComponent, ref} from "vue";
import {usePocketBase} from "stores/pocketbase";
import {useRoute} from "vue-router";

export default defineComponent({
  name: 'ManageServerPage',
  props: {
  },
  components: {
  },
  methods: {

  },
  mounted() {},
  setup () {
    // const servers: Ref<Server[], Server[]> = ref([])
    const sidebarOpen = ref(false)
    const pb = usePocketBase();
    const route = useRoute()
    const serverID: string = route.params.serverID ? route.params.serverID.toString() : ''
    const server = computed(() => pb.getServer(serverID))
    return { sidebarOpen, pb, serverID, server };
  }
});
</script>
