// Global Scripts

function dismissAlert() {
  document.getElementById("alert-message").remove();
}

export function storeInLocalForage(key, value) {
  localforage.setItem(key, value);
}

export function getFromLocalForage(key) {
  return localforage.getItem(key) || "";
}

export function removeFromLocalForage(key) {
  localforage.removeItem(key);
}

function storeWithExpiry(key, value, ttl) {
  const now = Date.now(); // ✅ Current timestamp
  const item = {
    value: value,
    expiry: now + ttl, // ✅ Expiry time (ttl in milliseconds)
  };
  storeInLocalForage(key, JSON.stringify(item));
}
