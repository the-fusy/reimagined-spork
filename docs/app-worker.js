const cacheName = "app-" + "15a1a4fba5738df68ca6e8f3b77b55faba84e5b8";

self.addEventListener("install", event => {
  console.log("installing app worker 15a1a4fba5738df68ca6e8f3b77b55faba84e5b8");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/reimagined-spork",
          "/reimagined-spork/app.css",
          "/reimagined-spork/app.js",
          "/reimagined-spork/manifest.webmanifest",
          "/reimagined-spork/wasm_exec.js",
          "/reimagined-spork/web/app.wasm",
          "https://storage.googleapis.com/murlok-github/icon-192.png",
          "https://storage.googleapis.com/murlok-github/icon-512.png",
          "https://www.w3schools.com/w3css/4/w3.css",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 15a1a4fba5738df68ca6e8f3b77b55faba84e5b8 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
