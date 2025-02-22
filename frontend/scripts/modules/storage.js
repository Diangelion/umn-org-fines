window.StorageModules = {
  storeInLocalForage: async (key, value) => {
    await localforage.setItem(key, value);
  },
  getFromLocalForage: async (key) => {
    const item = await localforage.getItem(key);
    return JSON.parse(item) || "";
  },
  removeFromLocalForage: async (key) => {
    await localforage.removeItem(key);
  },
  storeWithExpiry: async (key, value, ttl) => {
    const now = Date.now(); // ✅ Current timestamp
    const item = {
      value: value,
      expiry: now + ttl, // ✅ Expiry time (ttl in milliseconds)
    };
    await storeInLocalForage(key, JSON.stringify(item));
  },
};
