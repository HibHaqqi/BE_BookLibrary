let data;
  
      function fetchDataAndInitialize() {
          fetch("http://localhost:8000/api/books")
              .then((response) => response.json())
              .then((apiData) => {
                  data = apiData;
                  filterAndInitializeBooks();
              })
              .catch((error) => console.error("Error fetching data:", error));
      }
  
      function filterAndInitializeBooks() {
          const searchInput = document.getElementById("searchInput");
          const searchTerm = searchInput.value.toLowerCase();
  
          // Filter books based on title or author containing the search term
          const filteredBooks = data.Books.filter(
              (book) =>
                  book.Title.toLowerCase().includes(searchTerm) ||
                  book.Author.toLowerCase().includes(searchTerm)
          );
  
          // Clear the existing content in the table body
          const tableBody = document.getElementById("bookTableBody");
          tableBody.innerHTML = "";
  
          // Iterate through the filtered list of books and append rows to the table
          filteredBooks.forEach((book) => {
              const row = `
              <tr>
                  <td>${book.Title}</td>
                  <td>${book.Author}</td>
                  <td>${book.PublishYear}</td>
                  <td><div class="rateyo-readonly-widg" id="rateYo_${book.ID}"></div></td>
                  <td>
                  <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#bookDetailsModal">
                  Detail
                     </button>
                  </td>
              </tr>`;
              tableBody.innerHTML += row;
  
              // Initialize RateYo with options for each book
              $(`#rateYo_${book.ID}`).rateYo({
                rating: book.AverageRating,
                readOnly: false,
                starWidth: "20px",
                onChange: function (rating, rateYoInstance) {
                   // Handle rating change (e.g., update the rating in your data)
                   console.log(`Book ID ${book.ID} rated: ${rating}`);
          
                   // Send the rating to the server
                   submitRating(book.ID, rating);
                },
              });
          });
      }
  
      // Fetch data and initialize books when the page loads
      fetchDataAndInitialize();
  
      // Add an event listener to the search input for live filtering
      document.getElementById("searchInput").addEventListener("input", filterAndInitializeBooks);
  
      // Function to show book details (replace this with your logic)
      function showBookDetails(bookID) {
          // Implement your logic to show book details (e.g., open a modal)
          console.log("Show details for book ID:", bookID);
      }
      function submitRating(bookID, rating) {
        // Prepare data for the API request
        const requestData = {
           MemberId: 1, // Replace with the actual member ID (you may get this from your authentication system)
           BookId: bookID,
           Rating: rating,
        };
     
        // Send the rating to the server
        fetch("http://localhost:8000/api/rate", {
           method: "POST",
           headers: {
              "Content-Type": "application/json",
           },
           body: JSON.stringify(requestData),
        })
           .then((response) => response.json())
           .then((result) => {
              console.log("Rating submitted successfully:", result.rating);
              // You may update the UI or take other actions based on the server response
           })
           .catch((error) => console.error("Error submitting rating:", error));
     }
