window.AuthModules = {
  getAccessToken: async () =>
    await window.StorageModules.getWithExpiry("access_token"),
  getRefreshToken: async () =>
    await window.StorageModules.getWithExpiry("refresh_token"),
};
