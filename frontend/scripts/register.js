function triggerTogglePassword(inputValue, toggleId) {
  if (inputValue) {
    document.getElementById(toggleId).classList.remove("hidden");
  } else {
    document.getElementById(toggleId).classList.add("hidden");
  }
}

function togglePassword(inputId, toggleIcon) {
  const input = document.getElementById(inputId);
  input.type = input.type === "password" ? "text" : "password";
  const currentHxGet = toggleIcon.getAttribute("hx-get");
  toggleIcon.setAttribute(
    "hx-get",
    currentHxGet.includes("opened")
      ? "icons/eye_closed.html"
      : "icons/eye_opened.html",
  );
  htmx.ajax("GET", toggleIcon.getAttribute("hx-get"), {
    target: toggleIcon,
    swap: "innerHTML",
  });
}
