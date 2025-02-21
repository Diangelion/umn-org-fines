document.body.addEventListener("jwtReceived", function (event) {
  console.log("✅ `jwtReceived` event triggered!");

  const xhr = event.detail.xhr;
  const token = xhr.getResponseHeader("Authorization");

  if (token) {
    console.log("✅ Storing JWT Token:", token);
    localStorage.setItem("refresh_token", token);
  }
});
