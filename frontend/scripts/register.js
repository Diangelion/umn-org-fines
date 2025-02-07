function triggerTogglePassword(inputValue, toggleId) {
  // Check whether password value in input column exist or not
  if (inputValue) {
    document.getElementById(toggleId).classList.remove("hidden");
  } else {
    document.getElementById(toggleId).classList.add("hidden");
  }

  // Clear the error message
  const errorContainer = document.getElementById("register-error-message");
  if (errorContainer.innerText) errorContainer.innerText = "";
}

function togglePassword(inputId, toggleIcon) {
  // Get the password input tag
  const input = document.getElementById(inputId);

  // Change input type to text/password depends on current state
  input.type = input.type === "password" ? "text" : "password";

  // Get current hx-get state
  const currentHxGet = toggleIcon.getAttribute("hx-get");

  // Change hx-get attribute's value if toggle icon is clicked
  toggleIcon.setAttribute(
    "hx-get",
    currentHxGet.includes("opened")
      ? "icons/eye_closed.html"
      : "icons/eye_opened.html",
  );

  // Re-trigger the hx-get to render the reusable icon
  htmx.ajax("GET", toggleIcon.getAttribute("hx-get"), {
    target: toggleIcon,
    swap: "innerHTML",
  });
}

function matchPasswordAndConfirmationPassword() {
  // Get the needed tag
  const password = document.getElementById("register-password").value;
  const confirmPassword = document.getElementById(
    "register-confirm-password",
  ).value;
  const errorContainer = document.getElementById("register-error-message");

  // Clear previous errors
  errorContainer.innerText = "";

  // Validation check
  if (password !== confirmPassword) {
    errorContainer.innerText =
      "Password and Confirmation Password do not match!";
    return;
  }

  // Submit form to trigger hx-post
  document.getElementById("register-form").requestSubmit();
}
