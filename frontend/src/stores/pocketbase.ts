import { defineStore, acceptHMRUpdate } from 'pinia';
import PocketBase from "pocketbase";

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
const pb = new PocketBase(process.env.DEV?process.env.API_DEV:process.env.API)

export const usePocketBase = defineStore('pocketbase', {
  state: () => ({
    servers: [] as discordServer[],
    channelsMap: new Map<string, discordChannel>,
    channels: new Map<string, string[]>,
    channelGeneratorConfigsMap: new Map<string, channelGeneratorConfig>,
    channelGeneratorConfigs: new Map<string, string[]>
  }),
  getters: {
  },
  actions: {
    getServers() {
      pb.send("/api/custom/servers", {method:"GET"}).
        then((resp) => {
          this.servers = resp.servers
        }).
        catch((err) => {
          console.log(err)
      })
    },
    getChannels(serverID: string) {
      if (this.channels.has(serverID)) {
        return
      }
      pb.send(`/api/custom/servers/${serverID}/channels`, {method: "GET"}).
        then((resp) => {
          const channels: string[] = []
          for (const channel of resp.channels) {
            this.channelsMap.set(channel.id, channel)
            channels.push(channel.id)
          }
          this.channels.set(serverID, channels)
        }).
        catch((err) => {
          console.log(err)
        })
    },
    getChannelGeneratorConfigs(serverID: string) {
      if (this.channelGeneratorConfigs.has(serverID)) {
        return
      }
      pb.send(`/api/custom/servers/${serverID}/channel_generators`, {method: "GET"}).
      then((resp) => {
        const configs: string[] = []
        for (const config of resp.configs) {
          configs.push(config.id)
          this.channelGeneratorConfigsMap.set(config.id, config)
        }
        this.channelGeneratorConfigs.set(serverID, configs)
      }).
      catch((err) => {
        console.log(err)
      })
    },
    getServer(serverID: string) {
      for (const server of this.servers) {
        if (server.id === serverID) {
          return server
        }
      }
    }
  }
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(usePocketBase, import.meta.hot));
}
