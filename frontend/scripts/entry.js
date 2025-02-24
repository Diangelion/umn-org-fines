// Handle Login and Register page

function triggerTogglePassword(inputValue, toggleId) {
  // Check whether password value in input column exist or not
  if (inputValue) {
    document.getElementById(toggleId).classList.remove("hidden");
  } else {
    document.getElementById(toggleId).classList.add("hidden");
  }
}

function togglePassword(inputId, toggleIcon) {
  // Get the password input tag
  const input = document.getElementById(inputId);

  // Change input type to text/password depends on current state
  input.type = input.type === "password" ? "text" : "password";

  // Get current src value
  const currentSrc = toggleIcon.getAttribute("src");

  // Change src attribute's value if toggle icon is clicked
  toggleIcon.setAttribute(
    "src",
    currentSrc.includes("opened")
      ? currentSrc.replace("opened", "closed")
      : currentSrc.replace("closed", "opened"),
  );
}

// Custom events
document.addEventListener("resetForm", () => {
  ["toggle-password", "toggle-confirm-password"].forEach((toggleId) =>
    triggerTogglePassword("", toggleId),
  );
  document.getElementById("register-form").reset();
});

document.addEventListener("receiveJWT", async (e) => {
  // Get AccessToken and RefreshToken from hx-trigger in header response
  const accessToken = e?.detail?.["AccessToken"] || "";
  const refreshToken = e?.detail?.["RefreshToken"] || "";

  if (!accessToken || !refreshToken) return;

  // Store in local forage
  const ttl15Minutes = 15 * 60 * 1000; // ttl in ms
  await window.StorageModules.storeWithExpiry(
    "access_token",
    accessToken,
    ttl15Minutes,
  );

  const ttl7Days = 7 * 24 * 60 * 60 * 1000; // ttl in ms
  await window.StorageModules.storeWithExpiry(
    "refresh_token",
    refreshToken,
    ttl7Days,
  );

  document.getElementById("login-form").reset();
  window.location.href = "/home";
});

document.addEventListener("renewAccessToken", async (e) => {
  const newAccessToken = e?.detail?.value || "";
  if (!newAccessToken) {
    console.error("renewAccessToken | New access token is missing");
    return;
  }

  const ttl15Minutes = 15 * 60 * 1000; // ttl in ms
  await window.StorageModules.storeWithExpiry(
    "access_token",
    newAccessToken,
    ttl15Minutes,
  );
});
