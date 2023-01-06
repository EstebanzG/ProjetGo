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

  var trigoStrength = 3;
  var iteration = 11;
  var heure = 17;

  // RANDOM FUNCTIONS

  function getRandom() {
    var i = iteration;
    return (
      (Math.sin(i / trigoStrength) * (i / trigoStrength) +
        i / trigoStrength +
        1) *
      (trigoStrength * 2)
    );
  }

  function getRangeRandom(yrange) {
    return (
      Math.floor(Math.random() * (yrange.max - yrange.min + 1)) + yrange.min
    );
  }

  function generateMinuteWiseTimeSeries(baseval, count, yrange) {
    var i = 0;
    var series = [];
    while (i < count) {
      var x = baseval;
      var y =
        (Math.sin(i / trigoStrength) * (i / trigoStrength) +
          i / trigoStrength +
          1) *
        (trigoStrength * 2);

      series.push([x, y]);
      baseval += 300000;
      i++;
    }
    return series;
  }

  function getNewData(baseval, yrange) {
    var newTime = baseval + 300000;
    return {
      x: newTime,
      y: Math.floor(Math.random() * (yrange.max - yrange.min + 1)) + yrange.min,
    };
  }

  fetch(
    "http://localhost:8080/data/temperature/2023-01-06-14:00/2023-01-06-15:00"
  )
    .then((response) => {
      response.json();
      console.log(response);
    })
    .then((data) => {
      console.log(data);
      const valueT = Math.round(data[0].value * 10) / 10;
      // TEMPERATURE
      var optionsTemp = {
        chart: {
          height: 250,
          type: "radialBar",
        },
        series: [valueT + 10],
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
                  return value - 10 + "°C";
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
    })
    .catch((error) => {
      console.log("ERROR -> " + error);
    });

  fetch("http://localhost:8080/data/pressure")
    .then((response) => response.json())
    .then((data) => {
      const valueP = Math.round(data[0].value / 15);
      // PRESSURE
      var optionsPre = {
        chart: {
          height: 250,
          type: "radialBar",
        },

        series: [valueP],
        colors: ["seagreen"], // alert : indianred
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
                formatter: function (val) {
                  return val * 15 + " P";
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
      var chartPre = new ApexCharts(
        document.querySelector("#chartPre"),
        optionsPre
      );
      chartPre.render();
    })
    .catch((error) => {
      console.log("DATE INVALIDE -> " + error);
    });

  fetch("http://localhost:8080/data/wind")
    .then((response) => response.json())
    .then((data) => {
      const valueW = Math.round(data[0].value);
      // WIND
      var optionsWind = {
        chart: {
          height: 250,
          type: "radialBar",
        },

        series: [valueW],
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
                formatter: function (val) {
                  return val + " km/h";
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
    })
    .catch((error) => {
      console.log("DATE INVALIDE -> " + error);
    });

  // CHART LINE 1 = TEMPERATURE

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
      events: {
        animationEnd: function (chartCtx, opts) {
          const newData1 = chartCtx.w.config.series[0].data.slice();
          newData1.shift();

          // check animation end event for just 1 series to avoid multiple updates
          if (opts.el.node.getAttribute("index") === "0") {
            window.setTimeout(function () {
              chartCtx.updateOptions(
                {
                  series: [
                    {
                      data: newData1,
                    },
                  ],
                  subtitle: {
                    text: parseInt(getRandom() * Math.random()).toString(),
                  },
                },
                false,
                false
              );
            }, 300);
          }
        },
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
        data: generateMinuteWiseTimeSeries(
          new Date("12/12/2016 00:20:00").getTime(),
          12,
          {
            min: 30,
            max: 110,
          }
        ),
      },
    ],
    xaxis: {
      type: "datetime",
      range: 2700000,
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

  var chartLine = new ApexCharts(
    document.querySelector("#linechart"),
    optionsLine
  );
  chartLine.render();

  // CHART LINE 2 = PRESSURE

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
      events: {
        animationEnd: function (chartCtx, opts) {
          const newData1 = chartCtx.w.config.series[0].data.slice();
          newData1.shift();

          // check animation end event for just 1 series to avoid multiple updates
          if (opts.el.node.getAttribute("index") === "0") {
            window.setTimeout(function () {
              chartCtx.updateOptions(
                {
                  series: [
                    {
                      data: newData1,
                    },
                  ],
                  subtitle: {
                    text: parseInt(getRandom() * Math.random()).toString(),
                  },
                },
                false,
                false
              );
            }, 300);
          }
        },
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
        data: generateMinuteWiseTimeSeries(
          new Date("12/12/2016 00:20:00").getTime(),
          12,
          {
            min: 30,
            max: 110,
          }
        ),
      },
    ],
    xaxis: {
      type: "datetime",
      range: 2700000,
    },
    title: {
      text: "Pressure",
      align: "left",
      style: {
        fontSize: "12px",
      },
    },
    subtitle: {
      text: "Pascal",
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

  // CHART LINE 3 = WIND

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
      events: {
        animationEnd: function (chartCtx, opts) {
          const newData1 = chartCtx.w.config.series[0].data.slice();
          newData1.shift();

          // check animation end event for just 1 series to avoid multiple updates
          if (opts.el.node.getAttribute("index") === "0") {
            window.setTimeout(function () {
              chartCtx.updateOptions(
                {
                  series: [
                    {
                      data: newData1,
                    },
                  ],
                  subtitle: {
                    text: parseInt(getRandom() * Math.random()).toString(),
                  },
                },
                false,
                false
              );
            }, 300);
          }
        },
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
        data: generateMinuteWiseTimeSeries(
          new Date("12/12/2016 00:20:00").getTime(),
          12,
          {
            min: 30,
            max: 110,
          }
        ),
      },
    ],
    xaxis: {
      type: "datetime",
      range: 2700000,
    },
    title: {
      text: "Wind",
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

  var chartLine3 = new ApexCharts(
    document.querySelector("#linechart3"),
    optionsLine3
  );
  chartLine3.render();

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
      text: "44%",
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

  // REAL TIME FUNCTION

  window.setInterval(function () {
    iteration++;

    chartLine.updateSeries([
      {
        data: [
          ...chartLine.w.config.series[0].data,
          [chartLine.w.globals.maxX + 300000, getRandom()],
        ],
      },
    ]);

    chartLine2.updateSeries([
      {
        data: [
          ...chartLine2.w.config.series[0].data,
          [chartLine2.w.globals.maxX + 300000, getRandom()],
        ],
      },
    ]);

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
});
