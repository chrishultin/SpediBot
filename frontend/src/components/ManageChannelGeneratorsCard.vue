<template>
  <q-expansion-item
    expand-separator
    icon="settings"
    label="Channel Generators"
    default-opened
    >
    <q-card>
      <q-card-section>
        Settings get listed here, m'kay
      </q-card-section>
    </q-card>
  </q-expansion-item>
</template>

<script lang="ts">
import {defineComponent, ref} from 'vue';
import {useRouter} from "vue-router";
import {usePocketBase} from "stores/pocketbase";

export default defineComponent({
  name: 'ManageChannelGeneratorsCard',
  components: {},
  methods: {
  },
  props: {
    serverID: {
      type: String,
      required: true,
    },
  },
  setup () {
    const expanded = ref(true)
    const router = useRouter()
    const pb = usePocketBase()
    return {router, pb, expanded}
  }
});
</script>
