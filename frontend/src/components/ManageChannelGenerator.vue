<template>
  <q-expansion-item
    expand-separator
    default-opened
    :label="channelName"
  >
    <q-card>
      <q-card-section>
        <q-select :options="channelOptions"
                  v-model="channelChoice"
                  option-value="id"
                  option-label="name"
                  emit-value
                  map-options
                  label="Channel"
                  stack-label
                  filled
        />
      </q-card-section>
      <q-card-section>
        <q-select
          filled
          label="Name Format"
          stack-label
          :options="nameOptions"
          v-model="nameChoice"
          hint="How to name the generated channels"
          hide-hint
          map-options
          emit-value/>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn label="Save" color="primary" rounded :disable="saveDisabled"/>
      </q-card-actions>
    </q-card>
  </q-expansion-item>
</template>

<script lang="ts">
import {reactive, defineComponent, ref, computed, type ComputedRef} from 'vue';
import {useRouter} from "vue-router";
import {type channelGeneratorConfig, usePocketBase} from "stores/pocketbase";

interface channelChoice {
  name: string;
  id: string;
}

export default defineComponent({
  name: 'ManageChannelGenerator',
  components: {},
  methods: {},
  props: {
    serverID: {
      type: String,
      required: true,
    },
    configID: {
      type: String,
      required: true,
    }
  },
  watch: {
    config(newConf: channelGeneratorConfig) {
      const channel = this.pb.channelsMap.get(newConf.channelId)
      if (channel != undefined) {
        this.channelName = channel.name
      }
    },
    nameChoice(newFormat: string) {
      this.saveDisabled = (newFormat != this.startingNameFormat || this.startingChannel != this.channelChoice)
    },
    channelChoice(newChannel: string) {
      this.saveDisabled = (this.nameChoice != this.startingNameFormat || newChannel != this.channelChoice)
    }
  },
  mounted() {
    const channel = this.pb.channelsMap.get(this.config.channelId)
    if (channel != undefined) {
      this.channelName = channel.name
    }
  },
  setup(props) {
    const expanded = ref(true)
    const router = useRouter()
    const pb = usePocketBase()
    const config = reactive(pb.channelGeneratorConfigsMap.get(props.configID)!)
    const channelName = ref("Loading")
    const label = ref("")
    const channelChoice = ref(config.channelId)
    const channelOptions: ComputedRef<channelChoice[]> = computed(() => {
      const channels = [] as channelChoice[]
      for (const channel of pb.channels.get(props.serverID)!) {
        const name = pb.channelsMap.get(channel)!.name
        if (name != undefined) {
          channels.push({name: name, id: channel})
        }
      }

      return channels
    })
    const nameOptions = [{value: "owner", label: "By Owner's Name"}, {value: "index", label: "By Index"}]
    const nameChoice = ref("index")

    const startingChannel = config.channelId
    const startingNameFormat = config.nameFormat

    const saveDisabled = ref(true)

    return {
      router,
      pb,
      expanded,
      props,
      config,
      label,
      channelName,
      channelOptions,
      channelChoice,
      nameOptions,
      nameChoice,
      startingChannel,
      startingNameFormat,
      saveDisabled
    }
  }
});
</script>
