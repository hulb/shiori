// This is the "Offline copy of assets" service worker

const CACHE = "pwabuilder-offline";

// importScripts('https://storage.googleapis.com/workbox-cdn/releases/5.1.2/workbox-sw.js');
const version = 'latest'
importScripts(`https://cdn.jsdelivr.net/npm/workbox-sw@${version}`)
workbox.setConfig({
  modulePathCb(name, debug) {
    const env = debug ? 'dev' : 'prod'
    return `https://cdn.jsdelivr.net/npm/${name}@${version}/build/${name}.${env}.js`
  },
})

self.addEventListener("message", (event) => {
  if (event.data && event.data.type === "SKIP_WAITING") {
    self.skipWaiting();
  }
});

workbox.routing.registerRoute(
  new RegExp('/*'),
  new workbox.strategies.NetworkFirst({
    cacheName: CACHE
  })
);