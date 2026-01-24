<template>
  <q-item
    clickable
    @click="handleClick"
  >
    <q-item-section
      avatar
    >
      <q-icon :name="'img:'+ iconURL" size="32px" v-if="iconURL"/>
      <q-icon name="sports_esports" size="32px" v-else/>
    </q-item-section>

    <q-item-section>
      <q-item-label>{{ serverName }}</q-item-label>
      <q-item-label caption>{{ description }}</q-item-label>
    </q-item-section>
  </q-item>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import {useRouter} from "vue-router";

export default defineComponent({
  name: 'SidebarServerLink',
  components: {},
  methods: {
    async handleClick() {
      await this.router.push("/servers/" + this.serverID)
    }
  },
  props: {
    serverID: {
      type: String,
      required: true,
    },
    serverName: {
      type: String,
      required: true,
    },
    iconURL: {
      type: String,
      required: false,
    },
    description: {
      type: String,
      required: false,
    }
  },
  setup () {
    const router = useRouter()
    return {router}
  }
});
</script>
