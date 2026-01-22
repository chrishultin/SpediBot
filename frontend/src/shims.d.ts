import '@vue/runtime-core';
import type PocketBase from "pocketbase";

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $pb: PocketBase;
  }
}

// Important: Add an empty export statement to make the file a module
export {};
