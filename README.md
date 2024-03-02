<p align="center">
<img src="internal/assets/logo_medium.png" alt="logo" width="110" height="110">
</p>
<h1 align="center"><a href="https://pkg.go.dev/gitlab.com/Yinebeb-01/ethiopiancalendar">Ethiopian Calendar (ባሕረ-ሐሳብ)</a></h1>

## Description
The Ethiopian calendar(ባሕረ-ሀሳብ) is used to get Fasting and Holiday's specific date with in a year based on 
[EOTC](https://www.ethiopianorthodox.org/) calendar. It also designed to facilitate the conversion between Ethiopian dates (in the format yy-mm-dd) and 
Gregorian dates. Ethiopia follows its own calendar system, which consists of 13 months, each with 30 days. 

### Functionality
This API allows users to:
* Fetching all festivals with a year such as year, Basic data and Fasting dates.
* Convert Ethiopian dates to Gregorian dates.
* Convert Gregorian dates to Ethiopian dates.

### Usage
To utilize the API, simply send a date or year using the specified endpoints. For the date conversion,the API will 
respond with the converted date in JSON format whereas for the calendar, it will respond data in a json object.

#### Example Usage

1. Convert Gregorian Date to Ethiopian Date

```curl
GET /ad-to-et/{date}
```

* Parameters:

    `date:` The Gregorian date to convert (in _yy-mm-dd_ format).

**Example**:

```curl
GET /ad-to-et/2013-09-11
```

Response:

```json
{
"gregorian_date": "2021-05-22"
}
```

2. Convert Gregorian Date to Ethiopian Date

```curl
GET /et-to-ad/{date}
```

* Parameters:

    `date:` The Ethiopia date to convert (in _yy-mm-dd_ format).

**Example**:
```curl
GET /et-to-ad/2021-05-22
```
Response:
```json
{
"ethiopian_date": "2013-09-11"
}
```

3. Get Ethiopian Calendar such as Holidays and Fasting dates.

```curl
GET /bahire-hasab/{year}
```

* Parameters:

  `year:` The Ethiopian year.

**Example**:

```curl
GET /bahire-hasab/2016
```

Response:

```json
{
  "data": {
      "year": {
        "year":2016,
        "evangelist_num":0,
        "evangelist":"ዮሐንስ(John)",
        "day_of_the_week":"ማክሰኞ(Tuesday)"
      },
      "basic": {
        "medeb":11,
        "wenber":10,
        "abektie":20,
        "metiq":10,
        "beale_metiq":2,
        "mebaja_hamer":18,
        "nenewie":{
          "DateOfTheMonth":18,
          "MonthOfTheYear":6
        }
      },
    "fasting":{
      "abiy":{
        "DateOfTheMonth":2,
        "MonthOfTheYear":7
      },
      "debrezeit":{
        "DateOfTheMonth":29,
        "MonthOfTheYear":7
      },
      "hosanna":{
        "DateOfTheMonth":20,
        "MonthOfTheYear":8
      },
      "siklet":{
        "DateOfTheMonth":25,
        "MonthOfTheYear":8
      },
      "tinsaye":{
        "DateOfTheMonth":27,
        "MonthOfTheYear":8
      },
      "rkbekahnat":{
        "DateOfTheMonth":21,
        "MonthOfTheYear":9
      },
      "dihnet":{
        "DateOfTheMonth":19,
        "MonthOfTheYear":10
      },
      "hawariyat":{
        "DateOfTheMonth":17,
        "MonthOfTheYear":10
      },
      "erget":{
        "DateOfTheMonth":6,
        "MonthOfTheYear":10},
      "peraklitos":{
        "DateOfTheMonth":16,
        "MonthOfTheYear":10
      }
    }
  }
}
```

## Installation
Install using below go command:
```bash
go get gitlab.com/Yinebeb-01/ethiopiandateconverter
```

## Author
- [Yinebeb T.](https://github.com/Yinebeb-01/)