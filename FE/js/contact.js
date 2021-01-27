/* When the user clicks on the button, 
toggle between hiding and showing the dropdown content */
function myFunction() {
  document.getElementById("myDropdown").classList.toggle("show");
}

// Close the dropdown if the user clicks outside of it
window.onclick = function (e) {
  if (!e.target.matches('.dropbtn')) {
    var myDropdown = document.getElementById("myDropdown");
    if (myDropdown.classList.contains('show')) {
      myDropdown.classList.remove('show');
    }
  }
}

//Once submit form is pressed, this will happen
function contactUs() {
  event.preventDefault()
  validateForm();
}

//Validate first name, last name, and email
function validateForm() {
  let a = document.forms["contact-inputs"]["fname"].value;
  if (a == "") {
    alert("First name must be filled out");
    return false;
  }
  let b = document.forms["contact-inputs"]["lname"].value;
  if (b == "") {
    alert("Last name must be filled out");
    return false;
  }
  let c = document.forms["contact-inputs"]["email"].value;
  if(emailIsValid(c)==false) {
    alert("Must input a valid email");
    return false;
  }
}

//checks if email is s@s.s 
function emailIsValid (email) {
  return /\S+@\S+\.\S+/.test(email)
}