const BASEURL = "http://localhost:6969";
function debounce(callback, delay) {
  let timer;
  return function () {
    clearTimeout(timer);
    timer = setTimeout(() => {
      callback();
    }, delay);
  };
}

function search() {
  const searchInput = document.getElementById("searchInput");
  const searchOptions = document.getElementById("searchOptions");
  searchInput.addEventListener("input", function () {
    if (!searchInput.value.trim()) return;
    switch (searchOptions.value) {
      case "Books":
        var apiUrl = BASEURL + "/books?search=" + encodeURIComponent(searchInput.value);
        fetch(apiUrl)
          .then((response) => {
            if (!response.ok) {
              throw new Error("server response was not ok");
            }
            return response.json();
          })
          .then((data) => {
            const resultsList = document.getElementById("searchResults");
            if (data.length > 0) {
              for (let i = 0; i < data.length; i++) {
                let obj = data[i];
                resultsList.innerHTML += `
        <li class="list-row">
          <div>
            <img
              class="size-10 rounded-box"
              src="${obj.profileImage}" />
          </div>
          <div>
            <div>${obj.authors}</div>
            <div class="text-xs uppercase font-semibold opacity-60">
              ${obj.name}
            </div>
          </div>
          <a class="btn btn-square btn-ghost" href="/books/${obj.id}">go to</a>
          <button class="btn btn-square btn-ghost">remove</button>
        </li>
                
                `;
              }
            } else {
              resultsList.innerHTML += `
      <li class="p-4 text-sm opacity-50">No results found</li>
    `;
            }
          })
          .catch((error) => {
            console.error("Error:", error);
          });
        break;
      default:
        console.log("Not Implemented");
    }
  });
}

const debouncedSearch = debounce(search, 500);
debouncedSearch();