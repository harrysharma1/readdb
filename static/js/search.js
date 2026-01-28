const BASEURL = "http://localhost:6969";
function debounce(fn, delay) {
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => fn(...args), delay);
  };
}
const searchInput = document.getElementById("searchInput");
const resultsList = document.getElementById("searchResults");
const handleSearch = debounce(() => {
  const query = searchInput.value.trim();

  if (!query) {
    resultsList.innerHTML = "";
    return;
  }
  searchAll(query);
}, 500);
searchInput.addEventListener("input",handleSearch);

function searchAll(query) {
  Promise.all([
    fetch(`${BASEURL}/books?search=${encodeURIComponent(query)}`).then((r) =>
      r.json(),
    ),
    fetch(`${BASEURL}/authors?search=${encodeURIComponent(query)}`).then((r) =>
      r.json(),
    ),
  ])
    .then(([books, authors]) => {
      renderResults([
        ...books.map((b) => ({ ...b, _type: "book" })),
        ...authors.map((a) => ({ ...a, _type: "author" })),
      ]);
    })
    .catch(console.error);
}

function renderResults(items) {
  resultsList.innerHTML = "";
  if (!items.length) {
    resultsList.innerHTML = `<li class="p-4 text-sm opacity-50">No results found</li>`;
    return;
  }
  items.forEach((item) => {
    resultsList.innerHTML += renderCard(item);
  });
}

function renderCard(item) {
  if (item._type === "book") {
    const authors = Array.isArray(item.authors) && item.authors.length ? item.authors.map(a=>a.name).join(" and ") : "Unkown Creator"
    return `
          <li class="list-row">
        <div>
          <img class="size-10 rounded-box" src="${item.profileImage}" />
        </div>
        <div>
          <div>${authors}</div>
          <div class="text-xs uppercase font-semibold opacity-60">
            ${item.name}
          </div>
          <span class="badge badge-info badge-sm mt-1">Book</span>
        </div>
        <a class="btn btn-square btn-ghost" href="/books/${item.id}">
          <span class="material-symbols-outlined">arrow_forward</span>
        </a>
      </li>
    `;
  }

  if (item._type === "author") {
    return `
          <li class="list-row">
        <div>
          <img class="size-10 rounded-box" src="${item.profileImage}" />
        </div>
        <div>
          <div>${item.name}</div>
          <span class="badge badge-success badge-sm mt-1">${item.type}</span>
        </div>
        <a class="btn btn-square btn-ghost" href="/authors/${item.id}">
          <span class="material-symbols-outlined">arrow_forward</span>
        </a>
      </li>
    `;
  }
  return "";
}

