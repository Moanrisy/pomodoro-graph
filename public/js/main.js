// Add squares

const options = {
  year: "numeric",
  month: "short",
  day: "numeric",
};

// Get the current date and time
var currentDate = new Date();
var currentMonth = currentDate.getMonth();
var initialMonth = currentDate.getMonth();
var repeatedMonth = false;

// Set the day of the month to yesterday's date
currentDate.setDate(currentDate.getDate() - 1);

const months = document.querySelector(".months");
const squares = document.querySelector(".squares");
for (var i = 1; i < 367; i++) {
  const formattedDate = currentDate.toLocaleDateString("en-US", options);

  if (currentDate.getMonth() != currentMonth) {
    var getMonthFromDate = new Date(currentDate.valueOf());
    getMonthFromDate.setMonth(currentDate.getMonth() + 1);
    // currentDate.setMonth(currentDate.getMonth() - 2);

    months.insertAdjacentHTML(
      "beforeend",
      `<li>${getMonthFromDate.toLocaleDateString("en-US", {
        month: "short",
      })}</li>`
    );

    currentMonth = currentDate.getMonth();
    if (currentMonth == initialMonth) {
      repeatedMonth = true;
    }
  }

  // Move to the next day
  currentDate.setDate(currentDate.getDate() + 1);
  // If we are at the last day of the current month, move to the first day of the next month
  if (currentDate.getDate() === 1) {
    currentDate.setDate(1);
    currentDate.setMonth(currentDate.getMonth() - 2);
  }

  // If we have reached the 11th previous month, break out of the loop
  console.log(currentDate.getMonth());
  if (repeatedMonth && currentDate.getMonth() === initialMonth) {
    console.log("Hmmm");
    var getMonthFromDate = new Date(currentDate.valueOf());
    // getMonthFromDate.setMonth(currentDate.getMonth() + 1);
    months.insertAdjacentHTML(
      "beforeend",
      `<li>${getMonthFromDate.toLocaleDateString("en-US", {
        month: "short",
      })}</li>`
    );
    // break;
    repeatedMonth = false
  }

  const level = Math.floor(Math.random() * 3);
  squares.insertAdjacentHTML(
    "beforeend",
    `<li data-date="${i}" data-level="${level}">          <span class="tooltiptext">${formattedDate}</span> </li>`
  );
}

const hmms = document.querySelectorAll(".squares li");
hmms.forEach(function (sq) {
  sq.addEventListener("mouseover", function () {
    // console.log(sq.dataset.date)
    const tooltip = sq.querySelector(".tooltiptext");
    tooltip.style.visibility = "visible";
  });
  sq.addEventListener("mouseout", function () {
    const tooltip = sq.querySelector(".tooltiptext");
    tooltip.style.visibility = "hidden";
  });
});
