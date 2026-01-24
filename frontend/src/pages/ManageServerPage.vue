<template>
  <q-page class="row items-center justify-evenly">
    <q-card class="col-8">
      <q-card-section class="bg-primary text-h4 text-white">
        {{ server?.name }}
      </q-card-section>
      <q-card-section>
        <manage-channel-generators-card :serverID="serverID"/>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script lang="ts">
import {computed, defineComponent, reactive, ref} from "vue";
import {usePocketBase} from "stores/pocketbase";
import {useRoute} from "vue-router";
import ManageChannelGeneratorsCard from "components/ManageChannelGeneratorsCard.vue";

export default defineComponent({
  name: 'ManageServerPage',
  props: {
  },
  components: {
    ManageChannelGeneratorsCard
  },
  methods: {

  },
  mounted() {
  },
  setup () {
    const sidebarOpen = ref(false)
    const pb = usePocketBase();
    const route = useRoute()
    const serverID: string = route.params.serverID ? route.params.serverID.toString() : ''
    pb.getChannels(serverID)
    const server = computed(() => pb.getServer(serverID))
    const channels = reactive(pb.channels)
    return { sidebarOpen, pb, serverID, server, channels};
  }
});
</script>
