<template>
  <q-expansion-item
    expand-separator
    icon="settings"
    label="Channel Generators"
    default-opened
    >
    <q-card>
      <q-card-section>
        <q-list
          bordered
          >
          <manage-channel-generator
            v-for="config in pb.channelGeneratorConfigs.get(props.serverID)"
            v-bind:key="config"
            :serverID="props.serverID"
            :configID="config" />

        </q-list>
      </q-card-section>
    </q-card>
  </q-expansion-item>
</template>

<script lang="ts">
import {defineComponent, ref} from 'vue';
import {useRouter} from "vue-router";
import {usePocketBase} from "stores/pocketbase";
import ManageChannelGenerator from "components/ManageChannelGenerator.vue";

export default defineComponent({
  name: 'ManageChannelGeneratorsCard',
  components: {ManageChannelGenerator},
  methods: {
  },
  props: {
    serverID: {
      type: String,
      required: true,
    },
  },
  mounted () {
    this.pb.getChannelGeneratorConfigs(this.serverID)
  },
  setup (props) {
    const expanded = ref(true)
    const router = useRouter()
    const pb = usePocketBase()
    return {router, pb, expanded, props}
  }
});
</script>
