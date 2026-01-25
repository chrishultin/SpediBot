<template>
  <q-expansion-item
    expand-separator
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
      <q-card-section>
        <q-toggle v-model="enableRename" label="Allow Owner To Rename Channel" @click="interactedRename = true"/>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn label="Save" color="primary" rounded :disable="saveDisabled" @click="handleSave"/>
      </q-card-actions>
    </q-card>
  </q-expansion-item>
</template>

<script lang="ts">
import {defineComponent, ref, computed, type ComputedRef} from 'vue';
import {useRouter} from "vue-router";
import { type channelGeneratorConfig, usePocketBase} from "stores/pocketbase";
import {useQuasar} from "quasar";

interface channelChoice {
  name: string;
  id: string;
}

export default defineComponent({
  name: 'ManageChannelGenerator',
  components: {},
  methods: {
    handleSave() {
      this.pb.updateConfig({
        channelId: this.channelChoice,
        enableRename: this.enableRename,
        nameFormat: this.nameChoice,
        serverId: this.serverID,
        id: this.configID
      }).then(() => {
        this.config = this.pb.channelGeneratorConfigsMap.get(this.configID)!
        const channel = this.pb.channelsMap.get(this.config.channelId)
        this.channelName = channel!.name
        this.q.notify({
          type: "positive",
          message: "Channel Generator Saved",
          timeout: 5000
        })
      }).catch(() => {
        this.q.notify({
          type: "negative",
          message: "Something went wrong...",
          timeout: 5000
        })
      })
    }
  },
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
    nameChoice() {
      this.saveDisabled = (this.nameChoice == this.startingNameFormat && this.startingChannel == this.channelChoice) || this.interactedRename
    },
    channelChoice() {
      this.saveDisabled = (this.nameChoice == this.startingNameFormat && this.channelChoice == this.startingChannel) || this.interactedRename
    },
    enableRename() {
      this.saveDisabled = (this.nameChoice == this.startingNameFormat && this.startingChannel == this.channelChoice) || this.interactedRename
    },
  },
  async mounted() {
    await this.pb.getChannelGeneratorConfigs(this.serverID)
    await this.pb.getChannels(this.serverID)
    this.config = this.pb.channelGeneratorConfigsMap.get(this.configID)!
    const channel = this.pb.channelsMap.get(this.config.channelId)

    this.enableRename = this.config.enableRename
    this.channelName = channel!.name
    this.channelChoice = channel!.id
    this.nameChoice = this.config.nameFormat
    this.startingNameFormat = this.config.nameFormat
    this.startingChannel = this.config.channelId
  },
  setup(props) {
    const expanded = ref(true)
    const router = useRouter()
    const pb = usePocketBase()
    const config = ref({} as channelGeneratorConfig)
    const channelName = ref("Loading")
    const label = ref("")
    const channelChoice = ref(config.value.channelId)
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

    const startingChannel = ref("")
    const startingNameFormat = ref("")

    const saveDisabled = ref(true)

    const enableRename = ref(false)
    const interactedRename = ref(false)

    const q = useQuasar()

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
      saveDisabled,
      enableRename,
      interactedRename,
      q
    }
  }
});
</script>
