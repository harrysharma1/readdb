document.addEventListener("keydown", e =>{
    if ((e.metaKey || e.ctrlKey) && e.key.toLowerCase() === "k") {
        e.preventDefault();
        const searchModal = document.getElementById("bookSearch");
        const input = document.getElementById("searchInput");

        if (!searchModal) return;

        searchModal.showModal();

        setTimeout(() => input?.focus(), 50);
    }
});