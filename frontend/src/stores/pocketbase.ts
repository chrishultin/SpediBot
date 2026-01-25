import {defineStore, acceptHMRUpdate} from 'pinia';
import PocketBase from "pocketbase";
import {ref} from "vue";

export interface discordServer {
  name: string;
  id: string;
  description: string;
  icon: string;
}

export interface discordChannel {
  name: string;
  id: string;
  serverId: string;
}

export interface channelGeneratorConfig {
  id: string;
  serverId: string;
  channelId: string;
  nameFormat: string;
  enableRename: boolean;
}

const pb = new PocketBase(process.env.DEV ? process.env.API_DEV : process.env.API)

export const usePocketBase = defineStore('pocketbase', {
  state: () => ({
    servers: [] as discordServer[],
    channelsMap: new Map<string, discordChannel>,
    channels: new Map<string, string[]>,
    channelGeneratorConfigsMap: new Map<string, channelGeneratorConfig>,
    channelGeneratorConfigs: new Map<string, string[]>,
    channelsLoading: ref(false),
    channelGeneratorConfigsLoading: ref(false),
  }),
  getters: {},
  actions: {
    async getServers() {
      const resp = await pb.send("/api/custom/servers", {method: "GET"})
      this.servers = resp.servers
    },
    async getChannels(serverID: string) {
      if (this.channels.has(serverID)) {
        return
      }
      const resp = await pb.send(`/api/custom/servers/${serverID}/channels`, {method: "GET"})
      const channels: string[] = []
      for (const channel of resp.channels) {
        this.channelsMap.set(channel.id, channel)
        channels.push(channel.id)
      }
      this.channels.set(serverID, channels)
    },
    async getChannelGeneratorConfigs(serverID: string) {
      if (this.channelGeneratorConfigs.has(serverID)) {
        return
      }
      const resp = await pb.send(`/api/custom/servers/${serverID}/channel_generators`, {method: "GET"})
      const configs: string[] = []
      for (const config of resp.configs) {
        configs.push(config.id)
        this.channelGeneratorConfigsMap.set(config.id, config)
      }
      this.channelGeneratorConfigs.set(serverID, configs)
      this.channelGeneratorConfigsLoading = false
    },
    getServer(serverID: string) {
      for (const server of this.servers) {
        if (server.id === serverID) {
          return server
        }
      }
    },
    async updateConfig(config: channelGeneratorConfig) {
      return pb.send(`/api/custom/servers/${config.serverId}/channel_generators`, {
        method: "POST", body: JSON.stringify({
          id: config.id,
          channelId: config.channelId,
          nameFormat: config.nameFormat,
          enableRename: config.enableRename,
        })
      }).then((res) => {
        this.channelGeneratorConfigsMap.set(res.config.id, config)
      })
    }
  }
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(usePocketBase, import.meta.hot));
}
