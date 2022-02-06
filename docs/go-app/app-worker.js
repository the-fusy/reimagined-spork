const cacheName = "app-" + "ecb0bd2c34b9d759853bdb3854dfc27e5d8d407f";

self.addEventListener("install", event => {
  console.log("installing app worker ecb0bd2c34b9d759853bdb3854dfc27e5d8d407f");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/reimagined-spork/go-app",
          "/reimagined-spork/go-app/app.css",
          "/reimagined-spork/go-app/app.js",
          "/reimagined-spork/go-app/manifest.webmanifest",
          "/reimagined-spork/go-app/wasm_exec.js",
          "/reimagined-spork/go-app/web/app.wasm",
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
  console.log("app worker ecb0bd2c34b9d759853bdb3854dfc27e5d8d407f is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
