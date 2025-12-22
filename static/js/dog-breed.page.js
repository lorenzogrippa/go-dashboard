document.addEventListener("DOMContentLoaded", () => {
  localAllDogBreeds();
});

function localAllDogBreeds() {
  fetch("/api/dog-breeds", { method: "GET" })
    .then((response) => response.json())
    .then((dataResponse) => {
      if (!dataResponse || !dataResponse.length) {
        return;
      }
      window.datatable = new window.simpleDatatables.DataTable(
        "#dogBreedsDataTable",
        {
          data: {
            data: dataResponse.map((item) => Object.values(item)),
          },
          columns: [
            {
              select: 1,
              render: (data, id, rowIndex, cellIndex) => {
                debugger;
                return `<a href="/dog-breed/${dataResponse[rowIndex].id}">${data[0].data}</a>`;
              },
            },
            {
              select: 4,
              render: (data) =>
                `<div class="text-center">${data[0].data}</div>` || "N/A",
            },
            {
              select: 5,
              render: (data) =>
                `<div class="text-center">${data[0].data}</div>` || "N/A",
            },
            { select: [0, 2, 3, 6, 7, 8], hidden: true },
          ],
        }
      );
    })
    .catch((error) => {
      console.error("Error fetching dog breeds:", error);
    });
}
