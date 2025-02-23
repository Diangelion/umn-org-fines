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

// Handler function to receive JWT directly after log in
async function receiveJWT(e) {
  // Get Authorization & X-Refresh-Token from response headers
  const xhr = e?.detail?.xhr || {};
  const accessToken = xhr?.getResponseHeader("Authorization") || "";
  const refreshToken = xhr?.getResponseHeader("X-Refresh-Token") || "";

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
}

// Custom events
document.addEventListener("resetForm", () => {
  ["toggle-password", "toggle-confirm-password"].forEach((toggleId) =>
    triggerTogglePassword("", toggleId),
  );
  document.getElementById("register-form").reset();
});

document.addEventListener("refreshAccessToken", (e) => {
  console.log(e);
  // const xhr = e.detail.xhr;
  // const newAccessToken = xhr.getResponseHeader("Authorization");
  // if (newAccessToken) {
  //   console.log("✅ New Access Token Received:", newAccessToken);
  //   localStorage.setItem("accessToken", newAccessToken); // ✅ Store new token
  // }
});
