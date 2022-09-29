Feature: convert dates from BS to AD using an API
  As an app-developer in Nepal
  I want to be able to send BS dates to an API endpoint and receive the corresponding AD dates
  So that I have a simple way to convert BS to AD dates, that can be used in other apps

  Scenario Outline: converting a valid Et date
    When a "GET" request is sent to the endpoint "/ad-from-et/<et-date>"
    Then the HTTP-response code should be "<code>"
    And the response content should be "<response>"
    Examples:
      | et-date    | response    |  |code|
      | 2060-04-01 | 2067-12-11 |  |200|
      | 2015-01-18 | 2022-09-28 |  |200|
      | 2003-04-16 | 2010-12-25 |  |200|

  Scenario Outline: converting an invalid BS date
    When a "GET" request is sent to the endpoint "/ad-from-et/<et-date>"
    Then the HTTP-response code should be "400"
    And the response content should be "not a valid date"
    Examples:
      | et-date       |   
      | 2060-144-01   |   
      | 40-01-01      |   
      | 1987-04-02-01 |                  
  
  Scenario Outline: unhandled request types
    When a "<type>" request is sent to the endpoint "/ad-from-et/<et-date>"
    Then the HTTP-response code should be "404"
    And the response content should be "404 page not found"
    Examples:
      | type | et-date    |   
      | POST | 87-04      |   
      | PUT  | 87-04      |  
      | POST | 2040-12-30 |
      | PUT  | 2040-12-30 |


  Scenario Outline: Internal server error cases
    When a "GET" request is sent to the endpoint "/ad-from-et/<et-date>"
    Then the HTTP-response code should be "500"
    And the response content should be ""
    Examples:
      | et-date    |   
      | 2060-14    |   
      | 4478       |    
      | date-time  |   