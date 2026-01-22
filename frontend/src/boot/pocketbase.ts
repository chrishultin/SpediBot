import { defineBoot } from '#q-app/wrappers'

import PocketBase from 'pocketbase';


export default defineBoot(({app}) => {
  app.config.globalProperties.$pb = new PocketBase(process.env.DEV?process.env.API_DEV:process.env.API)
})
