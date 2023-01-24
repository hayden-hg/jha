### How to run:
* A valid API key is required to run this project, instructins on how to create one can be found [here](https://openweathermap.org/faq). 
* Once you have a valid API key, paste it into main.go on line 15 where the empty string currently is.


### Dependencies:
This project uses the [gorilla/mux](https://github.com/gorilla/mux) package. Download the dependency with:

    go get -u github.com/gorilla/mux

    
### Running the Project

First run:

    go mod tidy

Then, use command to run the API in local:

    go run . 
 
 To hit the endpoint, utilize the URI and host:
 
    http://localhost:8080/getWeather?lat={latitude}&long={longitude} 
  
  You can paste the following command into another tab of your terminal to try out the API: 
  
  
    curl http://localhost:8080/getWeather\?lat\=24.4\&long\=20.40



