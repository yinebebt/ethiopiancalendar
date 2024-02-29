Feature: convert dates from AD to ET using an API
  As an Ethiopian app-developer
  I want to send AD dates to an API endpoint and receive the corresponding Et dates
  So that I have a simple way to convert AD to Et dates, and use in other apps

  Scenario Outline: converting a valid AD date
    When a "GET" request is sent to the endpoint "/et-from-ad/<ad-date>"
    Then the HTTP-response code should be "200"
    And the response content should be "<et-date>"
    Examples:
      | ad-date    | et-date    |
      | 2060-4-01  | 2052-07-23 |
      | 1584-04-12 | 1576-08-07 |
      | 2020-02-29 | 2012-06-21 |

  Scenario Outline: converting an invalid AD date
    When a "GET" request is sent to the endpoint "/et-from-ad/<ad-date>"
    Then the HTTP-response code should be "400"
    And the response content should be "not a valid date"
    Examples:
      | ad-date       |   
      | 97-04-08      |
      | 1987-04-05-01 |
      | 22-22-29      |
  
  Scenario Outline: unhandled request types
    When a "<type>" request is sent to the endpoint "/et-from-ad/<ad-date>"
    Then the HTTP-response code should be "404"
    And the response content should be "404 page not found"
    Examples:
      | type | ad-date    |   
      | POST | 87-04      |   
      | PUT  | 87-04      |  
      | POST | 2040-12-30 |
      | PUT  | 2040-12-30 |

  Scenario Outline: Internal server error cases
    When a "GET" request is sent to the endpoint "/et-from-ad/<ad-date>"
    Then the HTTP-response code should be "500"
    And the response content should be ""
    Examples:
      | ad-date    |   
      | 2000-14    |   
      | 4478       |    
      | yinebeb    |   