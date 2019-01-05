window.onload = function() {
    var firstCall = displayLineChart.call();
    var timer =  setInterval(displayLineChart, 10000); 

}
var humidityArr1=[];
    var temperatureArr1=[];
    var batteryLevels1=[];
    var lightLevels1=[];
    var humidityArr2=[];
    var temperatureArr2=[];
    var batteryLevels2=[];
    var lightLevels2=[];
    var humidityArr3=[];
    var temperatureArr3=[];
    var batteryLevels3=[];
    var lightLevels3=[];
    var humidityArr4=[];
    var temperatureArr4=[];
    var batteryLevels4=[];
    var lightLevels4=[];
function getData(){
    fetch('http://localhost:7000/api/sensors')
    .then((resp) => resp.json()) // Transform the data into json
    .then(function(data) {
    for (i = 0; i < 40; ++i) {
        
         switch(data[i].sensorid)
         {
         case "a81758fffe031a82": {temperatureArr1.push(data[i].temperature); 
                                   humidityArr1.push(data[i].humidity);
                                   batteryLevels1.push(data[i].batterylevel);
                                   lightLevels1.push(data[i].light);
                                   break;
         }
         case "a81758fffe031a83": {temperatureArr2.push(data[i].temperature); 
                                   humidityArr2.push(data[i].humidity);
                                   batteryLevels2.push(data[i].batterylevel);
                                   lightLevels2.push(data[i].light);
                                   break;
         }
         case "a81758fffe031a81": {temperatureArr3.push(data[i].temperature); 
                                   humidityArr3.push(data[i].humidity);
                                   batteryLevels3.push(data[i].batterylevel);
                                   lightLevels3.push(data[i].light);
                                   break;
         }
         case "a81758fffe031a79": {temperatureArr4.push(data[i].temperature); 
                                   humidityArr4.push(data[i].humidity);
                                   batteryLevels4.push(data[i].batterylevel);
                                   lightLevels4.push(data[i].light); 
                                   break;
         } 
        }    
    } 
    })
}
var displayLineChart =function() {
    getData();
    var temperatureData = {
        labels: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
        datasets: [
            {
                label: "Sensor 1",
                fillColor: "rgba(220,220,220,0.2)",
                strokeColor: "rgba(220,220,220,1)",
                pointColor: "rgba(220,220,220,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(220,220,220,1)",
                data: temperatureArr1
            },
            {
                label: "Sensor 2",
                fillColor: "rgba(151,187,205,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: temperatureArr2
            },
            {
                label: "Sensor 3",
                fillColor: "rgba(121,100,162,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: temperatureArr3
            },
            {
                label: "Sensor 4",
                fillColor: "rgba(11,10,12,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: temperatureArr4
            }
        ]
    };

    var humidityData = {
        labels: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
        datasets: [
            {
                label: "Sensor 1",
                fillColor: "rgba(220,220,220,0.2)",
                strokeColor: "rgba(220,220,220,1)",
                pointColor: "rgba(220,220,220,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(220,220,220,1)",
                data: humidityArr1
            },
            {
                label: "Sensor 2",
                fillColor: "rgba(151,187,205,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: humidityArr2
            },
            {
                label: "Sensor 3",
                fillColor: "rgba(121,100,162,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: humidityArr3
            },
            {
                label: "Sensor 4",
                fillColor: "rgba(11,10,12,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: humidityArr4
            }
        ]
    };

    var batteryLevelsData = {
        labels: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
        datasets: [
            {
                label: "Sensor 1",
                fillColor: "rgba(220,220,220,0.2)",
                strokeColor: "rgba(220,220,220,1)",
                pointColor: "rgba(220,220,220,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(220,220,220,1)",
                data: batteryLevels1
            },
            {
                label: "Sensor 2",
                fillColor: "rgba(151,187,205,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: batteryLevels2
            },
            {
                label: "Sensor 3",
                fillColor: "rgba(121,100,162,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: batteryLevels3
            },
            {
                label: "Sensor 4",
                fillColor: "rgba(11,10,12,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: batteryLevels4
            }
        ]
    };

    var lightLevelsData = {
        labels: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
        datasets: [
            {
                label: "Sensor 1",
                fillColor: "rgba(220,220,220,0.2)",
                strokeColor: "rgba(220,220,220,1)",
                pointColor: "rgba(220,220,220,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(220,220,220,1)",
                data: lightLevels1
            },
            {
                label: "Sensor 2",
                fillColor: "rgba(151,187,205,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: lightLevels2
            },
            {
                label: "Sensor 3",
                fillColor: "rgba(121,100,162,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: lightLevels3
            },
            {
                label: "Sensor 4",
                fillColor: "rgba(11,10,12,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: lightLevels4
            }
        ]
    };
    
    var temperature = document.getElementById("temperatureChart").getContext("2d");
    var options = { };
    
    var tempChart = new Chart(temperature).Line(temperatureData, options);

    var humidity = document.getElementById("humidityChart").getContext("2d");
    var humidityChart = new Chart(humidity).Line(humidityData, options);

    var batteryLevel = document.getElementById("batteryLevelChart").getContext("2d");
    var humidityChart = new Chart(batteryLevel).Line(batteryLevelsData, options);

    var lightLevel = document.getElementById("lightLevelChart").getContext("2d");
    var humidityChart = new Chart(lightLevel).Line(lightLevelsData, options);

    temperatureArr1=[];
    batteryLevels1=[];
    lightLevels1=[];
    humidityArr2=[];
    temperatureArr2=[];
    batteryLevels2=[];
    lightLevels2=[];
    humidityArr3=[];
    temperatureArr3=[];
    batteryLevels3=[];
    lightLevels3=[];
    humidityArr4=[];
    temperatureArr4=[];
    batteryLevels4=[];
    lightLevels4=[];
  }

