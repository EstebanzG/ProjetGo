window.addEventListener("DOMContentLoaded", (event) => {
  // WINDOWS APEX
  window.Apex = {
    chart: {
      foreColor: "#fff",
      toolbar: {
        show: false,
      },
    },
    colors: ["#FCCF31", "#17ead9", "#f02fc2"],
    stroke: {
      width: 3,
    },
    dataLabels: {
      enabled: false,
    },
    grid: {
      borderColor: "#40475D",
    },
    xaxis: {
      axisTicks: {
        color: "#333",
      },
      axisBorder: {
        color: "#333",
      },
    },
    fill: {
      type: "gradient",
      gradient: {
        gradientToColors: ["#F55555", "#6078ea", "#6094ea"],
      },
    },
    tooltip: {
      theme: "dark",
      x: {
        formatter: function (val) {
          return moment(new Date(val)).format("HH:mm:ss");
        },
      },
    },
    yaxis: {
      decimalsInFloat: 2,
      opposite: true,
      labels: {
        offsetX: -10,
      },
    },
  };
  // FUNCTION DATE
  function generateDate() {
    var date = new Date();
    let day = date.getDate();
    if (day < 10) {
      day = "0" + day;
    }
    let month = date.getMonth() + 1;
    if (month < 10) {
      month = "0" + month;
    }
    let year = date.getFullYear();
    return (currentDate = `${year}-${month}-${day}`);
  }
  function generateCurrentDate() {
    var date = new Date();
    let day = date.getDate();
    if (day < 10) {
      day = "0" + day;
    }
    let month = date.getMonth() + 1;
    if (month < 10) {
      month = "0" + month;
    }
    let year = date.getFullYear();
    let hour = date.getHours();
    if (hour < 10) {
      hour = "0" + hour;
    }
    let min = date.getMinutes();
    if (min < 10) {
      min = "0" + min;
    }
    return (currentDate = `${year}-${month}-${day}-${hour}:${min}`);
  }
  function generatePrevioustDate() {
    var date = new Date();
    let day = date.getDate();
    if (day < 10) {
      day = "0" + day;
    }
    let month = date.getMonth() + 1;
    if (month < 10) {
      month = "0" + month;
    }
    let year = date.getFullYear();
    let hour = date.getHours();
    if (hour < 10) {
      hour = "0" + hour;
    }
    let min = date.getMinutes();
    if (min < 10) {
      min = "0" + min;
    }
    let minBefore = date.getMinutes() - 1;
    if (minBefore < 10) {
      minBefore = "0" + minBefore;
    }
    return (previousDate = `${year}-${month}-${day}-${hour}:${minBefore}`);
  }
  function generateCompleteDate() {
    var date = new Date();
    let day = date.getDate();
    if (day < 10) {
      day = "0" + day;
    }
    let month = date.getMonth() + 1;
    if (month < 10) {
      month = "0" + month;
    }
    let year = date.getFullYear();
    let hour = date.getHours();
    if (hour < 10) {
      hour = "0" + hour;
    }
    let min = date.getMinutes();
    if (min < 10) {
      min = "0" + min;
    }
    let sec = date.getSeconds();
    if (sec < 10) {
      sec = "0" + sec;
    }
    return (completeDate = `${year}-${month}-${day}-${hour}:${min}:${sec}`);
  }
  function generatePreviousCompleteDate() {
    var date = new Date();
    let day = date.getDate();
    if (day < 10) {
      day = "0" + day;
    }
    let month = date.getMonth() + 1;
    if (month < 10) {
      month = "0" + month;
    }
    let year = date.getFullYear();
    let hour = date.getHours();
    if (hour < 10) {
      hour = "0" + hour;
    }
    let min = date.getMinutes();
    if (min < 10) {
      min = "0" + min;
    }
    let sec = date.getSeconds();
    if (sec < 10) {
      sec = "0" + sec;
    }
    let secBefore = date.getSeconds() - 9;
    if (secBefore < 10) {
      secBefore = "0" + secBefore;
    }
    return (completeDate = `${year}-${month}-${day}-${hour}:${min}:${secBefore}`);
  }
  // TEMPERATURE GAUGE
  var optionsTemp = {
    chart: {
      height: 250,
      type: "radialBar",
    },
    series: [0],
    colors: ["sandybrown"],
    plotOptions: {
      radialBar: {
        startAngle: -90,
        endAngle: 90,
        hollow: {
          margin: 0,
          size: "70%",
        },
        track: {
          startAngle: -90,
          endAngle: 90,
          dropShadow: {
            enabled: true,
            top: 2,
            left: 0,
            blur: 4,
            opacity: 0.15,
          },
        },
        dataLabels: {
          name: {
            offsetY: -30,
            color: "#fff",
            fontSize: "13px",
          },
          value: {
            offsetY: -15,
            color: "#fff",
            fontSize: "30px",
            show: true,
            formatter: function (value) {
              if (value == 0) {
                return value + "°C";
              } else return value - 10 + "°C";
            },
          },
        },
      },
    },
    fill: {
      type: "gradient",
      gradient: {
        shade: "dark",
        type: "vertical",
        gradientToColors: ["peachpuff"],
        stops: [0, 100],
      },
    },
    stroke: {
      lineCap: "round",
    },
    labels: ["Temperature"],
  };
  var chartTemp = new ApexCharts(
    document.querySelector("#chartTemp"),
    optionsTemp
  );
  chartTemp.render();

  // WIND GAUGE
  var optionsWind = {
    chart: {
      height: 250,
      type: "radialBar",
    },
    series: [0],
    colors: ["deepskyblue"],
    plotOptions: {
      radialBar: {
        startAngle: -90,
        endAngle: 90,
        hollow: {
          margin: 0,
          size: "70%",
        },
        track: {
          startAngle: -90,
          endAngle: 90,
          dropShadow: {
            enabled: true,
            top: 2,
            left: 0,
            blur: 4,
            opacity: 0.15,
          },
        },
        dataLabels: {
          name: {
            offsetY: -30,
            color: "#fff",
            fontSize: "13px",
          },
          value: {
            offsetY: -15,
            color: "#fff",
            fontSize: "30px",
            show: true,
            formatter: function (value) {
              return value + " km/h";
            },
          },
        },
      },
    },
    fill: {
      type: "gradient",
      gradient: {
        shade: "dark",
        type: "vertical",
        gradientToColors: ["skyblue"],
        stops: [0, 100],
      },
    },
    stroke: {
      lineCap: "round",
    },
    labels: ["Wind"],
  };
  var chartWind = new ApexCharts(
    document.querySelector("#chartWind"),
    optionsWind
  );
  chartWind.render();

  // PRESSURE GAUGE
  var optionsPressure = {
    chart: {
      height: 250,
      type: "radialBar",
    },
    series: [0],
    colors: ["seagreen"],
    plotOptions: {
      radialBar: {
        startAngle: -90,
        endAngle: 90,
        hollow: {
          margin: 0,
          size: "70%",
        },
        track: {
          startAngle: -90,
          endAngle: 90,
          dropShadow: {
            enabled: true,
            top: 2,
            left: 0,
            blur: 4,
            opacity: 0.15,
          },
        },
        dataLabels: {
          name: {
            offsetY: -30,
            color: "#fff",
            fontSize: "13px",
          },
          value: {
            offsetY: -15,
            color: "#fff",
            fontSize: "30px",
            show: true,
            formatter: function (value) {
              return value * 20 + " P";
            },
          },
        },
      },
    },
    fill: {
      type: "gradient",
      gradient: {
        shade: "dark",
        type: "vertical",
        gradientToColors: ["lightgreen"],
        stops: [0, 100],
      },
    },
    stroke: {
      lineCap: "round",
    },
    labels: ["Pressure"],
  };
  var chartPressure = new ApexCharts(
    document.querySelector("#chartPressure"),
    optionsPressure
  );
  chartPressure.render();

  window.setInterval(() => {
    var dateComplete = generateCompleteDate();
    var dateCompletePrevious = generatePreviousCompleteDate();
    fetch("http://localhost:8080/data/temperature")
      .then((response) => response.json())
      .then((data) => {
        data.forEach((element) => {
          if (
            element.date > dateCompletePrevious &&
            element.date <= dateComplete
          ) {
            //console.log("Temperature : " + element.value);
            var val = Math.round(element.value * 10) / 10;
            chartTemp.updateSeries([val + 10]);
          }
        });
      });
    fetch("http://localhost:8080/data/wind")
      .then((response) => response.json())
      .then((data) => {
        data.forEach((element) => {
          if (
            element.date > dateCompletePrevious &&
            element.date <= dateComplete
          ) {
            //console.log("Vitesse du vent : " + element.value);
            chartWind.updateSeries([element.value]);
          }
        });
      });
    fetch("http://localhost:8080/data/pressure")
      .then((response) => response.json())
      .then((data) => {
        data.forEach((element) => {
          if (
            element.date > dateCompletePrevious &&
            element.date <= dateComplete
          ) {
            //console.log("Pression : " + element.value);
            chartPressure.updateSeries([Math.round(element.value / 20)]);
          }
        });
      });
  }, 3000);

  // AVERAGE CHART

  var optionsProgress1 = {
    chart: {
      height: 70,
      type: "bar",
      stacked: true,
      sparkline: {
        enabled: true,
      },
    },
    plotOptions: {
      bar: {
        horizontal: true,
        barHeight: "20%",
        colors: {
          backgroundBarColors: ["#40475D"],
        },
      },
    },
    stroke: {
      width: 0,
    },
    series: [
      {
        name: "Process 1",
        data: [44],
      },
    ],
    title: {
      floating: true,
      offsetX: -10,
      offsetY: 5,
      text: "Moyenne Vent",
    },
    subtitle: {
      floating: true,
      align: "right",
      offsetY: 0,
      text: "km/h",
      style: {
        fontSize: "20px",
      },
    },
    tooltip: {
      enabled: false,
    },
    xaxis: {
      categories: ["Process 1"],
    },
    yaxis: {
      max: 100,
    },
    fill: {
      opacity: 1,
    },
  };

  var chartProgress1 = new ApexCharts(
    document.querySelector("#progress1"),
    optionsProgress1
  );
  chartProgress1.render();

  var cd = generateDate();
  fetch(`http://localhost:8080/average/${cd}`)
    .then((response) => response.json())
    .then((data) => {
      chartProgress1.updateSeries([
        {
          name: "Process 1",
          data: [data.temperature_average],
        },
      ]);
    });

  var optionsProgress2 = {
    chart: {
      height: 70,
      type: "bar",
      stacked: true,
      sparkline: {
        enabled: true,
      },
    },
    plotOptions: {
      bar: {
        horizontal: true,
        barHeight: "20%",
        colors: {
          backgroundBarColors: ["#40475D"],
        },
      },
    },
    colors: ["#17ead9"],
    stroke: {
      width: 0,
    },
    series: [
      {
        name: "Process 2",
        data: [80],
      },
    ],
    title: {
      floating: true,
      offsetX: -10,
      offsetY: 5,
      text: "Moyenne Pressure",
    },
    subtitle: {
      floating: true,
      align: "right",
      offsetY: 0,
      text: "80%",
      style: {
        fontSize: "20px",
      },
    },
    tooltip: {
      enabled: false,
    },
    xaxis: {
      categories: ["Process 2"],
    },
    yaxis: {
      max: 100,
    },
    fill: {
      type: "gradient",
      gradient: {
        inverseColors: false,
        gradientToColors: ["#6078ea"],
      },
    },
  };

  var chartProgress2 = new ApexCharts(
    document.querySelector("#progress2"),
    optionsProgress2
  );
  chartProgress2.render();

  var optionsProgress3 = {
    chart: {
      height: 70,
      type: "bar",
      stacked: true,
      sparkline: {
        enabled: true,
      },
    },
    plotOptions: {
      bar: {
        horizontal: true,
        barHeight: "20%",
        colors: {
          backgroundBarColors: ["#40475D"],
        },
      },
    },
    colors: ["#f02fc2"],
    stroke: {
      width: 0,
    },
    series: [
      {
        name: "Process 3",
        data: [74],
      },
    ],
    fill: {
      type: "gradient",
      gradient: {
        gradientToColors: ["#6094ea"],
      },
    },
    title: {
      floating: true,
      offsetX: -10,
      offsetY: 5,
      text: "Moyenne Temperature",
    },
    subtitle: {
      floating: true,
      align: "right",
      offsetY: 0,
      text: "74%",
      style: {
        fontSize: "20px",
      },
    },
    tooltip: {
      enabled: false,
    },
    xaxis: {
      categories: ["Process 3"],
    },
    yaxis: {
      max: 100,
    },
  };

  var chartProgress3 = new ApexCharts(
    document.querySelector("#progress3"),
    optionsProgress3
  );
  chartProgress3.render();

  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  //
  /*
  // REAL TIME FUNCTION

  window.setInterval(function () {
    iteration++;

    //
    //
    // DATE
    //
    //
    var date = new Date();
    let day = date.getDate();
    if (day < 10) {
      day = "0" + day;
    }
    let month = date.getMonth() + 1;
    if (month < 10) {
      month = "0" + month;
    }
    let year = date.getFullYear();
    let hour = date.getHours();
    if (hour < 10) {
      hour = "0" + hour;
    }
    let min = date.getMinutes();
    if (min < 10) {
      min = "0" + min;
    }
    let minBefore = date.getMinutes() - 1;
    if (minBefore < 10) {
      minBefore = "0" + minBefore;
    }
    let currentDate = `${year}-${month}-${day}-${hour}:${min}`;
    console.log(currentDate);
    let previousDate = `${year}-${month}-${day}-${hour}:${minBefore}`;
    console.log(previousDate);
    //
    //
    // TEMPERATURE
    //
    //
    //
    //
    // PRESSURE
    //
    //
    chartLine.updateSeries([
      {
        data: [
          ...chartLine.w.config.series[0].data,
          [chartLine.w.globals.maxX + 300000, getRandom()],
        ],
      },
    ]);

    var tab = [];
    var urlChartLine2 = `http://localhost:8080/data/pressure/${previousDate}/${currentDate}`;
    fetch(urlChartLine2)
      .then((response) => response.json())
      .then((data) => {
        data.forEach((element) => {
          tab.push(element.value);
        });
        console.log(tab);
        chartLine2.updateSeries([{ data: tab }]);
      })
      .catch((error) => {
        console.log("LOAD -> " + error);
      });

    chartLine3.updateSeries([
      {
        data: [
          ...chartLine3.w.config.series[0].data,
          [chartLine3.w.globals.maxX + 300000, getRandom()],
        ],
      },
    ]);

    var p1Data = getRangeRandom({ min: 10, max: 100 });
    chartProgress1.updateOptions({
      series: [
        {
          data: [p1Data],
        },
      ],
      subtitle: {
        text: p1Data + "%",
      },
    });

    var p2Data = getRangeRandom({ min: 10, max: 100 });
    chartProgress2.updateOptions({
      series: [
        {
          data: [p2Data],
        },
      ],
      subtitle: {
        text: p2Data + "%",
      },
    });

    var p3Data = getRangeRandom({ min: 10, max: 100 });
    chartProgress3.updateOptions({
      series: [
        {
          data: [p3Data],
        },
      ],
      subtitle: {
        text: p3Data + "%",
      },
    });
  }, 3000);
  */
});
