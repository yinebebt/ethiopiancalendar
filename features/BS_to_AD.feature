Feature: convert dates from BS to AD using an API
  As an app-developer in Nepal
  I want to be able to send BS dates to an API endpoint and receive the corresponding AD dates
  So that I have a simple way to convert BS to AD dates, that can be used in other apps

  Scenario Outline: converting a valid Et date
    When a "GET" request is sent to the endpoint "/ad-from-et/<et-date>"
    Then the HTTP-response code should be "200"
    And the response content should be "<ad-date>"
    Examples:
      | et-date    | ad-date    |
      | 2060-04-01 | 2067-12-11 |
      | 2015-01-18 | 2022-09-28 |
      | 2003-04-16 | 2010-12-25 |

  Scenario Outline: converting an invalid BS date
    When a "GET" request is sent to the endpoint "/ad-from-et/<et-date>"
    Then the HTTP-response code should be "400"
    And the response content should be "not a valid date"
    Examples:
      | et-date       |
      | 2060-144-01   |
      | 40-01-01      |
      | date          |
      | 1987-04-05-01 |

  Scenario Outline: unhandled request types
    When a "<type>" request is sent to the endpoint "/ad-from-et/<et-date>"
    Then the HTTP-response code should be "400"
    And the response content should be "<response>"
    Examples:
      | type | et-date    | response                      |
      | POST | 87-04      | Could not create POST request |
      | PUT  | 87-04      | Could not create PUT request  |
      | POST | 2040-12-30 | Could not create POST request |
      | PUT  | 2040-12-30 | Could not create PUT request  |