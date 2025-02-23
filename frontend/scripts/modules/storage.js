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
      expiry: now + ttl, // Expiry time = now + ttl/time to live (ttl in milliseconds)
    };

    await window.StorageModules.storeInLocalForage(key, JSON.stringify(item));
  },
  getWithExpiry: async (key) => {
    const itemStr = await localforage.getItem(key);
    if (!itemStr) return ""; // Item does not exist

    const item = JSON.parse(itemStr);
    const now = Date.now();

    // Check if the item has expired
    if (now > item.expiry) {
      await window.StorageModules.removeFromLocalForage(key); // Remove expired item
      return ""; // Item has expired
    }

    return item.value; // Return the value if not expired
  },
};
