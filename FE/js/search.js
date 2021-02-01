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
            myDropdown.classList.removeLet('show');
        }
    }
}

console.log("JS is working");

const theDiv = document.getElementById("searchResults");

function search() {
    //event.preventDefault()
    console.log("button clicked");
    theDiv.innerHTML = ""
    let _mySearchField = document.getElementById("search").value;
    fetch('http://localhost:8000/api/product/' + _mySearchField)
        .then(response => response.json())
        .then(data => {
            for (i = 0; i < data.length; i++) {
                console.log(data);
                theDiv.innerHTML +=
                `<div class="filterDiv ${data[i].category}">
                    <div class="product">
                    <a class="product-URL" href="productResult.html"></a>
                    <div class="product-image"> <img src="${data[i].image}"> </div>
                    <div class="product-name">${data[i].name} </div>
                    <div class="product-description"> ${data[i].description}</div>
                    <div class="product-price"> $${data[i].price}</div>
                </div></div>`;
            }
        })
        document.getElementById("search").value=""
};

