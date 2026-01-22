import { defineStore, acceptHMRUpdate } from 'pinia';
import PocketBase from "pocketbase";

interface discordServer {
  name: string;
  id: string;
  description: string;
  icon: string;
}

const pb = new PocketBase(process.env.DEV?process.env.API_DEV:process.env.API)

export const usePocketBase = defineStore('pocketbase', {
  state: () => ({
    servers: [] as discordServer[]
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
