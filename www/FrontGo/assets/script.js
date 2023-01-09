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
          return "Valeur du " + val;
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
  function generatePreviousDate() {
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
    let minBefore = date.getMinutes() - 2;
    if (minBefore < 0) {
      minBefore = 60 + minBefore;
      hour = hour - 1;
    } else if (minBefore < 10) {
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
    let secBefore = date.getSeconds() - 10;
    if (secBefore < 0) {
      secBefore = 60 + secBefore;
      min = min - 1;
    } else if (secBefore < 10) {
      secBefore = "0" + secBefore;
    }
    return (completeDate = `${year}-${month}-${day}-${hour}:${min}:${secBefore}`);
  }
  function generateDateTitle() {
    var date = new Date();
    let today = date.getDay();
    switch (today) {
      case 0:
        today = "SUNDAY";
        break;
      case 1:
        today = "MONDAY";
        break;
      case 2:
        today = "TUESDAY";
        break;
      case 3:
        today = "WEDNESDAY";
        break;
      case 4:
        today = "THURSDAY";
        break;
      case 5:
        today = "FRIDAY";
        break;
      case 6:
        today = "SATURDAY";
        break;
      default:
        today = "SUNDAY";
    }
    let day = date.getDate();
    if (day < 10) {
      day = "0" + day;
    }
    let month = date.getMonth() + 1;
    let monthDOTD = "JAN";
    switch (month) {
      case 0:
        monthDOTD = "DEC";
        break;
      case 1:
        monthDOTD = "JAN";
        break;
      case 2:
        monthDOTD = "FEB";
        break;
      case 3:
        monthDOTD = "MAR";
        break;
      case 4:
        monthDOTD = "APR";
        break;
      case 5:
        monthDOTD = "MAY";
        break;
      case 6:
        monthDOTD = "JUN";
        break;
      case 7:
        monthDOTD = "JUL";
        break;
      case 8:
        monthDOTD = "AUG";
        break;
      case 9:
        monthDOTD = "SEP";
        break;
      case 10:
        monthDOTD = "OCT";
        break;
      case 11:
        monthDOTD = "NOV";
        break;
      default:
        monthDOTD = "JAN";
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
    let DOTD = `${today} ${day} ${monthDOTD} ${year} ${hour}:${min}`;
    document.getElementById("DOTD").innerText = DOTD;
  }
  // TEMPERATURE GAUGE
  var optionsTemp = {
    chart: {
      height: 250,
      type: "radialBar",
    },
    series: [0],
    colors: ["#e4bd36"],
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
              return "0°C";
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
        gradientToColors: ["#ea5454"],
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
    colors: ["#5f78e9"],
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
              return "0 km/h";
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
        gradientToColors: ["#22d9cd"],
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
    colors: ["#d834b4"],
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
              return "0 P";
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
        gradientToColors: ["#6592e9"],
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
        data: [0],
      },
    ],
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
      text: "0°C",
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
        data: [0],
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
      text: "0 km/h",
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
        data: [0],
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
      text: "Moyenne Pression",
    },
    subtitle: {
      floating: true,
      align: "right",
      offsetY: 0,
      text: "0 P",
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

  // LINE CHART

  // Generate default line chart values and date on X-axis
  var date = new Date();
  var tabDate = [];
  var tabData = [];
  for (let index = 8; index >= 0; index--) {
    var m = date.getMonth() + 1;
    var d = date.getDate() - index;
    if (index == 8) {
      tabDate.push("....");
      tabData.push(0);
    } else {
      if (d < 1) {
        d = 31 + d;
        if (m == 1) {
          m = 12;
        } else {
          m--;
        }
      } else if (d < 10) {
        d = "0" + d;
      }
      if (m < 10) {
        m = "0" + m;
      }
      tabDate.push(`${d}/${m}`);
      tabData.push(0);
    }
  }

  var optionsLine = {
    chart: {
      height: 250,
      width: 370,
      type: "line",
      stacked: true,
      animations: {
        enabled: true,
        easing: "linear",
        dynamicAnimation: {
          speed: 1000,
        },
      },
      dropShadow: {
        enabled: true,
        opacity: 0.3,
        blur: 5,
        left: -7,
        top: 22,
      },
      toolbar: {
        show: false,
      },
      zoom: {
        enabled: false,
      },
    },
    dataLabels: {
      enabled: false,
    },
    stroke: {
      curve: "straight",
      width: 5,
    },
    grid: {
      padding: {
        left: 0,
        right: 0,
      },
    },
    markers: {
      size: 0,
      hover: {
        size: 3,
      },
    },
    series: [
      {
        name: "Données",
        data: tabData,
      },
    ],
    xaxis: {
      categories: tabDate,
    },
    title: {
      text: "Vent",
      align: "left",
      style: {
        fontSize: "12px",
      },
    },
    subtitle: {
      text: "km/h",
      floating: true,
      align: "right",
      offsetY: 0,
      style: {
        fontSize: "22px",
      },
    },
    legend: {
      show: false,
    },
  };
  var chartLine = new ApexCharts(
    document.querySelector("#linechart"),
    optionsLine
  );
  chartLine.render();

  var optionsLine2 = {
    chart: {
      height: 250,
      width: 370,
      type: "line",
      stacked: true,
      animations: {
        enabled: true,
        easing: "linear",
        dynamicAnimation: {
          speed: 1000,
        },
      },
      dropShadow: {
        enabled: true,
        opacity: 0.3,
        blur: 5,
        left: -7,
        top: 22,
      },
      toolbar: {
        show: false,
      },
      zoom: {
        enabled: false,
      },
    },
    dataLabels: {
      enabled: false,
    },
    stroke: {
      curve: "straight",
      width: 5,
    },
    grid: {
      padding: {
        left: 0,
        right: 0,
      },
    },
    markers: {
      size: 0,
      hover: {
        size: 0,
      },
    },
    series: [
      {
        name: "Données",
        data: tabData,
      },
    ],
    xaxis: {
      categories: tabDate,
    },
    title: {
      text: "Temperature",
      align: "left",
      style: {
        fontSize: "12px",
      },
    },
    subtitle: {
      text: "°C",
      floating: true,
      align: "right",
      offsetY: 0,
      style: {
        fontSize: "22px",
      },
    },
    legend: {
      show: false,
    },
  };
  var chartLine2 = new ApexCharts(
    document.querySelector("#linechart2"),
    optionsLine2
  );
  chartLine2.render();

  var optionsLine3 = {
    chart: {
      height: 250,
      width: 370,
      type: "line",
      stacked: true,
      animations: {
        enabled: true,
        easing: "linear",
        dynamicAnimation: {
          speed: 1000,
        },
      },
      dropShadow: {
        enabled: true,
        opacity: 0.3,
        blur: 5,
        left: -7,
        top: 22,
      },
      toolbar: {
        show: false,
      },
      zoom: {
        enabled: false,
      },
    },
    dataLabels: {
      enabled: false,
    },
    stroke: {
      curve: "straight",
      width: 5,
    },
    grid: {
      padding: {
        left: 0,
        right: 0,
      },
    },
    markers: {
      size: 0,
      hover: {
        size: 0,
      },
    },
    series: [
      {
        name: "Données",
        data: tabData,
      },
    ],
    xaxis: {
      categories: tabDate,
    },
    title: {
      text: "Pression",
      align: "left",
      style: {
        fontSize: "12px",
      },
    },
    subtitle: {
      text: "P",
      floating: true,
      align: "right",
      offsetY: 0,
      style: {
        fontSize: "22px",
      },
    },
    legend: {
      show: false,
    },
  };
  var chartLine3 = new ApexCharts(
    document.querySelector("#linechart3"),
    optionsLine3
  );
  chartLine3.render();

  // DEFAULT AIRPORT
  console.log("Aéroport de Nantes - NTE");
  window.setInterval(() => {
    generateDateTitle();
    var cd = generateDate();
    var dateComplete = generateCompleteDate();
    var dateCompletePrevious = generatePreviousCompleteDate();
    //console.log("DATA From : " + dateCompletePrevious + " to : " + dateComplete);
    fetch(`http://localhost:8080/data/NTE/temperature`)
      .then((response) => response.json())
      .then((data) => {
        data.forEach((element) => {
          if (
            element.date > dateCompletePrevious &&
            element.date <= dateComplete
          ) {
            //console.log("Temperature : " + element.value);
            var val = Math.round(element.value * 100) / 100;
            chartTemp.updateSeries([val + 10]);
            chartTemp.updateOptions({
              plotOptions: {
                radialBar: {
                  dataLabels: {
                    value: {
                      formatter: function (value) {
                        return val + "°C";
                      },
                    },
                  },
                },
              },
            });
          }
        });
      });
    fetch(`http://localhost:8080/data/NTE/wind`)
      .then((response) => response.json())
      .then((data) => {
        data.forEach((element) => {
          if (
            element.date > dateCompletePrevious &&
            element.date <= dateComplete
          ) {
            //console.log("Vitesse du vent : " + element.value);
            chartWind.updateSeries([element.value]);
            chartWind.updateOptions({
              plotOptions: {
                radialBar: {
                  dataLabels: {
                    value: {
                      formatter: function (value) {
                        return element.value + " km/h";
                      },
                    },
                  },
                },
              },
            });
          }
        });
      });
    fetch(`http://localhost:8080/data/NTE/pressure`)
      .then((response) => response.json())
      .then((data) => {
        data.forEach((element) => {
          if (
            element.date > dateCompletePrevious &&
            element.date <= dateComplete
          ) {
            //console.log("Pression : " + Math.round(element.value * 10) / 10);
            var valeur = Math.round(element.value * 10) / 10;
            chartPressure.updateSeries([valeur / 20]);
            chartPressure.updateOptions({
              plotOptions: {
                radialBar: {
                  dataLabels: {
                    value: {
                      formatter: function (value) {
                        return valeur + " P";
                      },
                    },
                  },
                },
              },
            });
          }
        });
      });
    fetch(`http://localhost:8080/average/NTE/${cd}`)
      .then((response) => response.json())
      .then((data) => {
        chartProgress2.updateSeries([
          {
            name: "Process 1",
            data: [data.wind_average],
          },
        ]);
        chartProgress2.updateOptions({
          subtitle: {
            text: Math.round(data.wind_average * 100) / 100 + " km/h",
          },
        });
      });
    fetch(`http://localhost:8080/average/NTE/${cd}`)
      .then((response) => response.json())
      .then((data) => {
        chartProgress1.updateSeries([
          {
            name: "Process 1",
            data: [data.temperature_average + 10],
          },
        ]);
        chartProgress1.updateOptions({
          subtitle: {
            text: Math.round(data.temperature_average * 100) / 100 + "°C",
          },
        });
      });
    fetch(`http://localhost:8080/average/NTE/${cd}`)
      .then((response) => response.json())
      .then((data) => {
        chartProgress3.updateSeries([
          {
            name: "Process 1",
            data: [data.pressure_average / 18],
          },
        ]);
        chartProgress3.updateOptions({
          subtitle: {
            text: Math.round(data.pressure_average) + " Pascal",
          },
        });
      });

    var pre = generatePreviousDate();
    var cur = generateCurrentDate();
    //console.log(pre + " -> " + cur);
    fetch(`http://localhost:8080/data/NTE/wind/${pre}/${cur}`)
      .then((response) => response.json())
      .then((data) => {
        var tab = [];
        for (let index = 0; index < 9; index++) {
          tab.push(data[index].value);
        }
        chartLine.updateSeries([
          {
            data: tab,
          },
        ]);
      });
    fetch(`http://localhost:8080/data/NTE/temperature/${pre}/${cur}`)
      .then((response) => response.json())
      .then((data) => {
        var tab2 = [];
        for (let index = 0; index < 9; index++) {
          tab2.push(data[index].value);
        }
        chartLine2.updateSeries([
          {
            data: tab2,
          },
        ]);
      });
    fetch(`http://localhost:8080/data/NTE/pressure/${pre}/${cur}`)
      .then((response) => response.json())
      .then((data) => {
        var tab3 = [];
        for (let index = 0; index < 9; index++) {
          tab3.push(data[index].value);
        }
        chartLine3.updateSeries([
          {
            data: tab3,
          },
        ]);
      });
  }, 3000);
});
