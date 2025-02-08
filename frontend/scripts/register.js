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
  // Get the needed tag
  const password = document.getElementById("register-password").value;
  const confirmPassword = document.getElementById(
    "register-confirm-password",
  ).value;

  // Validation check
  if (password !== confirmPassword) {
  }

  // Submit form to trigger hx-post
  document.getElementById("register-form").requestSubmit();
}

function errorRegister(event) {
  const response = event.detail?.errorInfo?.xhr?.response;
  document.getElementById("register-error-container").outerHTML = response;
}
