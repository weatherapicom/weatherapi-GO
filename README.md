# Getting started


# Introduction

WeatherAPI.com provides access to weather and geo data via a JSON/XML restful API. It allows developers to create desktop, web and mobile applications using this data very easy.

We provide following data through our API:

  * Real-time weather

  * 10 day weather forecast

  * Astronomy

  * Time zone

  * Location data

  * Search or Autocomplete API

  * NEW: Historical weather

# Getting Started

You need to [signup](https://www.weatherapi.com/signup.aspx) and then you can find your API key under [your account](https://www.weatherapi.com/login.aspx), and start using API right away!

If you find any features missing or have any suggestions, please [contact us](https://www.weatherapi.com/contact.aspx).

# Authentication

API access to the data is protected by an API key. If at anytime, you find the API key has become vulnerable, please regenerate the key using Regenerate button next to the API key.

Authentication to the WeatherAPI.com API is provided by passing your API key as request parameter through an API .

  ##  key parameter
  key=YOUR_API_KEY


## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Weather%20API-GoLang&projectName=weatherapi_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=weatherapi_lib)

## How to Use

The following section explains how to use the WeatherapiLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=weatherapi_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=weatherapi_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=weatherapi_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=weatherapi_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=weatherapi_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Weather%20API-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=weatherapi_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=weatherapi_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| key | TODO: add a description |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [apis_pkg](#apis_pkg)

## <a name="apis_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".apis_pkg") apis_pkg

### Get instance

Factory for the ``` APIS ``` interface can be accessed from the package apis_pkg.

```go
aPIs := apis_pkg.NewAPIS()
```

### <a name="get_realtime_weather"></a>![Method: ](https://apidocs.io/img/method.png ".apis_pkg.GetRealtimeWeather") GetRealtimeWeather

> Current weather or realtime weather API method allows a user to get up to date current weather information in json and xml. The data is returned as a Current Object.Current object contains current or realtime weather information for a given city.


```go
func (me *APIS_IMPL) GetRealtimeWeather(
            q string,
            lang *string)(*models_pkg.CurrentJsonResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| q |  ``` Required ```  | Pass US Zipcode, UK Postcode, Canada Postalcode, IP address, Latitude/Longitude (decimal degree) or city name. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to learn more. |
| lang |  ``` Optional ```  | Returns 'condition:text' field in API in the desired language. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to check 'lang-code'. |


#### Example Usage

```go
q := "q"
lang := "lang"

var result *models_pkg.CurrentJsonResponse
result,_ = aPIs.GetRealtimeWeather(q, lang)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Error code 1003: Parameter 'q' not provided.Error code 1005: API request url is invalid.Error code 1006: No location found matching parameter 'q'Error code 9999: Internal application error. |
| 401 | Error code 1002: API key not provided.Error code 2006: API key provided is invalid |
| 403 | Error code 2007: API key has exceeded calls per month quota.<br />Error code 2008: API key has been disabled. |



### <a name="get_forecast_weather"></a>![Method: ](https://apidocs.io/img/method.png ".apis_pkg.GetForecastWeather") GetForecastWeather

> Forecast weather API method returns upto next 10 day weather forecast and weather alert as json. The data is returned as a Forecast Object.<br />Forecast object contains astronomy data, day weather forecast and hourly interval weather information for a given city.


```go
func (me *APIS_IMPL) GetForecastWeather(
            q string,
            days int64,
            dt *time.Time,
            unixdt *int64,
            hour *int64,
            lang *string)(*models_pkg.ForecastJsonResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| q |  ``` Required ```  | Pass US Zipcode, UK Postcode, Canada Postalcode, IP address, Latitude/Longitude (decimal degree) or city name. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to learn more. |
| days |  ``` Required ```  | Number of days of weather forecast. Value ranges from 1 to 10 |
| dt |  ``` Optional ```  | Date should be between today and next 10 day in yyyy-MM-dd format |
| unixdt |  ``` Optional ```  | Please either pass 'dt' or 'unixdt' and not both in same request.<br />unixdt should be between today and next 10 day in Unix format |
| hour |  ``` Optional ```  | Must be in 24 hour. For example 5 pm should be hour=17, 6 am as hour=6 |
| lang |  ``` Optional ```  | Returns 'condition:text' field in API in the desired language. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to check 'lang-code'. |


#### Example Usage

```go
q := "q"
days,_ := strconv.ParseInt("43", 10, 8)
dt := time.Now()
unixdt,_ := strconv.ParseInt("43", 10, 8)
hour,_ := strconv.ParseInt("43", 10, 8)
lang := "lang"

var result *models_pkg.ForecastJsonResponse
result,_ = aPIs.GetForecastWeather(q, days, dt, unixdt, hour, lang)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Error code 1003: Parameter 'q' not provided.Error code 1005: API request url is invalid.Error code 1006: No location found matching parameter 'q'Error code 9999: Internal application error. |
| 401 | Error code 1002: API key not provided.Error code 2006: API key provided is invalid |
| 403 | Error code 2007: API key has exceeded calls per month quota.<br />Error code 2008: API key has been disabled. |



### <a name="get_history_weather"></a>![Method: ](https://apidocs.io/img/method.png ".apis_pkg.GetHistoryWeather") GetHistoryWeather

> History weather API method returns historical weather for a date on or after 1st Jan, 2015 as json. The data is returned as a Forecast Object.


```go
func (me *APIS_IMPL) GetHistoryWeather(
            q string,
            dt *time.Time,
            unixdt *int64,
            endDt *time.Time,
            unixendDt *int64,
            hour *int64,
            lang *string)(*models_pkg.HistoryJsonResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| q |  ``` Required ```  | Pass US Zipcode, UK Postcode, Canada Postalcode, IP address, Latitude/Longitude (decimal degree) or city name. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to learn more. |
| dt |  ``` Required ```  | Date on or after 1st Jan, 2015 in yyyy-MM-dd format |
| unixdt |  ``` Optional ```  | Please either pass 'dt' or 'unixdt' and not both in same request.<br />unixdt should be on or after 1st Jan, 2015 in Unix format |
| endDt |  ``` Optional ```  | Date on or after 1st Jan, 2015 in yyyy-MM-dd format'end_dt' should be greater than 'dt' parameter and difference should not be more than 30 days between the two dates. |
| unixendDt |  ``` Optional ```  | Date on or after 1st Jan, 2015 in Unix Timestamp format<br />unixend_dt has same restriction as 'end_dt' parameter. Please either pass 'end_dt' or 'unixend_dt' and not both in same request. e.g.: unixend_dt=1490227200 |
| hour |  ``` Optional ```  | Must be in 24 hour. For example 5 pm should be hour=17, 6 am as hour=6 |
| lang |  ``` Optional ```  | Returns 'condition:text' field in API in the desired language. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to check 'lang-code'. |


#### Example Usage

```go
q := "q"
dt := time.Now()
unixdt,_ := strconv.ParseInt("43", 10, 8)
endDt := time.Now()
unixendDt,_ := strconv.ParseInt("43", 10, 8)
hour,_ := strconv.ParseInt("43", 10, 8)
lang := "lang"

var result *models_pkg.HistoryJsonResponse
result,_ = aPIs.GetHistoryWeather(q, dt, unixdt, endDt, unixendDt, hour, lang)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Error code 1003: Parameter 'q' not provided.Error code 1005: API request url is invalid.Error code 1006: No location found matching parameter 'q'Error code 9999: Internal application error. |
| 401 | Error code 1002: API key not provided.Error code 2006: API key provided is invalid |
| 403 | Error code 2007: API key has exceeded calls per month quota.<br />Error code 2008: API key has been disabled. |



### <a name="search_autocomplete_weather"></a>![Method: ](https://apidocs.io/img/method.png ".apis_pkg.SearchAutocompleteWeather") SearchAutocompleteWeather

> WeatherAPI.com Search or Autocomplete API returns matching cities and towns as an array of Location object.


```go
func (me *APIS_IMPL) SearchAutocompleteWeather(q string)([]*models_pkg.SearchJsonResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| q |  ``` Required ```  | Pass US Zipcode, UK Postcode, Canada Postalcode, IP address, Latitude/Longitude (decimal degree) or city name. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to learn more. |


#### Example Usage

```go
q := "q"

var result []*models_pkg.SearchJsonResponse
result,_ = aPIs.SearchAutocompleteWeather(q)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Error code 1003: Parameter 'q' not provided.Error code 1005: API request url is invalid.Error code 1006: No location found matching parameter 'q'Error code 9999: Internal application error. |
| 401 | Error code 1002: API key not provided.Error code 2006: API key provided is invalid |
| 403 | Error code 2007: API key has exceeded calls per month quota.<br />Error code 2008: API key has been disabled. |



### <a name="get_ip_lookup"></a>![Method: ](https://apidocs.io/img/method.png ".apis_pkg.GetIpLookup") GetIpLookup

> IP Lookup API method allows a user to get up to date information for an IP address.


```go
func (me *APIS_IMPL) GetIpLookup(q string)(*models_pkg.IpJsonResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| q |  ``` Required ```  | Pass IP address. |


#### Example Usage

```go
q := "q"

var result *models_pkg.IpJsonResponse
result,_ = aPIs.GetIpLookup(q)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Error code 1003: Parameter 'q' not provided.Error code 1005: API request url is invalid.Error code 1006: No location found matching parameter 'q'Error code 9999: Internal application error. |
| 401 | Error code 1002: API key not provided.Error code 2006: API key provided is invalid |
| 403 | Error code 2007: API key has exceeded calls per month quota.<br />Error code 2008: API key has been disabled. |



### <a name="get_time_zone"></a>![Method: ](https://apidocs.io/img/method.png ".apis_pkg.GetTimeZone") GetTimeZone

> Return Location Object


```go
func (me *APIS_IMPL) GetTimeZone(q string)(*models_pkg.TimezoneJsonResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| q |  ``` Required ```  | Pass US Zipcode, UK Postcode, Canada Postalcode, IP address, Latitude/Longitude (decimal degree) or city name. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to learn more. |


#### Example Usage

```go
q := "q"

var result *models_pkg.TimezoneJsonResponse
result,_ = aPIs.GetTimeZone(q)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Error code 1003: Parameter 'q' not provided.Error code 1005: API request url is invalid.Error code 1006: No location found matching parameter 'q'Error code 9999: Internal application error. |
| 401 | Error code 1002: API key not provided.Error code 2006: API key provided is invalid |
| 403 | Error code 2007: API key has exceeded calls per month quota.<br />Error code 2008: API key has been disabled. |



### <a name="get_astronomy"></a>![Method: ](https://apidocs.io/img/method.png ".apis_pkg.GetAstronomy") GetAstronomy

> Return Location and Astronomy Object


```go
func (me *APIS_IMPL) GetAstronomy(
            q string,
            dt *time.Time)(*models_pkg.AstronomyJsonResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| q |  ``` Required ```  | Pass US Zipcode, UK Postcode, Canada Postalcode, IP address, Latitude/Longitude (decimal degree) or city name. Visit [request parameter section](https://www.weatherapi.com/docs/#intro-request) to learn more. |
| dt |  ``` Required ```  | Date on or after 1st Jan, 2015 in yyyy-MM-dd format |


#### Example Usage

```go
q := "q"
dt := time.Now()

var result *models_pkg.AstronomyJsonResponse
result,_ = aPIs.GetAstronomy(q, dt)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Error code 1003: Parameter 'q' not provided.Error code 1005: API request url is invalid.Error code 1006: No location found matching parameter 'q'Error code 9999: Internal application error. |
| 401 | Error code 1002: API key not provided.Error code 2006: API key provided is invalid |
| 403 | Error code 2007: API key has exceeded calls per month quota.<br />Error code 2008: API key has been disabled. |



[Back to List of Controllers](#list_of_controllers)



