function triggerEditProfileForm() {
  console.log(document.getElementById("edit-profile-form").classList);
  if (
    !document.getElementById("edit-profile-form").classList.contains("hidden")
  )
    return;
  document.getElementById("edit-profile-form").classList.remove("hidden");
}

function removeEditProfileForm() {
  document.getElementById("edit-profile-form").classList.add("hidden");
}

function setPreviewPhoto(e, id) {
  const file = e.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = function (e) {
      document.getElementById(id).src = e.target.result;
    };
    reader.readAsDataURL(file);
  }
}
