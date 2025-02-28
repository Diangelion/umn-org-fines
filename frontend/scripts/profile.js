function triggerEditProfileForm() {
  if (document.getElementById("edit-profile-form").style.display === "block")
    return;
  document.getElementById("edit-profile-form").style.display = "block";
}

function removeEditProfileForm() {
  document.getElementById("edit-profile-form").style.display = "none";
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
