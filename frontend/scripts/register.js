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

function matchPasswordAndConfirmationPassword() {
  // Submit form to trigger hx-post
  document.getElementById("register-form").requestSubmit();
}

function resetForm() {
  ["toggle-password", "toggle-confirm-password"].forEach((toggleId) =>
    triggerTogglePassword("", toggleId),
  );
  document.getElementById("register-form").reset();
}
